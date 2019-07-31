package gitio

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	Header      map[string]string
	Body        []byte
	ContentType string
}

func NewHttpClient(header map[string]string, body []byte, contentType string) *HttpClient {
	return &HttpClient{Header: header, Body: body, ContentType: contentType}
}

func (hc HttpClient) Get(url string) (int, []byte) {
	return hc.Do(http.MethodGet, url)
}

func (hc HttpClient) Post(url string) (int, []byte) {
	return hc.Do(http.MethodPost, url)
}

func (hc HttpClient) Delete(url string) (int, []byte) {
	return hc.Do(http.MethodDelete, url)
}

func (hc HttpClient) Do(method string, url string) (int, []byte) {
	c := &http.Client{}
	//set body (json)
	req, _ := http.NewRequest(method, url, bytes.NewReader(hc.Body))
	//set header
	for k, v := range hc.Header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", hc.ContentType)
	//do request
	resp, err := c.Do(req)
	if err != nil {
		println(err.Error())
		return -1, nil
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil
	}
	return resp.StatusCode, data
}
