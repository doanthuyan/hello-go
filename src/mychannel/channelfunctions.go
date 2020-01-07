package mychannel

import (
	"fmt"
	"time"
	"golang.org/x/tour/tree"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func ChannelShow() {
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	 
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	newChannel := make(chan int, 10)
	go fibonacci(cap(newChannel), newChannel)
	for i := range newChannel {
		fmt.Println(i)
	}

	newChannel2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-newChannel2)
		}
		quit <- 0
	}()
	newFibonacci(newChannel2, quit)

	
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func newFibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
func walk(t * tree.Tree,ch *chan int){
	if t != nil{	
		walk (t.Left,ch)
		//fmt.Println(t.Value)
		*ch <- t.Value
		walk (t.Right,ch)
	}else{
		
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	channel1 := make (chan int)
	channel2 := make (chan int)

	go func (){
		
		walk(t1,&channel1)
		close(channel1)
	}()
	go func (){
		
		walk(t2,&channel2)
		close(channel2)
	}()
	var n1,n2 int
	var ok1,ok2 bool = true,true
	var isSame bool = true
	fmt.Println("Tree 1", "  " , "Tree 2")
	for ;ok1 && ok2 ;{
		n1,ok1 = <- channel1
		n2,ok2 = <- channel2
		fmt.Println ( n1,"       " , n2)
		
		isSame = isSame && (n1 == n2)
	}
	
	return isSame
}
func TreeShow() {
	
	tree1 := tree.New (1)
	tree2 := tree.New (1)
	tree3 := tree.New (2)
	fmt.Println(Same(tree1,tree2))
	fmt.Println(Same(tree1,tree3))
}