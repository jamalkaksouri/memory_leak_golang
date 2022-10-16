package main

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

/*
Typical memory leaks
1. Growing global variable
2. Hanging go routine
3. Open streams
*/

var globalSlice = make([]int64, 0)

// appendSlice: Growing global variable
func appendSlice(c *gin.Context) {
	globalSlice = append(globalSlice, time.Now().Unix())
	c.JSON(200, map[string]int{
		"sliceSize": len(globalSlice),
	})
}

// hangingGoRoutine: Hanging go routine
func hangingGoRoutine(ctx *gin.Context) {
	go time.Sleep(time.Hour * 24)
}

// makeHttpCall: Open streams
func makeHttpCall(ctx *gin.Context) {
	client := &http.Client{}
	file, err := os.ReadFile("LOREM-IPSUM.txt")
	if err != nil {
		return
	}
	bodyReader := bytes.NewReader(file)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", bodyReader)
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(bodyReader)
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//		return
	//	}
	//}(res.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	r := gin.Default()
	pprof.Register(r)

	r.GET("/append-slice", appendSlice)
	r.GET("/hanging", hangingGoRoutine)
	r.GET("/streams", makeHttpCall)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
