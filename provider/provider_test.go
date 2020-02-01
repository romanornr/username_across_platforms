package provider

import (
	"errors"
	"github.com/romanornr/username_across_platforms/server/client"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	getRequestFunc func(url string) (*http.Response, error)
)

type clientMock struct {}

// mocking client call
func (cm *clientMock) GetValue(url string) (*http.Response, error) {
	return getRequestFunc(url)
}

func TestCheckUrls_success(t *testing.T) {
	getRequestFunc = func(url string) (response *http.Response, e error) {
		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	}

	client.ClientCall = &clientMock{}

	url := "https://twitter.com/rnr_0"
	ch := make(chan string)
	go Checker.CheckUrl(url, ch)
	result := <-ch
	assert.NotNil(t, result)
	assert.EqualValues(t, "https://twitter.com/rnr_0", result)
}

func TestCheckUrls_Not_Existent_Url(t *testing.T) {
	getRequestFunc = func(url string) (response *http.Response, e error) {
		return nil, errors.New("there is an error here")
	}
	client.ClientCall = &clientMock{}

	url := "https://invalid_url/rnr_0"
	ch := make(chan string)
	go Checker.CheckUrl(url, ch)
	err := <-ch
	assert.NotNil(t, err)
	assert.EqualValues(t, "cant_access_resource", err)
}

func TestCheckUrls_Username_Dont_Exist(t *testing.T) {
	getRequestFunc = func(url string) (response *http.Response, e error) {
		return &http.Response {
			StatusCode: http.StatusNotFound,
		}, nil
	}
	client.ClientCall = &clientMock{}
	url := "https://twitter.com/random_xxx23343r"
	ch := make(chan string)
	go Checker.CheckUrl(url, ch)
	result := <- ch
	assert.NotNil(t, result)
	assert.EqualValues(t, "no_match", result)
}