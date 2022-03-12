package battle

import "fmt"

func Load() {
	m := &Deck{}

	err := m.Get("mathias")
	if err != nil {
		panic(err)
	}

	v := &Deck{}
	err = v.Get("vanessa")
	if err != nil {
		panic(err)
	}

	b := &Battle{
		Name:  "Paradigm Shift",
		Decks: []*Deck{m, v},
	}

	err = database.battles.write(b)
	if err != nil {
		panic(err)
	}

	s := &Sets{}
	err = s.Get()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", s)
}
