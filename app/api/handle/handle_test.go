package handle

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConsultAdress(t *testing.T) {
	reqBody := `{"cep": "01001000"}`
	req, err := http.NewRequest("POST", "/consult-address", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ConsultAdress)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	a := rr.Body.String()
	expected := "{\"cep\":\"01001--000\",\"rua\":\"Praça da Sé\",\"complemento\":\"lado ímpar\",\"bairro\":\"Sé\",\"cidade\":\"São Paulo\",\"estado\":\"SP\",\"frete\":7.85}\n"
	if a != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
