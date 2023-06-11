package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// go docs:https://go.dev/doc/
// go web docs：https://go.dev/doc/articles/wiki/
func main() {
	http.HandleFunc("/", handler)
	testSlice()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// slice: 使用子切片的天坑，避免改底层数组造成连锁反应
func testSlice() {
	i, j := make([]int, 0, 8), make([]int, 8)
	i = append(i, 1, 2, 3, 4, 5, 6, 7, 8)
	copy(j[0:2], i[1:])
	i[2] = 99
	fmt.Printf("i: %v\n", i)
	fmt.Printf("j: %v\n", j)
	switch i[0] {
	case 1:
	}

}
