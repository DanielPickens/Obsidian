package main

import (
	"testing"

	app_testing "github.com/DanielPickens/obsidian-client-cli/internal/testing"
)

func TestAppServiceCallsNoReflect(t *testing.T) {
	runAppServiceCalls(t, &startOpts{
		Target:        app_testing.TestServerNoReflectAddr(),
		Deadline:      15,
		IsInteractive: false,
		Protos:        []string{"../../testdata/test.proto"},
	})
}

func TestAppServiceCallsNoReflectInteractive(t *testing.T) {
	runAppServiceCalls(t, &startOpts{
		Target:        app_testing.TestServerNoReflectAddr(),
		Deadline:      15,
		IsInteractive: true,
		Protos:        []string{"../../testdata/test.proto"},
	})
}

func TestAppServiceCallsReflect(t *testing.T) {
	runAppServiceCalls(t, &startOpts{
		Target:        app_testing.TestServerReflectAddr(),
		Deadline:      15,
		IsInteractive: false,
		Protos:        []string{"../../testdata/test.proto"},
	})
}

func TestAppServiceCallsReflectInteractive(t *testing.T) {
	runAppServiceCalls(t, &startOpts{
		Target:        app_testing.TestServerReflectAddr(),
		Deadline:      15,
		IsInteractive: true,
		Protos:        []string{"../../testdata/test.proto"},
	})
}

func TestAppServiceCallsNoReflectMultiple(t *testing.T) {
	runAppServiceCalls(t, &startOpts{
		Target:        app_testing.TestServerNoReflectAddr(),
		Deadline:      15,
		IsInteractive: false,
		Protos:        []string{"../../testdata/test.proto", "../../testdata/test.proto"},
	})
}

func TestAppServiceCallsNoReflectMultipleInteractive(t *testing.T) {
	runAppServiceCalls(t, &startOpts{
		Target:        app_testing.TestServerNoReflectAddr(),
		Deadline:      15,
		IsInteractive: true,
		Protos:        []string{"../../testdata/test.proto", "../../testdata/test.proto"},
	})
}

func TestInactiveServiceCallsReflect(t *testing.T) {
	runAppServiceCalls(t, &startOpts{
		Target:        app_testing.TestServerReflectAddr(),
		Deadline:      15,
		IsInteractive: false,
		Protos:        []string{"../../testdata/test.proto"},
		Service:       "TestService",
	})
}