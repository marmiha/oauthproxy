package oauthtransport

import (
	"log"
	"net/http"
)

type (
	oauthTransport struct{}
)

func NewOauthTransport() http.RoundTripper {
	return &oauthTransport{}
}

func (t *oauthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Todo: logging
	// Body is read-once only. Buffer it into two buffers, one will be used for forwarding
	// the other for reading.
	/* bodyBuf, _ := ioutil.ReadAll(req.Body)
	bodyCopy := ioutil.NopCloser(bytes.NewBuffer(bodyBuf))
	bodyPreflight := ioutil.NopCloser(bytes.NewBuffer(bodyBuf))

	// Replaced the already read
	req.Body = bodyCopy

	//appId := req.FormValue("app_id")
	//toolId := req.FormValue("tool_id")

	req.Form.Add("test-field", "test-value")
	req.PostForm.Add("test-field", "test-value")
	req.Body = bodyPreflight

	// Send to the OAuth2 server. */
	//bytes, _ := httputil.DumpRequest(req, true)
	//log.Println("123\n", string(bytes))

	res, err := http.DefaultTransport.RoundTrip(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}
