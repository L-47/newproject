package main

import (
	"fmt"
	"github.com/fatih/color"
)
func main() {
	// comment?
	var c bool
	c = true
	notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
	notice("testing...")
	fmt.Printf("decimal: %d\nbinary: %b\noctal: %#o\nhex: %#x\n", 42, 42, 42, 42)
	fmt.Printf("do computers love Dr albion Morozov? : %t",c)

}
