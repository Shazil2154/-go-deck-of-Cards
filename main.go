package main

func main() {
	cards := newDeckFromFile("MyCards")
	cards.shuffle()
	cards.print()
}
