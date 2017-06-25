package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildURL(t *testing.T) {
	ci := &Gitlab{
		Host: "https://gitlab.com",
	}

	params := url.Values{
		"token": []string{"foo"},
		"ref":   []string{"bar"},
	}

	assert.Equal(t, "https://gitlab.com/api/v4/projects/1234/trigger/pipeline?ref=bar&token=foo", ci.buildURL("1234", params))
}

func TestHostNotFound(t *testing.T) {
	ci := &Gitlab{
		Host: "https://foo.bar",
	}

	params := url.Values{
		"token": []string{"foo"},
		"ref":   []string{"bar"},
	}

	err := ci.trigger("1234", params, nil)
	assert.NotNil(t, err)
}

func TestNilBody(t *testing.T) {
	ci := &Gitlab{
		Host: "https://gitlab.com",
	}

	params := url.Values{
		"token": []string{"foo"},
		"ref":   []string{"bar"},
	}

	err := ci.trigger("1234", params, nil)
	assert.Nil(t, err)
}

func TestResponse404Body(t *testing.T) {
	type body struct {
		Message string `json:"message"`
	}

	ci := &Gitlab{
		Host: "https://gitlab.com",
	}

	params := url.Values{
		"token": []string{"foo"},
		"ref":   []string{"bar"},
	}

	resp := &body{}
	err := ci.trigger("1234", params, resp)
	assert.Equal(t, "404 Not Found", resp.Message)
	assert.Nil(t, err)
}
