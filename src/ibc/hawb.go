package ibc

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Hawb struct {
	RecordType        string      `json:"record_type"`
	RecordVersion     int         `json:"record_version"`
	ProfileKey        string      `json:"profile_key"`
	House             string      `json:"hawb"`
	Reference         string      `json:"shipper_reference"`
	Reference2        string      `json:"internal_reference"`
	VendorReference   string      `json:"vendor_reference"`
	Origin            string      `json:"origin"`
	Destination       string      `json:"final_destination"`
	IsOutlying        string      `json:"outlying"`
	Provider          string      `json:"service_provider"`
	DlsStation        string      `json:"dls_station"`
	DlsFinalDest      string      `json:"dls_final_dest"`
	Pieces            string      `json:"pieces"`
	Weight            string      `json:"weight"`
	WeightUnit        string      `json:"weight_unit"`
	Content           string      `json:"content"`
	Currency          string      `json:"currency"`
	Value             string      `json:"value"`
	Insurance         string      `json:"insurance_amount"`
	Description       string      `json:"description"`
	HTS               string      `json:"hts_code"`
	FDANotice         string      `json:"fda_prior_notice"`
	Terms             string      `json:"terms"`
	Packaging         string      `json:"packaging"`
	ServiceType       string      `json:"service_type"`
	CollectAmount     string      `json:"collect_amount"`
	CustomerKey       string      `json:"cust_key"`
	ShipperAccount    string      `json:"shipper_account"`
	DlsAccount        string      `json:"dls_account"`
	ExtAccount        string      `json:"ext_account"`
	ShipperName       string      `json:"shipper_name"`
	ShipperAddress    string      `json:"shipper_address"`
	ShipperAddress2   string      `json:"shipper_address_2"`
	ShipperCity       string      `json:"shipper_city"`
	ShipperState      string      `json:"shipper_state"`
	ShipperZip        string      `json:"shipper_zip"`
	ShipperCountry    string      `json:"shipper_country"`
	ShipperPhone      string      `json:"shipper_phone"`
	ConsigneeName     string      `json:"consignee_name"`
	ConsigneeName2    string      `json:"consignee_company"`
	ConsigneeAddress  string      `json:"consignee_address"`
	ConsigneeAddress2 string      `json:"consignee_address_2"`
	ConsigneeCity     string      `json:"consignee_city"`
	ConsigneeState    string      `json:"consignee_state"`
	ConsigneeZip      string      `json:"consignee_zip"`
	ConsigneeCountry  string      `json:"consignee_country"`
	ConsigneePhone    string      `json:"consignee_phone"`
	ConsigneeEmail    string      `json:"consignee_email"`
	ConsigneeTaxId    string      `json:"consignee_tax_id"`
	Comments          string      `json:"comments"`
	GoodsOrigin       string      `json:"goods_origin_country"`
	IncomingContainer string      `json:"incoming_container"`
	MID               string      `json:"mid"`
	Goods             []Commodity `json:"commodities"`
	mu                sync.Mutex  `json:"-"`
}

func (e *Hawb) Defaults() {
	e.RecordType = "hawb"
	e.RecordVersion = 14
}

