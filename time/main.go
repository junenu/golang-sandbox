package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	tz, _ := time.LoadLocation("America/Los_Angeles")
	future := time.Date(2015, time.October, 21, 16, 29, 0, 0, tz)

	fmt.Println(now.String())
	fmt.Println(future.Format(time.RFC3339Nano))

	// Nanosecond, Millisecond, Second, Minute, Hour
	fiveMuinutes := 5 * time.Minute

	var seconds int = 10
	tenSeconds := time.Duration(seconds) * time.Second

	past := time.Date(1955, time.November, 5, 6, 0, 0, 0, time.UTC)
	dur := time.Now().Sub(past)

	fmt.Println(fiveMuinutes)
	fmt.Println(tenSeconds)
	fmt.Println(dur)
}
