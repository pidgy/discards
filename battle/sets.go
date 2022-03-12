package battle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pidgy/discards/options"
)

type Sets struct {
	Data []SetData `json:"data"`
}

type SetData struct {
	ID    string
	Name  string
	Total int
}

const (
	uriSets = "https://api.pokemontcg.io/v2/sets"
)

func (s *Sets) Get() error {
	buf, err := database.sets.read("sets")
	if err != nil {
		return err
	}

	if buf != nil {
		return s.decode(bytes.NewBuffer(buf))
	}

	req, err := http.NewRequest(http.MethodGet, uriSets, nil)
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

	err = s.decode(res.Body)
	if err != nil {
		return err
	}

	return database.sets.write(s)
}

func (s *Sets) id() string {
	return "all"
}

func (s *Sets) decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(s)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sets) encode() ([]byte, error) {
	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(s)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
