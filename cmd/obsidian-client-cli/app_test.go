package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	app_testing "github.com/Daniel/obsidian-client-cli/internal/testing"
	"github.com/DanielPickens/obsidian-client-cli/internal/caller"
	"github.com/spyzhov/ajson"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestMain(m *testing.M) {
	err := app_testing.SetupTestServer()
	if err != nil {
		panic(err)
	}

	defer app_testing.StopTestServer()
	os.Exit(m.Run())
}

func TestAppServiceCalls(t *testing.T) {
	runAppServiceCalls(t, &startOpts{
		Target:        app_testing.TestServerAddr(),
		Deadline:      15,
		IsInteractive: false,
	})
}

func runAppServiceCalls(t *testing.T, appOpts *startOpts) {
	buf := &bytes.Buffer{}
	appOpts.w = buf
	app, err := newApp(appOpts)
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("appCallUnaryServerError", func(t *testing.T) {
		appCallUnaryServerError(t, app)
	})

	t.Run("appCallUnary", func(t *testing.T) {
		buf.Reset()
		appCallUnary(t, app, buf)
	})

	t.Run("appCallStreamOutput", func(t *testing.T) {
		buf.Reset()
		appCallStreamOutput(t, app, buf)
	})

	t.Run("appCallStreamOutputError", func(t *testing.T) {
		appCallStreamOutputError(t, app)
	})

	t.Run("appCallClientStream", func(t *testing.T) {
		buf.Reset()
		appCallClientStream(t, app, buf)
	})

	t.Run("appCallClientStreamError", func(t *testing.T) {
		buf.Reset()
		appCallClientStreamError(t, app)
	})

	t.Run("appCallFullDuplexBidiStream", func(t *testing.T) {
		buf.Reset()
		appCallBidiStream(t, app, buf, "FullDuplexCall")
	})

	t.Run("appCallFullDuplexBidiStreamError", func(t *testing.T) {
		buf.Reset()
		appCallBidiStreamError(t, app, buf)
	})

	t.Run("appCallFullDuplexBidiStreamErrorProcessing", func(t *testing.T) {
		buf.Reset()
		appCallBidiStreamErrorProcessing(t, app, buf)
	})

	t.Run("appCallHalfDuplexBidiStream", func(t *testing.T) {
		buf.Reset()
		appCallBidiStream(t, app, buf, "HalfDuplexCall")
	})
}

func appCallUnaryServerError(t *testing.T, app *app) {
	m, ok := findMethod(t, app, "obsidian-client-cli.testing.TestService", "UnaryCall")
	if !ok {
		return
	}

	errCode := int32(codes.Internal)

	msgTmpl := `
{
  "response_status": {
    "code": %d
  }
}
`

	msg := []byte(fmt.Sprintf(msgTmpl, errCode))

	err := app.callClientStream(context.Background(), m, [][]byte{msg})
	if err == nil {
		t.Error("error expected, got nil")
		return
	}

	s, _ := status.FromError(errors.Unwrap(err))
	if s.Code() != codes.Code(errCode) {
		t.Errorf("expectd status code %v, got %v, err: %v", codes.Code(errCode), s.Code(), err)
	}
}
func appCallUnary(t *testing.T, app *app, buf *bytes.Buffer) {
	m, ok := findMethod(t, app, "obsidian_client_cli.testing.TestService", "UnaryCall")
	if !ok {
		return
	}

	userId := int32(123)
	userName := "testuser"

	msgTmpl := `
{
  "user": {"id": %d, "name": "%s"}
}
`
	msg := []byte(fmt.Sprintf(msgTmpl, userId, userName))

	if app.opts.InFormat == caller.Text {
		msgTmpl = `user { id: %d name: "%s" }`
		msg = []byte(fmt.Sprintf(msgTmpl, userId, userName))
	}

	err := app.callClientStream(context.Background(), m, [][]byte{msg})
	if err != nil {
		t.Errorf("error executing callClientStream(): %v", err)
		return
	}

	res := buf.Bytes()
	root, err := ajson.Unmarshal(res)
	if err != nil {
		t.Errorf("error unmarshaling result json: %v", err)
		return
	}

	if jsonInt32(root, "$.user.id") != userId {
		t.Errorf("payload type not found: %s", res)
		return
	}

	if jsonString(root, "$.user.name") != userName {
		t.Errorf("payload body not found: %s, %s", res, userName)
	}
}

