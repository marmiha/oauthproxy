package proxy

import (
	"github.com/gume1a/oauth-proxy/internal/oauthtransport"
	"github.com/gume1a/oauth-proxy/internal/target"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type (
	oauthProxy struct {
		target *url.URL
		proxy  *httputil.ReverseProxy
	}

	OAuthProxy interface {
		ServeHTTP(rw http.ResponseWriter, req *http.Request)
	}
)

func New(trgt *target.Provider) OAuthProxy {
	proxy := httputil.NewSingleHostReverseProxy(trgt.TokenEndpoint)
	defaultDirector := proxy.Director

	// Director handles the modification of the request before it is sent to the OAut2 server.
	proxy.Director = func(req *http.Request) {
		defaultDirector(req)
		req.Header.Add("X-Proxy", "oauth-proxy")
		req.Host = trgt.RequestHost
	}

	// Handling of the response.
	proxy.ModifyResponse = func(res *http.Response) error {
		return nil
	}

	// Override the round trip function.
	proxy.Transport = oauthtransport.NewOauthTransport()

	return &oauthProxy{
		target: trgt.TokenEndpoint,
		proxy:  proxy,
	}
}

func (p *oauthProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Drop the URL path, so that it does not append to the target token endpoint.
	// If you remove this line, the target path and proxy path will be joined.
	req.URL.Path = ""
	req.RequestURI = ""

	// Forward the request to the proxy
	p.proxy.ServeHTTP(rw, req)
}
