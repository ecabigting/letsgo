package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// check if args are present
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a valid github username.")
		os.Exit(1)
	}
	input := os.Args[1]

	// print the input
	fmt.Printf("You typed:%s\n", input)

	// build the url request with the provided gh username
	url := "https://api.github.com/users/" + input + "/events"
	// initiate the Get request
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Request Error with: %v", err)
	}
	// close the response body
	defer response.Body.Close() //etc/
	// check if reponse is status 200 OK
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Erroor with status code: %d", response.StatusCode)
	}
	// read thedyu response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading the response: %v", err)
	}

	// create a vairable to hold the events response
	var events []Event

	// unmarshall the json response into a variable
	if err := json.Unmarshal(body, &events); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// build the activity output
	var activities []Activity
	for _, event := range events {
		// check if repo and type exist in the activities slice
		exists := false
		for index, activity := range activities {
			if activity.Type == event.Type && activity.RepoName == event.Repo.Name {
				exists = true
				activities[index].Count += 1
				break
			}
		}

		if !exists {
			activities = append(activities, Activity{
				Count:    1,
				Type:     event.Type,
				RepoName: event.Repo.Name,
			})
		}

	}

	fmt.Println("Output:")
	for _, activity := range activities {
		switch {
		case activity.Type == "PushEvent":
			fmt.Println("- Pushed", activity.Count, "commits to", activity.RepoName)
		case activity.Type == "CreateEvent":
			fmt.Println("- Created new repo", activity.RepoName)
		case activity.Type == "WatchEvent":
			fmt.Println("- Starred", activity.RepoName)
		}
	}
}
