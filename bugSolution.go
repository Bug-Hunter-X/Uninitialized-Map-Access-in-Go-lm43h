```go
package main

import (

    "fmt"
)

func main() {
    var m map[string]int
    m = make(map[string]int)
    fmt.Println(m["a"]) //Prints 0 because it's initialized
    m["a"] = 10
    fmt.Println(m["a"]) //Prints 10
    value, ok := m["b"] //Properly check for a key
    if ok {
        fmt.Println(value) //Prints nothing because the key does not exist
    } else{
        fmt.Println("Key does not exist")
    }
}
```