# Store

Simple configuration storage for your CLIs. Places configuration in the right place depending on your OS.

## Example

```go
db, err := store.New("app")
defer db.Close()

err = db.Put("user", "matt")

var v string
err = db.Get("user", &v)
fmt.Println("got", v)
```

## Installation

```
go get -u github.com/matthewmueller/store
```

## Thanks

- [env-paths](https://github.com/sindresorhus/env-paths): all the research into where to place configuration files depending on your OS was done here.
- [skv](https://github.com/rapidloop/skv): this package simplified the storage piece a lot.

## License

MIT
