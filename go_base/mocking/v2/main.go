package main

import (
	"fmt"
	"io"
	"time"
)
const num = 3
const FinalWord = "Go!"

func Countdown(w io.Writer) {
	for i := num; i > 0; i-- {
		time.Sleep(1 *time.Second)
		fmt.Fprintf(w, "%d\n", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(w, FinalWord)
}	