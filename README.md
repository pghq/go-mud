# go-mud
Golang trend forecasting library.

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
client, err := mud.NewClient()
if err != nil{
    panic(err)
}

err := client.Trends.Insert("foo", "bar", []float64{0.5, 0.1})
if err != nil{
    panic(err)
}
```
