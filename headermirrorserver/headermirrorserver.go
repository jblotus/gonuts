package headermirrorserver

import "fmt"
import "io"
import "net/http"

func HeaderMirror(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Here are your headers\n")

	for key, value := range req.Header {
		//todo output to screen
		fmt.Println(key, value)
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
