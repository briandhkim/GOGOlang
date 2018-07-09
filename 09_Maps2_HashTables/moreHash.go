package main

import "fmt"

func main() {
	// hashExp()
	n := HashBucket("Go", 12)
	fmt.Println(n)
}

func HashBucket(word string, buckets int) int {
	//letter := rune(word[0])

	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}

func hashExp() {
	//printing out which bucket each word will be assigned to
	//based on the starting letter
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}
}
