package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// create data structures to hold JSON data
// this needs to match the fields in the JSON feed
// see sample data fetch at bottom of this file

// create an individual entry data structure
// this does not need to hold every field, just the ones we want
type Entry struct {
	Title     string
	Author    string
	URL       string
	Permalink string
}

// the feed is the full JSON data structure
// this sets up the array of Entry types (defined above)
type Feed struct {
	Data struct {
		Children []struct {
			Data Entry
		}
	}
}

func main() {
	// url of REST endpoint we are grabbing data from
	url := "https://www.reddit.com/r/golang/new.json"

	// fetch url
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error fetching:", err)
	}
	// defer response close
	defer resp.Body.Close()

	// confirm we received an OK status
	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Error Status not OK:", resp.StatusCode)
	}

	// read the entire body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading body:", err)
	}

	// create an empty instance of Feed struct
	// this is what gets filled in when unmarshaling JSON
	var entries Feed
	if err := json.Unmarshal(body, &entries); err != nil {
		log.Fatalln("Error decoing JSON", err)
	}

	// loop through the children and create entry objects
	for _, ed := range entries.Data.Children {
		entry := ed.Data
		log.Println(">>>")
		log.Println("Title   :", entry.Title)
		log.Println("Author  :", entry.Author)
		log.Println("URL     :", entry.URL)
		log.Printf("Comments: http://reddit.com%s \n", entry.Permalink)
	}
}

// 2022/05/13 20:27:12 Error fetching: Get "https://www.reddit.com/r/golang/new.json": EOF
// exit status 1
// pass
