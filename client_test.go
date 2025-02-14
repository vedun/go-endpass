package endpass

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2"
)

type TestSuite struct {
	suite.Suite

	// test client
	c *Client

	// test server
	srv *httptest.Server
}

func (ts *TestSuite) SetupSuite() {
	ts.srv = httptest.NewServer(ts)
	ts.c = NewClient("clientID", "clientSecret", []string{"1111"}, "12345", "12345")
	ts.c.token = &oauth2.Token{}
	ts.c.baseClient = &http.Client{}
	ts.c.clientWithTokenSource = &http.Client{}
	ts.c.baseUrl = ts.srv.URL
}

// handler for API requests
func (ts *TestSuite) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fp := fmt.Sprintf("./testdata%s.json", r.URL.Path)
	body, err := ioutil.ReadFile(fp)
	ts.NoError(err)
	w.Write(body)
}

func TestClient(t *testing.T) {
	ts := &TestSuite{}
	suite.Run(t, ts)
}
