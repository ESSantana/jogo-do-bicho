package entities

type BetGroup struct {
	ID          int64  `db:"id" json:"id"`
	GroupName   string `db:"group_name" json:"group_name"`
	GroupNumber string `db:"group_number" json:"group_number"`
}
