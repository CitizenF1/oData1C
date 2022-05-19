package main

import (
	"fmt"

	go_odoo "github.com/skilld-labs/go-odoo"
)

// 9781b37f-cc0d-11e5-9653-3085a93ddca2
// (guid'8cfb288f-cc0d-11e5-9653-3085a93ddca2')
// ChartOfCharacteristicTypes_ВидыСубконтоХозрасчетные(guid'8cfb288f-cc0d-11e5-9653-3085a93ddca2')?$format=json
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
	updates := make(map[string]interface{})
	lines := &go_odoo.AccountAccountSubcontoLine{}
	for _, account := range *accounts {
		if account.Name.Get() == "Нематериальные активы" {
			acc, err := ClientPM.GetAccountAccount(account.Id.Get())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(acc.Id)
			id := Clietn1cConnect("Нематериальные активы")
			in := GetGuid(id, "ChartOfAccounts_Хозрасчетный")
			for _, v := range in.ExtDimensionTypes {
				sub := GetGuid(v.ExtDimensionTypeKey, "ChartOfCharacteristicTypes_ВидыСубконтоХозрасчетные")

				fmt.Println(sub.Description, "DESCRIPTION")
				lines.AccountId = go_odoo.NewMany2One(acc.Id.Get(), "A")
				lines.Name = go_odoo.NewString(sub.Description)
				lines.Sum = go_odoo.NewBool(v.Sum)
				id, err := ClientPM.CreateAccountAccountSubcontoLine(lines)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(id, "ID SUBLINES")
			}
			// fmt.Println(sub.Description)
			updates["tax"] = in.TaxAccounting
			err = ClientPM.Update("account.account", []int64{acc.Id.Get()}, updates)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}
