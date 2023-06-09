package module1

import "fmt"

type Factory interface {
  Type()
}

func Foo() {
  fmt.Println("Foo")
}