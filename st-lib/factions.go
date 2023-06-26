package stlib

type Faction struct {
	Symbol       string
	Name         string
	Description  string
	Headquarters string
	Traits       []Trait
	IsRecruiting bool `json:"isRecruiting"`
}

type Trait struct {
	Symbol      string
	Name        string
	Description string
}

type FactionShort struct {
	Symbol string
}

type FactionData struct {
	Faction Faction `json:"data"`
}
