package fileprinter

import "io"
import "os"
import "fmt"

const BUFFER_LENGTH = 2

func IfErrorPrint(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func PrintFile() {
	file, err := os.Open("fileprinter/demo.txt")
	IfErrorPrint(err)

	defer file.Close()

	for {

		buffer := make([]byte, BUFFER_LENGTH)
		_, err := file.Read(buffer)

		if err == io.EOF {
			fmt.Println("\n")
			break
		}
		IfErrorPrint(err)
		for _, i := range buffer {
			fmt.Printf("%c", i)
		}
	}
}
