package ibc

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

type Commodity struct {
	RecordType          string `json:"record_type"`
	RecordVersion       int    `json:"record_version"`
	Quantity            string `json:"quantity"`
	Description         string `json:"description"`
	URL                 string `json:"url"`
	HTS                 string `json:"hts_code"`
	OriginCountry       string `json:"origin_country"`
	ItemValue           string `json:"item_value"`
	ItemCurrency        string `json:"item_currency"`
	ItemWeight          string `json:"item_weight"`
	ItemWeightUnit      string `json:"item_weight_unit"`
	OrderId             string `json:"order_id"`
	OrderDt             string `json:"order_date"`
	ListPrice           string `json:"list_price"`
	ListPriceCurrency   string `json:"list_price_currency"`
	RetailPrice         string `json:"retail_price"`
	RetailPriceCurrency string `json:"retail_price_currency"`
	ManufacturerId      string `json:"manufacturer_id"`
}

func (e *Commodity) Defaults() {
	e.RecordType = "commodity"
	e.RecordVersion = 1
}

func (e *Commodity) ReadLine(input string) error {
	rdr := csv.NewReader(strings.NewReader(input))
	rc, err := rdr.Read()
	if err != nil {
		return err
	}
	if err := e.Load(rc); err != nil {
		return err
	}
	return nil
}

func (e *Commodity) Load(rc []string) error {
	e.RecordType = rc[0]
	if x, err := strconv.Atoi(rc[1]); err != nil {
		return err
	} else {
		e.RecordVersion = x
	}
	e.Quantity = rc[2]
	e.Description = rc[3]
	e.URL = rc[4]
	e.HTS = rc[5]
	e.OriginCountry = rc[6]
	e.ItemValue = rc[7]
	e.ItemCurrency = rc[8]
	e.ItemWeight = rc[9]
	e.ItemWeightUnit = rc[10]
	e.OrderId = rc[11]
	e.OrderDt = rc[12]
	e.ListPrice = rc[13]
	e.ListPriceCurrency = rc[14]
	e.RetailPrice = rc[15]
	e.RetailPriceCurrency = rc[16]
	e.ManufacturerId = rc[17]
	return nil
}

func (e *Commodity) ToString() string {
	return fmt.Sprintf("%s [%d] commodity: '%s' hts: %s from: %s value: %s-%s",
		e.RecordType, e.RecordVersion, e.Description, e.HTS, e.OriginCountry, e.Quantity, e.ItemValue)
}
