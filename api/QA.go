package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func gerartokenqa() (tokenresponse, error) {
	requestBody, err := json.Marshal(map[string]string{
		"grant_type": "client_credencials",
	})
	if err != nil {
		return tokenresponse{}, fmt.Errorf("erro ao gerar JSON de requisição: %w", err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "", bytes.NewBuffer(requestBody)) //url do token
	req.Header.Add("Authorization", "Basic "+basicAuth("", ""))           //req que manda o user e senha pra api do  token
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição pra api:", resp.StatusCode)
	}
	defer resp.Body.Close()

	fmt.Print("Enviando Token....")
	var tokenResp tokenresponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return tokenresponse{}, fmt.Errorf("Erro ao decodificar a resposta JSON: %w", err)
	}
	return tokenResp, nil
}
