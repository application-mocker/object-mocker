package utils

import (
	"github.com/google/uuid"
	"object-mocker/config"
)

func init() {
	uuid.SetNodeID([]byte(config.Config.Application.NodeId))
	uuid.SetClockSequence(config.Config.Application.ClockSequence)
}

func NewUUIDString() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}
