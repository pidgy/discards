package battle

import (
	"bytes"
	"encoding/json"
	"io"
)

type Battle struct {
	Name  string
	Decks []*Deck
}

func (b *Battle) Get(name string) error {
	buf, err := database.battles.read(name)
	if err != nil {
		return err
	}

	return b.decode(bytes.NewBuffer(buf))
}

func (b *Battle) decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(b)
	if err != nil {
		return err
	}

	return nil
}

func (b *Battle) encode() ([]byte, error) {
	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(b)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (b *Battle) id() string {
	return b.Name
}
