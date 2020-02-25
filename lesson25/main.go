package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID int64
	Name string
	Age int
	Active bool
}

func main() {
	u1 := User{
		ID:     123,
		Name:   "q1mi",
		Age:    18,
		Active: true,
	}
	b, err := json.Marshal(&u1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
