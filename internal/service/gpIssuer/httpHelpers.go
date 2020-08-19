package gpIssuer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

var auth = struct {
	AccessToken string
	ExpiresAt   time.Time
	TokenType   string
}{}

func post(url string, body interface{}) *reqExecutor {
	return baseRequest("POST", url, body)
}

func get(url string) *reqExecutor {
	return baseRequest("GET", url, nil)
}

func baseRequest(method string, url string, body interface{}) *reqExecutor {
	var req *http.Request
	var err error
	if body == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		// marshal de json body
		bytesArr, _ := json.Marshal(body)
		// Create a new request using http
		req, err = http.NewRequest(method, url, bytes.NewBuffer(bytesArr))
	}

	if err == nil {
		req.Header.Add("Content-Type", "application/json")
		// add authorization header to the req
		req.Header.Add("Authorization", "Bearer " + auth.AccessToken)
		// Gp requires to add a RequestId in the header
		req.Header.Add("RequestId", time.Now().String()) // todo vary reqId
	}
	return &reqExecutor{req: req, err: err}
}

type reqExecutorInterface interface {
	execute(interface{}) error
}

type reqExecutor struct {
	req *http.Request
	err error
}

func (re *reqExecutor) execute(resultPtr interface{}) error {
	if re.err != nil {
		return re.err
	}
	// we check the if token expires during next minute (entire minute just as a precaution)
	// 		(has expired)       (will expire in the next minute)       not expired
	// |----------------------|---------------------------------|-----------------------|
	//                       Now							Now + 1 min
	// |-------------------Refresh token------------------------|--No need to refresh---|
	if auth.ExpiresAt.Before(time.Now().Add(time.Minute)) {
		// if so we generate a new one
		getToken()
	}
	// we execute the requests
	resp, err := client.Do(re.req)
	if err != nil {
		return err
	}
	// check the status code
	if (resp.StatusCode / 100) != 2 {
		gpError := &GPError{}
		if err := json.NewDecoder(resp.Body).Decode(&gpError); err != nil {
			return err
		}
		return gpError
	}
	// if nothing was expected, nothing will be returned
	if resultPtr == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(&resultPtr)
}
