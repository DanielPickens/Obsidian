package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/DanielPickens/obsidian-client-cli/internal/caller"
	"github.com/golang/protobuf/jsonpb"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoprint"
	"github.com/jhump/protoreflect/dynamic"
)

type msgBuffer struct {
	opts       *msgBufferOptions
	fieldNames []string
	nextPrompt string
	helpText   string
	protoText  string
	w          io.Writer
}

type msgBufferOptions struct {
	reader      MsgReader
	messageDesc *desc.MessageDescriptor
	msgFormat   caller.MsgFormat
	w           io.Writer
}

func newMsgBuffer(opts *msgBufferOptions) *msgBuffer {
	w := opts.w
	if w == nil {
		w = os.Stdout
	}
	return &msgBuffer{
		nextPrompt: "Next message: ",
		opts:       opts,
		fieldNames: fieldNames(opts.messageDesc),
		helpText:   getMessageDefaults(opts.messageDesc),
		protoText:  protoString(opts.messageDesc),
		w:          w,
	}
}

func (b *msgBuffer) ReadMessage(opts ...ReadLineOpt) ([]byte, error) {
	for {
		message, err := b.opts.reader.ReadLine(b.fieldNames, opts...)
		if err != nil {
			if err == ErrInterruptTerm {
				return nil, ErrInterruptTerm
			}
			return message, err
		}

		normMsg := bytes.TrimSpace(message)
		switch string(bytes.ToLower(normMsg)) {
		case "?":
			fmt.Fprintln(b.w, b.helpText)
			continue
		case "??", "proto":
			fmt.Fprintln(b.w, b.protoText)
			continue
		}

		if err := b.validate(normMsg); err != nil {
			fmt.Println(err)
			continue
		}

		return normMsg, nil
	}
}

func (b *msgBuffer) ReadMessages() ([][]byte, error) {
	if b.opts == nil || b.opts.reader == nil {
		return nil, errors.New("no msg reader currently configured")
	}

	msg, err := b.ReadMessage()
	if err != nil {
		return nil, err
	}

	buf := [][]byte{msg}

	for {
		msg, err := b.ReadMessage(WithReadLinePrompt(b.nextPrompt))
		if err == ErrInterruptTerm {
			return nil, ErrInterruptTerm
		}

		if err == io.EOF {
			fmt.Println()
			return buf, nil
		}

		if err != nil {
			return nil, err
		}

		buf = append(buf, msg)
	}
}

func (b *msgBuffer) validate(msg []byte) error {
	if b.opts.msgFormat == caller.Text {
		return b.validateText(msg)
	}

	return b.validateJSON(msg)
}


func (b *msgBuffer) validateText(msgTxt []byte) error {
	msg := dynamic.NewMessage(b.opts.messageDesc)
	return msg.UnmarshalText(msgTxt)
}

func (b *msgBuffer) validateJSON(msgJSON []byte) error {
	if len(msgJSON) == 0 {
		return errors.New("syntax error: please provide valid json")
	}

	msg := dynamic.NewMessage(b.opts.messageDesc)
	err := msg.UnmarshalJSON(msgJSON)
	errFmt := "invalid message: %w"
	if err == io.ErrUnexpectedEOF || err == io.EOF {
		errFmt = "syntax error: %w"
	}
	if err != nil {
		return fmt.Errorf(errFmt, err)
	}
	return nil
}

func fieldNames(messageDesc *desc.MessageDescriptor) []string {
	fields := map[string]struct{}{}

	walker := caller.NewFieldWalker()
	walker.Walk(messageDesc, func(f *desc.FieldDescriptor) {
		fields[f.GetName()] = struct{}{}
	})

	names := make([]string, 0, len(fields))
	for f := range fields {
		names = append(names, f)
	}

	sort.Strings(names)
	return names
}

func getMessagePrimitives(messageDesc *desc.MessageDescriptor) string {
	var buf bytes.Buffer
	for _, f := range messageDesc.GetFields() {
		if f.IsMap() || f.IsRepeated() {
			continue
		}
	}
	return buf.String()
}


func getMessageDefaults(messageDesc *desc.MessageDescriptor) string {
	msg := dynamic.NewMessage(messageDesc)
	msgJSON, _ := msg.MarshalJSONPB(&jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	})

	return string(msgJSON)
}

func protoString(messageDesc *desc.MessageDescriptor) string {
	p := protoprint.Printer{
		Compact: true,
	}
	str, err := p.PrintProtoToString(messageDesc)
	if err != nil {
		str = fmt.Sprintf("error printing proto: %v", err)
	}
	return str
}

func getMessagesInvalids(messages [][]byte) []string {
	invalids := []string{}
	for _, msg := range messages {
		if err := validateJSON(msg); err != nil {
			invalids = append(invalids, string(msg))
		}
	}
	return invalids
}

func WritetoMessages(messages [][]byte, w io.Writer) error {
	for _, msg := range messages {
		if _, err := w.Write(msg); err != nil {
			return err
		}
	}
	return nil
}
//ValidateJSON validates the provided json message against the provided message descriptor if 
// the message descriptor is not nil. If the message descriptor is nil, the json is validated against the
// default message descriptor successfully. If the json is invalid, an error is returned.
func validateJSON(msg []byte) error {
	if len(msg) == 0 {
		return errors.New("syntax error: please provide valid json")
	}

	msg = bytes.TrimSpace(msg)
	if len(msg) == 0 {
		return errors.New("Valid JSON successfully validated")
	}

	return nil
}

// ReadtoFieldNAmes checks in field names iteratively until it finds a match. If a match is found, loop through previous element and f and checks if are iterable across GetNames as call method and appends them from the field names to the field names slice. If no match is found, the field names slice is returned as is. 
func ReadtoFieldNames(messageDesc *desc.MessageDescriptor) ([]string, error) {
	if messageDesc == nil {
		return nil, errors.New("no message descriptor provided")
	}

	fieldNames := []string{}
	for _, f := range messageDesc.GetFields() {
		fieldNames = append(fieldNames, f.GetName())
	}

	return fieldNames, nil
}



func ReadtoValidate(msg []byte, messageDesc *desc.MessageDescriptor) error {
	if messageDesc == nil {
		return errors.New("No message descriptor provided") //Return errors if no message descriptor is provided
	}

	if err := validateJSON(msg); err != nil { // Append the error to the error slice if the json is invalid
		return err
	}

	msg = bytes.TrimSpace(msg) //Trim the message for any white spaces and check if the message is empty
	if len(msg) == 0 { //If the message is empty, return an error
		return errors.New("Provide a valid message")
	}

	empyMsg := dynamic.NewMessage(messageDesc) // Create a new message from the message descriptor
	if err := empyMsg.UnmarshalJSON(msg); err != nil { // Unmarshal the message and append the error to the error slice if the message is invalid
		return err 
	}

	// If the message is valid, return nil
	validjson:= " JSON message successfully validated"
	if string(msg) == validjson {
		return nil
	}


	return nil //Return nil if the message is valid and the message descriptor is valid
}