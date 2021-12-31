package servos

import (
	"github.com/cebbaker/roscontroller/config"
	"github.com/cebbaker/roscontroller/schemas"
	log "github.com/sirupsen/logrus"
	"periph.io/x/periph"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/experimental/devices/pca9685"
	"periph.io/x/periph/host"
)

var boardServos *pca9685.ServoGroup

type ServoStructure struct {
	Servo0 *pca9685.Servo
	Servo1 *pca9685.Servo
}

var RawServos = map[int]*pca9685.Servo{}

var lHost *periph.State
var bus i2c.BusCloser
var pca *pca9685.Dev

/******************** Init Call hardware interface and serv configurations **********************/
func InitHardwareServoInterface() {
	var err error

	lHost, err = host.Init()
	if err != nil {
		log.WithFields(log.Fields{"Error": err.Error()}).Fatal("Error While attempting to load servo Drivers")
	} else {
		log.Info("Servo Drivers loaded")
	}

	bus, err = i2creg.Open("")
	if err != nil {
		log.WithFields(log.Fields{"Error": err.Error()}).Fatal("Error Opening I2C Bus")
	} else {
		log.Info("Created I2C instance")
	}

	pca, err = pca9685.NewI2C(bus, pca9685.I2CAddr)
	if err != nil {
		log.WithFields(log.Fields{"Error": err.Error()}).Fatal("Error creating I2C link to servo card")
	} else {
		log.Info("Linked I2C to Servo Card.")
	}

	if err := pca.SetPwmFreq(60 * physic.Hertz); err != nil {
		log.WithFields(log.Fields{"Error": err.Error()}).Fatal("Error Setting PWA Freq")
	} else {
		log.Info("Pwm Frequency set")
	}
	if err := pca.SetAllPwm(0, 0); err != nil {
		log.WithFields(log.Fields{"Error": err.Error()}).Fatal("Error Setting PWA Freq for each Servo")
	} else {
		log.Info("PWA set for all Servos")
	}
	boardServos = pca9685.NewServoGroup(pca, 145, 605, 0, 180)
	log.Info("Servo Group Configured ")
	RawServos[0] = boardServos.GetServo(0)
	RawServos[0].SetMinMaxAngle(0, 110)
	RawServos[1] = boardServos.GetServo(0)
	RawServos[1].SetMinMaxAngle(0, 180)

}

func MoveServo(servoRequest *schemas.ExtentionAmount) {
	log.Info("Begining Move Servo Hardware Call")
	lServo := Servos[servoRequest.Servo]

	if config.Configuration.DisablePinCalls {
		log.WithFields(log.Fields{"Servo Name": servoRequest.Servo,
			"Servo Nbr": lServo.Config.ServoNbr}).Info("Servo Pins Inactive for testing.")
		return
	}
	if lHost == nil {
		InitHardwareServoInterface()
	}
	log.Info("Init of Hardware Board is complete.")
	MoveServoAngle(RawServos[lServo.Config.ServoNbr], physic.Angle(int(servoRequest.Input)))

}

func MoveServoAngle(lF *pca9685.Servo, lAng physic.Angle) error {
	log.Info("attempting Phyical Servo Movment")
	if err := lF.SetAngle(lAng); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
