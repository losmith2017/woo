package url

import (
	"fmt"
	"io"
	"net/http"
)

type HttpFiler interface {
	Get(w io.Writer) (resp *http.Response, err error)
	Stat() (resp *http.Response, err error)
}

type HttpFile struct {
	Url string
	Header map[string][]string
}

func (f *HttpFile) Stat() (resp *http.Response, err error) {
	req, err := http.NewRequest("HEAD", f.Url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = f.Header

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, fmt.Errorf("bad response code: %d", resp.StatusCode)
	}

	return resp, err
}

func (f *HttpFile) Get(w io.Writer) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", f.Url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = f.Header

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, fmt.Errorf("bad response code: %d", resp.StatusCode)
	}

	bytes, err := io.Copy(w, resp.Body)
	fmt.Printf("total bytes: %d\n", bytes)
	return resp, err
}



// https://coderwall.com/p/uz2noa/fast-parallel-downloads-in-golang-with-accept-ranges-and-goroutines