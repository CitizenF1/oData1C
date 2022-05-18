package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetRequestOdata() {

}

//account.account
func Clietn1cConnect() {
	url := url1c + "ChartOfAccounts_Хозрасчетный?$format=json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.SetBasicAuth(login1c, pass1c) //axpert22@gmail.com // fex6Etu
	// req.Header.Add("Authorization", "Basic YXhwZXJ0MjJAZ21haWwuY29tOmZleDZFdHU=")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	metadata1c := Metadata{}
	err = json.Unmarshal(body, &metadata1c)
	if err != nil {
		fmt.Println(err)
	}
}

func GetGuid(id string) {
	//ChartOfAccounts_Хозрасчетный(guid'9781b499-cc0d-11e5-9653-3085a93ddca2')
	url := url1c + id + "?$format=json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.SetBasicAuth(login1c, pass1c)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	in := Value{}
	err = json.Unmarshal(body, &in)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(in.QuickPickCode)
}
