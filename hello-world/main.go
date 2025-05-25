/*
My First Go Program
A simple Go demo that prints user input, current timestamp, and variable memory addresses.
Reads input with fmt.Scanln, captures time via time.Now().Unix(), and uses reflection to inspect pointer addresses.
Output includes formatted messages and debug info.
*/

package main

import (
	"fmt"
	"reflect"
	"time"
)

var timestamp int64
var s string

func helloworld(user_prompt string, timenow int64) {
	fmt.Printf("Hello world! \nThe user said %s\n", user_prompt)
	fmt.Println("Time from 1970.1.1 in second:", timenow)
}

func print_var_address(variables ...interface{}) {
	for i, v := range variables {
		if reflect.ValueOf(v).Kind() == reflect.Ptr {
			fmt.Printf("Variable %d address: %p\n", i, v)
		}
	}
}

func main() {
	_, err := fmt.Scanln(&s)
	if err != nil {
		return
	}
	timestamp = time.Now().Unix()
	helloworld(s, timestamp)
	print_var_address(&err, &timestamp, &s)
}