func (e *Hawb) ReadLine(input string) error {
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

func (e *Hawb) Load(rc []string) error {
	e.RecordType = rc[0]
	if x, err := strconv.Atoi(rc[1]); err != nil {
		return err
	} else {
		e.RecordVersion = x
	}
	e.ProfileKey = rc[2]
	e.House = rc[3]
	e.Reference = rc[4]
	e.Reference2 = rc[5]
	e.VendorReference = rc[6]
	e.Origin = rc[7]
	e.Destination = rc[8]
	e.IsOutlying = rc[9]
	e.Provider = rc[10]
	e.DlsStation = rc[11]
	e.DlsFinalDest = rc[12]
	e.Pieces = rc[13]
	e.Weight = rc[14]
	e.WeightUnit = rc[15]
	e.Content = rc[16]
	e.Currency = rc[17]
	e.Value = rc[18]
	e.Insurance = rc[19]
	e.Description = rc[20]
	e.HTS = rc[21]
	e.FDANotice = rc[22]
	e.Terms = rc[23]
	e.Packaging = rc[24]
	e.ServiceType = rc[25]
	e.CollectAmount = rc[26]
	e.CustomerKey = rc[27]
	e.ShipperAccount = rc[28]
	e.DlsAccount = rc[29]
	e.ExtAccount = rc[30]
	e.ShipperName = rc[31]
	e.ShipperAddress = rc[32]
	e.ShipperAddress2 = rc[33]
	e.ShipperCity = rc[34]
	e.ShipperState = rc[35]
	e.ShipperZip = rc[36]
	e.ShipperCountry = rc[37]
	e.ShipperPhone = rc[38]
	e.ConsigneeName = rc[39]
	e.ConsigneeName2 = rc[40]
	e.ConsigneeAddress = rc[41]
	e.ConsigneeAddress2 = rc[42]
	e.ConsigneeCity = rc[43]
	e.ConsigneeState = rc[44]
	e.ConsigneeZip = rc[45]
	e.ConsigneeCountry = rc[46]
	e.ConsigneePhone = rc[47]
	e.ConsigneeEmail = rc[48]
	e.ConsigneeTaxId = rc[49]
	e.Comments = rc[50]
	e.GoodsOrigin = strings.TrimSpace(rc[51])

	fmt.Println(len(rc))

	if len(rc) > 52 {
		e.IncomingContainer = rc[52]

		if e.RecordVersion > 13 {
			e.MID = rc[53]
		}
	}
	return nil
}

func (e *Hawb) AddGood(g Commodity) {
	e.Goods = append(e.Goods, g)
}

func (e *Hawb) ToString() string {
	return fmt.Sprintf("%s [%d] track: '%s' hts: %s from: %s to: %s",
		e.RecordType, e.RecordVersion, e.House, e.HTS, e.ShipperName, e.ConsigneeName)
}

func (e *Hawb) Write(file *os.File, m Manifest) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	wrtr := csv.NewWriter(file)

	out := make([]string, 650)
	mwb := strings.ReplaceAll(m.Mawb, "-", "")

	out[0] = "11"                                     // ENTRY TYPE
	out[1] = ""                                       // CUSTOMER_ID
	out[4] = m.Date                                   // ENTRY_DT
	out[5] = m.Date                                   // IMPORT DT
	out[6] = m.Date                                   // ARRIVAL DT
	out[7] = "1200"                                   // ARRIVAL TM
	out[8] = m.Destination                            // PORT OF ENTRY  -- MIA  ORD  DFW  LAX
	out[9] = e.ConsigneeState                         // CON state:Write()
	out[10] = "40"                                    // MODE OF TRANSPORT
	out[12] = "IBC"                                   // Importer of record
	out[24] = "0"                                     // Bond type 0 for t11
	out[30] = m.Date                                  // Preliminary statment dt
	out[37] = m.IsExpress                             // N  ECCF y CFS n  express consignement
	out[32] = ToFirmsCode(m.Destination, m.IsExpress) // firms
	out[34] = "Y"                                     // Known Importer
	out[35] = "N"                                     // Perishable Goods
	out[41] = e.ConsigneeName                         // CONS Name
	out[42] = fmt.Sprintf("%s %s", e.ConsigneeAddress, e.ConsigneeAddress2)
	out[43] = e.ConsigneeCity
	out[44] = e.ConsigneeState
	out[45] = e.ConsigneeZip[0:5]
	out[46] = e.ConsigneeCountry
	out[47] = "M" // Bill Type M master
	out[48] = mwb
	out[49] = e.House
	out[52] = "Y"                                    // GROUP Y
	out[55] = m.Flight[0:2]                          // Flight carrier
	out[56] = m.Flight[3:]                           // Flight no
	out[57] = m.Destination                          //
	out[58] = m.Origin                               //
	out[59] = m.Flight[0:2]                          // Flight carrier
	out[60] = m.Flight[3:]                           // Flight no
	out[63] = "PCS"                                  // Flight no
	out[65] = m.Flight                               // Flight no
	out[94] = m.Date                                 // ARRIVAL DT
	out[95] = toPortCode(m.Destination, m.IsExpress) // PORT OF ENTRY  -- MIA  ORD  DFW  LAX
	out[97] = e.GoodsOrigin
	out[98] = e.Currency
	out[103] = e.Pieces
	out[104] = "PCS"
	out[105] = e.Description
	out[106] = e.House
	out[108] = e.GoodsOrigin
	out[117] = e.HTS
	out[121] = e.Description
	out[122] = e.GoodsOrigin
	out[124] = e.Currency
	out[126] = e.Currency
	out[127] = e.Value
	out[128] = e.Pieces
	out[133] = e.Weight
	out[134] = e.WeightUnit
	// 177 Manufacturer ID Shipper information
	out[205] = "CONSIGNEE"
	out[207] = e.ConsigneeName // CONS Name
	out[208] = fmt.Sprintf("%s %s", e.ConsigneeAddress, e.ConsigneeAddress2)
	out[210] = e.ConsigneeCity
	out[211] = e.ConsigneeState
	out[212] = out[45]
	out[213] = e.ConsigneeCountry
	out[218] = "CONSIGNEE"
	out[220] = e.ConsigneeName // CONS Name
	out[221] = fmt.Sprintf("%s %s", e.ConsigneeAddress, e.ConsigneeAddress2)
	out[223] = e.ConsigneeCity
	out[224] = e.ConsigneeState
	out[225] = out[45]
	out[226] = e.ConsigneeCountry

	// pcs, err := strconv.Atoi( e.Pieces )
	// if err != nil {
	// 	pcs = 1
	// }

	if len(e.Goods) == 0 {
		if err := writeLine(wrtr, out); err != nil {
			return err
		}
	} else {
		for i := range e.Goods {
			out[103] = e.Goods[i].Quantity
			out[128] = e.Goods[i].Quantity
			out[105] = e.Goods[i].Description
			out[108] = e.Goods[i].OriginCountry
			out[117] = e.Goods[i].HTS
			out[124] = e.Goods[i].ItemCurrency
			out[126] = e.Goods[i].ItemCurrency
			out[127] = e.Goods[i].ItemValue
			if e.Goods[i].ItemWeight != "" {
				out[133] = e.Goods[i].ItemWeight
			}

			//
			if err := writeLine(wrtr, out); err != nil {
				return err
			}
		}
	}

	return nil
}

