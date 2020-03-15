package endpoint

import "bank/model"

type Product struct {
	Category model.Category `json:"category"`
	Price    float64        `json:"price"`
}

type Evaluation struct {
	Evaluation  string `json:"evaluation"`
	Description string `json:"description"`
}

type RiskEvaluate struct {
	Evaluation Evaluation `json:"evaluation"`
	Blog       model.Blog `json:"blog"`
}

type IAResponse struct {
	Response string `json:"response"`
	Status   string `json:"status"`
}
