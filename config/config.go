package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configurations struct {
	MasterROSAddress string
	ROSNodeName      string
	DisablePinCalls  bool
}

var Configuration Configurations

func ImportConfig() {
	viper.SetConfigName("config.yml")
	viper.AddConfigPath("../config/")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&Configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	fmt.Println("ROS Value %s", Configuration.MasterROSAddress)
	fmt.Println("ROS Node Name %s", Configuration.ROSNodeName)
	fmt.Println("Disable Pin Triggers  %b", Configuration.DisablePinCalls)
}
