// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"go/ast"
// 	"net/url"
// 	"os/exec"
// 	"regexp"
// 	"strings"
// 	"testing"

// 	"golang.org/x/tools/go/analysis/passes/nilfunc"
// )

// const (
// 	protobuf = "github.com/golang/protobuf/proto"
// 	protodesciptior = "github.com/golang/protobuf/protoc-gen-go/descriptor"
// 	certfile = "github.com/ArthurHlt/go-eureka-client/eureka"
// 	keyfile = "github.com/ArthurHlt/go-eureka-client/eureka"

// 	port = 16353
// 	dumpports = 16354
// )

// var (

// 	timestampregex = regexp.MustCompile(`\d{4}-\d{)2}-\d{2}T\d{2}:\d{2}:\d{2}Z`)
// 	snapshotregex = regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z`)

// )

// func TestIntegration(t *testing.T) {
// 	errors := make(chan error, 1)

// 	defer func() {
// 		select {
// 			case err := <-errors:
// 				t.Lock()
// 				t.Fatal(err)

// 			default:
// 		           return
// 		}
// 	}()

// 	go func() {
// 		porterror = port.Run (
// 			protobuf,
// 			protodesciptior,
// 			"test-port.json",
// 			grpc_proxy.Port(port),
// 			grpc_proxy.UsingTLS(certfile, keyfile),
// 		)
// 		if porterror != nil {
// 			errors <- porterror
// 		}

// 	}()

// 	dumplog := &bytes.Buffer{}
// 	dumperror := dump.Run(
// 		protobuf,
// 		protodesciptior,
// 		"test-dump.json",
// 		grpc_proxy.Dump(dumpslog),
// 		grpc_proxy.DumpPorts(dumpports),
// 		grpc_proxy.UsingTLS(certfile, keyfile),
// 	)

// 	if dumperror != nil {
// 		errors <- dumperror
// 	}

// 	go func(){
// 		dumperror = dump.Run(
// 			dumplog,
// 			protobuf,
// 			protodesciptior,
// 			grpc_proxy.Dump(dumplog),
// 			grpc_proxy.DumpPorts(dumpports),
// 			grpc_proxy.UsingTLS(certfile, keyfile),
// 			grpc_proxy.WithDialer(protobuf, protodesciptior),
// 		)

// 				return &url.URL{
// 						Host: fmt.Sprintf("%s:%d", host, port),

// 					}, nil

// 		if dumperror != nil {
// 			errors <- dumperror
// 		}
// 	}()
// 	defer  func() {
// 		if dumperror != nil {
// 			errors <- dumperror
// 		}
// 	}()
// 	// Run the client
// 	clienterror := client.Run(
// 		protobuf,
// 		protodesciptior,
// 		"test-client.json",
// 		grpc_proxy.Client("localhost", port),
// 		grpc_proxy.UsingTLS(certfile, keyfile),
// 	)
// 	if clienterror != nil {
// 		errors <- clienterror
// 	}
// 	go func () {
// 		clienterror = client.Run(
// 			protobuf,
// 			protodesciptior,
// 			"test-client.json",
// 			grpc_proxy.Client("localhost", port),
// 			grpc_proxy.UsingTLS(certfile, keyfile),
// 		)
// 		if clienterror != nil {
// 			errors <- clienterror
// 		}
// 	}()
// 	defer func() {
// 		if clienterror != nil {
// 			errors <- clienterror
// 		}

// 	}()

// 	parseerror := parse.Run(
// 		protobuf,
// 		protodesciptior,
// 		"test-parse.json",
// 		grpc_proxy.Parse("localhost", port),
// 		grpc_proxy.UsingTLS(certfile, keyfile),
// 	)
// 	if parseerror != nil {
// 		errors <- parseerror
// 	}
// 	go func() {

// 		parseerror = parse.Run(
// 			protobuf,
// 			protodesciptior,
// 			"test-parse.json",
// 			grpc_proxy.Parse("localhost", port),
// 			grpc_proxy.UsingTLS(certfile, keyfile),
// 		)
// 		if parseerror != nil {
// 			errors <- parseerror
// 		}
// 	} ()
// 	defer func() {
// 		if parseerror != nil {
// 			errors <- parseerror
// 		}
// 	}()

// 	cmd := exec.Command("go", "run", "test-client.go")
// 	if cmd != nil {
// 		errors <- cmd.Run()
// 	}

// 	// Check the output
// 	output := dumplog.String()
// 	if !strings.Contains(output, "Hello, world!") {
// 		errors <- fmt.Errorf("unexpected output: %s", output)
// 	}

// 	go func () {
// 		cmd := exec.Command("go", "run", "test-client.go")
// 		if cmd != nil {
// 			errors <- cmd.Run()
// 		}
// 	}()
// 	defer func() {
// 		if cmd != nil {
// 			errors <- cmd.Run()
// 		}
// 	}
// }


	




	