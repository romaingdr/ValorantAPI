package backend

type Element struct {
	Nom      string `json:"nom"`
	ImageURL string `json:"imageUrl"`
}

type ElementMap struct {
	Nom string `json:"nom"`
	Id  string `json:"id"`
}
type Data struct {
	Agents []Element    `json:"agent"`
	Maps   []ElementMap `json:"map"`
}
