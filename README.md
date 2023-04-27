# fieldmask

Fieldmask provides a convention for marking fields of a struct.
There are several reasons one may want to mark fields on a struct.
The main usecase I have are as follows:
1) Mark inclusion or exclusing in a set in a dynamic manner
2) 

Including fields in a hash:
Here, we want to automatically implement etags for this struct, but
we need some control over which fields are included, field order,
and the etag fieldname
```go
type Resource struct {
    etagMask FieldMask 

    Id string `json:"id,omitempty" fieldMask"exclude"` // Generate by server, output only, hide from etag mask
    Name string `json:"id,omitempty" fieldMask"1"`  // reorder mask fields
    Size int `json:"id,omitempty" fieldMask"2"`   // reorder mask fields

    Etag  string `json:"etag,omitempty" fieldMask:"exclude"`    // exclude mask, and mark as etag field


}
```

Only include fields that are updated:
Some APIs use the same struct for update requests. In go, you can 
omit empty values, which are nil, but some types use a zero value instead,
and then it can be hard to determine if a type was explicit set to the 
zero value or if it was unset.
```go
type Resource struct {
    updateMask FieldMask 

    Id string `json:"id,omitempty" fieldMask:"exclude"` // Generate by server, output only, hide from etag mask
    Name string `json:"id,omitempty" fieldMask:"1"`  // reorder mask fields
    Size int`json:"id,omitempty" fieldMask:"2"`   // reorder mask fields

    Etag  string `json:"etag,omitempty" fieldMask:"exclude"`    // exclude mask, and mark as etag field


}

type ResouceOption func(*Resource)

func Name(value string) ResourceOption {
	return func(target *Resource) { target.Name = value; target.fieldMask.Set(target.Name) }
}
func Size(value int) ResourceOption {
	return func(target *Resource) { target.Size = value; target.fieldMask.Set(target.Size) }
}

func UpdateResource(resourceId string, updates ...ResourceOption) Resource {
    resouce := &Resource{Id: resourceId} 
    for _, update := range updates {
        update(resource)
    }
    json := CustomJSON.Marshall(resource, RecalculateEtag(), Include(resource.updateMask))

    // make request


}
func GetResource(id string) Resource {
    return Resource{id: id, Size: 32, Name: "Resource 1", Etag: "837983798327987"}
}

func main() {
    resource := GetResource("resource-0001")
    updated := UpdateResource(resource.id, Name("Resource 2"), Size(64))
}
```