package utils

import "fmt"

type RequestData struct {
	Cep string `json:"cep" validate:"required,cep"`
}

func FormatCEP(cep string) string {
	return fmt.Sprintf("%s-%s", cep[:5], cep[5:])
}

func CalculateFreight(uf string) float32 {
	switch uf {
	case "SP", "RJ", "ES", "MG":
		return 7.85
	case "MT", "MS", "GO", "DF":
		return 12.50
	case "MA", "PI", "CE", "RN", "PE", "PB", "SE", "AL", "BA":
		return 15.98
	case "RS", "SC", "PR":
		return 17.30
	case "AM", "RR", "AP", "PA", "TO", "RO", "AC":
		return 20.83
	default:
		return 0.0
	}
}

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type Address struct {
	Cep         string  `json:"cep"`
	Rua         string  `json:"rua"`
	Complemento string  `json:"complemento,omitempty"`
	Bairro      string  `json:"bairro"`
	Cidade      string  `json:"cidade"`
	Estado      string  `json:"estado"`
	Frete       float32 `json:"frete"`
}
