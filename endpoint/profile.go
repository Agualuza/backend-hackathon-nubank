package endpoint

type Profile struct {
	Id   int `json:"id"`
	Payment float64 `json:"payment"`
}

type RiskEvaluate struct {
	Rate float64 `json:"rate"`
	Profile Profile `json:"user_profile"`
}