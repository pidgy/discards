package battle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/pidgy/discards/options"
)

type Card struct {
	Data `json:"data"`
}

type Attack struct {
	Name                string
	Cost                []string
	ConvertedEnergyCost int
	Damage              string
	Text                string
}

type Weakness struct {
	Type  string
	Value string
}

type Set struct {
	ID           string
	Name         string
	Series       string
	PrintedTotal int
	Total        int
	Legalities   map[string]string
	PtcgoCode    string
	ReleaseDate  string
	UpdatedAt    string
	Images       map[string]string
}

type Data struct {
	ID                   string
	Name                 string
	Supertype            string
	Subtypes             []string
	HP                   string
	Types                []string
	EvolvesFrom          string
	Rules                []string
	Attacks              []Attack
	Weaknesses           []Weakness
	RetreatCost          []string
	ConvertedRetreatCost int
	Set
	Number                 string
	Artist                 string
	Rarity                 string
	NationalPokedexNumbers []int
	Legalities             map[string]string
	RegulationMark         string
	Images                 map[string]string
}

const (
	uriCards = "https://api.pokemontcg.io/v2/cards/"
)

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
	r := rand.Intn(len(s.Data))

	prefix, postfix := "", ""

	for _, set := range s.Data {
		if i == r {
			prefix = set.ID
			postfix = strconv.Itoa(rand.Intn(set.Total))
		}
		i++
	}

	return c.Get(prefix + "-" + postfix)
}
