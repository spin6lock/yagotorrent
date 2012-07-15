package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main(){
	h := sha1.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
}
