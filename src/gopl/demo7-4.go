package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

type Student struct {
}

func main() {
	var sile1 int = 10

	var a interface{} = sile1
	fmt.Println(a == a)

}

func fmtDemo() {
	var i, j int
	n, _ := fmt.Sscanf("23", "%d %d", &i, &j)
	fmt.Println(n)
}

func demoPeriod() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
