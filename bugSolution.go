```go
package main

import "fmt"

func main() {
    var m map[string]int
    value, ok := m["a"]
    if ok {
        fmt.Println(value) // Access the value if the key exists
    } else {
        fmt.Println("Key not found") //Handle the case where the key is absent
    }
}
```