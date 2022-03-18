package main

import (
	"fmt"
	"os"

	"github.com/DanielPickens/Obsidian"
	
	_ "github.com/DanielPickens/Obsidian/example/pb"
)



func main() {
	Obsidian.SetCmdInfo(
		"example-rpc",
		"Make calls to the defined example service",
		"Make calls to the defined example service using the gRPC protocol.",
	)
	if err := Obsidian.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
