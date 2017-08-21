package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68, -74.3996},
	"Google":    {37.422, -122.08},
}

func main() {
	fmt.Println(m)
}
