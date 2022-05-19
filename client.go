package main

// oData1C config
type ClientOdata struct {
	Login    string
	Password string
	URL      string
}

func (c *ClientOdata) NewClient(login string, pass string, url string) ClientOdata {
	return ClientOdata{
		Login:    login,
		Password: pass,
		URL:      url,
	}
}
