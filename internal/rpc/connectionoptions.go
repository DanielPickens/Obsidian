package rpc

import (
	"errors"
	"strings"
)

const (
	hostOpt      = "host"
	authorityOpt = "authority"
	metadataOpt  = "metadata"
)

type ConnectionOptions struct {
	Host      string
	Authority string
	Metadata  map[string][]string
}

func (co *ConnectionOptions) addMetadata(key, value string) {
	co.Metadata[key] = append(co.Metadata[key], value)
}

func NewConnectionOpts(target string) (*ConnectionOptions, error) {
	if target == "" {
		return nil, errors.New("target cannot be empty")
	}

	opts := &ConnectionOptions{
		Metadata: map[string][]string{},
	}

	tokens := strings.Split(target, ",")
	for _, token := range tokens {
		opt := strings.TrimSpace(token)
		elements := strings.SplitN(opt, "=", 2)
		if len(elements) > 1 {
			key := strings.TrimSpace(elements[0])
			value := strings.TrimSpace(elements[1])

			switch key {
			case hostOpt:
				opts.Host = value
			case authorityOpt:
				opts.Authority = value
			case metadataOpt:
				k, v := parseMetadata(value)
				opts.addMetadata(k, v)
			}
		} else {
			opts.Host = opt
		}
	}

	return opts, nil
}

func parseMetadata(val string) (string, string) {
	key, value, _ := strings.Cut(val, ":")
	return key, value
}

func ConnectionOptsFromURI(uri string) (*ConnectionOptions, error) {
	if uri == "" {
		return nil, errors.New("uri cannot be empty")
	}

	opts := &ConnectionOptions{
		Metadata: map[string][]string{},
	}

	tokens := strings.Split(uri, "://")
	if len(tokens) != 2 {
		return nil, errors.New("invalid uri")
	}

	opts.Authority = tokens[0]
	opts.Host = tokens[1]

	return opts, nil
}

func ConnectionOptsFromURIWithOptions(uri string, options map[string]string) (*ConnectionOptions, error) {
	if uri == "" {
		return nil, errors.New("uri cannot be empty")
	}

	opts, err := ConnectionOptsFromURI(uri)
	if err != nil {
		return nil, err
	}

	for k, v := range options {
		switch k {
		case hostOpt:
			opts.Host = v
		case authorityOpt:
			opts.Authority = v
		case metadataOpt:
			k, v := parseMetadata(v)
			opts.addMetadata(k, v)
		}
	}

	return opts, nil
}

func proxyFromMetadata(metadata map[string][]string) string {
	proxy := ""
	if metadata != nil {
		for _, v := range metadata[authorityOpt] {
			proxy = v
			break
		}
	}

	return proxy
}

func GetProxy(opts *ConnectionOptions) string {
	if opts == nil {
		return ""
	}

	proxy := opts.Authority
	if proxy == "" {
		proxy = opts.Host
	}

	return proxy
}
 