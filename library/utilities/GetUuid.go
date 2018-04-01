package utilities

import "github.com/twinj/uuid"

/*
 * get uuid
 */
func GetUuid() uuid.Uuid {
	theUuid := uuid.NewV4()
	//return uuid.Formatter(theUuid, uuid.Format())
	return theUuid
}
