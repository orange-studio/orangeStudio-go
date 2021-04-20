package main

import (
	"fmt"
)

func main() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "stockcode=%d&enddate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
}
