OAuth Proxy
===========

![OAuth Proxy Banner](./assets/static/logo_banner_big.png)
[![CI Worker (Golang)](https://github.com/gume1a/oauth-proxy/actions/workflows/ci-worker.yaml/badge.svg)](https://github.com/gume1a/oauth-proxy/actions/workflows/ci-worker.yaml)
[![Maintainability Rating](https://sonarqube.gume1a.com/api/project_badges/measure?project=gume1a_oauth-proxy_AYZm3vqmPvHIgh4c6jWj&metric=sqale_rating&token=sqb_ad10b28fe89f2a16c2c9f7899e50e2c6cf192074)](https://sonarqube.gume1a.com/dashboard?id=gume1a_oauth-proxy_AYZm3vqmPvHIgh4c6jWj)
[![Lines of Code](https://sonarqube.gume1a.com/api/project_badges/measure?project=gume1a_oauth-proxy_AYZm3vqmPvHIgh4c6jWj&metric=ncloc&token=sqb_ad10b28fe89f2a16c2c9f7899e50e2c6cf192074)](https://sonarqube.gume1a.com/dashboard?id=gume1a_oauth-proxy_AYZm3vqmPvHIgh4c6jWj)
[![Technical Debt](https://sonarqube.gume1a.com/api/project_badges/measure?project=gume1a_oauth-proxy_AYZm3vqmPvHIgh4c6jWj&metric=sqale_index&token=sqb_ad10b28fe89f2a16c2c9f7899e50e2c6cf192074)](https://sonarqube.gume1a.com/dashboard?id=gume1a_oauth-proxy_AYZm3vqmPvHIgh4c6jWj)

A simple ready to go service, that reverse-proxies your Token endpoint requests to configured OAuth2 providers and
attaches the client secret to the request. This is useful for when the Authorization servers don't support the non
client secret authorization flows but the client application is required to be run on the end-user device.

### Installation

A valid [installation of Go](https://go.dev/doc/install) is required. This installs the latest version of
the `oauth-proxy` cmd tool from the master branch.

```
go install github.com/gume1a/oauth-proxy@latest
```

### Configuration

The configuration is done via a yaml file and environment variables. The default path is `./oauthconfig.yaml`, if none 
provided the server will start on `localhost:8081` with no configured providers. Example configuration is as follows:

```yaml config/.template.oauthconfig.yaml
# config/.template.oauthconfig.yaml
host: localhost
port: 8081
providers:
  supported:
    - id: github
      client_secret: GITHUB_SECRET
    - id: google
      client_secret: GOOGLE_SECRET
  custom:
    - id: custom
      client_secret: CUSTOM_CLIENT_SECRET
      token_endpoint: https://example.com/oauth2/authorize
```

As seen above, the configuration is split into two parts. The first part is the `supported` providers. These are the 
providers that are already configured in the code and can be used without any additional configuration. The second part
is the `custom` providers. These are the definitions of the custom providers. The supported providers just have the token
endpoint set.

The `client_secret` is the name of the environment variable that contains the client secret for the provider. It supports
loading from the dotenv file `.env` but it's not required.

```sh config/.template.env
# config/.template.env
GITHUB_SECRET=github_secret
GOOGLE_SECRET=google_secret
CUSTOM_SECRET=custom_secret
```
With this configuration, the proxy will be able to handle requests for the `github`, `google` and the `custom` provider.
Arbitrarily many providers can be configured.

### Usage

After installation the  server can be started by running the `oauth-proxy` command. The proxy will start listening on 
the configured host and port.

```ansi 
$ aouth-proxy           
                   _   _
  ___   __ _ _   _| |_| |__  _ __  _ __ _____  ___   _
 / _ \ / _` | | | | __| '_ \| '_ \| '__/ _ \ \/ / | | |
| (_) | (_| | |_| | |_| | | | |_) Who let the secrets OUT?
 \___/ \__,_|\__,_|\__|_| |_| .__/|_|  \___/_/\_\\__, |
                            |_|           v0.2.0 |___/


2023/02/20 08:04:17 INIT .env loaded
2023/02/20 08:04:17 PROVIDERS [github google custom]
2023/02/20 08:04:17 SERVER starting listening on http://localhost:8081
```

### Endpoints

The proxy currently exposes two endpoints:
- `/oauth/<client_id>` - This endpoint is used to get the token for the client with the given id. The client id is the 
  id of the provider in the configuration file. The request is forwarded to the configured token endpoint and the 
  client secret is attached to the request. The response is then returned to the client.
- `/supported` - This endpoint returns a list of the supported providers. The list is the same as the list of the 
  providers in the configuration file.