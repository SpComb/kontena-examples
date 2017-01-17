package main

const LISTEN = ":8000"
const STATIC = "static"

func main() {
	udpStart(LISTEN)
	httpMain(LISTEN, STATIC)
}
