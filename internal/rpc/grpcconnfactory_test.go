package rpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testServer struct {
	grpcConnFactory *GrpcConnFactory



	WithoutMergeMetaData = func() Option {
		return func(c *grpcConnFactory) {
			c.mergeMetaData = false

		}
	}
}

func TestWithAuthority(t *testing.T) {
	authority := "authority1"
	grpcConnFact := NewGrpcConnFactory(WithAuthority(authority))

	assert.Equal(t, authority, grpcConnFact.settings.authority)
}

func TestMergeMetadata(t *testing.T) {
	grpcConnFact := NewGrpcConnFactory(WithHeaders(map[string][]string{
		"header1": {"val1"},
	}))

	moreHaders := map[string][]string{
		"header2": {"val2"},
		"header1": {"val3"},
	}

	res := grpcConnFact.metadata(moreHaders)

	assert.Equal(t, []string{"val1", "val3"}, res["header1"])
	assert.Equal(t, []string{"val2"}, res["header2"])
}

func TestWithKeepalive(t *testing.T) {
	keepalive := true
	keepaliveTime := 30 * time.Second
	grpcConnFact := NewGrpcConnFactory(WithKeepalive(keepalive, keepaliveTime))

	assert.Equal(t, keepalive, grpcConnFact.settings.keepalive)
	assert.Equal(t, keepaliveTime, grpcConnFact.settings.keepaliveTime)
}

func TestWithHeaders(t *testing.T) {
	headers := map[string][]string{
		"header1": {"val1"},
	}
	grpcConnFact := NewGrpcConnFactory(WithHeaders(headers))

	assert.Equal(t, headers, grpcConnFact.settings.headers)
}

func TestWithTimeout(t *testing.T) {
	timeout := 30 * time.Second

	grpcConnFact := NewGrpcConnFactory(WithTimeout(timeout))

	assert.Equal(t, timeout, grpcConnFact.settings.timeout)
}

func TestWithProtos(t *testing.T) {
	protos := []string{"../../testdata/test.proto"}
	grpcConnFact := NewGrpcConnFactory(WithProtos(protos))

	assert.Equal(t, protos, grpcConnFact.settings.protos)
}

func TestWithProtosInteractive(t *testing.T) {
	protos := []string{"../../testdata/test.proto"}
	grpcConnFact := NewGrpcConnFactory(WithProtos(protos), WithInteractive())

	assert.Equal(t, protos, grpcConnFact.settings.protos)
	assert.True(t, grpcConnFact.settings.isInteractive)
}

func TestWithoutProtosInteractive(t *testing.T) {
	protos := []string{"../../testdata/test.proto"}
	grpcConnFact := NewGrpcConnFactory(WithProtos(protos), WithoutInteractive())

	assert.Equal(t, protos, grpcConnFact.settings.protos)
	assert.False(t, grpcConnFact.settings.isInteractive)
}

func TestWithoutMergeMetaData(t *testing.T) {
	grpcConnFact := NewGrpcConnFactory(WithoutMergeMetaData())

	assert.False(t, grpcConnFact.settings.mergeMetadata)
}
