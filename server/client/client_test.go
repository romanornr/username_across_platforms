package client

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func NewFakeClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestGetWithRoundTripper_Success(t *testing.T) {
	client := NewFakeClient(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200, //the real api status code may be 404, 422, 500. But we dont care
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}, nil
	})
	api := clientCall{*client}
	url := "https://twitter.com/rnr_0"
	body, err := api.GetValue(url)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.EqualValues(t, http.StatusNotFound, body.StatusCode)
}

func TestGetWithRoundTripper_No_Match(t *testing.T) {
	client := NewFakeClient(func(req *http.Request) (response *http.Response, e error) {
		return &http.Response{
			StatusCode: 404,
			Header:     make(http.Header),
		}, nil
	})
	api := clientCall{*client}
	url := "https://twitter.com/rnr_0"
	body, err := api.GetValue(url)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.EqualValues(t, http.StatusNotFound, body.StatusCode)
}
