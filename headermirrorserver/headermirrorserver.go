package headermirrorserver

import "fmt"
import "io"
import "net/http"

func HeaderMirror(w http.ResponseWriter, req *http.Request) {
	fmt.Println("sent header to client")

	for key, value := range req.Header {
		var output string

		output += key
		output += ": "

		for _, eachHeader := range value {
			output += eachHeader
			output += "\n"
		}
		io.WriteString(w, output)
	}
}

func NewServer() {
	fmt.Println("Starting a new headermirrorserver")

	http.HandleFunc("/", HeaderMirror)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
