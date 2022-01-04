package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var interval time.Duration

func init() {
	flag.DurationVar(&interval, "interval", time.Minute*10, "Check Interval")
	flag.Parse()

}
func main() {
	fmt.Println(interval)
	for c := time.Tick(interval); ; <-c {
		log.Println(c)
	}
}
