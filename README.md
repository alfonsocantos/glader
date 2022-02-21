# g.l.a.d.er

## Get, List, Add, Delete things


Just a tiny implementation of the [glader interface](interfaces.go) with a memory store.

Start a Glader

```go
myglader := glader.New()
```

### Get

Gets an item from the store

```go
item := myglader.Get("my-item")
```

### List

Lists all items from the store. It returns only a list of ids.

```go
itemList := myglader.List()
```

### Add

It adds an item to the store by id.

```go
item := struct {
	field1 string
}{
	"field1": "my data"
}

myglader.Add("my-item", item)
```

### Add with TTL

It adds an item to the store by id, with Time To Live. To do so, use ttl glader using myglader as store

```go
item := struct {
	field1 string
}{
	"field1": "my data"
}

myTTLglader := ttl.Neew(myglader, time.Second)

myTTLglader.Add("my-item", item)
```

Item will be deleted after 1 second

### Delete

Deletes an item to the store by id.

```go
myglader.Delete("my-item")
```

Item will be deleted after 1 second