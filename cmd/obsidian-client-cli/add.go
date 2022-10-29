package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/urfave/cli"
)

func add(c *cli.Context) error {
	if c.NArg() < 2 {
		return errors.New("Please add a test case name and a test case file")
		cli.ShowCommandHelp(c, "add")
		return nil
	}

	targetname, outputpath := c.Args().Get(0), c.Args().Get(1)
	protoroot = c.String("protoroot")
	protofile = c.String("protofile")
	protoimport = c.String("protoimport")
	protoimportpath = c.String("protoimportpath")

	if protoroot == "" {
		protoroot = "./"
	}
	_err := os.MkdirAll(protoroot, 0755)
	if _err != nil {
		return _err
	}
	for _, _import := range strings.Split(protoimport, ",") {
		if _import != "" {
			protoimportpath = protoimportpath + ":" + _import
		}
	}
	_, err = os.Stat(outPath)
	if err != nil {
		fmtPrintln(" output directory already exists " + outPath)
		return nil
	}

	cuefile := protoroot + "/" + targetname + ".cue"
	cue, err := GenerateCueFile(cuefile, protofile, protoimportpath)
	if err != nil {
		if err.Error() == "No proto file found" {
			fmtPrintln("No file exists. Generating Proto file")
			return nil
		}
	}
	tmplt = template.New("testcase")
	tmplt, err = tmplt.Parse(testcase)
	m := make(map[string]string)
	m["name"] = targetname
	m["outputpath"] = outputpath
	m["cuefile"] = cuefile
	m["protofile"] = protofile
	m["protoimportpath"] = protoimportpath
	m["protoroot"] = protoroot
	m["protoimport"] = protoimport
	m["cue"] = cue
	f, err := os.Create(protoroot + "/" + targetname + ".go")

	if err != nil {
		return err
	}
	f, err := os.Create(outputpath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = tmplt.Execute(f, cue)
	if err != nil {
		return err
	}
	return nil
}

func GenerateCueFile(cuefile, protofile, protoimportpath string) (string, error) {
	proto, err := os.Open(protofile)
	if err != nil {
		return "", err
	}
	defer proto.Close()

	// Read the proto file
	protocontent, err := ioutil.ReadAll(proto)
	if err != nil {
		return "", err
	}

	// Generate the cue file
	cue, err := os.Create(cuefile)
	if err != nil {
		return "", err
	}
	defer cue.Close()

	// Read the proto file
	cuecontent, err := ioutil.ReadAll(cue)
	if err != nil {
		return "", err
	}
	return string(cuecontent), nil
}
