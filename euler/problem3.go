package euler

import "fmt"
import "sort"

//http://projecteuler.net/problem=3
//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?

const TARGET int = 600851475143

// Generate numbers until the limit max.
// after the 2, all the prime numbers are odd
// Send a channel signal when the limit is reached
func Generate(max int, ch chan<- int) {
	ch <- 2
	for i := 3; i <= max; i += 2 {
		ch <- i
	}
	ch <- -1 // signal that the limit is reached
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for i := <-in; i != -1; i = <-in {
		if i%prime != 0 {
			out <- i
		}
	}
	out <- -1
}

func CalcPrimeFactors(number_to_factorize int) []int {
	rv := []int{}
	ch := make(chan int)
	go Generate(number_to_factorize, ch)
	for prime := <-ch; (prime != -1) && (number_to_factorize > 1); prime = <-ch {
		for number_to_factorize%prime == 0 {
			number_to_factorize = number_to_factorize / prime
			rv = append(rv, prime)
		}
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	return rv
}

func GetLargestInt(a []int) int {
	sort.Ints(a)
	return a[len(a)-1]
}

func Problem3() int {
	factors := CalcPrimeFactors(TARGET)
	largest := GetLargestInt(factors)
	fmt.Println(largest)
	return largest
}
