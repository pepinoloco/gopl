// Package unsafeptr demonstrates basic use of unsafe.Pointer.
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var x struct {
        a bool
        b int16
        c []int
    }

    // equivalent to pb := &x.b
    pb := (*int16)(unsafe.Pointer(
        uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
    *pb = 42

    fmt.Println(x.b) // "42"
}

/*
    // NOTE: subtly incorrect!
    tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
    pb := (*int16)(unsafe.Pointer(tmp))
    *pb = 42
*/
