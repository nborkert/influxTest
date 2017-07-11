package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	//"time"
)

//We will query random.org for data and store in a local InfluxDB database

func main() {
	max, _ := strconv.Atoi(os.Args[1])
	targetPage := os.Args[2]


	responseCode := ""
	for count := 0; count != max; count++ {
		responseCode = GetHTTPCode(targetPage)
		fmt.Printf("%v\n", responseCode)
	}
}

func GetHTTPCode(u string) string {
	//t := time.Now()
	res, err := http.Get(u)
	//t2 := time.Now()

	/*if err != nil {
		fmt.Printf(err.Error()+",%v\n", t2.Sub(t))
		return ""
	} */
	if err == nil {
		defer res.Body.Close()
		//fmt.Printf(res.Status+",%v\n", t2.Sub(t))
		return res.Status
	} else {
		return ""
	}
}
