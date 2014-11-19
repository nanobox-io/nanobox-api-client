// package client consists of a core api client struct with methods broken into
// related calls, for interacting and communicating with the Nanobox API.
package client

//
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"regexp"
)

//
const (
	DefaultAPIURL      = "http://localhost:8080"
	DefaultContentType = "application/json"
	Version            = "0.0.1"
)

//
type (

	// Client represents a Nanobox API client. Its zero value is a default
	// http.Client. A Client should be created only once and reused:
	Client struct {
		APIURL     string       // The URL to which the API client will make requests.
		Debug      bool         // If debug mode is enabled.
		HTTPClient *http.Client // The HTTP.Client to use when making requests.
		Password   string       // The password to use when requesting an auth_token.
		Username   string       // The Username to use when requesting an auth_token.
	}

	// Error represents a pagoda-client error
	Error struct {
		error             // The entire error (ex. {"error":"404 Not Found"})
		Code       int    // The number status code (ex. 404)
		Status     string // The status code and text (ex. "404 Not Found")
		StatusCode string // The status code (ex. "404")
		StatusText string // The status text (ex. "Not Found")
		Body       string // The error body (ex. "Not Found")
	}

	// Email represents an email that can be attached to objects like cron jobs or
	// invoices
	Email struct {
		Email string
	}
)

// NewClient initializes a new API client.
func NewClient() *Client {

	// create a new client
	c := &Client{}

	// set the HTTP client
	c.HTTPClient = http.DefaultClient

	// if no APIRUL provided, use default.
	if c.APIURL == "" {
		c.APIURL = DefaultAPIURL
	}

	return c
}

// post handles standard POST operations to the Nanobox API
func (c *Client) post(v interface{}, path string, body interface{}) error {
	return c.doAPIRequest(v, "POST", path, body)
}

// get handles standard GET operations to the Nanobox API
func (c *Client) get(v interface{}, path string) error {
	return c.doAPIRequest(v, "GET", path, nil)
}

// patch handles standard PATH operations to the Nanobox API [NOT SUPPORTED]
func (c *Client) patch(v interface{}, path string, body interface{}) error {
	return c.doAPIRequest(v, "PATCH", path, body)
}

// put handles standard PUT operations to the Nanobox API
func (c *Client) put(v interface{}, path string, body interface{}) error {
	return c.doAPIRequest(v, "PUT", path, body)
}

// delete handles standard DELETE operations to the Nanobox API
func (c *Client) delete(path string) error {
	return c.doAPIRequest(nil, "DELETE", path, nil)
}

// TODO: take a look at DoRawRequest, doAPIRequest, and newRequest. see if there
// is a better way to split these up that is a little cleaner now that we have to
// have raw requests

// DoRawRequest creates and perform a standard HTTP request, allowing for the
// addition of custom headers
func (c *Client) DoRawRequest(v interface{}, method, path string, body interface{}, headers map[string]string) error {

	req, err := c.newRequest(method, path, body, headers)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// doAPIRequest creates and perform a standard HTTP request.
func (c *Client) doAPIRequest(v interface{}, method, path string, body interface{}) error {

	// the fully qualified URL includes the apiURL + path + auth_token
	fullPath := c.APIURL + path

	req, err := c.newRequest(method, fullPath, body, nil)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// newRequest creates an HTTP request for the Pagoda Box API, but does not perform
// it.
func (c *Client) newRequest(method, path string, body interface{}, headers map[string]string) (*http.Request, error) {

	var rbody io.Reader

	//
	switch t := body.(type) {
	case string:
		rbody = bytes.NewBufferString(t)
	case io.Reader:
		rbody = t
	default:
		rbody = nil
	}

	// an HTTP request
	req, err := http.NewRequest(method, path, rbody)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", DefaultContentType)
	req.Header.Set("Content-Type", DefaultContentType)

	// add additional headers
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}

// Do performs an http.NewRequest
func (c *Client) do(req *http.Request, v interface{}) error {

	// debugging
	if c.Debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(`
Request:
--------------------------------------------------------------------------------
` + string(dump))
	}

	//
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	// debugging
	if c.Debug {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(`
Response:
--------------------------------------------------------------------------------
` + string(dump))
	}

	//
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// p := make([]byte, 2048)
	// for {
	// 	_, err := res.Body.Read(p)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Println("EOF!!")
	// 			break
	// 		}
	// 	}
	// 	fmt.Println(string(p))
	// }

	// rw.Write([]byte(out))
	// rw.(http.Flusher).Flush()

	defer res.Body.Close()

	// check the response
	if err = checkResponse(res, b); err != nil {
		return err
	}

	//
	if err := json.Unmarshal(b, v); err != nil {
		panic(err)
	}

	return err
}

// checkResponse determins if the response is !20* and return a custom api error
func checkResponse(res *http.Response, b []byte) error {

	if res.StatusCode/100 != 2 {

		//
		var subMatch []string

		// pull the body out of the error
		reFindErrorBody := regexp.MustCompile(`^\{\s*\"*errors?\"*\s*\:(.*)\}$`)
		subMatch = reFindErrorBody.FindStringSubmatch(string(b))
		if subMatch == nil {
			return errors.New("Unable to parse error body: " + string(b))
		}

		body := subMatch[1]

		// separate the status code and text
		reFindStatusText := regexp.MustCompile(`^\s*(\d+)(\D+)$`)
		subMatch = reFindStatusText.FindStringSubmatch(res.Status)
		if subMatch == nil {
			return errors.New("Unable to parse error status: " + res.Status)
		}

		statusCode := subMatch[1]
		statusText := subMatch[2]

		return Error{error: errors.New(string(b)), Code: res.StatusCode, Status: res.Status, StatusCode: statusCode, StatusText: statusText, Body: body}
	}

	return nil
}

// toJSON converts an interface (v) into JSON bytecode
func toJSON(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return b
}
