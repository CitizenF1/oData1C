package main

import (
	"fmt"

	go_odoo "github.com/skilld-labs/go-odoo"
)

func main() {
	odooClientPM, err := go_odoo.NewClient(&go_odoo.ClientConfig{
		Admin:    odooLoginPM,
		Password: odooPasswordPM,
		Database: odooDatabasePM,
		URL:      odooUrlPM,
	})

	if err != nil {
		fmt.Println("[Odoo connection] error: ", err)
		// return err
	}
	ClientPM = *odooClientPM

	accounts, err := ClientPM.FindAccountAccounts(nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	for _, account := range *accounts {
		if account.Name.Get() == "Нематериальные активы" {
			acc, err := ClientPM.GetAccountAccount(account.Id.Get())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(acc.DisplayName, acc.UserTypeId, account.Parent)
		}
	}
	// Clietn1cConnect()
}
