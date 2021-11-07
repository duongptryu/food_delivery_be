package common

import (
	"errors"
	"fmt"
)

var (
	RecordNotFound = errors.New("record not foud")
)

func AppRecovery() {
	if err := recover(); err != nil {
		fmt.Println("Recovery error - ", err)
	}
}
