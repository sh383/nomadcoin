package main

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func main() {
	chain := blockchain{}
}
