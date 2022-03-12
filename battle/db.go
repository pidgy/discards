package battle

import (
	"strings"

	"github.com/peterbourgon/diskv/v3"
)

type cache struct {
	*diskv.Diskv
}

type db struct {
	battles cache
	cards   cache
	decks   cache
	sets    cache
}

type encoder interface {
	id() string
	encode() ([]byte, error)
}

var (
	database = db{
		battles: cache{
			Diskv: diskv.New(diskv.Options{
				BasePath:     "battle/cache/battles",
				Transform:    func(s string) []string { return []string{} }, // flat.
				CacheSizeMax: 1024 * 1024,
			}),
		},
		cards: cache{
			Diskv: diskv.New(diskv.Options{
				BasePath:     "battle/cache/cards",
				Transform:    func(s string) []string { return []string{} }, // flat.
				CacheSizeMax: 1024 * 1024,
			}),
		},
		decks: cache{
			Diskv: diskv.New(diskv.Options{
				BasePath:     "battle/cache/decks",
				Transform:    func(s string) []string { return []string{} }, // flat.
				CacheSizeMax: 1024 * 1024,
			}),
		},
		sets: cache{
			Diskv: diskv.New(diskv.Options{
				BasePath:     "battle/cache/sets",
				Transform:    func(s string) []string { return []string{} }, // flat.
				CacheSizeMax: 1024 * 1024,
			}),
		},
	}
)

func (c *cache) read(key string) ([]byte, error) {
	buf, err := c.Read(sanitize(key))
	if err != nil {
		return nil, nil
	}

	return buf, nil
}

func (c *cache) write(e encoder) error {
	buf, err := e.encode()
	if err != nil {
		return err
	}

	return c.Write(sanitize(e.id()), buf)
}

func sanitize(key string) string {
	return strings.ToLower(strings.ReplaceAll(key, " ", "-")) + ".json"
}
