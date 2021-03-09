package main

import (
	"fmt"
	"gotest/add"
	"math"
	"time"
	"github.com/jinzhu/configor"
)

type person struct {

	name string
	age int
}

type rectangle struct {
	length, width int
}

func (rect rectangle) area() int {
	return rect.length * rect.width
}

func (rect *rectangle) cir() int {
	return 2 * (rect.length + rect.width)
}

func newPerson(name string) *person {
	p := person{name : name}
	p.age = 33
	return &p

}

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

func plus2(a, b int) int {
	return a + b
}
func plus3(a, b, c int) int {
	return a + b + c
}

func re2() (int, int) {
	return 38, 88
}

func sumf(nums ...int) {
	fmt.Print(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 10
	fmt.Println("i now is:", i)
	return func() int {
		i += 3
		return i
	}
}

func fab(n int) int {
	if n == 0 {
		return 1
	}
	return n * fab(n - 1)
}

func zeroval(i int) {
	i = 0
}

func zeroptr(i *int) {
	*i = 0
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

	ma1 := make(map[string]int)

	ma1["k1"] = 5
	ma1["k2"] = 12

	fmt.Println(ma1, ma1["k1"], ma1["k2"])

	v2 := ma1["k2"]
	fmt.Println("v2 :", v2)
	fmt.Println("length :", len(ma1))

	vv, retur := ma1["k1"]
	fmt.Println("retur :", retur, "vv :", vv)

	n := map[int]string{1: "xx", 2: "oo", 3: "xxoo"}
	fmt.Println(n)

	fmt.Println("############### Range ################")

	nums := []int{3, 5, 8, 13}
	sum := 0
	for index, t := range nums {
		sum += t
		fmt.Println(index, ":", t)
	}
	fmt.Println("sum is:", sum)

	for index, t := range nums {
		if t == 8 {
			fmt.Println("num 8 index is:", index)
		}
	}
	kvs := map[string]string{"a" : "apple", "b" : "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range  kvs {
		fmt.Println("key :", k)
	}
	for i, c := range "go" {
		fmt.Println(i,c)
	}

	fmt.Println("############# func #############")

	var in int = 3
	fmt.Printf("in is %d\n", in)

	res2 := plus2(3, 8)
	fmt.Println("3 + 8 =", res2)
	res3 := plus3(2, 4, 6)
	fmt.Println("2 + 4 + 8 =", res3)

	r1, r2 := re2()
	fmt.Println(r1, r2)
	ree1, _ := re2()
	fmt.Println(ree1)

	fmt.Println("####### variadic functions ######")
	sumf(1,3,6,8)

	num1 := []int{2,6,19,14}
	sumf(num1...)           // call slice argument as funcname(slice...)

	fmt.Println("########### Closures ###########")
	closure := intSeq()

	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())

	closurenew := intSeq()
	fmt.Println(closurenew())

	fmt.Println("############# fab-recursive ###########")
	fmt.Println("fab(7) is:", fab(7))

	fmt.Println("################ pointer ##############")
	ia := 3

	fmt.Println("before func:", &ia, ia)

	zeroval(ia)
	fmt.Println("after zeroval:", &ia, ia)

	zeroptr(&ia)
	fmt.Println("after zeroptr:", &ia, ia)

	fmt.Println("############### struct ##############")

	fmt.Println(person{name : "burgess", age : 66})
	fmt.Println(person{name : "zhu"})
	fmt.Println(&person{age : 38})
	fmt.Println(newPerson("burgess-zhu"))

	ss := person{name : "wangling", age : 26}
	fmt.Println(ss)
	fmt.Print(ss.name, "\n", ss.age, "\n")
	ssp := &ss
	ssp.age = 100
	fmt.Println(ssp)
	fmt.Println(ssp.age)

	fmt.Println("############### methods #############")
	pp := rectangle{length:8, width : 23}
	fmt.Println(pp.area())
	fmt.Println(pp.cir())

	pptr := &pp
	fmt.Println(pptr.area())
	fmt.Println(pptr.cir())


	fmt.Println("############### channel ##############")
	messages := make(chan string)

	go func() {messages <- "go channel"}()
	msg := <- messages
	fmt.Println(msg)

	messages2 := make(chan string, 3)

	messages2 <- "burgess"
	messages2 <- "zhu"
	messages2 <- "!"

	fmt.Println(<- messages2)
	fmt.Println(<- messages2)
	fmt.Println(<- messages2)

	fmt.Println(">>>>>>>>>>>>>>> call func from funcTocall <<<<<<<<<<<<<<<<")
//	call("burgess zhu")
	ccc := add.Add(3.1, 4.4)
	fmt.Println(ccc)
	fmt.Println("执行github的外部函数 ：", configor.Config{})
}

