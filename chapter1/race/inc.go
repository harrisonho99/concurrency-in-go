package race

import "fmt"

func Increase() {

	var data int

	go func() {
		data++
	}()
	fmt.Println(data)
}
