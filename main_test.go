package main

import (
	"app/TestYourSpeed/router"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGivenUnexpectedParameterReturnsBadRequestStatusCode(t *testing.T) {

	// Arrange
	testerID := "7"
	ts := httptest.NewServer(router.SetupRouter())
	defer ts.Close()
	request := fmt.Sprintf("%s/test/%s", ts.URL, testerID)

	// Act
	resp, err := http.Get(request)
	if err != nil {
		panic(err)
	}

	// Assert
	assert.Equal(t, 400, resp.StatusCode)
}

func TestGiven1ParameterFastTestsSpeedSuccesfully(t *testing.T) {

	// Arrange
	testerID := "1"
	ts := httptest.NewServer(router.SetupRouter())
	defer ts.Close()
	request := fmt.Sprintf("%s/test/%s", ts.URL, testerID)

	// Act
	resp, err := http.Get(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Assert
	assert.Equal(t, 200, resp.StatusCode)

	var response map[string]float64
	err = json.Unmarshal(body, &response)
	dwnld, dwnldExists := response["Download"]
	_, upldExists := response["Upload"]

	assert.Nil(t, err)
	assert.True(t, dwnldExists)
	assert.True(t, upldExists)
	assert.True(t, dwnld > 0)
	// assert.True(t, upld > 0)
}

func TestGiven2ParameterSpeedTestTestsSpeedSuccesfully(t *testing.T) {

	// Arrange
	testerID := "2"
	ts := httptest.NewServer(router.SetupRouter())
	defer ts.Close()
	request := fmt.Sprintf("%s/test/%s", ts.URL, testerID)

	// Act
	resp, err := http.Get(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Assert
	assert.Equal(t, 200, resp.StatusCode)

	var response map[string]float64
	err = json.Unmarshal(body, &response)
	dwnld, dwnldExists := response["Download"]
	upld, upldExists := response["Upload"]

	assert.Nil(t, err)
	assert.True(t, dwnldExists)
	assert.True(t, upldExists)
	assert.True(t, dwnld > 0)
	assert.True(t, upld > 0)
}

func TestFastHasHigherPerformanceThanSpeedTest(t *testing.T) {

	// Arrange
	testerID1 := "1"
	testerID2 := "2"
	ts := httptest.NewServer(router.SetupRouter())
	defer ts.Close()
	request1 := fmt.Sprintf("%s/test/%s", ts.URL, testerID1)
	request2 := fmt.Sprintf("%s/test/%s", ts.URL, testerID2)

	// Act
	start1 := time.Now()
	_, err1 := http.Get(request1)
	if err1 != nil {
		panic(err1)
	}
	duration1 := time.Since(start1)

	start2 := time.Now()
	_, err2 := http.Get(request2)
	if err2 != nil {
		panic(err2)
	}
	duration2 := time.Since(start2)

	// Assert
	assert.True(t, duration1 < duration2)
}

func TestCallingNotMappedEndpointReturnsNotFoundStatusCodeWithExpectedMessage(t *testing.T) {

	// Arrange
	ts := httptest.NewServer(router.SetupRouter())
	request := fmt.Sprintf("%s/dummy", ts.URL)
	defer ts.Close()

	// Act
	resp, err := http.Get(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Assert
	assert.Equal(t, 404, resp.StatusCode)
	assert.Equal(t, string(body), "Call \".../test/1\" for fast\nCall \".../test/2\" for testspeed")
}
