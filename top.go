package main

import (
	"fmt"
	"math/rand"
	"strings"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	m := BuildDB()

	word := randomWord(m)
	a := app.New()
	w := a.NewWindow("あ 	い 	う 	え 	お")

	key := widget.NewLabel(word.key)
	trans := widget.NewLabel(word.jp)
	def := widget.NewLabel(strings.Join(word.def, " "))
	search := widget.NewEntry()

	w.SetContent(widget.NewVBox(
		key, trans, def,
		widget.NewButton("Random Word", func() {
			word = randomWord(m)
			key.SetText(word.key)
			trans.SetText(word.jp)
			def.SetText(strings.Join(word.def, ";"))
		}), search,
		widget.NewButton("Search", func() {
			temp, wurd := lookUpWord(search.Text, m)
			if temp > 0 {
				m = BuildDB()
				fmt.Println("rebuillding hash table...")
			} else {
				key.SetText(wurd.key)
				trans.SetText(wurd.jp)
				def.SetText(wurd.def[0])
			}

		}),
	))

	w.ShowAndRun()

}

func randomWord(m map[string]*word) *word {
	k := rand.Intn(len(m))
	for key := range m {
		if k == 0 {
			return m[key]
		}
		k--
	}
	return nil
}

func printWord(w *word) {
	fmt.Printf("%s %s %s\n", w.key, w.jp, w.def)
}

func lookUpWord(s string, m map[string]*word) (int, *word) {
	out, _ := m[s]
	err := 0
	if out == nil {
		fmt.Println("I don't know that word yet, just a moment...")
		learnWord(s)
		err++
	} else {
		printWord(out)
	}
	return err, out
}
