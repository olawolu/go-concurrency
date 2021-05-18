package main

func doNothing() string {
	return "nothing"
}

func main() {
	var channel chan string = make(chan string)
	channel <- doNothing()
}
