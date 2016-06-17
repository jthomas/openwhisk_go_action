package main

import "os"
import "encoding/json"
import "log"

type Params struct {
	// fill in this with input parameters
}

type Result struct {
	// fill this in with output parameters
}

func params() Params {
	var params Params
	source := os.Args[1]
	buf := []byte(source)
	if err := json.Unmarshal(buf, &params); err != nil {
		log.Fatal(err)
	}
	return params
}

func return_result(result Result) {
	buf, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(buf)
}

func main() {
	result := Result{}
	return_result(result)
}
