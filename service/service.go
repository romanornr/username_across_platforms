package service

import "github.com/romanornr/username_across_platforms/provider"

type usernameCheck struct {}

type usernameService interface {
	UsernameCheck(url []string) []string
}
var UsernameService usernameService = &usernameCheck{}

func (u *usernameCheck) UsernameCheck(urls []string) []string {
	c := make(chan string)
	var links []string
	matchingLinks := []string{}

	for _, url := range urls {
		go provider.Checker.CheckUrl(url, c)
	}
	for i := 0; i < len(urls); i++ {
		links = append(links, <-c)
	}

	// Remove the "no_match" and "cant_access_resource" values from links array
	for _, v := range links {
		if v == "cant_access_resource" {
			continue
		}
		if v == "no_match" {
			continue
		}
		matchingLinks = append(matchingLinks, v)
	}
	return matchingLinks
}

