// Package assert is a set of helper functions for testing to make it less verbose and repetitive
package assert

import (
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Condition assertion fails the test if the condition is false, you must provide a message explaining the condition
func Condition(t *testing.T, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("\n\033[31m%s:%d: "+msg+"\033[39m\n↓\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		t.FailNow()
	}
}

// Nil fails the test if an err is not nil.
func Nil(t *testing.T, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("\n\033[31m%s:%d: unexpected error: %s\033[39m\n↓\n", filepath.Base(file), line, err.Error())
		t.FailNow()
	}
}

// NotNil fails the test if an err is nil
func NotNil(t *testing.T, err error) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("\n\033[31m%s:%d: expected error but got nil\033[39m\n↓\n", filepath.Base(file), line)
		t.FailNow()
	}
}

// Equal fails the test if exp is not equal to act
func Equal(t *testing.T, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("\n\033[31m%s:%d: \n\texp: %#v\n\n\tgot: %#v\033[39m\n↓\n", filepath.Base(file), line, exp, act)
		t.FailNow()
	}
}

// Panic fails the test if the code does not panic
func Panic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("The code did not panic")
	}
}
