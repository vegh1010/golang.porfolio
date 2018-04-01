package logHelper

import "time"

//Latin for Mode is Modus
//handles debug mode changes for LogManager
//validateFunc will return a bool based on table query or json file
//if true call enableFunc else disableFunc
func Modus(validateFunc func() bool, enableFunc func(), disableFunc func()) {
	for {
		if validateFunc() {
			enableFunc()
		} else {
			disableFunc()
		}
		time.Sleep(1 * time.Minute)
	}
}
