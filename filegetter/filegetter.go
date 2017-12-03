package filegetter

import (
	"fmt"
	"io"
	"net/http"
)

type HttpFileGetter struct {
	Url string
}

func (f *HttpFileGetter) Stat(w io.Writer) (*http.Response, error)  {
	resp, err := http.Head(f.Url)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (f *HttpFileGetter) WriteTo(w io.Writer) (*http.Response, error) {
	resp, err := http.Get(f.Url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, fmt.Errorf("bad response code: %d", resp.StatusCode)
	}

	_, err = io.Copy(w, resp.Body)
	return resp, err
}