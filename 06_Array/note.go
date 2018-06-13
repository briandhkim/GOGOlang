/***
array
	a numbered sequence of elements of a single type
	do not change in size
	length can be discovered by using the built-in function len
	not dynamic (length grows or shortens as content chagnes)

slice
	a list;
	a slice is a descriptor for contiguous segment of an underlying array and provides access to a numbered
		sequence of elements from that array. A slice type denotes the set of all slices of array of its
		element type. The value of an uninitialized slice is nil.
	change in size
	have a length and a capacity

map
	key / value storage
	a 'dictionary'
	a map is an unordered group of elements of one type, called the element type, indexed by a set of unique
		keys of another type, called the key type. The value of an unintialized map is nil.

struct
	a data structure
	a composite type
	allows us to collect properties together

***/

package main