func appCallStreamOutput(t *testing.T, app *app, buf *bytes.Buffer) {
	m, ok := findMethod(t, app, "obsidian_client_cli.testing.TestService", "StreamingOutputCall")
	if !ok {
		return
	}

	respSize1 := 3
	respSize2 := 5
	userName := "testuser"

	msgTmpl := `
{
  "user": {
	"name": "%s"
  },
  "response_parameters": [{
	"size": %d
  },{
	"size": %d
  }]
}
`

	getEncBody := func(c int) string {
		body := strings.Repeat(userName, c)
		return body
	}

	msg := []byte(fmt.Sprintf(msgTmpl, userName, respSize1, respSize2))

	if app.opts.InFormat == caller.Text {
		msgTmpl = `user { name: "%s"} response_parameters: {size: %d} response_parameters: {size: %d}`
		msg = []byte(fmt.Sprintf(msgTmpl, userName, respSize1, respSize2))
	}

	err := app.callStream(context.Background(), m, [][]byte{msg})
	require.NoError(t, err, "error executing callStream()")

	res := buf.Bytes()
	root, err := ajson.Unmarshal(res)
	require.NoError(t, err, "error unmarshaling result json")

	if len(root.MustArray()) < 2 {
		t.Fatalf("expected %d elements", 2)
	}

	assert.Equal(t, jsonString(root, "$[0].user.name"), getEncBody(respSize1))
	assert.Equal(t, jsonString(root, "$[1].user.name"), getEncBody(respSize2))
}

// appCallStreamOutputError is a test for a streaming output call with an error 
// when the response is not a valid json and the output format is json is requesting json from the server
// then the output is expected to be a json with the error in the response_status field and the response_parameters field is not present in the response
// which is the same as the output of the server and is expecting a failed call for the client side to handle
func appCallBidiStreamErrorProcessing(t *testing.T, app *app, buf *bytes.Buffer) {
	m, ok := findMethod(t, app, "obsidian-client-cli.testing.TestService", "FullDuplexCall")
	if !ok {
		return
	}
	errCode := int32(codes.Aborted)
	respSize1 := 3
	respSize2 := 5
	userName := "testuser"

	msgTmpl := `
{
  "user": {
    "name": "%s"
  },
  "response_parameters": [{
    "size": %d
  },{
    "size": %d
  }]
}
`

	msgTmplErr := `
{
  "response_status": {
    "code": %d
  }
}
`

	getEncBody := func(c int) string {
		return strings.Repeat(userName, c)
	}

	msg := []byte(fmt.Sprintf(msgTmpl, getEncBody(1), respSize1, respSize2))
	msgErr := []byte(fmt.Sprintf(msgTmplErr, errCode))

	messages := [][]byte{msg, msgErr}
	err := app.callStream(context.Background(), m, messages)
	if err == nil {
		t.Fatal("error expected, got nil")
	}

	s, _ := status.FromError(errors.Unwrap(err))
	if s.Code() != codes.Code(errCode) {
		t.Fatalf("expected status code %v, got %v", codes.Code(errCode), s.Code())
	}

	res := buf.Bytes()
	root, err := ajson.Unmarshal(res)
	if err != nil {
		t.Fatalf("error unmarshaling result json: %v", err)
	}

	if len(root.MustArray()) < 2 {
		t.Fatalf("expected %d elements, got %d", 2, len(root.MustArray()))
	}
}

// appCallBidiStream is a test for a bidi streaming call
// if the fullduplex call is successful the output is expected to be a json with the response_parameters field
// then the call is successful the response fields will allow 
// If there is not a response_parameters field the call is expected to be failed
// if the call is not successful the response_status field will contain error messages and the response_parameters field will not be present
func appCallBidiStreamError(t *testing.T, app *app, buf *bytes.Buffer) {
	m, ok := findMethod(t, app, "obsidian-client-cli.testing.TestService", "FullDuplexCall")
	if !ok {
		return
	}

	errCode := int32(codes.Internal)
	respSize1 := 3
	respSize2 := 5
	userName := "testuser"

	msgTmpl := `
{
  "user": {
    "name": "%s"
  },
  "response_parameters": [{
    "size": %d
  },{
    "size": %d
  }]
}
`

	getEncBody := func(c int) string {
		return strings.Repeat(userName, c)
	}

	msg := []byte(fmt.Sprintf(msgTmpl, getEncBody(1), respSize1, respSize2))

	ctx := metadata.AppendToOutgoingContext(context.Background(), app_testing.MethodExitCode, fmt.Sprintf("%d", errCode))
	messages := [][]byte{msg, msg}
	err := app.callStream(ctx, m, messages)
	if err == nil {
		t.Error("error expected, got nil")
		return
	}

	s, _ := status.FromError(errors.Unwrap(err))
	if s.Code() != codes.Code(errCode) {
		t.Errorf("expected status code %v, got %v, err: %v", codes.Code(errCode), s.Code(), err)
		return
	}

	resp := strings.TrimSpace(buf.String())
	if resp != "[]" {
		t.Errorf("expected `[]` response, got %s", resp)
	}
}