package main

import(
	"fmt"
	"time"
)

var (
	Debug bool
	LogLevel string
	startUpTime time.Time
)

func main()  {
	var(
		test string
	)
	query, limit, offset := "bat", 10, 0
	query, limit, offset = "ball", offset, 20
	test = "hi"
	fmt.Println(query, limit, offset,test)
	
	Debug, LogLevel, startUpTime := getTestValues()
	fmt.Println(Debug, LogLevel, startUpTime)
}

func getTestValues()(bool, string, time.Time)  {
	return true, "example", time.Now()
}