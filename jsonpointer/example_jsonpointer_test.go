package jsonpointer_test

import (
	"encoding/json"
	"fmt"

	"github.com/mattn/go-jsonpointer"
)

func Example_set() {
	input := `{"foo":[1,true,2]}`
	var obj interface{}
	if err := json.Unmarshal([]byte(input), &obj); err != nil {
		panic(err)
	}
	if err := jsonpointer.Set(obj, "/foo/1", false); err != nil {
		panic(err)
	}
	got, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", got)
	if err != nil {
		panic(err)
	}

	// Output:
	// {"foo":[1,false,2]}
}

func Example_nil() {
	input := `{"foo":[1,true,2]}`
	var obj interface{}
	if err := json.Unmarshal([]byte(input), &obj); err != nil {
		panic(err)
	}
	var nilPointer *interface{}
	if err := jsonpointer.Set(obj, "/bar", nilPointer); err != nil {
		panic(err)
	}
	got, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", got)

	// Output:
	// {"bar":null,"foo":[1,true,2]}
}
