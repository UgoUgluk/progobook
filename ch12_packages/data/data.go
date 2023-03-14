package data

import "fmt"

func init() {
	fmt.Println(("data.go init function invoked"))
}

//GetData return some data
func GetData() []string {
	return []string{"Kayak", "Lifejacket", "Paddle", "Soccer Ball"}
}
