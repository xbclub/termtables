// Copyright 2013 Apcera Inc. All rights reserved.

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/xbclub/termtables/term"
)

func main() {
	size, err := term.GetSize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Lines %d  Columns %d\n", size.Lines, size.Columns)
}
