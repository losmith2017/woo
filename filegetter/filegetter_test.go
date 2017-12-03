package filegetter

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func setupServer(t *testing.T, fdata string) (*httptest.Server, func(t *testing.T)) {
	t.Log("setup server")
	mtime := time.Unix(1512216000, 0).UTC()
	fname := "woo.txt"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, fname, mtime, strings.NewReader(fdata))
	}))
	return server, func(t *testing.T) {
		t.Log("teardown server")
		server.Close()
	}
}

func TestHttpFileGetter_Stat(t *testing.T) {
	fdata := "woooooo"
	tserv, teardownServer := setupServer(t, fdata)
	defer teardownServer(t)

	fget := HttpFileGetter{tserv.URL}
	resp, err := fget.Stat(os.Stdout)
	if err != nil {
		t.Fatalf("response.StatusCode: %s err %s", resp.StatusCode, err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("response.StatusCode = %s", resp.StatusCode)
	}

	cl, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if int64(len(fdata)) != cl {
		t.Fatalf("content-length = %v; want %v", cl, len(fdata))
	}
}

func TestHttpFileGetter_Write(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "filegetter_test")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	fdata := "woooooo"
	tserv, teardownServer := setupServer(t, fdata)
	defer teardownServer(t)

	fget := HttpFileGetter{tserv.URL}
	_, err = fget.WriteTo(tmpfile)
	if err != nil {
		t.Fatalf("write failed %", err)
	}

	fstat, err := os.Stat(tmpfile.Name())
	if int64(len(fdata)) != fstat.Size() {
		t.Fatalf("size = %v; want %v", fstat.Size(), len(fdata))
	}
}
