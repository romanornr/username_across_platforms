package service

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

func (cm *clientMock) GetValue(url string) (*http.Response, error) {
	return getRequestFunc(url)
}

func TestUsernameCheck_Success(t *testing.T) {
	urls := []string {
		"http://twitter.com/stevensunflash",
		"http://instagram.com/stevensunflash",
		"http://dev.to/stevensunflash",
	}

	getRequestFunc = func(url string) (response *http.Response, e error) {
		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	}
	client.ClientCall = &clientMock{}

	result := UsernameService.UsernameCheck(urls)
	assert.NotNil(t, result)
	assert.EqualValues(t, len(result), 3)
}

func TestUsername_No_Match(t *testing.T) {
	urls := []string {
		"http://twitter.com/no_match_username",
		"http://instagram.com/no_match_username",
		"http://dev.to/no_match_username",
	}
	getRequestFunc = func(url string) (response *http.Response, e error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
		}, nil
	}
	client.ClientCall = &clientMock{}
	result := UsernameService.UsernameCheck(urls)

	assert.EqualValues(t, len(result),0)
}

func TestUsernameCheck_Url_Invalid(t *testing.T) {
	urls := []string{
		"http://wrong.com/rnr_0",
		"http://wrong.com/rnr_0",
		"http://wrong.to/rnr_0",
	}
	getRequestFunc = func(url string) (response *http.Response, e error) {
		return nil, errors.New("cant_access_resource")
	}
	client.ClientCall = &clientMock{}
	result := UsernameService.UsernameCheck(urls)
	assert.EqualValues(t, len(result), 0)
}