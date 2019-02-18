Typed json store generator
--------------------------

### How to use

```
jstore TypeName store/typefile.go
```

```go
import "github.com/youname/youproject/store"

type TypeName struct {
    ID   string
    Name string
    Age  int
}

data, err := NewTypeNameCollection("./data");
if err != nil {
    log.Fatal(err)
}

//get record
some, _ := data.Get("some-id")

//get all records
for _, rec := range data.GetAll(){
    //do something
}

//save record
data.Save( TypeName{} )

//delete record
data.Remove("some-id")
```

### Limitations

- not thread safe

### License 

MIT