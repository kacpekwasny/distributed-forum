package utils

import "log"

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
