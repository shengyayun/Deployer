package tools

import (
	"fmt"
	"time"
)

//Out is to log a message
func Out(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " " + msg)
}
