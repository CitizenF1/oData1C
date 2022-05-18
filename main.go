package main

import (
	"fmt"

	go_odoo "github.com/skilld-labs/go-odoo"
)

func main() {
	odooClientPM, err := go_odoo.NewClient(&go_odoo.ClientConfig{
		Admin:    odooLoginMro,
		Password: odooPasswordMro,
		Database: odooDatabaseMro,
		URL:      odooUrlMro,
	})

	if err != nil {
		fmt.Println("[Odoo connection] error: ", err)
		// return err
	}
	ClientPM = *odooClientPM

	acc, err := ClientPM.FindAccountAccounts(nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range *acc {
		if v.Name.Get() == "Нематериальные активы" {
			c, err := ClientPM.GetAccountAccount(v.Id.Get())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(c.DisplayName, c.UserTypeId, v.Parent)
		}
	}
	// Clietn1cConnect()
}
