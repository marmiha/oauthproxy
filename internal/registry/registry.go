package registry

import (
	"github.com/gume1a/oauth-proxy/internal/proxy"
	"github.com/gume1a/oauth-proxy/internal/target"
	"github.com/gume1a/oauth-proxy/pkg/identity"
	"net/http"
)

var handler Registry

func init() {
	handler = &registry{
		providerProxyMap: map[identity.ProviderId]proxy.OAuthProxy{
			identity.GITHUB: proxy.New(&target.GITHUB),
			identity.GOOGLE: proxy.New(&target.GOOGLE),
		},
	}
}

type (
	Registry interface {
		ProxyServeHTTP(rw http.ResponseWriter, req *http.Request)
		Providers() []identity.ProviderId
	}
	registry struct {
		providerProxyMap map[identity.ProviderId]proxy.OAuthProxy
	}
)

func init() {
	handler = &registry{
		providerProxyMap: map[identity.ProviderId]proxy.OAuthProxy{
			identity.GITHUB: proxy.New(&target.GITHUB),
			identity.GOOGLE: proxy.New(&target.GOOGLE),
		},
	}
}

// GetRegistry returns the configured registry of configured (provided) proxies.
func GetRegistry() Registry {
	return handler
}

// Providers returns the identity.ProviderId-s of the configured providers.
func (r *registry) Providers() []identity.ProviderId {
	providers := make([]identity.ProviderId, len(r.providerProxyMap))

	i := 0
	for k := range r.providerProxyMap {
		providers[i] = k
		i++
	}

	return providers
}

// ProxyServeHTTP accepts the incoming request and forwards the request to the right
// target provider given its identity.ProviderId scraped from the request.
func (r *registry) ProxyServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Get the provider from the request.
	appId := r.getAppId(req)
	// toolId := r.getToolId(req)

	// Get the proxy for the provider.
	proxy, ok := r.providerProxyMap[appId]
	if !ok {
		rw.WriteHeader(http.StatusNotAcceptable)
		rw.Write([]byte("Provider not supported"))
	}

	proxy.ServeHTTP(rw, req)
}

func (r *registry) getAppId(req *http.Request) identity.ProviderId {
	// Todo: configure switching based on the request
	return identity.GITHUB
}

func (r *registry) getToolId(req *http.Request) identity.ToolId {
	// Todo: configure switching based on the request
	return identity.FELTNA
}
