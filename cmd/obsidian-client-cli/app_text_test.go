package main

import (
	"bytes"
	"testing"

	"github.com/DanielPickens/obsidian-client-cli/internal/caller"
	app_testing "github.com/DanielPickens/obsidian-client-cli/internal/testing"
)

func TestAppServiceCallsProtoText(t *testing.T) {
	appOpts := &startOpts{
		Target:        app_testing.TestServerAddr(),
		Deadline:      15,
		IsInteractive: false,
		InFormat:      caller.Text,
	}

	buf := &bytes.Buffer{}
	appOpts.w = buf
	app, err := newApp(appOpts)

	if err != nil {
		t.Error(err)
		return
	}

	t.Run("appCallUnary", func(t *testing.T) {
		buf.Reset()
		appCallUnary(t, app, buf)
	})

	t.Run("appCallStreamOutput", func(t *testing.T) {
		buf.Reset()
		appCallStreamOutput(t, app, buf)
	})
}

func TestAppServiceCallsProtoBinary(t *testing.T) {
	appOpts := &startOpts{
		Target:        app_testing.TestServerAddr(),
		Deadline:      15,
		IsInteractive: false,
		InFormat:      caller.Binary,
	}

	buf := &bytes.Buffer{}
	appOpts.w = buf
	app, err := newApp(appOpts)

	if err != nil {
		t.Error(err)
		return
	}

	t.Run("appCallUnary", func(t *testing.T) {
		buf.Reset()
		appCallUnary(t, app, buf)
	}
	)

	t.Run("appCallStreamOutput", func(t *testing.T) {
		buf.Reset()
		appCallStreamOutput(t, app, buf)
	}
	)
}

func TestAppServiceCallsProtoJSON(t *testing.T) {
	appOpts := &startOpts{
		Target:        app_testing.TestServerAddr(),
		Deadline:      15,
		IsInteractive: false,
		InFormat:      caller.JSON,
	}

	buf := &bytes.Buffer{}
	appOpts.w = buf
	app, err := newApp(appOpts)

	if err != nil {
		t.Error(err)
		return
	}

	t.Run("appCallUnary", func(t *testing.T) {
		buf.Reset()
		appCallUnary(t, app, buf)
	}
	)

	t.Run("appCallStreamOutput", func(t *testing.T) {
		buf.Reset()
		appCallStreamOutput(t, app, buf)
	}
	)
}


func 