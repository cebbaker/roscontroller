package actions

import (
	"os"
	"os/signal"
	"time"

	"github.com/aler9/goroslib"
	"github.com/cebbaker/roscontroller/schemas"
	"github.com/cebbaker/roscontroller/servos"
	log "github.com/sirupsen/logrus"
)

/*****************************************************************************/
/* Init All Actiond defined for Object                                       */
/*****************************************************************************/
func InitAllActions(lNode *goroslib.Node) {
	log.Infoln("Initiating All Actions")
	c := make(chan os.Signal, 1)
	servos.Init()
	RegisterROSServo(schemas.RIGHT_ARM, c, lNode)

	signal.Notify(c, os.Interrupt)
	<-c
}

/*****************************************************************************/
/* Register Each Server Defined                                              */
/*****************************************************************************/
func RegisterROSServo(actionName string, c chan os.Signal, lNode *goroslib.Node) {
	lNode.Log(goroslib.LogLevelInfo, "Registering Action Service "+actionName)

	sas, err := goroslib.NewSimpleActionServer(goroslib.SimpleActionServerConf{
		Node:      lNode,
		Name:      actionName,
		Action:    &schemas.ServoAction{},
		OnExecute: ServoAction,
	})
	if err != nil {
		panic(err)
	}
	defer sas.Close()
	<-c

}

func ServoAction(sas *goroslib.SimpleActionServer, goal *schemas.ExtentionAmount) {

	switch goal.Servo {
	case schemas.SERVO_UPPER_BICEP:
		log.WithFields(log.Fields{"Servo Name": goal.Servo,
			"Movement Request ": goal.Input}).Infoln("Upper Bicep Servo Triggering")
		servos.MoveServo(goal)

		time.Sleep(500 * time.Millisecond)

		// set the goal as succeeded
		sas.SetSucceeded(&schemas.ExtentionFinal{
			Output: goal.Input,
		})
	case schemas.SERVO_ELBO:
		log.WithFields(log.Fields{"Servo Name": goal.Servo,
			"Movement Request ": goal.Input}).Infoln("Elbo Triggering")
		servos.MoveServo(goal)

		time.Sleep(500 * time.Millisecond)

		// set the goal as succeeded
		sas.SetSucceeded(&schemas.ExtentionFinal{
			Output: goal.Input,
		})
	}

}
