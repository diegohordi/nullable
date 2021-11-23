#Nullable

Very simple Go module to handle nullable fields. Basically, it adds to `sql` package types the JSON marshal and
unmarshal features. It has 100% of coverage and includes also `Scan` tests to make sure that empty values from 
database are properly assigned.

## How to use

### Add module
`go get github.com/diegohordi/nullable`

### Import when needed
`import "github.com/diegohordi/nullable"`

## TODO

- [ ] Fuzzy tests
- [ ] Benchmark tests
