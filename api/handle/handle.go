package handle

import (
	"cep-consult/api/service"
	"cep-consult/api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Consulta endereço
// @Description Consulta o endereço a partir do CEP informado
// @Tags Endereço
// @Accept json
// @Produce json
// @Param cep body string true "CEP a ser consultado"
// @Success 200 {object} utils.Address
// @Failure 400 "Requisição inválida"
// @Failure 404 "CEP não encontrado"
// @Failure 500 "Erro ao consultar o ViaCEP ou decodificar a resposta"
// @Router /v1/consult-address [post]
func ConsultAddress(c echo.Context) error {
	requestData := new(utils.RequestData)

	if err := c.Bind(requestData); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	viaCepData := service.CheckCep(requestData.Cep, c.Response())

	if viaCepData.Cep == "" {
		return c.String(http.StatusNotFound, "CEP not found")
	}

	address := &utils.Address{
		Cep:         utils.FormatCEP(viaCepData.Cep),
		Rua:         viaCepData.Logradouro,
		Complemento: viaCepData.Complemento,
		Bairro:      viaCepData.Bairro,
		Cidade:      viaCepData.Localidade,
		Estado:      viaCepData.Uf,
		Frete:       utils.CalculateFreight(viaCepData.Uf),
	}

	return c.JSON(http.StatusOK, address)
}
