package model

// ThankYou represents a thank you note
type ThankYou struct {
	From string `json:"from" binding:"required"`
	To   string `json:"to" binding:"required"`
	For  string `json:"for" binding:"required"`
}
