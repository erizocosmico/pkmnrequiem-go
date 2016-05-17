package models

type Battle struct {
	ID       string       `json:"id"`
	Finished bool         `json:"finished"`
	Log      []string     `json:"log"`
	Users    []BattleUser `json:"users"`
	Turn     int          `json:"turn"`
}

type BattleMovement struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BattlePokemon struct {
	ID                  string           `json:"id"`
	Attack              int              `json:"attack"`
	Defense             int              `json:"defense"`
	Experience          int              `json:"experience"`
	Health              int              `json:"health"`
	Kind                string           `json:"kind"`
	Level               int              `json:"level"`
	Movements           []BattleMovement `json:"movements"`
	Name                string           `json:"name"`
	Shiny               bool             `json:"shiny"`
	SpecialAttack       int              `json:"special_attack"`
	SpecialDefense      int              `json:"special_defense"`
	Speed               int              `json:"speed"`
	TotalAttack         int              `json:"total_attack"`
	TotalDefense        int              `json:"total_defense"`
	TotalHealth         int              `json:"total_health"`
	TotalSpecialAttack  int              `json:"total_special_attack"`
	TotalSpecialDefense int              `json:"total_special_defense"`
	TotalSpeed          int              `json:"total_speed"`
}

type BattleUser struct {
	ID             string          `json:"id"`
	CurrentPokemon int             `json:"current_pokemon"`
	Party          []BattlePokemon `json:"party"`
	Position       int             `json:"position"`
}
