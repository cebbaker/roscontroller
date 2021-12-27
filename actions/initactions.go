package actions

import (
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/aler9/goroslib"
	"github.com/cebbaker/roscontroller/schemas"
	"github.com/cebbaker/roscontroller/servos"
)

func RegisterConsumers(lNode *goroslib.Node) {
	lNode.Log(goroslib.LogLevelInfo, "Registering Servo 0")

	sas, err := goroslib.NewSimpleActionServer(goroslib.SimpleActionServerConf{
		Node:   lNode,
		Name:   "Servo0",
		Action: &schemas.ServoAction{},
		OnExecute: func(sas *goroslib.SimpleActionServer, goal *schemas.ExtentionAmount) {
			lNode.Log(goroslib.LogLevelInfo, "Servo action Triggered "+strconv.Itoa(int(goal.Input)))
			// publish a feedback

			// wait some time
			servos.MoveServo(goal.Input)

			time.Sleep(500 * time.Millisecond)

			// set the goal as succeeded
			sas.SetSucceeded(&schemas.ExtentionFinal{
				Output: 12222222,
			})
		},
	})
	if err != nil {
		panic(err)
	}
	defer sas.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}
