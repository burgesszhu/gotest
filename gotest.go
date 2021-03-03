package main

import (
	"fmt"
	"math"
	"time"
)

func isStatus(status string) string {
	switch status {
	case "0" :
		return "not start"
	case "1" :
		return "on going"
	case "2" :
		return "finished"
	default:
		return  "false status"
	}

}

func main() {
	mystatus := isStatus("2")
	fmt.Println(mystatus)


	fmt.Println("############### Go by examples ##########")

	fmt.Println("############### 1. values ###############")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println("7/3 =", 7/3)
	fmt.Println("go" + "lang")
	fmt.Println(true || false)

	fmt.Println("################ 2. variables ############")
	var a = "initail"
	fmt.Println(a)
	var b, c int = 2, 9
	fmt.Println(b,c)
	var d = true
	fmt.Println(d)
	e := "xyz"
	fmt.Println(e)

	fmt.Println("################ 3. const #################")
	const aa string = "const a"
	fmt.Println(aa)

	const bb = 600000
	const cc = 3e20 / bb
	fmt.Println(cc)
	fmt.Println(int64(cc))
	fmt.Println("math.Sin(cc) =",math.Sin(cc))

	fmt.Println("################ 4. for  ##################")
	for i := 2; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	time.Sleep(1 *time.Second)
	fmt.Println("second for circle: ")
	for {
		fmt.Println("for end by break")
		break
	}
	fmt.Println(time.Now().Unix())

	fmt.Println("*************** if/else ****************")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 ==0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has at least 2 digits")
	}

	fmt.Println("################# switch ###############")
	ii := 2
	fmt.Println( "write ", ii, " as ")
	switch ii {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println("time.Now().Weekday()", time.Now().Weekday())
	fmt.Println("time.Saturday", time.Saturday)
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend")
	default:
		fmt.Println("It is the weekday")
	}

	t := time.Now()
	switch  {
	case t.Hour() <12 :
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool :
			fmt.Println("I am a bool")
		case int :
			fmt.Println("I am int")
		default:
			fmt.Printf("Donot know the type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("burgess")

	fmt.Println("################ array ###############")
	var ar [5]int
	fmt.Println("emp:",ar)

	ar[4] = 100
	fmt.Println("set:",ar,"get:",ar[4],"ar-len:",len(ar))

	br := [7]string{"burgess","zhu","xx","oo","nf","ll","ei"}
	fmt.Println("dcl:",br,br[3])

	var twoD[3][5] int
	for xx := 0; xx < 3; xx++ {
		for yy := 0; yy < 5; yy++ {
			twoD[xx][yy] = xx * yy
		}
	}
	fmt.Println("twoD is:", twoD)

	for o := 0; o <3; o++ {
		fmt.Println(twoD[o][o+1])
	}

	fmt.Println("################ slice ##################")

	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0] = "a b"
	s[1] = "c "
	s[2] = "de!"

	fmt.Println("set :", s)
	fmt.Println("get s[2] :", s[2])
	fmt.Println("len :", len(s))

	s = append(s, "4e")
	s = append(s,"5f", "6g")
	fmt.Println("after_append:", s, len(s))

	l1 := s[2:5]
	fmt.Println("l1 := s[2:5],l1 =", l1)

	sc := make([]string, len(s))
	copy(sc, s)
	fmt.Println("slice sc is:", sc, "slice s is:", s)

	slice_twoD := make([][]int, 3)
	for r := 0; r < 3; r++ {
		innerlength := r + 5
		slice_twoD[r] = make([]int, innerlength)
		for x := 0; x < innerlength; x++ {
			slice_twoD[r][x] = r + x
		}
	}

	fmt.Println("slice_twoD is:", slice_twoD)

	fmt.Println("############## Maps ###############")



}