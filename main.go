package main

import (
	"fmt"
	"time"
)

func main() {
	const layout = "3:04 PM"
	tm, _ := time.Parse(layout, "5:43 PM")
	fmt.Println(tm)
}