func writeLine(wrtr *csv.Writer, output []string) error {
	if err := wrtr.Write(output); err != nil {
		log.Printf("could not write to file")
		return err
	}
	wrtr.Flush()

	if err := wrtr.Error(); err != nil {
		log.Printf("Error detected writing to file %+v", err)
		return err
	}
	return nil
}

// FIRMS CODE
//         CFS -       ECCF
// JFK     E200        F707
// LAX     W978        Y926
// MIA     LAD0        N828
// ORD     HAL6        HAO2
//

func ToFirmsCode(station, is_express string) string {
	if is_express == "Y" {
		// ECCF
		switch station {

		case "LAX":
			return "Y926"
		case "MIA":
			return "N828"
		case "ORD":
			return "HAO2"
		default: // JFK
			return "F707"
		}
	} else {
		switch station {

		case "LAX":
			return "W978"
		case "MIA":
			return "LAD0"
		case "ORD":
			return "HAL6"
		default: // JFK
			return "E200"
		}
	}
}

// func toPortCode(station string) string {
// 	switch station {
//
// 	case "LAX":
// 		return "2720"
// 	case "MIA":
// 		return "5206"
// 	case "ORD":
// 		return "3901"
// 	case "DFW":
// 		return "5501"
// 	default: // JFK
// 		return "4701"
// 	}
// }

func toPortCode(station string, is_express string) string {
	if is_express == "Y" {
		// ECCF
		switch station {

		case "LAX":
			return "2776"
		case "MIA":
			return "5298"
		case "ORD":
			return "3972"
		case "DFW":
			return "5501"
		default: // JFK
			return "4774"
		}
	} else {
		// CFS
		switch station {

		case "LAX":
			return "2720"
		case "MIA":
			return "5206"
		case "ORD":
			return "3901"
		case "DFW":
			return "5501"
		default: // JFK
			return "4701"
		}
	}
}
