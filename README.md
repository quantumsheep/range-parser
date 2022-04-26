# Go range parser

A parser for `Range` header. Port of [jshttp/range-parser](https://github.com/jshttp/range-parser) to Go.

## Installation

```bash
go get github.com/quantumsheep/range-parser
```

## Example

Parse the given header string where `size` is the size of the selected representation that is to be partitioned into subranges. An array of subranges will be returned or negative numbers indicating an error parsing.

```go
import (
	"fmt"
	"net/http"

	range_parser "github.com/quantumsheep/range-parser"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ranges, err := range_parser.Parse(int64(len(arr)), r.Header.Get("Range"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, r := range ranges {
			values := arr[r.Start : r.End+1]
			w.Write([]byte(fmt.Sprintf("%v\n", values)))
		}
	})

	http.ListenAndServe(":8080", nil)
}
```

Header `Range: bytes=0-3,5-7` will return `[1 2 3 4] [6 7 8]`.
