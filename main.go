package main

import (
	"fmt"

	"github.com/hashstore/hashlogic/base"
)

func main() {
	p, err := base.ParseTagMatch(`// xya
	"text & \tabc" & ( !a | 3 | "c | b" | d ) x`)
	fmt.Println(p)
	fmt.Println(err)
}
