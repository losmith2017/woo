package url

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

func setupServer(t *testing.T, fdata string) (*httptest.Server, func(t *testing.T, s *httptest.Server)) {
	t.Log("setup server")
	mtime := time.Unix(1512216000, 0).UTC()
	fname := "woo.txt"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, fname, mtime, strings.NewReader(fdata))
	}))
	return server, func(t *testing.T, s *httptest.Server) {
		t.Log("teardown server")
		s.Close()
	}
}

func setupTmpfile(t *testing.T,  prefix string) (*os.File,  func(t *testing.T,  f *os.File)) {
	tmpfile, err := ioutil.TempFile("", prefix)
	if err != nil {
		t.Fatal(err)
	}
	return tmpfile, func(t *testing.T, f *os.File) {
		t.Log("teardown tmpfile")
		f.Close()
		os.Remove(f.Name())
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
	tmpfile, teardownTmpfile := setupTmpfile(t, "httpfiler")
	defer teardownTmpfile(t, tmpfile)

	fdata := "woooooo"
	tserv, teardownServer := setupServer(t, fdata)
	defer teardownServer(t, tserv)

	hdr := map[string][]string{}
	var filer HttpFiler = &HttpFile{tserv.URL, hdr }
	_, err := filer.Get(tmpfile)
	if err != nil {
		t.Fatalf("write failed %v", err)
	}

	fstat, err := os.Stat(tmpfile.Name())
	if err != nil {
		t.Fatalf("write failed %v", err)
	}
	if int64(len(fdata)) != fstat.Size() {
		t.Fatalf("size = %v; want %v", fstat.Size(), len(fdata))
	}
}


func TestHttpFiler_Authenticate(t *testing.T) {
	tmpfile, teardownTmpfile := setupTmpfile(t, "httpfiler")
	defer teardownTmpfile(t, tmpfile)

	token, err := getGithubToken()
	if err != nil {
		t.Fatalf("cannot retrieve token", err)
	}

	// url := "https://raw.githubusercontent.com/138over/woo-docs/master/src/img/2006-poc.png"
	// accept := "application/vnd.github.v3.raw"

	url := "https://api.github.com/user/repos"
	accept := "application/vnd.github.v3+json"

	var hdr  = map[string][]string{
		"Authorization":  {token},
		"Accept": {accept},
	}

	var filer HttpFiler = &HttpFile{url, hdr }

	_, err = filer.Get(tmpfile)
	if err != nil {
		t.Fatalf("write failed %v", err)
	}


	resp, err := filer.Stat()
	cl, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if cl < 1000 {
		t.Fatalf("content-length = %v; want > 1000", cl)
	}

	cr := resp.Header.Get("Content-Range")
	ar := resp.Header.Get("Accept-Ranges")
	t.Log(resp.Header)
	t.Logf("Content-Length: %v\n", cl)
	t.Logf("Content-Range: %s\n", cr)
	t.Logf("Accept-Ranges: %s\n", ar)


}

// Github Header https://developer.github.com/v3/media/
