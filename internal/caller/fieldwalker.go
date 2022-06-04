package caller

import "github.com/jhump/protoreflect/desc"

// FieldWalker walks fields message fields tree calling func for every field
type FieldWalker struct {
	processed map[string]struct{}
}

func NewFieldWalker() *FieldWalker {
	return &FieldWalker{
		processed: map[string]struct{}{},
	}
}

func (fw *FieldWalker) Walk(md *desc.MessageDescriptor, walkFn func(*desc.FieldDescriptor)) {
	if md == nil {
		return
	}
	if _, ok := fw.processed[md.GetName()]; ok {
		return
	}
	fw.processed[md.GetName()] = struct{}{}
	for _, f := range md.GetFields() {
		fw.Walk(f.GetMessageType(), walkFn)
		walkFn(f)
	}
}

func (fw *FieldWalker) WalkEnums(md *desc.MessageDescriptor, walkFn func(*desc.EnumDescriptor)) {
	if md == nil {
		return
	}
	if _, ok := fw.processed[md.GetName()]; ok {
		return
	}
	fw.processed[md.GetName()] = struct{}{}
	for _, f := range md.GetFields() {
		fw.WalkEnums(f.GetMessageType(), walkFn)
		if f.GetEnumType() != nil {
			walkFn(f.GetEnumType())
		}
	}
}

func (fw *FieldWalker) WalkMaps(md *desc.MessageDescriptor, walkFn func(*desc.FieldDescriptor)) {
	if md == nil {
		return
	}
	if _, ok := fw.processed[md.GetName()]; ok {
		return
	}
	fw.processed[md.GetName()] = struct{}{}
	for _, f := range md.GetFields() {
		fw.WalkMaps(f.GetMessageType(), walkFn)
		if f.GetMapType() != nil {
			walkFn(f)
		}
	}
}

func (fw *FieldWalker) WalkOneOfs(md *desc.MessageDescriptor, walkFn func(*desc.OneOfDescriptor)) {
	if md == nil {
		return
	}
	if _, ok := fw.processed[md.GetName()]; ok {
		return
	}
	fw.processed[md.GetName()] = struct{}{}
	for _, f := range md.GetFields() {
		fw.WalkOneOfs(f.GetMessageType(), walkFn)
		if f.GetOneOf() != nil {
			walkFn(f.GetOneOf())
		}
	}
}

func (fw *FieldWalker) WalkOneOfFields(md *desc.MessageDescriptor, walkFn func(*desc.FieldDescriptor)) {
	var oneOfs []*desc.OneOfDescriptor
	for _, f := range md.GetFields() {
		if f.GetOneOf() != nil {
			oneOfs = append(oneOfs, f.GetOneOf())
		}
	}
	for _, oneOf := range oneOfs {
		for _, f := range oneOf.GetChoices() {
			fw.WalkOneOfFields(f.GetMessageType(), walkFn)
			walkFn(f)
		}
	}
}

// 	if md == nil {
// 		return
// 	}
// 	if _, ok := fw.processed[md.GetName()]; ok {
// 		return
// 	}
// 	fw.processed[md.GetName()] = struct{}{}
// 	for _, f := range md.GetFields() {
// 		fw.WalkOneOfFields(f.GetMessageType(), walkFn)
// 		if f.GetOneOf() != nil {
// 			for _, of := range f.GetOneOf().GetChoices() {
// 				walkFn(of)
// 			}
// 		}
// 	}
// }