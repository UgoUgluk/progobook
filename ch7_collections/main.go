package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	fmt.Println("Hello, Collections")

	//arrays
	var names [3]string
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)

	//literal init arrays
	names2 := [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names2)
	names3 := [...]string{"Kayak", "Lifejacket"}
	names3[1] = "Paddle3"
	fmt.Println(names3)

	//copy and link arrays
	names4 := [3]string{"Kayak", "Lifejacket", "Paddle"}
	otherArray := names4
	linkArray := &names4
	names4[0] = "Canoe"
	fmt.Println("names:", names4)
	fmt.Println("otherArray:", otherArray)
	fmt.Println("linkArray:", *linkArray)

	//compare arrays
	fmt.Println("comparison:", names == names2)

	//range
	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}
	//range without index
	for _, value := range names {
		fmt.Println("Value:", value)
	}

	//slices
	slice := make([]string, 3)
	slice[0] = "Kayak"
	slice[1] = "Lifejacket"
	slice[2] = "Paddle"
	fmt.Println(slice)

	slice2 := []string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(slice2)

	//add el to slice
	slice2 = append(slice2, "Hat", "Gloves")
	fmt.Println(slice2)

	//extand slice
	slice3 := append(slice2, "Hat2", "Gloves2")
	fmt.Println(slice3)

	//capacity slice
	slice4 := make([]string, 3, 6)
	slice4[0] = "Kayak"
	slice4[1] = "Lifejacket"
	slice4[2] = "Paddle"
	fmt.Println("len:", len(slice4))
	fmt.Println("cap:", cap(slice4))

	//slice with capacity is not copied but linked
	appendedSlice4 := append(slice4, "Hat", "Gloves")
	slice4[0] = "Canoe"
	fmt.Println("slice4:", slice4)
	fmt.Println("appendedSlice4:", appendedSlice4)

	//add slice to slice
	appendedSlices := append(slice4, appendedSlice4...)
	fmt.Println("appendedNames:", appendedSlices)

	//create slices from arrays
	//someNames 	array					allNames
	//				[0][Kayak		]	<----	[0]
	//[0]	---->	[1][Lifejacket	]	<----	[1]
	//[1]	---->	[2][Paddle		]	<----	[2]
	//[2]	---->	[3][Hat			]	<----	[3]
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]

	//someNames = append(someNames, "Gloves")

	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))

	//slice capacity
	someNames2 := products[1:3:3]
	someNames2 = append(someNames2, "Gloves")
	fmt.Println("someNames2:", someNames2)
	fmt.Println("someNames2 len:", len(someNames2), "cap:", cap(someNames2))

	//create slice from other slice
	allNames2 := products[1:]
	fmt.Println("allNames2", allNames2)
	someNames3 := allNames2[1:3]
	allNames2 = append(allNames2, "Gloves")
	allNames2[1] = "Canoe"
	fmt.Println("someNames3:", someNames3)
	fmt.Println("allNames2", allNames2)

	//copy slices
	products2 := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames3 := products2[1:]      //[Lifejacket Paddle Hat]
	someNames4 := make([]string, 2) //dont use - var someNames []string - it out [] after copy
	copy(someNames4, allNames3)
	fmt.Println("someNames4:", someNames4) //[Lifejacket Paddle]
	fmt.Println("allNames3", allNames3)    //[Lifejacket Paddle Hat]

	//delete el from slice [Kayak Lifejacket Paddle Hat]
	deleted := append(products2[:2], products2[3:]...) //delete Paddle
	fmt.Println("Deleted:", deleted)                   //[Kayak Lifejacket Hat]

	//range slice
	products3 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	fmt.Println("products:", products3)
	for index, value := range products3[2:] {
		fmt.Println("Index:", index, "Value:", value)
	}

	//sort slice
	sort.Strings(products3)
	for index, value := range products3 {
		fmt.Println("Index:", index, "Value:", value)
	}
	//Equal (only with "reflect")
	p2 := products3
	fmt.Println("Equal:", reflect.DeepEqual(products3, p2))

	//get array which contains in slice
	arrayPtr := (*[3]string)(p2)
	array := *arrayPtr
	fmt.Println(array) //[Hat Kayak Lifejacket] (only first 3 el)
}
