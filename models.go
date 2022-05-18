package main

type Metadata struct {
	OdataMetadata string     `json:"odata.metadata"`
	Values        []struct{} `json:"value"`
}

type Value struct {
	RefKey                    string `json:"Ref_Key"`
	DataVersion               string `json:"DataVersion"`
	DeletionMark              bool   `json:"DeletionMark"`
	ParentKey                 string `json:"Parent_Key"`
	Code                      string `json:"Code"`
	Description               string `json:"Description"`
	Order                     string `json:"Order"`
	Type                      string `json:"Type"`
	OffBalance                bool   `json:"OffBalance"`
	ProhibitUseInTransactions bool   `json:"ЗапретитьИспользоватьВПроводках"`
	QuickPickCode             string `json:"КодБыстрогоВыбора"`
	Currency                  bool   `json:"Валютный"`
	Quantitative              bool   `json:"Количественный"`
	AccountingByDepartment    bool   `json:"УчетПоПодразделениям"`
	TaxAccounting             bool   `json:"НалоговыйУчет"`
	ExtDimensionTypes         []struct {
		LineNumber          string `json:"LineNumber"`
		ExtDimensionTypeKey string `json:"ExtDimensionType_Key"`
		Predefined          bool   `json:"Predefined"`
		TurnoversOnly       bool   `json:"TurnoversOnly"`
		Sum                 bool   `json:"Суммовой"`
		Currency            bool   `json:"Валютный"`
		Quantitative        bool   `json:"Количественный"`
	} `json:"ExtDimensionTypes"`
	Predefined              bool   `json:"Predefined"`
	PredefinedDataName      string `json:"PredefinedDataName"`
	ParentNavigationLinkURL string `json:"Parent@navigationLinkUrl,omitempty"`
}
