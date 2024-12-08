# Uninitialized Map Access in Go

This example demonstrates a subtle issue in Go regarding accessing elements in an uninitialized map.  In many other languages, attempting to access an element in an uninitialized map would result in a runtime error or exception.  However, Go handles this differently.

The `bug.go` file shows how accessing a key in an uninitialized map will not cause a panic.  Instead, Go will return the zero value for the map's value type (in this case, 0 for `int`). This behavior can lead to unexpected results and hidden bugs.