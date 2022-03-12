package battle

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type Deck struct {
	Name  string
	Cards []*Card
}

type file struct {
	Name  string   `json:"name"`
	Cards []string `json:"cards"`
}

func (d *Deck) Get(name string) error {
	cached, err := database.decks.read(name)
	if err != nil {
		return err
	}

	if cached != nil {
		return d.decode(bytes.NewBuffer(cached))
	}

	f, err := os.Open("battle/decks/" + name + ".json")
	if err != nil {
		return err
	}

	file := &file{}
	err = json.NewDecoder(f).Decode(file)
	if err != nil {
		return err
	}

	for _, id := range file.Cards {
		c := &Card{}
		err := c.Get(id)
		if err != nil {
			return err
		}

		d.Cards = append(d.Cards, c)
	}

	d.Name = file.Name

	return database.decks.write(d)
}

func (d *Deck) decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(d)
	if err != nil {
		return err
	}

	return nil
}

func (d *Deck) encode() ([]byte, error) {
	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(d)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (d *Deck) id() string {
	return d.Name
}
