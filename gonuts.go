package main

import "fmt"
import "./timeserver/"

func main() {
    
    for true {
        fmt.Println("Press 1 to launch timeserver")

        var choice string

        _, err := fmt.Scanln(&choice)

        if err != nil {
            fmt.Println("Error: ", err)
        }

        if choice == "1" {
            timeserver.NewServer();   
        } else {
            fmt.Println("That is not a valid choice")
        }
    }
    
}