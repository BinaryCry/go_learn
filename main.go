package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	x uint8 = 255
)


func main() {

	/*var input string
	fmt.Printf("Hello user #%d!\n", rand.Int())
	fmt.Scanf("%s", &input)
	fmt.Println(input)*/


	/*for i:= 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Printf("%d even\n", i)
		} else {
			fmt.Printf("%d odd\n", i)
		}
	}*/

	var seed int64 = time.Now().UnixNano()
	rand.Seed(seed)

	var wDay int = rand.Intn(7)

	switch wDay {
		case 0: fmt.Printf("Mon\n"); break
		case 1: fmt.Printf("Tu\n"); break
		case 2: fmt.Printf("Wh\n"); break
		case 3: fmt.Printf("Th\n"); break
		case 4: fmt.Printf("Fr\n"); break
		case 5: fmt.Printf("St\n"); break
		case 6: fmt.Printf("Sd\n"); break
	}

	loc, err := time.LoadLocation("Europe/Rome")

	fmt.Printf("%d\n%d\n", time.Now().Unix(), time.Now().UnixNano())

	tomorrow:= time.Date(2017, 11, 26, 0, 0, 0, 0, time.Local)
	fmt.Println(time.Until(tomorrow))

	//customTimeValue:= time.Unix( time.Now().Unix() + 4000, 0 )
	//fmt.Println(time.Until(customTimeValue))

	fmt.Println(time.Unix(1511641749, 808104745))


	//dUTC:= time.Date(2017, 11, 25, 23, 30, 10, 0, time.UTC)
	//dLOcal:= time.Date(2017, 11, 25, 23, 30, 10, 0, time.Local)

	if err != nil {
		panic(err)
	}

	fmt.Println(time.Now().In(loc).String())

	fmt.Println(tomorrow.Unix())


	fmt.Println(time.Local)

}


