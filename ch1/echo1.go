// Echo1 exibe seus argumentos de linha de comando.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + strconv.Itoa(i) + " - " + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
}
