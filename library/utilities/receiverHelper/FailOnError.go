package receiverHelper

import (
	"log"
)

//show error message
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
