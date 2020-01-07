package mymap

type GPS struct {
	Lat, Long float64
}
func WordCount(s string) map[string]int {
	var wordMap map[string] int
	wordMap = make (map[string] int)
	keys := strings.Fields(s)
	for _,k := range keys{
		wordMap[k] = wordMap[k] + 1
	}
	
	return wordMap
}
func mapShow(){
	var m map[string]GPS
	m = make(map[string]GPS)
	m["Bell Labs"] = GPS{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	var gpsList = map[string]GPS{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
		"AAVN": {
			10.8090057, 106.6643066,
		},
	}
	v, ok :=gpsList["AADN"]
	fmt.Println("AADN",v,"Present?",ok)
	v, ok =gpsList["AAVN"]
	fmt.Println("AAVN",v,"Present?",ok)
}

func timeShow(){
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
