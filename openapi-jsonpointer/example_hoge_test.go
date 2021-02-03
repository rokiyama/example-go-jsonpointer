package openapi_jsonpointer_test

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/jsonpointer"
)

func Example_set() {
	input := `{"foo":[1,true,2]}`
	var obj interface{}
	if err := json.Unmarshal([]byte(input), &obj); err != nil {
		panic(err)
	}
	p, err := jsonpointer.New("/foo/1")
	if err != nil {
		panic(err)
	}
	if _, err := p.Set(obj, false); err != nil {
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
	p, err := jsonpointer.New("/bar")
	if err != nil {
		panic(err)
	}
	if _, err := p.Set(obj, nil); err != nil {
		panic(err)
	}
	got, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", got)

	// Output:
	// {"foo":[1,true,2]}
}
