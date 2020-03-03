package endpoint

import "github.com/cheekybits/genny/generic"

type jsonReponse struct {
	Status   string `json:"status"`
	Response []generic.Type `json:"response"`
	Message string `json:"message"`
}

