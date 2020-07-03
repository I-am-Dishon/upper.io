# Get a collection by name

Use the `Collection` method on a `db.Session` to get a reference to an specific
collection by its name:

```go
col := sess.Collection("books")
```

A collection reference satisfies [db.Collection][2], it gives you access to a
set of methods that can be used to retrieve and manipulate data, such as `Find`
(to search for specific items in the collection) and `Insert` (to add more
items to a collection).

Note that if you create a reference to a collection that doesn't exist, you'll
get WARNING message:


```
2020/07/01 00:11:33 upper/db: log_level=WARNING file=/go/src/git...
	Session ID:     00001
	Query:          SELECT "pg_attribute"."attname" AS "pkey" ...
	Error:          pq: relation "fake_collection" does not exist
	Time taken:     0.00129s
	Context:        context.Background
```

If you'd prefer to not see WARNING messages, set a higher logging level:

```
db.Log().SetLevel(db.LogLevelError)
```

Use the `Exists` method to check whether a collection exists or not:

```
ok, err := nonExistentCollection.Exists()
if errors.Is(err, db.ErrCollectionDoesNotExist) {
  log.Print("Collection does not exist")
}
```

[1]: https://godoc.org/github.com/upper/db#Session
[2]: https://godoc.org/github.com/upper/db#Collection
