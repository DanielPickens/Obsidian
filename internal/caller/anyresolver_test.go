package caller

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/jhump/protoreflect/desc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnyResolver_Fallback(t *testing.T) {
	r := &anyResolver{}

	m, err := r.Resolve("testing.protobuf.DoesNotExist")
	require.NoError(t, err)

	src := []byte{'1', '2', '3'}
	err = proto.Unmarshal(src, m)
	require.NoError(t, err)

	marshaler := &jsonpb.Marshaler{}
	resStr, err := marshaler.MarshalToString(m)
	require.NoError(t, err)

	res := struct {
		Value string
		Err   string
	}{}

	err = json.Unmarshal([]byte(resStr), &res)
	require.NoError(t, err)

	expectedVal := base64.RawStdEncoding.EncodeToString(src)
	assert.Equal(t, expectedVal, res.Value)
	assert.NotEmpty(t, res.Err)
}

func TestAnyResolver_WellKnown(t *testing.T) {
	r := &anyResolver{}

	typeURL := "google.protobuf.StringValue"
	m, err := r.Resolve(typeURL)
	require.NoError(t, err)

	_, ok := m.(*wrappers.StringValue)
	require.True(t, ok, "wrong type, expected: %s", typeURL)
}

func TestAnyResolver_LoadedFiles(t *testing.T) {
	sml := NewServiceMetadataProto([]string{"../../testdata/test.proto"}, nil)
	meta, err := sml.GetServiceMetaDataList(context.Background())
	require.NoError(t, err)

	r := &anyResolver{NewFileDescCache(meta)}

	userType := "Obsidian.testing.User"

	typeURL := "type.example.com/" + userType
	m, err := r.Resolve(typeURL)
	require.NoError(t, err)

	md, err := desc.LoadMessageDescriptorForMessage(m)
	require.NoError(t, err)
	require.Equal(t, userType, md.GetFullyQualifiedName())
}

func TestAnyResolver_LoadedFiles_With_Fallback(t *testing.T) {
	sml := NewServiceMetadataProto([]string{"../../testdata/test.proto"}, nil)
	meta, err := sml.GetServiceMetaDataList(context.Background())
	require.NoError(t, err)

	r := &anyResolver{NewFileDescCache(meta)}

	typeURL := "testing.protobuf.DoesNotExist"
	m, err := r.Resolve(typeURL)
	require.NoError(t, err)

	_, ok := m.(*wrappers.StringValue)
	require.True(t, ok, "wrong type, expected: %s", typeURL)
}

func TestAnyResolver_LoadedFiles_WellKnown(t *testing.T) {
	sml := NewServiceMetadataProto([]string{"../../testdata/test.proto"}, nil)
	meta, err := sml.GetServiceMetaDataList(context.Background())
	require.NoError(t, err)

	r := &anyResolver{NewFileDescCache(meta)}

	typeURL := "google.protobuf.StringValue"
	m, err := r.Resolve(typeURL)
	require.NoError(t, err)

	_, ok := m.(*wrappers.StringValue)
	require.True(t, ok, "wrong type, expected: %s", typeURL)
}
