package main

import "fmt"

func errorHandler(err error) {
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}
