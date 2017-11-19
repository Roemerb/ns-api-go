package ns

import (
	"context"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"

	queryUtils "github.com/google/go-querystring/query"
)

const (
	// Version is the version of the client
	Version = "0.1.0"
	// HTTPScheme is the HTTP scheme used
	HTTPScheme = "https"
	// Host is the base URL of the API
	Host = "webservices.ns.nl"
	// Method is the HTTP method used. For the NS API it's always GET
	Method = http.MethodGet
	// UserAgent is the value of the UserAgent header that's being send
	UserAgent = "NSGoAPI/" + Version
)

//NS is an instance of the API library
type NS struct {
	client   *http.Client
	BaseURL  *url.URL
	Username string
	Password string
	Callback func(*http.Request, *http.Response)

	// Services
	TravelOptions *TravelOptionServiceImpl
	Stations      *StationServiceImpl
	Journeys      *JourneyServiceImpl
}

// APIResponse is a response from the API
type APIResponse struct {
	Response *http.Response
	Result   interface{}
	Success  bool
}

// APIError represents an error response from the NS API
type APIError struct {
	XMLName xml.Name `xml:"NS"`
	Error   struct {
		XMLName xml.Name `xml:"error"`
		Message string   `xml:"message"`
	}
}

// Init initiates the library and returns an NS instance pointer
func Init(username string, password string) *NS {
	url := url.URL{
		Scheme: HTTPScheme,
		Host:   Host,
		Path:   "/",
	}

	ns := &NS{
		client:   &http.Client{},
		BaseURL:  &url,
		Username: username,
		Password: password,
	}
	ns.TravelOptions = &TravelOptionServiceImpl{ns: ns}
	ns.Stations = &StationServiceImpl{ns: ns}
	ns.Journeys = &JourneyServiceImpl{ns: ns}

	return ns
}

// DoRequest will format and execute the
func (ns *NS) DoRequest(ctx context.Context, path string, query interface{}, auth bool) (*http.Response, error) {
	url, err := buildURL(path, query)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(Method, url.String(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if auth {
		req.SetBasicAuth(ns.Username, ns.Password)
	}
	req.Header.Add("User-Agent", UserAgent)

	res, err := ns.client.Do(req)
	return res, nil
}

// ParseResponse will unmarshall the response XML to an interface
func (ns *NS) ParseResponse(response *http.Response, target interface{}, isCollection bool) error {
	responseString, err := ns.getResponseString(response)
	if err != nil {
		return err
	}
	if isCollection {
		responseString = "<NS>" + responseString + "</NS>"
	}
	err = xml.Unmarshal([]byte(responseString), target)
	if err != nil {
		var apiErr APIError
		err = xml.Unmarshal([]byte(responseString), &apiErr)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (ns *NS) getResponseString(rawRes *http.Response) (string, error) {
	buff, ioerr := ioutil.ReadAll(rawRes.Body)
	defer rawRes.Body.Close()
	if ioerr != nil {
		return "", ioerr
	}

	return string(buff), nil
}

func buildURL(path string, query interface{}) (*url.URL, error) {
	var url url.URL
	url.Scheme = HTTPScheme
	url.Host = Host
	url.Path = path

	if query != nil {
		q, err := queryUtils.Values(query)
		if err != nil {
			return nil, err
		}
		url.RawQuery = q.Encode()
	}

	return &url, nil
}
