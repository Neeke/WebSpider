package main

import (
	"fmt"
	"github.com/Neeke/WebSpider/Spider"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("start...")

	lea := spider.NewWebSpider()
	lea.Fetch("http://image.baidu.com/channel/index", "testpath")

	fmt.Printf("time cost %v\n", time.Now().Sub(start))
}
