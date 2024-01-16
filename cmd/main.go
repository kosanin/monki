package main

import (
	"fmt"
	"time"
)

var replWelcomeMessage string = `Monki language REPL
time = %s
>>>
`

func main() {
	fmt.Println(fmt.Sprintf(replWelcomeMessage, time.Now()))
}
