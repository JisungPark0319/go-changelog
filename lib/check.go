package lib

import (
	"fmt"
	"os"
)

func IfError(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
