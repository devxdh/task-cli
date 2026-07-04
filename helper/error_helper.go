package helper

import (
	"fmt"
)

func HandleErr(err error, msg ...string) bool {
	if err == nil {
		return false
	}

	finalErrMsg := "An error occurred"
	if len(msg) > 0 && msg[0] != "" {
		finalErrMsg = msg[0]
	}

	fmt.Printf("%s: %v\n", finalErrMsg, err)
	return true
}
