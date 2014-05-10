package main

import "time"
import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {    
      t := time.Now()
      return "The time is: " + t.Format(time.UnixDate)
  })
  m.Run()
}