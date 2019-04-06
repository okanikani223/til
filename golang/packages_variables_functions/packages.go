/**
 * プログラムはmainパッケージから実行される
 */
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("私の好きな数は", rand.Intn(10))
}
