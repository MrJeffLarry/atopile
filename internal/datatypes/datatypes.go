package datatypes

import (
"fmt"
"strconv"
"strings"
)

// IsInt checks if a string or int represents an integer
func IsInt(name interface{}) bool {
if name == nil {
return false
}

switch v := name.(type) {
case int:
return true
case string:
_, err := strconv.Atoi(v)
return err == nil
default:
return false
}
}

// TypeRef represents a reference type as a slice of strings
type TypeRef []string

// NewTypeRef creates a new TypeRef
func NewTypeRef(parts ...string) TypeRef {
return TypeRef(parts)
}

// Empty returns an empty TypeRef
func Empty() TypeRef {
return TypeRef([]string{})
}

// FromOne creates a TypeRef with a single item
func FromOne(name string) TypeRef {
return TypeRef([]string{name})
}

// FromPathStr creates a TypeRef from a dot-separated path string
func FromPathStr(path string) TypeRef {
if path == "" {
return Empty()
}
return TypeRef(strings.Split(path, "."))
}

// AddName returns a new TypeRef with the given name appended
func (t TypeRef) AddName(name string) TypeRef {
result := make([]string, len(t)+1)
copy(result, t)
result[len(t)] = name
return TypeRef(result)
}

// String returns the dot-separated string representation
func (t TypeRef) String() string {
return strings.Join(t, ".")
}

// ReferencePartType represents a part of a field reference
type ReferencePartType struct {
Name            string
Key             interface{} // Can be string or int
IsNodeReference bool
}

// NewReferencePartType creates a new ReferencePartType
func NewReferencePartType(name string, key interface{}) *ReferencePartType {
return &ReferencePartType{
Name:            name,
Key:             key,
IsNodeReference: true,
}
}

// HasKey returns true if this part has a key
func (r *ReferencePartType) HasKey() bool {
return r.Key != nil
}

// String returns the string representation
func (r *ReferencePartType) String() string {
if r.Key != nil {
return fmt.Sprintf("%s[%v]", r.Name, r.Key)
}
return r.Name
}

// FieldRef represents a field reference (e.g., app.modules[0].resistors[1])
type FieldRef struct {
Parts []*ReferencePartType
}

// NewFieldRef creates a new FieldRef
func NewFieldRef(parts ...*ReferencePartType) *FieldRef {
return &FieldRef{Parts: parts}
}

// Append returns a new FieldRef with the part appended
func (f *FieldRef) Append(part *ReferencePartType) *FieldRef {
newParts := make([]*ReferencePartType, len(f.Parts)+1)
copy(newParts, f.Parts)
newParts[len(f.Parts)] = part
return &FieldRef{Parts: newParts}
}

// Stem returns a new FieldRef without the last part
func (f *FieldRef) Stem() *FieldRef {
if len(f.Parts) == 0 {
return f
}
return &FieldRef{Parts: f.Parts[:len(f.Parts)-1]}
}

// Last returns the last part of the reference
func (f *FieldRef) Last() *ReferencePartType {
if len(f.Parts) == 0 {
return nil
}
return f.Parts[len(f.Parts)-1]
}

// String returns the dot-separated string representation
func (f *FieldRef) String() string {
parts := make([]string, len(f.Parts))
for i, part := range f.Parts {
parts[i] = part.String()
}
return strings.Join(parts, ".")
}

// ToTypeRef converts to a TypeRef if possible (no keys)
func (f *FieldRef) ToTypeRef() *TypeRef {
for _, part := range f.Parts {
if part.HasKey() {
return nil
}
}

names := make([]string, len(f.Parts))
for i, part := range f.Parts {
names[i] = part.Name
}

ref := TypeRef(names)
return &ref
}

// FromTypeRef creates a FieldRef from a TypeRef
func FromTypeRef(typeRef TypeRef) *FieldRef {
parts := make([]*ReferencePartType, len(typeRef))
for i, name := range typeRef {
parts[i] = NewReferencePartType(name, nil)
}
return &FieldRef{Parts: parts}
}
