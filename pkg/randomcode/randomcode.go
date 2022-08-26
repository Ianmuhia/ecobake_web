package randomcode

import (
	"log"
	"strconv"
	"time"
)

func Code(length int) string {

	max := time.Now().UnixNano()
	d := strconv.FormatInt(max, 10)
	log.Println(d)
	return d[len(d)-length:]
}

//number := 111555
//slice := strconv.Itoa(number)
//fmt.Println(slice[:3]) // first 3 digits (111)
//fmt.Println(slice[len(slice)-2:]) // and the last 2 digits (55)
