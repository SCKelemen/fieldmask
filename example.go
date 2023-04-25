package test

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

type Resouce struct {
	fieldMask FieldMask

	Name string
	Id   string
}

func ListResources(orgId, resourceId string, filters ...Filter) []Resource {

}
