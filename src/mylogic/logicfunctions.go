package mylogic
func packageShow(){
	fmt.Printf(stringutils.Reverse("\nhello, world\n"))
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("PI:",math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("Dummy", "practices")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(split(19))
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
func funcRefShow(){
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	curNum := 0
	prevNum := 0
	return func () int{
		fibNum := curNum + prevNum
		prevNum = curNum
		if fibNum ==0 {
			curNum = 1
		}else {
			curNum = fibNum
		}
		return fibNum
	}
}