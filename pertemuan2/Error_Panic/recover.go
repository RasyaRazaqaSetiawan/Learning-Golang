package main

// import (
// 	"errors"
// 	"fmt"
// )

// func cover(){
// 	err := recover()
// 	fmt.Println("Error Found", err)
// }

// func divide(nilai, divider int) (int, error) {
// 	if divider == 0 {
// 		return 0, errors.new("divider cannot be zero")
// 	}

// 	return nilai / divider, nil
// }

// func main(){
// 	defer cover()

// 	result, err := divide(10, 0)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("result", result)
// }