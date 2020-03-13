package main

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestStringLiteralResponse(t *testing.T) {
	result := businesslogic("dx")
	if result != "{\"repos\":[\"https://github.com/getndazn/one\",\"https://github.com/getndazn/two\"]}" {
		t.Errorf("businesslogic(dx) = %s; want JSON", result)
	}
}

func TestFirstRepoRespones(t *testing.T) {
	result := businesslogic("dx")
	var myStoredVariable response

	fmt.Println(result)
	json.Unmarshal([]byte(result), &myStoredVariable)

	if myStoredVariable.Repos[0] != "https://github.com/getndazn/one" {
		t.Errorf("businesslogic(dx) = %s; want JSON", result)
	}
}
