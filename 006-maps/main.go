package main

import "fmt"

func main() {
	am := map[string]int{
		"Almond Biscotti Café":       1,
		"Banana Pudding":             2,
		"Balsamic Strawberry (GF)":   3,
		"Bittersweet Chocolate (GF)": 4,
	}

	// fmt.Println("The id for Banana Pudding is", am["Banana Pudding"])
	fmt.Println(am)
	// fmt.Printf("%#v\n", am)

	m := make(map[string]int)

	m["Lucas"] = 28
	m["Steph"] = 37
	m["George"] = 78

	// FOR range over a map
	// get the key and value
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// // get just the values
	// for _, v := range m {
	// 	fmt.Println(v)
	// }

	// // get just the keys
	// for k := range m {
	// 	fmt.Println(k)
	// }

	// // for range with a slice
	// xi := []int{1, 2, 3, 4, 5}

	// for i, v := range xi {
	// 	fmt.Println(i, v)
	// }

	// delete an element in a map
	// pass the map and the key to delete
	delete(m, "Lucas")

	// comma ok idiom aka statement-statement idiom
	// v, ok := m["Lucass"]

	// if ok {
	// 	fmt.Println("Value:", v)
	// } else {
	// 	fmt.Println("Value not found")
	// }

	// alternative way: limits the scope of v and ok
	if v, ok := m["Steph"]; !ok {
		fmt.Println("Value not found")
	} else {
		fmt.Println("Value:", v)
	}

	// fmt.Println(m["Lucas"])
	// fmt.Println(m)
	// fmt.Printf("%#v\n", m)

	// fmt.Println(len(m))

	// word frequency using maps
	a := 0
	fmt.Println(a)

	a++
	fmt.Println(a)

	n := make(map[string]int)
	fmt.Println(n["Lucas"])

}
