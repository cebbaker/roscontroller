package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/aler9/goroslib"
	"github.com/cebbaker/roscontroller/config"
	"github.com/cebbaker/roscontroller/schemas"
)

// define a custom action.
// unlike the standard library, an .action file is not needed.

func main() {
	// create a node and connect to the master
	config.ImportConfig(true)
	n, err := goroslib.NewNode(goroslib.NodeConf{
		Name:          "test_Node",
		MasterAddress: config.Configuration.MasterROSAddress,
	})
	if err != nil {
		panic(err)
	}
	defer n.Close()

	// create a simple action client
	sac, err := goroslib.NewSimpleActionClient(goroslib.SimpleActionClientConf{
		Node:   n,
		Name:   "Servo0",
		Action: &schemas.ServoAction{},
	})
	if err != nil {
		panic(err)
	}
	defer sac.Close()

	// wait for the server
	sac.WaitForServer()

	done := make(chan struct{})

	// send a goal
	err = sac.SendGoal(goroslib.SimpleActionClientGoalConf{
		Goal: &schemas.ExtentionAmount{
			Input: 60,
		},
		OnDone: func(state goroslib.SimpleActionClientGoalState, res *schemas.ExtentionFinal) {
			fmt.Println("result:", res)
			close(done)
		},
		OnFeedback: func(fb *schemas.ExtentionFeedBack) {
			fmt.Println("feedback", fb)
		},
	})
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	select {
	// goal is done
	case <-done:

	// handle CTRL-C
	case <-c:
	}
}
