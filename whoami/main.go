package main

const LISTEN = ":8000"

func main() {
	udpStart(LISTEN)
	httpMain(LISTEN)
}
