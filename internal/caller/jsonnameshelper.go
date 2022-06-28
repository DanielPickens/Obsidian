package caller

import "github.com/jhump/protoreflect/desc"

type jsonNamesUpdater struct {
	walker *FieldWalker
}

func newJsonNamesUpdater() *jsonNamesUpdater {
	return &jsonNamesUpdater{
		walker: NewFieldWalker(),
	}
}

// updateJSONNames sets JsonName property to camelCased version of field name if it is empty.
func (u *jsonNamesUpdater) updateJSONNames(t *desc.MessageDescriptor) {
	u.walker.Walk(t, func(f *desc.FieldDescriptor) {
		u.updateJSONName(f)
	})
}

func (u *jsonNamesUpdater) updateJSONName(f *desc.FieldDescriptor) {
	fdp := f.AsFieldDescriptorProto()
	if fdp.GetJsonName() == "" {
		cc := toLowerCamelCase(f.GetName())
		fdp.JsonName = &cc
	}
}

func newJsonNamesHelper() *jsonNamesHelper {
	return &jsonNamesHelper{
		updater: newJsonNamesUpdater(),
	}
}

func newJsonNamesHelperFieldWalker() *jsonNamesHelperFieldWalker {
	fdp := NewFieldDescriptorProto()
	return &jsonNamesHelperFieldWalker{
}
