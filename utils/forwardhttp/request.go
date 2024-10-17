package forwardhttp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GET(URL string, headers ...map[string]string) ([]byte, int, error) {
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, 500, err
	}
	req.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for key, val := range headers[0] {
			req.Header.Set(key, val)
		}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 500, err
	}
	defer res.Body.Close()
	r, _ := io.ReadAll(res.Body)
	return r, res.StatusCode, nil
}

func POST(URL string, body any, headers ...map[string]string) ([]byte, int, error) {
	byteData, err := json.Marshal(body)
	if err != nil {
		return nil, 500, err
	}
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(byteData))
	if err != nil {
		return nil, 500, err
	}
	req.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for key, val := range headers[0] {
			req.Header.Set(key, val)
		}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 500, err
	}
	defer res.Body.Close()
	r, _ := io.ReadAll(res.Body)
	return r, res.StatusCode, nil
}
func PUT(URL string, body any, headers ...map[string]string) ([]byte, int, error) {
	byteData, err := json.Marshal(body)
	if err != nil {
		return nil, 500, err
	}
	req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(byteData))
	if err != nil {
		return nil, 500, err
	}
	req.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for key, val := range headers[0] {
			req.Header.Set(key, val)
		}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 500, err
	}
	defer res.Body.Close()
	r, _ := io.ReadAll(res.Body)
	return r, res.StatusCode, nil
}
func DELETE(URL string, headers ...map[string]string) ([]byte, int, error) {
	req, err := http.NewRequest(http.MethodDelete, URL, nil)
	if err != nil {
		return nil, 500, err
	}
	req.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for key, val := range headers[0] {
			req.Header.Set(key, val)
		}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 500, err
	}
	defer res.Body.Close()
	r, _ := io.ReadAll(res.Body)
	return r, res.StatusCode, nil
}
