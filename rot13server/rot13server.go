package rot13server

import "github.com/joshlf13/rot13"
import "github.com/go-martini/martini"

func NewServer() *martini.ClassicMartini {
    m := martini.Classic()
    
    m.Get("/", func() string {
        return "enter a route like /foo to get it transformed into rot13"    
    })
    
    m.Get("/:text", func(params martini.Params) string {
        text := params["text"]
        textBytes := []byte(text)
        
        for i, v := range textBytes {
            textBytes[i] = rot13.Rot13(v)
        }
        return string(textBytes)
    })
    m.Run()
    return m
}