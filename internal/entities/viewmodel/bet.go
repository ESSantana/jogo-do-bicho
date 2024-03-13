package viewmodel

type Bet struct {
	ID        int32     `json:"id,omitempty"`
	Gambler   Gambler `json:"gambler,omitempty"`
	BetType   string  `json:"bet_type,omitempty"`
	BetPrice  float64 `json:"bet_price,omitempty"`
	BetChoice string  `json:"bet_choice,omitempty"`
}
