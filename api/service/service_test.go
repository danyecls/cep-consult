package service

import (
	"net/http/httptest"
	"testing"
)

func TestCheckCep(t *testing.T) {
	rw := httptest.NewRecorder()
	result := CheckCep("01310-100", rw)

	if result.Cep != "01310-100" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Cep, "01310-100")
	}
	if result.Logradouro != "Avenida Paulista" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Logradouro, "Avenida Paulista")
	}
	if result.Localidade != "São Paulo" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Localidade, "São Paulo")
	}
	if result.Uf != "SP" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Uf, "SP")
	}

	if result.Ibge != "3550308" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Ibge, "3550308")
	}

	if result.Gia != "1004" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Gia, "1004")
	}

	if result.Ddd != "11" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Ddd, "11")
	}

	if result.Siafi != "7107" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Siafi, "7107")
	}

	result = CheckCep("12345678", rw)

	if result.Cep != "" {
		t.Errorf("CheckCep(\"12345678\") = %+v; want %+v", result.Cep, "")
	}
	if result.Logradouro != "" {
		t.Errorf("CheckCep(\"12345678\") = %+v; want %+v", result.Logradouro, "")
	}
	if result.Localidade != "" {
		t.Errorf("CheckCep(\"12345678\") = %+v; want %+v", result.Localidade, "")
	}
	if result.Uf != "" {
		t.Errorf("CheckCep(\"12345678\") = %+v; want %+v", result.Uf, "")
	}

	if result.Ibge != "" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Ibge, "")
	}

	if result.Gia != "" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Gia, "")
	}

	if result.Ddd != "" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Ddd, "")
	}

	if result.Siafi != "" {
		t.Errorf("CheckCep(\"01310-100\") = %+v; want %+v", result.Siafi, "")
	}
}
