package main

import (
	"fmt"
	"io"
)

func CountDown(w io.Writer) {
	fmt.Fprint(w, "3")
} 