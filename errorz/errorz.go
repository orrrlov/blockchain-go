package errorz

import "log"

func Panic(err error) {
	log.Panic(err)
}