package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"strconv"
	"testing"
	"time"
)

func TestFileGetter_Write(t *testing.T) {
	mtime := time.Unix(1512216000, 0).UTC()
	fname := "woo.txt"
	fdata := "woooooo"

	tserv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, fname, mtime, strings.NewReader(fdata))
	}))
	defer tserv.Close()

	tmpfile, err := ioutil.TempFile("", "filegetter_test")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	fget := FileGetter{tserv.URL+"/"+fname}
	resp, err := fget.Write(tmpfile)
	if err != nil {
		t.Fatalf("response.StatusCode: %s err %s", resp.StatusCode, err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("response.StatusCode = %s", resp.StatusCode)
	}

	fsize, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	fstat, err := os.Stat(tmpfile.Name())
	if int64(len(fdata)) != fstat.Size() {
		t.Fatalf("size = %v; want %v", fstat.Size(), len(fdata))
	}
	if fsize != fstat.Size() {
		t.Fatalf("size = %v; want %v", fstat.Size(), fsize)
	}

	modstr := resp.Header.Get("Last-Modified")
	modtime, err := http.ParseTime(modstr)
	if err != nil {
		t.Fatalf("err: %q %s", modtime, err)
	}
	if !modtime.Equal(mtime) {
		t.Fatalf("modtime = %v; want %v", modtime, mtime)
	}
}