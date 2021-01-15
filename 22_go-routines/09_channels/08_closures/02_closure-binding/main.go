package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		fmt.Println(v)  // print a b c
		go func(u string) {
			fmt.Println(u) // 可能是乱需输出的 a b c
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}

/*
To bind the current value of v to each closure as it is launched,
one must modify the inner loop to create a new variable each iteration.
One way is to pass the variable as an argument to the closure.

In this example, the value of v is passed as an argument to the anonymous function.
That value is then accessible inside the function as the variable u.

SOURCE:
https://golang.org/doc/faq#closures_and_goroutines
*/
