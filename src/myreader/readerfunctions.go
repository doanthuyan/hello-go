package myreader
import(
	"fmt"
	"io"
	"strings"
	"golang.org/x/tour/reader"
	"unicode"
)
type MyReader struct{}
type Rot13Reader struct {
	R io.Reader
}
func ReaderShow() {
	r := strings.NewReader("Hello golang!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	fmt.Println()
	reader.Validate(MyReader{})
	
	r13Reader := Rot13Reader{strings.NewReader("HELLO   hello   ")}
	for {
		n, err := r13Reader.Read(b)
		
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
func (r MyReader) Read(b []byte) (int, error){
	var i int =0
	for ; i< len(b) ;i++{
		b[i] = 'A'
	}
	return i, nil
}
func (rot13 Rot13Reader) Read(b []byte) (int, error){
	n, err := rot13.R.Read(b)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	fmt.Printf("n = %v err = %v b = %q\n\n", n, err, b[:n])
	if err == io.EOF {
		
	}else{
		for i:= 0; i < n;i++{
			if unicode.IsLetter(rune(b[i])){
				if b[i] < 'a'{
					b[i] += 13
					if b[i] > 'Z'{
						b[i] = b[i] -'Z'+'@'
					}
				}else{
					b[i] += 13
					if b[i] > 'z'{
						b[i] = b[i] -'z'+'`'
					}
				}
			}
		}
	}
	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	fmt.Printf("n = %v err = %v b = %q\n\n", n, err, b[:n])
	return n, err
}
