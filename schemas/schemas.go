package schemas

import (
	"github.com/aler9/goroslib/pkg/msg"
)

type ExtentionAmount struct {
	Input uint32
}

type ExtentionFinal struct {
	Output uint32
}

type ExtentionFeedBack struct {
	PercentComplete float32
}

type ServoAction struct {
	msg.Package `ros:"shared_actions"`
	ExtentionAmount
	ExtentionFinal
	ExtentionFeedBack
}

type ServoConfig struct {
	Name     string
	Min      int
	Max      int
	Rest     int
	ServoNbr int
}

type ServoState struct {
	CurrentPointer int
}

type ServoS struct {
	Config  ServoConfig
	Pointer ServoState
}

const RIGHT_UPPER_BICEP = "rub"
