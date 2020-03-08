package endpoint

import "bank/model"

type Profile struct {
	Id      int     `json:"id"`
	Payment float64 `json:"payment"`
	Factor  string  `json:"factor"`
}

type Product struct {
	Category model.Category `json:"category"`
	Price float64 `json:"price"`
}

type RiskEvaluate struct {
	Rate    float64 `json:"rate"`
	Profile Profile `json:"user_profile"`
}
