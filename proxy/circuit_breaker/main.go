package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
)


func main()  {
	hystrix.Go("my_command", func() error {
		// talk to other services
		return nil
	}, nil)
	fmt.Println("hello world")
}
