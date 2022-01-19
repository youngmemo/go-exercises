package main

import (
	"fmt"
	"github.com/youngmemo/demo-is105-oppgave1-del2/del2pakke"
)

func main() {
	fmt.Println(del2pakke.SjekkQuote())
	fmt.Println(del2pakke.SjekkTilstand())
	del2pakke.EndreTilstand()
	fmt.Println("Putter inn HS og kylling i båten")
	fmt.Println(del2pakke.SjekkTilstand())

}
