package provider

import "github.com/romanornr/username_across_platforms/server/client"

type checkInterface interface {
	CheckUrl(string, chan string)
}
type checker struct {}

var Checker checkInterface = &checker{}

func (check *checker) CheckUrl(url string, c chan string) {
	resp, err := client.ClientCall.GetValue(url)
	// we could not access that endpoint
	if err != nil {
		c <- "cant_access_resource"
		return
	}

	// if the url exist but the username is not found, sent no_match to channel
	if resp.StatusCode > 229 {
		c <- "no_match"
	}

	// if desired username is found, send to channel
	if resp.StatusCode == 200 {
		c <- url
	}
}