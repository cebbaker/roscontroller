package main

import (
	"log"

	"github.com/cebbaker/roscontroller/actions"

	"github.com/aler9/goroslib"
)

func main() {
	// create a node and connect to the master
	lNode, err := goroslib.NewNode(goroslib.NodeConf{
		Name:          "Right_Servo_Controller",
		MasterAddress: "ceb-All-Series:11311",
	})
	if err != nil {
		panic(err)
	}
	defer lNode.Close()
	lNode.Log(goroslib.LogLevelInfo, "Created New Node")
	log.Println("Created New Node")
	actions.RegisterConsumers(lNode)

}
