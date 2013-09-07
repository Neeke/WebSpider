package main

import (
	"fmt"
	"github.com/Neeke/WebSpider/Spider"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("start...")

	spider := spider.NewWebSpider()
	spider.Fetch("http://image.baidu.com/channel/index", "testpath")

	fmt.Printf("time cost %v\n", time.Now().Sub(start))
}
