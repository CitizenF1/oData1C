package odoo

import (
	"fmt"
)

// AccountAccountSubcontoLine represents account.account.subconto.line model.
type AccountAccountSubcontoLine struct {
	AccountId   *Many2One `xmlrpc:"account_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Currency    *Bool     `xmlrpc:"currency,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Onlyrevs    *Bool     `xmlrpc:"onlyrevs,omptempty"`
	Sum         *Bool     `xmlrpc:"sum,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountAccountSubcontoLines represents array of account.account.subconto.line model.
type AccountAccountSubcontoLines []AccountAccountSubcontoLine

// AccountAccountSubcontoLineModel is the odoo model name.
const AccountAccountSubcontoLineModel = "account.account.subconto.line"

// Many2One convert AccountAccountSubcontoLine to *Many2One.
func (aasl *AccountAccountSubcontoLine) Many2One() *Many2One {
	return NewMany2One(aasl.Id.Get(), "")
}

// CreateAccountAccountSubcontoLine creates a new account.account.subconto.line model and returns its id.
func (c *Client) CreateAccountAccountSubcontoLine(aasl *AccountAccountSubcontoLine) (int64, error) {
	return c.Create(AccountAccountSubcontoLineModel, aasl)
}

// UpdateAccountAccountSubcontoLine updates an existing account.account.subconto.line record.
func (c *Client) UpdateAccountAccountSubcontoLine(aasl *AccountAccountSubcontoLine) error {
	return c.UpdateAccountAccountSubcontoLines([]int64{aasl.Id.Get()}, aasl)
}

// UpdateAccountAccountSubcontoLines updates existing account.account.subconto.line records.
// All records (represented by ids) will be updated by aasl values.
func (c *Client) UpdateAccountAccountSubcontoLines(ids []int64, aasl *AccountAccountSubcontoLine) error {
	return c.Update(AccountAccountSubcontoLineModel, ids, aasl)
}

// DeleteAccountAccountSubcontoLine deletes an existing account.account.subconto.line record.
func (c *Client) DeleteAccountAccountSubcontoLine(id int64) error {
	return c.DeleteAccountAccountSubcontoLines([]int64{id})
}

// DeleteAccountAccountSubcontoLines deletes existing account.account.subconto.line records.
func (c *Client) DeleteAccountAccountSubcontoLines(ids []int64) error {
	return c.Delete(AccountAccountSubcontoLineModel, ids)
}

// GetAccountAccountSubcontoLine gets account.account.subconto.line existing record.
func (c *Client) GetAccountAccountSubcontoLine(id int64) (*AccountAccountSubcontoLine, error) {
	aasls, err := c.GetAccountAccountSubcontoLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if aasls != nil && len(*aasls) > 0 {
		return &((*aasls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.account.subconto.line not found", id)
}

// GetAccountAccountSubcontoLines gets account.account.subconto.line existing records.
func (c *Client) GetAccountAccountSubcontoLines(ids []int64) (*AccountAccountSubcontoLines, error) {
	aasls := &AccountAccountSubcontoLines{}
	if err := c.Read(AccountAccountSubcontoLineModel, ids, nil, aasls); err != nil {
		return nil, err
	}
	return aasls, nil
}

// FindAccountAccountSubcontoLine finds account.account.subconto.line record by querying it with criteria.
func (c *Client) FindAccountAccountSubcontoLine(criteria *Criteria) (*AccountAccountSubcontoLine, error) {
	aasls := &AccountAccountSubcontoLines{}
	if err := c.SearchRead(AccountAccountSubcontoLineModel, criteria, NewOptions().Limit(1), aasls); err != nil {
		return nil, err
	}
	if aasls != nil && len(*aasls) > 0 {
		return &((*aasls)[0]), nil
	}
	return nil, fmt.Errorf("account.account.subconto.line was not found")
}

// FindAccountAccountSubcontoLines finds account.account.subconto.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountSubcontoLines(criteria *Criteria, options *Options) (*AccountAccountSubcontoLines, error) {
	aasls := &AccountAccountSubcontoLines{}
	if err := c.SearchRead(AccountAccountSubcontoLineModel, criteria, options, aasls); err != nil {
		return nil, err
	}
	return aasls, nil
}

// FindAccountAccountSubcontoLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountSubcontoLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAccountSubcontoLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAccountSubcontoLineId finds record id by querying it with criteria.
func (c *Client) FindAccountAccountSubcontoLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccountSubcontoLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.account.subconto.line was not found")
}
