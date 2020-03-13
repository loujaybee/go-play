package main

import (
	"encoding/json"
)

type response struct {
    Repos []string
}

type repos struct {
    Repos []response
}

func businesslogic(team string) string {

	/*
		1. Reformat response to be JSON:API response format
		2. Swap from using input arguments to ones passed from an API event
	*/

	repos := make(map[string][]string)

	/* Store repo's by team */
	repos["dx"] = append(repos["dx"], "https://github.com/getndazn/one")
	repos["dx"] = append(repos["dx"], "https://github.com/getndazn/two")

	/* Get repo by team, and store in a map */
	result := make(map[string][]string)
	result["repos"] = repos[team]

	jsonString, _ := json.Marshal(result)
	jsonStr := string(jsonString)

	return jsonStr
}
