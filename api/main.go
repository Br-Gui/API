package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var escolha int

type tokenresponse struct {
	IdToken string `json:"id_token"`
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func execurl() tokenresponse {
	tokenResp, err := gerartokenqa() || gerartokenprod()
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("GET", "", bytes.NewBuffer([]byte(tokenResp.IdToken))) //inserir a url da api
	if err != nil {
		fmt.Println("Erro ao fazer o requerimento", err)
	}
	req.Header.Add("Authorization", "Bearer "+tokenResp.IdToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("erro ao fazer o requerimento:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erro ao fazer o requerimento da URL", resp.StatusCode)
	}
	fmt.Println("Requerimento feito com	sucesso, Status:", resp.StatusCode)
	fmt.Println("Carregando conteudo...")
	time.Sleep(1 * time.Second)
	post, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Aconteceu um erro", resp.StatusCode)
	}

	fmt.Println("O conteudo é:", "\n", string(post))
	return tokenResp
}

func requerimento() {
	fmt.Println("Escolha o ambiente ")
	fmt.Println("Ambiente de QA, Digite 1")
	fmt.Println("Ambiente de PROD, digite 2")
	fmt.Scan(&escolha)
	switch escolha {
	case 1:
		gerartokenqa()
		execurl()

	case 2:
		gerartokenprod()
		execurl()
	}
}

func main() {
	requerimento()
}
