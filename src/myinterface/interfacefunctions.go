package myinterface

type I interface {
	M()
}
type Abser interface {
	Abs() float64
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type FloatVertex struct {
	X, Y float64
}

func (v FloatVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
func interfaceShow(){
	var ok bool 
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
	
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := FloatVertex{3, 4}
	
	a = f  // a MyFloat implements Abser
	_,ok = a.(MyFloat)
	fmt.Println(a.Abs(), "is MyFloat?", ok)

	a = &v // a *FloatVertex implements Abser
	fmt.Println(a.Abs())
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	_,ok = a.(MyFloat)
	fmt.Println(a.Abs(), "is MyFloat?", ok)
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}