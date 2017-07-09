package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

//We will query random.org for data and store in a local InfluxDB database

func main() {
	max, _ := strconv.Atoi(os.Args[1])
	targetPage := os.Args[2]



	for count := 0; count != max; count++ {
		GetHTTPCode(targetPage, count)
	}
}

func GetHTTPCode(u string, num int) {
	t := time.Now()
	res, err := http.Get(u)
	t2 := time.Now()

	if err != nil {
		fmt.Printf("%v,"+err.Error()+",%v\n", num, t2.Sub(t))
	}
	if err == nil {
		defer res.Body.Close()
		fmt.Printf("%v,"+res.Status+",%v\n", num, t2.Sub(t))
	}
}
