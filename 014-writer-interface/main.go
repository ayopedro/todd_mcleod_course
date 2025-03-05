package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, Ayotunde")
	fmt.Fprintln(os.Stdout, "Hello, Ayotunde")
	io.WriteString(os.Stdout, "Hello, Ayotunde")
}

/*
	Understanding the io.Writer Interface in Go

	The io.Writer interface is a standard output abstraction in Go. It is used to write data to different destinations (e.g., files, network, buffers, stdout).

	1️⃣ Definition of io.Writer

	type Writer interface {
		Write(p []byte) (n int, err error)
	}

	🔹 It has a single method:
		•	Write(p []byte) (n int, err error): Writes bytes and returns the number of bytes written + any error.

	2️⃣ Writing to Different Destinations

	Example: Writing to os.Stdout (Console)

	package main

	import (
		"fmt"
		"os"
	)

	func main() {
		fmt.Fprintln(os.Stdout, "Hello, Writer!") // Uses io.Writer
	}

	✔ os.Stdout implements io.Writer, so fmt.Fprintln can write to it.

	Example: Writing to a File

	package main

	import (
		"os"
	)

	func main() {
		file, _ := os.Create("output.txt")
		defer file.Close()

		file.Write([]byte("Hello, File!")) // Writing directly
	}

	✔ os.File implements io.Writer, allowing us to write to files.

	Example: Writing to a Buffer (bytes.Buffer)

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		var buf bytes.Buffer
		buf.Write([]byte("Hello, Buffer!"))

		fmt.Println(buf.String()) // Convert buffer to string
	}

	✔ bytes.Buffer is useful for in-memory writes.

	3️⃣ Creating a Custom io.Writer

	You can implement io.Writer for custom behavior.

	Example: Custom Writer That Converts Text to Uppercase

	package main

	import (
		"fmt"
		"strings"
	)

	// Custom writer
	type UpperWriter struct{}

	func (uw UpperWriter) Write(p []byte) (int, error) {
		fmt.Print(strings.ToUpper(string(p)))
		return len(p), nil
	}

	func main() {
		writer := UpperWriter{}
		fmt.Fprintln(writer, "hello, custom writer!") // Output: HELLO, CUSTOM WRITER!
	}

	✔ Custom behavior (uppercase conversion) while following the io.Writer contract.

	Key Takeaways

	Feature	Description
	io.Writer	Standard interface for writing data
	os.Stdout	Writes to the console
	os.File	Writes to a file
	bytes.Buffer	Stores data in memory
	Custom Writers	You can implement Write for custom output behavior

*/
