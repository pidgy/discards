package api

type Card struct {
	Data `json:"data"`
}

type Data struct {
	ID                     string
	Name                   string
	Supertype              string
	Subtypes               []string
	HP                     string
	Types                  []string
	EvolvesFrom            string
	Rules                  []string
	Attacks                []Attack
	Weaknesses             []Weakness
	RetreatCost            []string
	ConvertedRetreatCost   int
	Set                    Set
	Number                 string
	Artist                 string
	Rarity                 string
	NationalPokedexNumbers []int
	Legalities             map[string]string
	RegulationMark         string
	Images                 map[string]string
	CardMarket             CardMarket
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

type CardMarket struct {
	URL       string
	UpdatedAt string
	Prices    map[string]float64
}
