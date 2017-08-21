package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.1, -74.1,
	},
	"Google": Vertex{
		37.4, -122.1,
	},
}

func main() {
	fmt.Println(m)
}
