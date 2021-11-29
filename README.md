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

To create a new client:

```
cls := mud.NewClassifier()
if err := cls.Insert("foo", "bar", []float64{0.5, 0.1}); err != nil{
    panic(err)
}
```
