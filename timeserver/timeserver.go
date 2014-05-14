package timeserver

import "time"
import "github.com/go-martini/martini"

func NewServer() *martini.ClassicMartini {
	m := martini.Classic()
	m.Get("/", func() string {
		t := time.Now()
		return "The time is: " + t.Format(time.UnixDate)
	})
	m.Run()
	return m
}
