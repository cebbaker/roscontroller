package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/aler9/goroslib"
	"github.com/cebbaker/roscontroller/actions"
	"github.com/cebbaker/roscontroller/config"
)

func main() {
	// create a node and connect to the master
	config.ImportConfig(false)

	lNode, err := goroslib.NewNode(goroslib.NodeConf{
		Name:          config.Configuration.ROSNodeName,
		MasterAddress: config.Configuration.MasterROSAddress,
	})
	if err != nil {
		log.WithFields(log.Fields{"Name": config.Configuration.ROSNodeName,
			"MasterAddress": config.Configuration.MasterROSAddress,
			"Error":         err.Error()}).Panicln("Cannot Create a New Node.")
		panic(err)
	}
	defer lNode.Close()
	lNode.Log(goroslib.LogLevelInfo, "Created New Node"+config.Configuration.ROSNodeName)

	actions.InitAllActions(lNode)
}
