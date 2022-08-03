package db

import "time"

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json :"owner"`
	Balance   int64     `json: "balance"`
	Currency  string    `json: "currency"`
	CreatedAt time.Time `json: "created_at"`
}

type Entry struct {
	ID        int64     `json:"id"`
	AccountId int64     `json:"accountid"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID            int64     `json:"id"`
	FromAccountID int64     `json:"accountid"`
	ToAccountID   int64     `json:"accountid"`
	Amount        int       `json:" amount"`
	CreatedAt     time.Time `json:"created_at"`
}
