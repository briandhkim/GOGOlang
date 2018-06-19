package main

import "fmt"

func main() {
	//slice containing slice of strings
	records := make([][]string, 0)

	//student 1
	student1 := make([]string, 4)
	student1[0] = "Smith"
	student1[1] = "John"
	student1[2] = "95.00"
	student1[3] = "76.00"
	//store to the record
	records = append(records, student1)

	//student 2
	student2 := make([]string, 4)
	student2[0] = "Doe"
	student2[1] = "Jane"
	student2[2] = "80.00"
	student2[3] = "94.00"
	records = append(records, student2)

	fmt.Println(records)

	// doubleSliceInt()
	loopRecords(records)
}

func loopRecords(rec [][]string) {
	for i := 0; i < len(rec); i++ {
		for m := 0; m < len(rec[i]); m++ {
			fmt.Println("looping records: i - ", i, "; m - ", m)
			fmt.Println("record[i] - ", rec[i], "; records[i][m] - ", rec[i][m])
		}
	}
}

//extra slice of slice of int
func doubleSliceInt() {
	transactions := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		transaction := make([]int, 0)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}
