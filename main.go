package main

import (
	"fmt"
	"time"
)

func main() {
	var LOC, _ = time.LoadLocation("Asia/Ho_Chi_Minh")
	fmt.Println(time.Now().In(LOC).UnixMilli())
}
