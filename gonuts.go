package main

import "fmt"
import "./timeserver/"
import "./rot13server/"
import "./headermirrorserver/"
import "./sha256server/"

func main() {

	for true {
		fmt.Println("Press 1 to launch timeserver")
		fmt.Println("Press 2 to launch rot13server")
		fmt.Println("Press 3 to launch headermirrorserver")
		fmt.Println("Press 4 to launch sha256server")

		var choice string

		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		if choice == "1" {
			timeserver.NewServer()
		} else if choice == "2" {
			rot13server.NewServer()
		} else if choice == "3" {
			headermirrorserver.NewServer()
		} else if choice == "4" {
			sha256server.NewServer()
		} else {
			fmt.Println("That is not a valid choice")
		}
	}

}
