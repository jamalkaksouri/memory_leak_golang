package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
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
func appendSlice(c *gin.Context) error {
	globalSlice = append(globalSlice, time.Now().Unix())
	c.JSON(200, map[string]int{
		"sliceSize": len(globalSlice),
	})
	return nil
}

// hangingGoRoutine: Hanging go routine
//func hangingGoRoutine() {
//	go time.Sleep(time.Hour * 24)
//}

// makeHttpCall: Open streams
//func makeHttpCall() {
//	client := &http.Client{}
//	file, err := os.ReadFile("LOREM-IPSUM.txt")
//	if err != nil {
//		return
//	}
//	bodyReader := bytes.NewReader(file)
//
//	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", bodyReader)
//	if err != nil {
//		panic(err)
//	}
//	res, err := client.Do(req)
//	if err != nil {
//		panic(err)
//	}
//	body, err := io.ReadAll(bodyReader)
//	//defer func(Body io.ReadCloser) {
//	//	err := Body.Close()
//	//	if err != nil {
//	//		return
//	//	}
//	//}(res.Body)
//
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(body))
//}

func main() {
	r := gin.Default()
	pprof.Register(r)
	r.GET("/", func(ctx *gin.Context) {
		for {
			err := appendSlice(ctx)
			if err != nil {
				return
			}
			//hangingGoRoutine()
			//makeHttpCall()
		}
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("failed to running app")
	}
}
