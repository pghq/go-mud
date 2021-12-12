# go-mud
Golang supervised classification library.

## Installation

go-mud may be installed using the go get command:

```
go get github.com/pghq/go-mud
```
## Usage

```
import "github.com/pghq/go-mud"
```

To create a new graph:

```
g := mud.New()
if err := g.Plot([]byte("foo"), "bar", []float64{0.5, 0.1}); err != nil{
    panic(err)
}

if err := g.Plot([]byte("baz"), "qux", []float64{1.5, 7.1, 5}, "tag1"); err != nil{
    panic(err)
}

var values []string
if err := g.Nearest([]float64{7, 3}, &values); err != nil{
    panic(err)
}
```
