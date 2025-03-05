package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":28}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// people := []person{}
	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nAll of the data:", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}

/*
	- func Unmarshal(data []byte, v interface{}) error
	- Unmarshal parses the JSON encoded data and stores the result in the value pointed to by v

	- JSON to Go website: https://mholt.github.io/json-to-go/

	Understanding Struct Tags in Go

	1️⃣ What Are Struct Tags?

	Struct tags provide metadata about struct fields. The json:"First" tag tells Go’s JSON package how to map JSON keys to struct fields.

	Example

	type Person struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Age   int    `json:"age"`
	}

	🔹 Without tags: JSON uses field names as-is.
	🔹 With tags: JSON keys can be customized (e.g., lowercase "first" instead of "First").

	Usage in JSON Encoding & Decoding

	import (
		"encoding/json"
		"fmt"
	)

	func main() {
		p := Person{"John", "Doe", 30}

		// Convert struct to JSON
		jsonData, _ := json.Marshal(p)
		fmt.Println(string(jsonData)) // Output: {"first":"John","last":"Doe","age":30}
	}

	2️⃣ Handling Unknown JSON Fields

	If the JSON has extra fields not defined in the struct, Go ignores them by default.
	But to capture unknown fields, use map[string]interface{} or json.RawMessage.

	Option 1: Use map[string]interface{}

	✔ Allows any JSON structure
	✔ But loses type safety

	var data map[string]interface{}
	json.Unmarshal([]byte(`{"name": "Alice", "age": 25, "city": "NY"}`), &data)
	fmt.Println(data["name"]) // "Alice"

	Option 2: Use json.RawMessage

	✔ Captures extra fields while keeping known fields structured

	type PersonWithExtras struct {
		First  string          `json:"first"`
		Last   string          `json:"last"`
		Extras json.RawMessage `json:"-"` // Store unknown fields as raw JSON
	}

	var p PersonWithExtras
	json.Unmarshal([]byte(`{"first":"John","last":"Doe","nickname":"JD"}`), &p)
	fmt.Println(string(p.Extras)) // `{"nickname":"JD"}`

	Key Takeaways

	Feature	Behavior
	Struct Tags	Maps JSON keys to struct fields
	Unknown Fields	Ignored by default
	map[string]interface{}	Captures dynamic fields (loses type safety)
	json.RawMessage	Stores extra fields as raw JSON

*/
