package main

import (
	"fmt"

	"github.com/realtemirov/tasks/project5/bigint"
)

func main() {
	a, err := bigint.NewInt("10000")

	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("10000")
	if err != nil {
		panic(err)
	}

	res := bigint.Add(a, b)

	fmt.Println(res.Value())
}
