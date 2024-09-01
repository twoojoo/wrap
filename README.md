# wrap 

Golang wrappers for pointers, slices and maps that grants secure access, useful methods and JSON / XML serialization support 

## usage

```go
package main

import (
  "fmt"
  "github.com/twoojoo/wrap"
)

func main() {
  
  // wrapped pointer
  ptr := wrap.NewPointer(10)
  ptr.SetValue(20)

  v, exists := ptr.GetValue()
  fmt.Println(v, exists) // 20, true

  prt.Clear()
  v, exists = ptr.GetValue() // 0, false

  // wrapped slice
  slice := wrap.NewSlice([]int{1, 2, 3})
  slice.Append(4, 5, 6)

  v, exists = slice.ValueAt(3) // 4, true
  v, exists = slice.ValueAt(10) // 0, false

  // wrapped map
  m := wrap.NewMap(map[string]int{"a": 1, "b": 2})
  m.Set("c", 3) 

  m.Keys() // ["a", "b", "c"]
  m.Values() // [1, 2, 3]
}
```