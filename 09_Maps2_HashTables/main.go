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
	// hashingWordsOne()
	// hashingMobyDick()
	hashingMobyDickMoreEven()
}

func hashingMobyDickMoreEven() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	//scan page
	scanner := bufio.NewScanner(res.Body)
	//defer will execute `res.Body.Close` before exiting this function
	defer res.Body.Close()
	//set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	//create slice to hold counts
	buckets := make([][]string, 12)
	//create slices to hold words words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	//loop over words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucketEven(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	//print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
	//print the words in one of the buckets
	fmt.Println(buckets[6])
}

func HashBucketEven(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}

func hashingMobyDick() {
	//get Moby Dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	//scan page
	scanner := bufio.NewScanner(res.Body)
	//defer will execute `res.Body.Close` before exiting this function
	defer res.Body.Close()
	//set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	//create slice to hold counts
	buckets := make([]int, 12)
	//loop over words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}

	//prints buckets A-z
	//[1790 1260 910...4 1 2272]
	//1790 words starting with capital A; 2272 words starting with lower case z
	//uneven bucket distribution
	// fmt.Println(buckets[65:122])

	fmt.Println(buckets)

	// fmt.Println("*******************")
	// for i := 28; i <= 126; i++ {
	// 	fmt.Println("%v - %c - %v \n", i, i, buckets[i])
	// }

}

func HashBucket(word string, buckets int) int {
	// return int(word[0])

	//improving uneven hash bucket
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
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
