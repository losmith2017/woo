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
/*
	Accept: application/octet-stream
	Accept: application/vnd.github.v3+json
	Accept: application/vnd.github.v3.raw

	X-GitHub-Media-Type

	https://github.com/138over/woo-docs/blob/master/src/img/2006-poc.png
	https://raw.githubusercontent.com/138over/woo-docs/master/src/img/2006-poc.png

	application/octet-stream tells the browser that it's a generic binary file
	which will cause it to save to disk
	https://stackoverflow.com/questions/20508788/do-i-need-content-type-application-octet-stream-for-file-download


	TODO: X-Prefix convention is not recommended anymore... what is the recommendation?
	https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
		https://specs.openstack.org/openstack/api-wg/guidelines/headers.html
		ok... they are recommendating... instead of X-Github-Media-Type... go with Github-Media-Type

	Accept
		Informs the server about the types of data that can be sent back. It is MIME-type.

		The Accept request HTTP header advertises which content types, expressed as MIME types,
		the client is able to understand. Using content negotiation, the server then selects one of
		the proposals, uses it and informs the client of its choice with the Content-Type response header.
		Browsers set adequate values for this header depending of the context where the request is done:
		when fetching a CSS stylesheet a different value is set for the request than when fetching an
		image, video or a script

	Content Negotiation https://developer.mozilla.org/en-US/docs/Web/HTTP/Content_negotiation

	Accept-Encoding
		Informs the server about the encoding algorithm, usually a compression algorithm, that can be used on the resource sent back.

	Default Mime Types
		https://developer.mozilla.org/en-US/docs/Web/HTTP/Content_negotiation/List_of_default_Accept_values

	curl https://linuxacademy.com/howtoguides/posts/show/topic/13852-understanding-curl-and-http-headers
*/

/*
	Scenario-1
		Given a url that returns Accept-Ranges

	Scenario-2
		Given a url that oes not return Accept-Ranges

	type HttpFile struct {
		Url string
		Header map[string]string
		Maxcpu int
		Maxgoroutine int
	}

	Get url Header
	Get Content-Length from Header
	Get Accept-Ranges from Header

	if Accept-Ranges call
	Get(url, content-length, maxcpu, maxgoroutines)
	if !Accept-Ranges cal Get(url, content-length)

https://coderwall.com/p/uz2noa/fast-parallel-downloads-in-golang-with-accept-ranges-and-goroutines
https://github.com/huydx/hget/blob/master/main.go
	sizeInMb := float64(len) / (1024 * 1024)
		if clen == "1" {
		Printf("Download size: not specified\n")
	} else if sizeInMb < 1024 {
		Printf("Download target size: %.1f MB\n", sizeInMb)
	} else {
		Printf("Download target size: %.1f GB\n", sizeInMb/1024)
	}

	conn    := flag.Int("n", runtime.NumCPU(), "connection")

 */

/*
        HttpFile {
            HttpFileInputs
            HttpFileOutputs
        }

        HttpFileInputs struct {
            Url string
            Destdir string
        }

        HttpFileOptions struct {
            TokenFile string
            NumCpu int
            NumGoroutines int
            Filename string
            PersistMetaFile string
            MakeDestdir boolean
            Overwrite boolean
            RemoveOnFail boolean
        }

        HttpFileLog {
            State []HttpFileState
        }

        HttpFileState struct { 
            StartTime date
            StopTime date
            TokenFile boolean
            Token boolean
            Authenticated boolean
            FileLength int
            BytesRead int
            BytesWritten int
            FileAssembled boolean
            DownloadProgress float
            DownloadComplete boolean
            MetaDataRead boolean
            MetaDataWritten boolean
            DestdirExist boolean
            FileExist boolean
        }

        HttpFileOutput Struct {
            File string
            ExitStatus int
        }

        ===========
        if implemented as a downloader for a predefined cache

        module would sit on top of the above module

        implement policy in the code, and/or thru configuration...



    retries: https://github.com/sethgrid/pester HAHAHAH look at hget... you have that url already
        also has some good code for concurrency and testcode 
        java example http://codegist.net/code/golang%20download%20file%20http/

    upload resumable https://github.com/23/resumable.js#how-do-i-set-it-up-with-my-server



 */

/*
    ok... don't design anything... 
    just implement granular functions... 
    then refactor into structs with methods...

    lets just say this is the simple pipeline

        http.GetPartial(range)
        displayReadPartial()
        writePartial(range, tmpfile)
        displayWrittenPartial()
        displayProgress()
        assemblePartial(tmpfile...)

    then we need to figure out some code patterns... that will let us
        * add value into any step of the pipeline
        * extract value from any step of the pipeline
        * test any step of the pipeline
*/

/*
    #1 get a base eco-system up ASAP
        * use sde-dev/packer stuff, with Makefile, simpflify the makefile... not need to be so granular
        * a few nodes, some basic docker, etc...

    #2 manually setup vault, jenkins, jira, gerrit, nexus... use the makefile, whatever...

    #3 go thru and ramp on the api for doing a shitload of tasks using golan test cases
            i.e jenkins, github, git, jira, vault, docker

            build a catalog of tasks via the apis to these services :)

    #4 setup an S3 bucket to keep stuff...


    Coding wise...
        (_) duck typing
        (_) mocking
        (_) middleware
                https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
                https://www.nicolasmerouze.com/middlewares-golang-best-practices-examples/
        (_) dependency injection...
        (_) what else?

    hmmm... the target is to build a library... that is the end result
        * the library depends on objs... 
        * objs are defined by a list of intputs... .c's
        


    OBJS = $(SRC.c=$(OBJDIR)/%.o)

    $(LIBRARY): $(OBJS)

    $(OBJDIR)/%.o: $(SRCDIR)/%.c
        $(DOT.c.o.rule)

    state was simply... did src exist, what was timestamp... did dest exist, what was timestamp

    we have a pipeline of tasks.... the end result is a file that has been downloaded from 
    a url

    that says... if the URL.zip is newer than the FILE.zip... then apply the DOT.zip.zip.rule

    $(FILE).zip: $(URL).zip
        $(DOT.zip.zip.rule)

    so I would have a collection of zip files that need to be donwloaded to instantiate a workspace

    ZIPS = $(FILE.zip=$(OBJDIR)%.zip)

    $(WORKSPACE): $(ZIPS)

    what would that look like in code...

    type Deps struct {
        Src string
        Dst string
    }

    type DepReader interface {
    }

    type DepWriter interface {
    }

    deps = []Deps{}

    for dep in range deps 
        deps.Rule() // rule, would minimally do what... we want to be able to wrap it, middleware... etc...

        i.e.. bare bones minimum mechanism... now what else do you want to do to
            add value ?
            extract value of that bare bones minimum mechanism?


        bare bones is:
            get src and copy to dest ... i.e curl -O
            pass/fail

        renaming dest... option
        copying anywhere than current directory... option
        progress... option
        metadata.. option
        copy with partials... option
        resume... option

        so.. how would we share data between these options.. that are tasks?

        we... we did in in the conf by have common properties and then resolving at each task
        and creating a dependency between each task...

        so what if instead of middleware... we simply had... dependencies? and we built the 
        structs up the same way we build the conf up...?

        so the arguments to a task... is a list of deps right...

        so some task are plugins, and some tasks are deps...

      
    configuration would have to know...
        src
        dest
        if authentication
        if progress
        if metadata

        upon execution... 
        if accept-ranges



 */
