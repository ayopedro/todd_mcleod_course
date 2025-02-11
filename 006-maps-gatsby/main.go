package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// open a file
	f, err := os.Open("great-gatsby.txt")

	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer f.Close()

	// the frequency of words in a file
	words, err := freq(f)

	if err != nil {
		log.Fatalf("Error getting word frequency: %s", err)
	}

	// display the word frequencies
	// for word, frq := range words {
	// 	fmt.Printf("%s: %d\n", word, frq)
	//	}

	// sort the word frequencies
	pairs := sortWordFrequency(words)

	// print the sorted word frequencies
	for _, pair := range pairs {
		log.Printf("%s: %d\n", pair.word, pair.frequency)
	}

	// word with the highest frequency and it's frequency
	w, n, err := maxWord(words)
	if err != nil {
		log.Fatalf("Error getting word with highest frequency: %s", err)
	}

	fmt.Printf("%#v has a frquency %d", w, n)
}

func freq(r io.Reader) (map[string]int, error) {
	// create a map to store the frequency of words
	wordFreq := make(map[string]int)

	// create a scanner to read the file
	s := bufio.NewScanner(r)
	s.Split((bufio.ScanWords))

	// read each word and update the frequency
	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordFreq[word]++
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("error scanning words: %s", err)
	}

	return wordFreq, nil
}

// for sorting the words
type Pair struct {
	word      string
	frequency int
}

// implement the Len, Less and Swap methods on the PairList
// to satisfy the sort.Interface interface
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].frequency > p[j].frequency
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func sortWordFrequency(words map[string]int) PairList {
	// convert the map to a PairList
	pl := make(PairList, len(words))
	i := 0
	for k, v := range words {
		pl[i] = Pair{k, v}
		i++
	}

	// sort the word frequencies
	sort.Sort(pl)

	return pl
}

// word frequency in go
// https://chat.openai.com/share/6cb2b004-4cfa-4aad-a2d8-47b0eacd36dd

// sort maps in go
// turn them into structs
// get a slice of those structs
// implement the sort.Interface interface
// https://chat.openai.com/share/03a44e91-fc0d-4cdb-884a-c8acd8f440d8

func maxWord(m map[string]int) (string, int, error) {
	if len(m) == 0 {
		return "", 0, fmt.Errorf("no words in the map")
	}

	maxW := "" // word with the highest frequency
	maxF := 0  // frequency of the word

	for w, f := range m {
		if f > maxF {
			maxW = w
			maxF = f
		}
	}

	return maxW, maxF, nil
}

/*
In the context of the `bufio` package in Go, a "token" refers to a sequence of characters that represents a meaningful unit of data within an input stream. The `bufio` package provides functionality for reading and writing data in buffered streams, and it includes a `Scanner` type that is designed to scan and tokenize input.

The `Scanner` type in the `bufio` package has a method called `Scan()`, which reads the next token from the input stream and advances the scanner to the next position. The exact definition of a token depends on how the `Scanner` is configured.

By default, the `Scanner` uses white space (spaces, tabs, and newlines) as the delimiter to split the input into tokens. So, when you call `Scan()`, it will read the next token until it encounters a white space character. For example, if you have the input string "Hello World", calling `Scan()` twice would return "Hello" and "World" as separate tokens.

However, you can customize the tokenization behavior by setting a different split function using the `Scanner`'s `Split()` method. You can provide a function that defines your own criteria for what constitutes a token. This allows you to tokenize the input based on specific patterns or delimiters that are relevant to your use case.

In summary, in the context of the `bufio` package, a token represents a meaningful unit of data that is obtained by scanning an input stream using a `Scanner`. It can be a word, a sentence, or any other logical division of data based on the defined tokenization rules.

*/

/*
write code for me in the go programming language to read all of the words from a text file and count the frequency of each word in the text file
*/
// https://chat.openai.com/share/6cb2b004-4cfa-4aad-a2d8-47b0eacd36dd

/*
in the go programming language, I have a map with a string and an int. How do I sort the int values from largest to smallest and see their corresponding string value?
*/
// https://chat.openai.com/share/03a44e91-fc0d-4cdb-884a-c8acd8f440d8
