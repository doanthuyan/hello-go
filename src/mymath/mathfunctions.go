package mymath
import(
	
	"fmt"
	"math"
	"math/cmplx"
)
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	fmt.Printf("cannot Sqrt negative number: %f\n", float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %f\n", float64(e))
}

func MathShow(){
	var ToBe   bool       = false
	var MaxInt uint64     = 1<<64 - 1
	var z      complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	var x, y int = 3, 4
	f = math.Sqrt(float64(x*x + y*y))
	var k uint = uint(f)
	fmt.Println(x, y, k)
}
func Pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Add(x, y int) int {
	return x + y
}
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func DeferShow(){
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
func Sqrt(x float64) (float64, error) {
	fmt.Println("Calculating square root of:",x)
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
func SqrtShow(){
	var x  float64
	for x = -4; x < 10; x++{
		n,e := Sqrt(x)
		if e == nil{
			fmt.Println("Square root of",x,":",n)
		}else{
			fmt.Printf( "Err: %v %T\n",e,e)
		}
	}
}