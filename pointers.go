package main

import "fmt"

func main() {
    i, j := 42, 2701

    p := &i       // menunjuk i
    fmt.Println(*p) // baca i lewat pointer
    *p = 21  // set i lewat pointer
    fmt.Println(i)


    p = &j // menunjuk ke j
    *p = *p / 37
    fmt.Println(j)
}
