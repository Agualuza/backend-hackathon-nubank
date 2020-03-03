package endpoint

type Profile struct {
	Id   int
	Payment float64
}

type RiskEvaluate struct {
	Rate float64 `json:"rate"`
	Profile Profile `json:"user_profile"`
}