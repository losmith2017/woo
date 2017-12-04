package filer

import (
	"fmt"
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

func getGithubToken() (string, error) {
	f := os.Getenv("HOME")+"/.github/token"
	dat, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("token %s", strings.TrimSpace(string(dat))), err
}

func TestHttpFiler_Get(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "filegetter_test")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	fdata := "woooooo"
	tserv, teardownServer := setupServer(t, fdata)
	defer teardownServer(t)

	hdr := map[string][]string{}
	var filer HttpFiler = HttpFile{tserv.URL, hdr }
	_, err = filer.Get(tmpfile)
	if err != nil {
		t.Fatalf("write failed %", err)
	}

	fstat, err := os.Stat(tmpfile.Name())
	if int64(len(fdata)) != fstat.Size() {
		t.Fatalf("size = %v; want %v", fstat.Size(), len(fdata))
	}
}


func TestHttpFiler_Authenticate(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "filegetter_test")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	token, err := getGithubToken()
	if err != nil {
		t.Fatalf("cannot retrieve token", err)
	}

	var hdr  = map[string][]string{
		"Authorization":  {token},
		"Accept": {"application/vnd.github.v3+json"},
	}

	var filer HttpFiler = HttpFile{"https://api.github.com/user/repos", hdr }
	_, err = filer.Get(tmpfile)
	if err != nil {
		t.Fatalf("write failed %", err)
	}

	resp, err := filer.Stat()
	cl, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if cl < 1000 {
		t.Fatalf("content-length = %v; want > 1000", cl)
	}
}
