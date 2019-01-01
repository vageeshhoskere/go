package main

import "fmt"

func tryConst() int {
	const (
		MONDAY int = iota
		TUESDAY
		WEDNESDAY
		THURSDAY
		FRIDAY
		SATURDAY
	)
	return WEDNESDAY
}

func tryScan() {
	var num int
	fmt.Println("Enter input number")
	n, err := fmt.Scan(&num)
	fmt.Println(num, n, err)
}

func arrayCheck() {
    var x [5]int
    //index from 0
    //init with zero value
    x[0] = 2
    fmt.Println(x[0], x[1])

    //array literal
    var y = [5]int{1,2,3,4,5}
    fmt.Println(y[0])
    //... creates array
    z := [...]int{1,2,3,4}
    fmt.Println(z[0], len(z))

    for i,v := range z {
        fmt.Println("ind", i , "value", v)
    }
}

type Vertex struct {
    X int
    Y int
}

func structCheck() {
    //zero value -> individual zero value
    var v1 Vertex
    fmt.Println(v1)
    v2 := new(Vertex)
    fmt.Println(v2)
    fmt.Println(Vertex{1,2})
    v := Vertex{1,2}
    v.X = 4
    fmt.Println(v)
    fmt.Println(v.X, v.Y)
    //pointer to struct
    p := &v
    //ideally it should be (*p).Y
    p.Y = 5
    fmt.Println(*p)
    //struct literal
    v = Vertex{X:9}
    fmt.Println(v)
}

func slicesCheck() {
    x := [6]int{1,2,3,4,5,6}
    y := x[1:4]
    fmt.Println(y)
    z := x[2:5]
    z[0] = 2
    fmt.Println(y)
    fmt.Println(z)
    a := []bool{true,false}
    fmt.Println(a)
    fmt.Println(reflect.TypeOf(a))
    // capacity is counting from the first element in the slice
    b := x[:5]
    fmt.Println(b, len(b), cap(b))
    b = b[1:4]
    fmt.Println(b, len(b), cap(b))
    //zero value is nil
    var c []int
    fmt.Println(c, len(c), cap(c))
    //default capacity = length
    d := make([]int, 5)
    fmt.Println(d, len(d), cap(d))
    d = make([]int, 5, 10)
    fmt.Println(d, len(d), cap(d))
    d = append(d, 1,1,1)
    fmt.Println(d, len(d), cap(d))
}

func mapCheck() {
    //zero value is nil
    //no keys can be added without using make
    var m map[string]int
    m = make(map[string]int)
    fmt.Println(m)
    //define map literal
    n := map[string]int {"a":1}
    fmt.Println(n["a"])
    n["b"] = 2
    fmt.Println(n)
    //not nil -> returns zero value instead
    fmt.Println(n["c"])
    _, k := n["c"]
    fmt.Println(k)
    fmt.Println(len(n))
    for key, val := range(n) {
        fmt.Println(key, val)
    }
}

func main() {
	arrayCheck()
    structCheck()
    slicesCheck()
    mapCheck()
}
