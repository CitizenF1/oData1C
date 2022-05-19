package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetRequestOdata() {}

//account.account
func Clietn1cConnect(account string) string {
	var id string
	url := url1c + "ChartOfAccounts_Хозрасчетный?$format=json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth(login1c, pass1c) //axpert22@gmail.com // fex6Etu
	// req.Header.Add("Authorization", "Basic YXhwZXJ0MjJAZ21haWwuY29tOmZleDZFdHU=")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	metadata1c := &Metadata{}
	err = json.Unmarshal(body, &metadata1c)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range metadata1c.Values {
		if account == v.Description {
			id = v.RefKey
		}
	}
	return id
}

func GetGuid(id string, entities string) Value {
	//ChartOfAccounts_Хозрасчетный(guid'9781b499-cc0d-11e5-9653-3085a93ddca2')
	url := url1c + entities + "(guid'" + id + "'?$format=json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth(login1c, pass1c)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	in := Value{}
	err = json.Unmarshal(body, &in)
	if err != nil {
		fmt.Println(err)
	}
	return in
}
