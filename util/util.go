package util

import (
	"log"
	"reflect"
)

func IsNotZero(v interface{}) bool {
	t := reflect.TypeOf(v)
	if !t.Comparable() {
		log.Printf("type is not comparable: %v", t)
		return false
	}
	return v != reflect.Zero(t).Interface()
}
