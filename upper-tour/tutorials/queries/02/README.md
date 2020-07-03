# Build Common Queries

The `Collection` method takes the name of a table in the database and returns a
value that satisfies the [db.Collection][1] interface:

```go
booksTable := sess.Collection("books")
```

One of the methods defined by the `db.Collection` interface is `Find`.

We'll use `Find` to search for specific objects within the collection
hierarchy. `Find` returns a [db.Result][2] (which is delimited by the condition
passed to `Find` and can contain zero, one, or many items.)

> The `db.Result` API works the same on all supported databases, this is known
> as Common Result-Oriented Syntax (CROS) and comes in handy when you want to
> query a collection.

For instance, the following is a CROS query that fetches and maps all the items
from the "books" table:

```go
var books []Book

res := booksTable.Find()
err := res.All(&books)
```

You can build the query to return items in different ways, such as sorted by
title (descending order):

```go
var books []Book

res := booksTable.Find().OrderBy("-title")
err := res.All(&books)
```

Use `One` instead of `All` if you want to retrieve a single item from the set:

```go
var book Book

res := booksTable.Find(db.Cond{"id": 4})
err := res.One(&book)
```

You can also determine the total number of items in the result-set with
`Count`:


```go
res := booksTable.Find()

total, err := res.Count()
...
```

There are many options for you to define queries depending on your database
type. Take a look
[here](https://upper.io/db.v3/getting-started#defining-a-result-set-with-code-find-code).

## Query builder and raw SQL

In the particular case of adapters for SQL databases, you can also choose to
use a query builder (for more control over your query):

```go
q := sess.Select().From("books")

var books []Book
err := q.All(&books)
```

... or raw SQL (for absolute control over your query):

```
rows, err := sess.Query("SELECT * FROM books")
// rows is a regular *sql.Rows object.
```

Given that the example in this tour is based on a SQL database, we'll elaborate
hereunder on the use of both a) the SQL builder and b) raw SQL.

[1]: https://godoc.org/upper.io/db.v3#Collection
[2]: https://godoc.org/upper.io/db.v3#Result
