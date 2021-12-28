package actions

import (
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/aler9/goroslib"
	"github.com/cebbaker/roscontroller/schemas"
	"github.com/cebbaker/roscontroller/servos"
	log "github.com/sirupsen/logrus"
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

/*****************************************************************************/
/* Init All Actiond defined for Object                                       */
/*****************************************************************************/
func InitAllActions(lNode *goroslib.Node) {
	log.Infoln("Initiating All Actions")
	c := make(chan os.Signal, 1)
	servos.Init()
	for _, element := range servos.Servos {
		go RegisterROSServo(element, c, lNode)
	}

	signal.Notify(c, os.Interrupt)
	<-c
}

/*****************************************************************************/
/* Register Each Server Defined                                              */
/*****************************************************************************/
func RegisterROSServo(lServo schemas.ServoS, c chan os.Signal, lNode *goroslib.Node) {
	lNode.Log(goroslib.LogLevelInfo, "Registering Servo "+strconv.Itoa(lServo.Config.ServoNbr))

	sas, err := goroslib.NewSimpleActionServer(goroslib.SimpleActionServerConf{
		Node:   lNode,
		Name:   lServo.Config.Name,
		Action: &schemas.ServoAction{},
		OnExecute: func(sas *goroslib.SimpleActionServer, goal *schemas.ExtentionAmount) {
			lNode.Log(goroslib.LogLevelInfo, "Servo action Triggered "+strconv.Itoa(int(goal.Input)))
			// publish a feedback

			// wait some time
			servos.MoveServo(goal.Input)

			time.Sleep(500 * time.Millisecond)

			// set the goal as succeeded
			sas.SetSucceeded(&schemas.ExtentionFinal{
				Output: 444,
			})
		},
	})
	if err != nil {
		panic(err)
	}
	defer sas.Close()
	<-c

}
