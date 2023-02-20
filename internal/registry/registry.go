package registry

import (
	"github.com/gume1a/oauthproxy/internal/provider"
	"github.com/gume1a/oauthproxy/internal/proxy"
	"github.com/gume1a/oauthproxy/pkg/identity"
	"net/http"
)

type (
	Registry interface {
		ProxyServeHTTP(providerID identity.ProviderId, rw http.ResponseWriter, req *http.Request)
		Providers() []identity.ProviderId
	}
	registry struct {
		providerProxyMap map[identity.ProviderId]proxy.OAuthProxy
	}
)

// NewRegistry returns the configured registry of configured (provided) proxies.
func NewRegistry(providers []provider.Provider) Registry {
	proxyMap := make(map[identity.ProviderId]proxy.OAuthProxy)
	for _, p := range providers {
		proxyMap[p.Id] = proxy.New(p)
	}
	return &registry{providerProxyMap: proxyMap}
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
// provider given its identity.ProviderId scraped from the request.
func (r *registry) ProxyServeHTTP(providerId identity.ProviderId, rw http.ResponseWriter, req *http.Request) {
	// Get the proxy for the provider.
	proxy, ok := r.providerProxyMap[providerId]
	if !ok {
		rw.WriteHeader(http.StatusNotAcceptable)
		rw.Write([]byte("provider not supported"))
		return
	}

	proxy.ServeHTTP(rw, req)
}
