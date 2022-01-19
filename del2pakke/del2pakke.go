package del2pakke

import (
	"rsc.io/quote"
)

func SjekkQuote() string {
	return quote.Glass() + "\n" + quote.Go()
}

var tilstand = "[kylling rev korn HS ---\\ \\__/ ______________/---]"

func SjekkTilstand() string {
	return tilstand
}

func EndreTilstand() {
	tilstand = "[ rev korn  ---\\___________________ \\__ kylling HS__/ ____/---]"
}
