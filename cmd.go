package cmd

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{}

var addr = new(string)
var input = new(string)

func init() {
	rootCmd.PersistentFlags().StringVar(addr, "addr", "", "gRPC server address.")
	rootCmd.PersistentFlags().StringVar(input, "input", "", "JSON representation of the input data for the method.")
}

func SetCmdInfo(name, short, long string) {
	rootCmd.Use = fmt.Sprintf("%s [command]", name)
	rootCmd.Short = short
	rootCmd.Long = long
}

func RegisterServiceCmd(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func Execute() error {
	return rootCmd.Execute()
}

// RunD is a convenience function that runs the requested cmd and returns the error.
func runD(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing method name")
	}

	if *addr == "" {
		return fmt.Errorf("missing address")
	}

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := NewObsidianClient(conn)

	methodName := args[0]
	method := reflect.ValueOf(client).MethodByName(methodName)
	if !method.IsValid() {
		return fmt.Errorf("unknown method: %s", methodName)
	}

	inputType := method.Type().In(2)
	inputValue := reflect.New(inputType)
	if *input != "" {
		if err := json.Unmarshal([]byte(*input), inputValue.Interface()); err != nil {
			return err
		}
	}

	outputType := method.Type().In(3)
	outputValue := reflect.New(outputType)

	if err := method.Call([]reflect.Value{
		reflect.ValueOf(context.Background()),
		inputValue.Elem(),
		outputValue.Elem(),
	}); err != nil {
		return err
	}

	output, err := json.MarshalIndent(outputValue.Interface(), "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}
func RunE(
	method,
	inT string,
	newClient func(*grpc.ClientConn) interface{},
) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		conn, err := dial(*addr)
		if err != nil {
			return err
		}
		defer conn.Close()
		c := newClient(conn)
		cv := reflect.ValueOf(c)
		method := cv.MethodByName(method)
		if method.IsValid() {

			in := reflect.New(proto.MessageType(inT).Elem()).Interface()
			if len(*input) > 0 {
				if err := json.Unmarshal([]byte(*input), in); err != nil {
					return err
				}
			}

			result := method.Call([]reflect.Value{
				reflect.ValueOf(context.Background()),
				reflect.ValueOf(in),
			})
			if len(result) != 2 {
				panic("service methods should always return 2 values (error and result)")
			}
			if !result[1].IsNil() {
				return result[1].Interface().(error)
			}
			out := result[0].Interface()
			data, err := json.MarshalIndent(out, "", "    ")
			if err != nil {
				return err
			}
			fmt.Println(string(data))
		}

		return nil
	}
}
// func dial connects to the server using a secure transport and insecure and then returns a grpc.ClientConn.
func dial(addr string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	return grpc.Dial(addr, opts...)  
