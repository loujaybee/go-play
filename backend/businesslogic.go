package main

type response struct {
    Repos []string
}

type repos struct {
    Repos []response
}

func businesslogic(team string) string {

	/*
		1. Needs to take a string "team" and then grab the related repo's to that string
		2. Need an object of teams with repo's underneath
		3. Reformat response to be JSON:API response format
	*/

	return "{ \"repos\": [\"one\"] }"
}
