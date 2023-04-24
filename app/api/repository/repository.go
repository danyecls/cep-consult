package repository

import "app/api/utils"

func CacheCep(cep string) (utils.ViaCepResponse, bool) {
	// consulta cep no banco
	// se tiver cadastrado retorna response
	// se nao tiver retorna false

	t := utils.ViaCepResponse{}

	return t, true
}
