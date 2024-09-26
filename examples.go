package music_lib

import (
	"fmt"
	"strings"
)

func main() {
	//m, err := migrate.New(
	//	"file://internal/db/migrations",
	//	"postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//m.Up()
	text := "But it's just somethin that we have no control over and that's what destiny is\nBut no more worries, rest your head and go to sleep\nMaybe one day we'll wake up and this'll all just be a dream\n\n[Chorus]\nNow hush little baby, don't you cry\nEverything's gonna be alright\nStiffen that upper lip up little lady, I told ya\nDaddy's here to hold, ya through the night\nI know mommy's not here right now and we don't know why\nWe feel how we feel inside\nIt may seem a little crazy, pretty baby\nBut I promise, momma's gon' be alright"
	cuplets := strings.Split(text, "\n\n")
	fmt.Println(cuplets[0])
}
