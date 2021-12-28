package servos

import (
	"github.com/cebbaker/roscontroller/schemas"
	log "github.com/sirupsen/logrus"
)

var Servos = map[string]schemas.ServoS{}

func Init() {
	lServo := schemas.ServoS{
		Config:  schemas.ServoConfig{Name: schemas.RIGHT_UPPER_BICEP, Min: 60, Max: 110, Rest: 85, ServoNbr: 0},
		Pointer: schemas.ServoState{CurrentPointer: 0}}
	CreateServo(lServo)
}

func CreateServo(lServo schemas.ServoS) {
	if _, ok := Servos[lServo.Config.Name]; ok {
		Servos[lServo.Config.Name] = lServo
		log.WithFields(log.Fields{"Servo": lServo}).Info("Updateing an existing Servo")
	} else {
		Servos[lServo.Config.Name] = lServo
		log.WithFields(log.Fields{"Servo": lServo}).Info("Creating a new Servo")
	}
}
