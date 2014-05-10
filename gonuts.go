package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    println("someone hit /")
    return "Hello World"
  })
  m.Run()
}