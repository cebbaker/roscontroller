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
