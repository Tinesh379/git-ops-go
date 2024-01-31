package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetProjectDetails() {
	projectId := "46132905"
	url := "https://gitlab.com/api/v4/projects/" + projectId

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
		return
	}

	// Add headers to the request
	req.Header.Set("PRIVATE-TOKEN", "glpat-zwhxXS2sM1yQdhjCmAdC")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
		return
	}

	// Print the response body
	//fmt.Println("Response Body:", string(body))

	// parse json data and get name and path
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
		return
	}

	info := []string{"name", "path_with_namespace"}
	var myValues []string

	for _, key := range info {
		value, ok := data[key].(string)
		if !ok {
			log.Fatal("cannot find any key with name")
		}
		myValues = append(myValues, value)

	}

	projectSlug := strings.Replace(myValues[1], myValues[0], "", 1)
	fmt.Println("name of the project is:", myValues[0])
	fmt.Println("path of the project is:", projectSlug)

}
