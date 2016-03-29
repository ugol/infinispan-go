package main

import "fmt"
import "github.com/ugol/infinispan-go/infinispan"

func main() {

	fmt.Println("Infinispan GO client")
	infinispan.TestConnect()
}
