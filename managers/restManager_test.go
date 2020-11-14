package managers

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func TestRestManager_GetMilestones(t *testing.T) {
	client := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(getMilestones())),
			Header:     make(http.Header),
		}
	})

	ms, err := getRestManager(client).GetMilestones()
	require.NoError(t, err)
	m := *ms
	require.Equal(t, 2, len(m))
	require.Equal(t, 12, m[0].Id)
	require.Equal(t, 3, m[0].Iid)
	require.Equal(t, "10.0", m[0].Title)
	require.Equal(t, 13, m[1].Id)
	require.Equal(t, 4, m[1].Iid)
	require.Equal(t, "11.0", m[1].Title)
}

func TestRestManager_UploadFile(t *testing.T) {
	client := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(getUploads())),
			Header:     make(http.Header),
		}
	})
	i := "image.png"
	rm := getRestManager(client)
	_, err := rm.UploadFile(i)
	require.Error(t, err)

	_, err = fm.manager.Create(i)
	require.NoError(t, err)

	md, err := rm.UploadFile(i)
	require.NoError(t, err)
	require.Equal(t, "![image](/uploads/225bf851d24565ee9dfaeac81f585399/image.png)", md)
}

func TestRestManager_Create(t *testing.T) {
	client := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(getResponse())),
			Header:     make(http.Header),
		}
	})

	rm := getRestManager(client)
	iid, err := rm.Create("")
	require.NoError(t, err)
	require.Equal(t, 7, iid)
}

func TestRestManager_Update(t *testing.T) {
	client := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(getResponse())),
			Header:     make(http.Header),
		}
	})

	rm := getRestManager(client)
	iid, err := rm.Update(1, "")
	require.NoError(t, err)
	require.Equal(t, 7, iid)
}

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func getRestManager(c *http.Client) *RestManager {
	return NewRestManager(123, "abc", c, fm)
}

func getMilestones() string {
	return `
	[
		{
			"id": 12,
			"iid": 3,
			"title": "10.0"
		},
		{
			"id": 13,
			"iid": 4,
			"title": "11.0"
		}
	]`
}

func getUploads() string {
	return `
	{
		"art": "image.png",
		"url": "/uploads/225bf851d24565ee9dfaeac81f585399/image.png",
		"markdown": "![image](/uploads/225bf851d24565ee9dfaeac81f585399/image.png)"
	}`
}

func getResponse() string {
	return `
	{
		"iid" : 7
	}`
}
