package api

type Sets struct {
	Data []SetData `json:"data"`
}

type SetData struct {
	ID    string
	Name  string
	Total int
}
