package battle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"

	"github.com/pidgy/discards/battle/api"
	"github.com/pidgy/discards/options"

	crypto "crypto/rand"
	math "math/rand"
)

const (
	uriCards = "https://api.pokemontcg.io/v2/cards/"
)

type Card api.Card

func (c *Card) Get(id string) error {
	if id == "" {
		return c.random()
	}

	buf, err := database.cards.read(id)
	if err != nil {
		return err
	}

	if buf != nil {
		return c.decode(bytes.NewBuffer(buf))
	}

	req, err := http.NewRequest(http.MethodGet, uriCards+id, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Api-Key", options.APIKey)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with code: %s", res.Status)
	}

	err = c.decode(res.Body)
	if err != nil {
		return err
	}

	return database.cards.write(c)
}

func (c *Card) id() string {
	return c.ID
}

func (c *Card) decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Card) encode() ([]byte, error) {
	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(c)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *Card) random() error {
	s := &Sets{}
	err := s.Get()
	if err != nil {
		return err
	}

	i := 0

	r := rand(len(s.Data))

	prefix, postfix := "", ""

	for _, set := range s.Data {
		if i == r {
			prefix = set.ID
			postfix = strconv.Itoa(rand(set.Total))
		}
		i++
	}

	return c.Get(prefix + "-" + postfix)
}

func rand(max int) int {
	b, err := crypto.Int(crypto.Reader, big.NewInt(int64(max)))
	if err != nil {
		return math.Intn(max)
	}

	return int(b.Int64())
}
