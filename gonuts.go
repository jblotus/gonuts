package main

import "fmt"
import "./timeserver/"
import "./rot13server/"
import "./headermirrorserver/"
import "./sha256server/"
import "./euler/"

func main() {

	for true {
		fmt.Println("Press 1 to launch timeserver")
		fmt.Println("Press 2 to launch rot13server")
		fmt.Println("Press 3 to launch headermirrorserver")
		fmt.Println("Press 4 to launch sha256server")
		fmt.Println("Press 5 to solve euler #1")
		fmt.Println("Press 6 to solve euler #2")

		var choice string

		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		switch choice {
		case "1":
			timeserver.NewServer()
		case "2":
			rot13server.NewServer()
		case "3":
			headermirrorserver.NewServer()
		case "4":
			sha256server.NewServer()
		case "5":
			//move these euler probs to a diff menu
			euler.Problem1()
		case "6":
			//move these euler probs to a diff menu
			euler.Problem2()
		default:
			fmt.Println("That is not a valid choice")
		}
	}

}
