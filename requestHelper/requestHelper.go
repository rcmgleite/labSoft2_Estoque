package requestHelper

import (
	"bytes"
	"net/http"
)

//MakeRequest ...
func MakeRequest(httpMethod string, url string, requestObj []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(requestObj))
	addHeaders(req, headers)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func addHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}
