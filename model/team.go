package structs

type Player struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
	Age  string `json:"age"`
}

type Team struct {
	Players []Player `json:"players"`
	Name    string   `json:"name"`
}

type DataStruct struct {
	Team Team `json:"team"`
}
