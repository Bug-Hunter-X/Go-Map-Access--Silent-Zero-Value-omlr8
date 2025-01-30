```go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["a"]) // This will not cause a runtime error, but will print 0.
}
```