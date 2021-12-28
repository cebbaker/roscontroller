package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configurations struct {
	MasterROSAddress string
	ROSNodeName      string
	DisablePinCalls  bool
}

var Configuration Configurations

func ImportConfig(isTester bool) {
	viper.SetConfigName("config.yml")
	if isTester {
		viper.AddConfigPath("../config/")
	} else {
		viper.AddConfigPath("./config/")
	}

	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("Unable to Register Node using IP")
	}
	err := viper.Unmarshal(&Configuration)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("Unable to Read Configuration File")
	}
	log.WithFields(log.Fields{
		"ROS_Master_Address": Configuration.MasterROSAddress}).Infoln("Imported Env Var")
	log.WithFields(log.Fields{
		"ROS_Node_Name": Configuration.ROSNodeName}).Infoln("Imported Env Var")
	log.WithFields(log.Fields{
		"Disable_Pin_Triggers": Configuration.DisablePinCalls}).Infoln("Imported Env Var")

}
