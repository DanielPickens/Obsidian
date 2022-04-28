package Obsidian

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type Obsidian struct {
	gen *generator.Generator
}

func init() {
	generator.RegisterPlugin(new(Obsidian))
}

// Name returns the name of this plugin, "Obsidian".
func (g *Obsidian) Name() string {
	return "Obsidian"
}

// P forwards to g.gen.P.
func (g *Obsidian) P(args ...interface{}) { g.gen.P(args...) }

func (g *Obsidian) Init(gen *generator.Generator) {
	g.gen = gen
}
func (g *Obsidian) GenerateImports(file *generator.FileDescriptor) {
	g.P("// Obsidian imports")
	g.P(`import (
	"github.com/DanielPickens/Obsidian"
	"github.com/spf13/cobra"
)`)
}

func (g *Obsidian) Generate(file *generator.FileDescriptor) {
	g.P("// Begin Obsidian ")

	g.P("var _ = Obsidian.RunE")

	for _, f := range g.gen.Request.ProtoFile {
		for _, s := range f.GetService() {
			var methodVars []string
			name := s.GetName()

			g.P("// ", name)
			serviceCmdVar := fmt.Sprintf("_%sCmd", name)
			g.P("var ", serviceCmdVar, " = &cobra.Command{")
			g.P(`Use: "`, lowerFirst(name), ` [method]",`)
			g.P(`Short: "Subcommand for the `, name, ` service.",`)
			g.P("}")
			g.P()

			for _, m := range s.GetMethod() {
				methodName := m.GetName()
				methodCmdVar := fmt.Sprintf("_%s_%sCmd", name, methodName)
				methodVars = append(methodVars, methodCmdVar)
				g.P("var ", methodCmdVar, " = &cobra.Command{")
				g.P(`Use: "`, lowerFirst(methodName), `",`)
				g.P(fmt.Sprintf(
					`Short: "Make the %s method call, input-type: %s output-type: %s",`,
					methodName,
					toTypeName(m.GetInputType()),
					toTypeName(m.GetOutputType()),
				))
				g.P(fmt.Sprintf(
					`RunE: Obsidian.RunE(
						"%s",
						"%s",
						func(c *grpc.ClientConn) interface{} {
						return New%sClient(c)
					},
				),`,
					methodName,
					toTypeName(m.GetInputType()),
					name,
				))
				g.P("}")
				g.P()

			}

			g.P("// Register commands with the root command and service command")
			g.P("func init() {")
			g.P("Obsidian.RegisterServiceCmd(", serviceCmdVar, ")")
			g.P(serviceCmdVar, ".AddCommand(")
			for _, n := range methodVars {
				g.P(n, ",")
			}
			g.P(")")
			g.P("}")
			g.P()
		}
	}

	g.P("// End Obsidian")
}

func upperFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}
func toTypeName(t string) string {
	
	return strings.TrimPrefix(t, ".")
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}


func GenerateExports(gen *generator.Generator) {
	o := &Obsidian{gen: gen}
	o.Init(gen)
	o.GenerateImports(gen.Request.FileToGenerate[0])
	o.Generate(gen.Request.FileToGenerate[0])
}

