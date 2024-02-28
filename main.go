package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Order struct {
	ResourceType string              `json:"resourceType"`
	Parameter    []map[string]string `json:"parameter"`
}

func main() {
	file, err := os.Open("orders.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		order := scanner.Text()
		cancelJsonOrder(order)
	}

}

func cancelJsonOrder(valueString string) {
	serviceURL := "http://10.128.66.207:2226/exlab/api/fhir/$cancelorder?_format=json"
	param := []map[string]string{}
	mapKey := make(map[string]string)
	name := "OrderId"
	mapKey["name"] = name
	mapKey["valueString"] = valueString
	param = append(param, mapKey)
	j := Order{
		ResourceType: "Parameters",
		Parameter:    param,
	}
	b, err := json.Marshal(j)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "") // token N3
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}
