package utils

import (
	"errors"
	"log"
)

// Panic if error
func Pife(err error) {
	if err != nil {
		panic(err)
	}
}

// Log if the err != nil
func Loge(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ErrIfNotOk(ok bool, msg string) error {
	if ok {
		return nil
	}
	return errors.New(msg)
}
