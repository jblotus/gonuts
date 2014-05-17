package concurrentrng

import "fmt"
import "time"
import "math/rand"

const MAX_ARRAY_SIZE = 10000

var rands [MAX_ARRAY_SIZE]int

func OutputRandomNumber(i int) {

	rands[i] = rand.Int()
	fmt.Println("fired iteration #", i)
	if i == MAX_ARRAY_SIZE-1 {
		fmt.Println("last one fired")
	}
}
func SeedRng() {
	rand.Seed(time.Now().UTC().UnixNano())
}
func RunRng() {
	fmt.Println("generating numbers")
	SeedRng()
	for i := 0; i < MAX_ARRAY_SIZE; i++ {
		go OutputRandomNumber(i)
	}
	fmt.Scanln()
	fmt.Println(rands)

}
