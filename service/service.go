package service

import (
	"app/api/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CheckCep(cep string, w http.ResponseWriter) utils.ViaCepResponse {
	responseCep := utils.ViaCepResponse{}

	//responseCache, ok := repository.CacheCep(cep)

	// if ok == true {
	// 	responseCep = responseCache
	// } else {
	responseCep = getViaCep(cep, w)
	//}

	return responseCep
}

func getViaCep(cep string, w http.ResponseWriter) utils.ViaCepResponse {
	cepResponse := utils.ViaCepResponse{}

	viaCepURL := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	viaCepResp, err := http.Get(viaCepURL)
	if err != nil {
		http.Error(w, "Erro ao consultar o ViaCEP", http.StatusInternalServerError)
		return cepResponse
	}
	defer viaCepResp.Body.Close()

	if viaCepResp.StatusCode != http.StatusOK {
		http.Error(w, "CEP não encontrado", http.StatusNotFound)
		return cepResponse
	}

	var viaCepData utils.ViaCepResponse
	err = json.NewDecoder(viaCepResp.Body).Decode(&viaCepData)
	if err != nil {
		http.Error(w, "Erro ao decodificar a resposta do ViaCEP", http.StatusInternalServerError)
		return cepResponse
	}

	return viaCepData
}
