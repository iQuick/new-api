package service

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
	"one-api/constant"
	"strings"
)

func GetProxiedHttpClient(proxyUrl string, proxyUsername string, proxyPassword string) (*http.Client, error) {
	if "" == proxyUrl {
		return &http.Client{}, nil
	}

	if proxyUsername != "" && proxyPassword != "" {
		u, err := url.Parse(proxyUrl)
		if err == nil {
			proxyUrl = fmt.Sprintf("%s://%s:%s@%s", u.Scheme, proxyUsername, proxyPassword, u.Host)
		}
	}

	u, err := url.Parse(proxyUrl)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(proxyUrl, "http") {

		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(u),
			},
		}, nil
	} else if strings.HasPrefix(proxyUrl, "socks") {
		dialer, err := proxy.FromURL(u, proxy.Direct)
		if err != nil {
			return nil, err
		}

		return &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return dialer.(proxy.ContextDialer).DialContext(ctx, network, addr)
				},
			},
		}, nil
	}

	return nil, errors.New("unsupported proxy type")
}

func ProxiedHttpGet(url string) (*http.Response, error) {
	client, err := GetProxiedHttpClient(constant.OutProxyUrl, constant.OutProxyUsername, constant.OutProxyPassword)
	if err != nil {
		return nil, err
	}

	return client.Get(url)
}

func ProxiedHttpHead(url string) (*http.Response, error) {
	client, err := GetProxiedHttpClient(constant.OutProxyUrl, constant.OutProxyUsername, constant.OutProxyPassword)
	if err != nil {
		return nil, err
	}

	return client.Head(url)
}
