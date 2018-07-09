package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

/**
Hash table:
	type - unordered associative array
	Search, Insert, Delete, Space -> O(n)
hash tables advantages:
	fast search

in order of slowest to fastest in searching:
	unsorted
	sorted
	categorized (buckets/slots)
		(e.g. - words in As, Bs, Cs... Zs
			in this case, some 'buckets' may have more
			data then others; [Ss might have more than
			Zs])
		-need even distribution between 'buckets'
			this is done by running the data through a hash algorithm
			which would distribute the content evenly across buckets
			looking up data would require running it through the algorithm
		-hash will take an input and always produce a certain output
			if the input changes even a little bit, the output will be
			drastically different; every input produces a unique output

- maps are an implementation of hash table
**/

func main() {
	hashingWordsOne()
}

func hashingWordsOne() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}

	i := 0
	for k, _ := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}

}
