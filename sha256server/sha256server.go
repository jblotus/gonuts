package sha256server

import "crypto/sha256"
import "encoding/base64"
import "github.com/go-martini/martini"

func NewServer() *martini.ClassicMartini {
	m := martini.Classic()

	m.Get("/", func() string {
		return "enter a route like /foo to get it transformed into sha256"
	})

	m.Get("/:text", func(params martini.Params) string {
		text := params["text"]
		textBytes := []byte(text)

		hasher := sha256.New()
		hasher.Write(textBytes)

		sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		return string(sha)
	})
	m.Run()
	return m
}
