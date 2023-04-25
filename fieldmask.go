package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	inset := InsetStruct{Int: 32}
	str := Struct{Int: 32, String: "s", Struct: inset}
	test := EncodeTest{
		Int:    0,
		String: "",
		Struct: str,
	}
	set := EncodeTest{
		Int:    32,
		String: "helo",
		Struct: str,
	}
	nilset := EncodeTest{
		Int:    nil,
		String: nil,
		Struct: str,
	}

	zero, _ := json.Marshal(test)
	fmt.Println(string(zero))
	exp, _ := json.Marshal(set)
	fmt.Println(string(exp))
	nils, _ := json.Marshal(nilset)
	fmt.Println(string(nils))
	ex := Example{}
	fs := FilterSet{}
	Inspect(fs)
	Inspect(ex)
}
func printField(field reflect.StructField) {
	fmt.Printf("\tFieldName:\t%s\r\n", field.Name)
	fmt.Printf("\tFieldIndex:\t%d\r\n", field.Index)
	fmt.Printf("\tTypeName:\t%s\r\n", field.Type.Name())
	fmt.Printf("\tTypePath:\t%s\r\n", field.Type.PkgPath())
	printTags(field)

}
func printTags(field reflect.StructField) {
	fmt.Printf("\tTags:\t%s\r\n", field.Tag)

}

func Inspect(item any) bool {
	kind := reflect.ValueOf(item).Kind()
	switch kind {
	case reflect.Struct:
		fmt.Printf("struct:\t%s\r\n", kind) // type of i is type of x (interface{})
	default:
		fmt.Printf("not struct:\t%s\r\n", kind) // type of i is type of x (interface{})
		return false
	}
	typ := reflect.TypeOf(item)
	mask := typ.Field(0)
	fmt.Println(mask.Type.PkgPath())
	fmt.Println(mask.Type.Name())

	field_count := typ.NumField()
	fmt.Printf("fields: %d\r\n", field_count)
	fmt.Println("Fields:")
	i := 0
	for i < typ.NumField() {

		printField(typ.Field(i))
		i++
	}

	i = 0
	maskIndex := make(map[int]reflect.StructField)

	for i < field_count {
		field := typ.Field(i)
		tags, ok := field.Tag.Lookup("fieldmask")

		if ok {
			index, err := strconv.Atoi(tags)
			if err == nil {
				maskIndex[index] = field
				fmt.Printf("%s\t%s\r\n", index, field.Name)
			}
		}

		i++
	}
	return true

}

type FieldMask uint16

func (mask FieldMask) SetIndex(field int) {
	mask |= (1 << field)

}
func (mask FieldMask) ClearIndex(field int) {
	mask &= ^(1 << field)
}

func (mask FieldMask) Set(field any) {

}
func (mask FieldMask) Clear(field any) {

}
func (mask FieldMask) Zero() {
	mask = 0
}

// func Tags(item any) (map[int]reflect.StructField, bool) {
// 	//var str struct := item.(struct)
// 	switch _type := item.(type) {
// 	case Struct:
// 		fmt.Printf("struct: %s\r\n", _type)
// 	default:
// 		fmt.Printf("not struct: %s\r\n", _type)
// 		return nil, false
// 	}

// 	typ := reflect.TypeOf(item)
// 	kind := reflect.Kind(item)
// 	fmt.Printf("type: %s\tkind: %s\r\n", typ, kind)
// 	return nil, false
// 	mask := kind.Field(0)
// 	fmt.Println(mask.Type.PkgPath())
// 	fmt.Println(mask.Type.Name())

// 	field_count := kind.NumField()
// 	i := 0
// 	maskIndex := make(map[int]reflect.StructField)

// 	for i <= field_count {
// 		field := kind.Field(i)
// 		tags, ok := field.Tag.Lookup("fieldmask")

// 		if ok {
// 			index, err := strconv.Atoi(tags)
// 			if err == nil {
// 				maskIndex[index] = field
// 				fmt.Printf("%s\t%s\r\n", index, field.Name)
// 			}
// 		}

// 		i++
// 	}

// 	return maskIndex, true
// }

type ExampleOption func(*Example)
type Example struct {
	fieldMask FieldMask

	String string `json:"example_string,omitempty" 	fieldmask:"1"`
	Int    int32  `json:"example_int,omitempty 		fieldmask:"2"`
	Bool   bool   `json:"example_bool,omitempty		fieldmask:"3"`
	Field  string `json:"example_field,omitempty 	fieldmask:"4"`
}

func String(value string) ExampleOption {
	return func(target *Example) { target.String = value; target.fieldMask.Set(1) }
}
func Int(value int32) ExampleOption {
	return func(target *Example) { target.Int = value; target.fieldMask.Set(2) }
}
func Bool(value bool) ExampleOption {
	return func(target *Example) { target.Bool = value; target.fieldMask.Set(3) }
}
func Field(value string) ExampleOption {
	return func(target *Example) { target.Field = value; target.fieldMask.Set(4) }
}

// Example usage
type Filter func(*FilterSet)
type FilterSet struct {
	fieldMask FieldMask

	TargetFilter        string
	NodesFilter         string
	ExcludedNodesFilter string
	EdgesFilter         string
	ExcludedEdgesFilter string
}

func TargetFilter(value string) Filter {
	return func(target *FilterSet) { target.TargetFilter = value; target.fieldMask.Set(1) }
}
func NodesFilter(value string) Filter {
	return func(target *FilterSet) { target.NodesFilter = value; target.fieldMask.Set(2) }
}
func ExcludedNodesFilter(value string) Filter {
	return func(target *FilterSet) { target.ExcludedNodesFilter = value; target.fieldMask.Set(3) }
}
func EdgesFilter(value string) Filter {
	return func(target *FilterSet) { target.EdgesFilter = value; target.fieldMask.Set(4) }
}
func ExcludedEdgesFilter(value string) Filter {
	return func(target *FilterSet) { target.ExcludedEdgesFilter = value; target.fieldMask.Set(5) }
}

type Resource struct {
	fieldMask FieldMask

	Name string
	Id   string
}

func ListResources(orgId, resourceId string, filters ...Filter) []Resource {
	return nil
}

type InsetStruct struct {
	Int int `json:"int,omitempty"`
}
type Struct struct {
	Int    int         `json:"int,omitempty"`
	String string      `json:"string,omitempty"`
	Struct InsetStruct `json:"struct,omitempty"`
}
type EncodeTest struct {
	Int    int    `json:"int,omitempty"`
	String string `json:"string,omitempty"`
	Struct Struct `json:"struct,omitempty"`

	UnsetInt    int    `json:"unset_int,omitempty"`
	UnsetString string `json:"unset_string,omitempty"`
	UnsetStruct Struct `json:"unset_struct,omitempty"`
}
