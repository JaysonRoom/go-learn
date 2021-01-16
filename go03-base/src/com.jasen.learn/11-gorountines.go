package main

func main() {
	go sayWhat("B")
	println("AA")
	for i:=0;i<10;i++{
		sayWhat("C")
	}
}

func sayWhat(w string)  {
	println(w)
}