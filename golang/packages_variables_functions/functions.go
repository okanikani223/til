package main

import "fmt"

/**
 * 関数は0個以上の引数を取れる。
 * 型の指定は変数名の後に行う。
 */
func add(x int, y int) int {
	return x + y
}

/**
 * 引数の型が共通である場合には、下記のように最後に型指定するだけでも良い。
func add(x, y int) int{
	return x + y
}
*/

func main() {
	fmt.Println(add(42, 13))
}
