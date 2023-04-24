package handle

import (
	"app/api/service"
	"app/api/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// @Summary Consulta endereço
// @Description Consulta o endereço a partir do CEP informado
// @Tags Endereço
// @Accept json
// @Produce json
// @Param cep body string true "CEP a ser consultado"
// @Success 200 {object} Address
// @Failure 400 "Requisição inválida"
// @Failure 404 "CEP não encontrado"
// @Failure 500 "Erro ao consultar o ViaCEP ou decodificar a resposta"
// @Router /v1/consulta-endereco [post]
func ConsultAdress(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &utils.RequestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	viaCepData := service.CheckCep(utils.RequestData.Cep, w)

	address := utils.Address{
		Cep:         utils.FormatCEP(viaCepData.Cep),
		Rua:         viaCepData.Logradouro,
		Complemento: viaCepData.Complemento,
		Bairro:      viaCepData.Bairro,
		Cidade:      viaCepData.Localidade,
		Estado:      viaCepData.Uf,
		Frete:       utils.CalculateFreight(viaCepData.Uf),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)
}
