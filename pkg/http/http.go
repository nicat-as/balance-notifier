package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func DoRequest[REQ comparable, RES any](url string, method Method, headers map[string]string, body REQ) (RES, error) {
	var bodyReader io.Reader
	if !IsZero(body) {
		b, err := json.Marshal(&body)
		if err != nil {
			log.Fatalf("error %v", err)
			return Zero[RES](), err
		}
		bodyReader = bytes.NewReader(b)
	}
	request, err := http.NewRequest(string(method), url, bodyReader)
	var responseType RES
	if err != nil {
		return responseType, err
	}
	if headers != nil {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return responseType, err
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		b, err := ioutil.ReadAll(response.Body)
		//log.Println(string(b))
		err = json.Unmarshal(b, &responseType)
		if err != nil {
			log.Fatalf("parsing error: %v", err)
			return responseType, err
		}
		return responseType, nil
	} else {
		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("bad request parsing error: %v", err)
			return responseType, err
		}
		return responseType, errors.New(string(b))
	}
}

type Method string

const (
	Get  Method = "GET"
	Post Method = "POST"
)

func IsZero[T comparable](v T) bool {
	return v == *new(T)
}

func Zero[T any]() T {
	return *new(T)
}
