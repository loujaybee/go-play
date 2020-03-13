package main

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestStringLiteralResponse(t *testing.T) {
	result := businesslogic("Hello")
	if result != "{ \"repos\": [\"one\"] }" {
		t.Errorf("businesslogic(Hello) = %s; want JSON", result)
	}
}

type response struct {
    Repos []string
}

func TestFirstRepoRespones(t *testing.T) {
	result := businesslogic("Hello")
	var myStoredVariable response

	fmt.Println(result)
	json.Unmarshal([]byte(result), &myStoredVariable)

	if myStoredVariable.Repos[0] != "one" {
		t.Errorf("businesslogic(Hello) = %s; want JSON", result)
	}
}
