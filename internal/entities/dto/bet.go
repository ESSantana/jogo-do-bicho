package dto

type Bet struct {
	GamblerID int     `json:"gambler_id"`
	BetType   string  `json:"bet_type"`
	BetPrice  float64 `json:"bet_price"`
	BetChoice string  `json:"bet_choice"`
}
