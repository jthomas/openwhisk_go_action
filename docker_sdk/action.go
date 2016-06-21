package main

import "os"
import "encoding/json"
import "log"

type Params struct {
	Payload string `json:"payload"`
}

type Result struct {
	Reversed string `json:"reversed"`
}

// extract invocation parameters, passed as JSON string argument on command-line.
func params() Params {
	var params Params
	source := os.Args[1]
	buf := []byte(source)
	if err := json.Unmarshal(buf, &params); err != nil {
		log.Fatal(err)
	}
	return params
}

// convert struct back to JSON for response
func return_result(result Result) {
	buf, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(buf)
}

func main() {
	input := params()

	// reverse the string passed from invocation parameters
	chars := []rune(input.Payload)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	result := Result{
		Reversed: string(chars),
	}

	return_result(result)
}
