package main

func main() {
	go StartGRPC()
	StartHTTP()
}
