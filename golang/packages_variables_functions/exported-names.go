package main

import (
	"fmt"
	"math"
)

func main() {
	// math.Piのように大文字で始まる名前は外部から参照可能
	// 以下のように書くとエラー
	// fmt.Println(math.pi)
	fmt.Println(math.Pi)
}
