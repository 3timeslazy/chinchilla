package main

import (
	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
