package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Pentagon"] = Vertex{14.88, -13.77}
	fmt.Println(m["Pentagon"])

	var m2 = map[string]Vertex{
		"Kremlin":     Vertex{1.35, 2.36},
		"White House": Vertex{-10.0, -20.5},
	}
	fmt.Println(m2)

	var m3 = map[string]Vertex{
		"Pskov Kremlin": {57.82223, 28.32881},
		"White Tower":   {51.508056, -0.076111},
	}
	fmt.Println(m3)

	fmt.Println("How to retrieve element: elem = mapname[key]:", m3["Pskov Kremlin"])
	m3["Red square"] = Vertex{55.75, 37.62}
	fmt.Println("Insert or update element: mapname[key] = elem:", m3)
	delete(m3, "White Tower")
	fmt.Println("How to delete element: delete(mapname, key):", m3)
	elem, ok := m3["White Tower"]
	fmt.Println("How to test is element in map: elem, ok := mapname[key]:", elem, ok)
	elem, ok = m3["Pskov Kremlin"]
	fmt.Println(elem, ok)
}
