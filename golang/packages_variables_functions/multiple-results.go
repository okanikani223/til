package main

import "fmt"

/**
 * 関数は複数の戻り値を返す事が可能。
 */
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("こんにちわ", "世界")
	fmt.Println(a, b)
}
