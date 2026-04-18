package main

import (
	"fmt"
	"os/user"
	"time"
)

func helloWorld() {
	fmt.Println("Hello, World!", time.Now())
	currentUser, err := user.Current()

	if err != nil {
		fmt.Println("Error fetching user:", err)
		return
	}

	fmt.Println("Current user:", currentUser.Username)
}
