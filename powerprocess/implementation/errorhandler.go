package implementation

import (
	"log"
)

func ErrorHandle(err error, errorMessage string) {

	if err != nil {
		log.Fatal(err, errorMessage)
	}
}
func ErrorHandleThow(err error) error {
	if err != nil {
		return err
	}
	return nil
}
