package skewd

import (
	"errors"
)

// TODO:
// This naming is not Go style?

type TaskInterface interface {
	Name() string
	Run(*Logger, interface{}) string
	CanReschedule(error) bool
}
