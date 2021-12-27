package servos

import (
	"fmt"
	"log"

	"github.com/cebbaker/roscontroller/config"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/experimental/devices/pca9685"
	"periph.io/x/periph/host"
)

var servos *pca9685.ServoGroup

type ServoStructure struct {
	Servo0 *pca9685.Servo
	Servo1 *pca9685.Servo
}

func MoveServo(amount uint32) {
	if config.Configuration.DisablePinCalls {
		fmt.Println("Pin Functionality is turns off")
		return
	}

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}

	pca, err := pca9685.NewI2C(bus, pca9685.I2CAddr)
	if err != nil {
		log.Fatal(err)
	}

	if err := pca.SetPwmFreq(60 * physic.Hertz); err != nil {
		log.Fatal(err)
	}
	if err := pca.SetAllPwm(0, 0); err != nil {
		log.Fatal(err)
	}
	servos = pca9685.NewServoGroup(pca, 145, 605, 0, 180)
	log.Println("Servo Group Created")

	lH := ServoStructure{}
	lH.Servo0 = servos.GetServo(0)
	lH.Servo0.SetMinMaxAngle(0, 110)
	lH.Servo1 = servos.GetServo(1)
	lH.Servo1.SetMinMaxAngle(0, 180)

	MoveServoAngle(lH.Servo0, physic.Angle(int(amount)))

}

func MoveServoAngle(lF *pca9685.Servo, lAng physic.Angle) error {
	if err := lF.SetAngle(lAng); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
