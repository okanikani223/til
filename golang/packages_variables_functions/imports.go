package main

/**
 * import "fmt"
 * import "mat/rand"
 * でもOK
 * ただし下記の記載が推奨される様子。
 * (factored import statementと呼ばれる)
 */
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("あなたは%gの問題を抱えています", math.Sqrt(4))
}
