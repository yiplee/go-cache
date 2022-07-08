# go-cache

go-cache is a simple, fast, and efficient in-memory cache written in Go.

### Installation

```bash
go get github.com/yiplee/go-cache
```

### Example

```go 
package cache-example

import (
    "fmt"
    "github.com/yiplee/go-cache"
)

func main() {
    cache := cache.New[string]()
    cache.Set("key", "value", cache.WithTTL(10 * time.Second))
    value, ok := cache.Get("key")
    fmt.Println(value, ok)
}

```
