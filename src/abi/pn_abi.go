package abi

import (
	// "fmt"
	// "log/slog"
	// "net/url"
	// "slices"
	// "strings"

	// "ibcinc.com/ibcdb/conf"
	// prof "ibcinc.com/ibcdb/pactrak/profiles"
	// db "ibcinc.com/ibcdb/query"
)

func toAirportPort(ibc_port string) string {
	switch ibc_port {
	case "MIA":
		return "5206"
	case "ORD":
		return "3901"
	case "DFW":
		return "5501"
	case "LAX":
		return "2720"
	default: // NYC
		return "4701"
	}
}

// func BuildPNRequest(key int, datestr string, url *url.URL) (string, error) {
// 	ef := func(msg string, err error) (string, error) {
// 		slog.Warn("BuildPNRequest issue reported", "message", msg, "cuase", err)
// 		return "", err
// 	}
// 	var prof prof.IBCProfile
// 	err := LoadIBCProfile(key, &prof)
// 	if err != nil {
// 		return ef(fmt.Sprintf("Profile for key could not be loaded %d", key), err)
// 	}
// 	// // now := time.Now()
// 	// // dateString := now.Format("2006-01-02")
// 	// dateString := "2025-12-24"
// 	q := url.Query()
// 	q.Set("dateFrom", datestr)
// 	q.Set("dateTo", datestr)
// 	q.Set("houseBOLNumber", trim(prof.Profile.BillNum.String))
// 	q.Set("masterBOLNumber", prof.Manifest.GetMaster())
// 	q.Set("type", "fda-pn")
// 	q.Set("skip", "0")
// 	url.RawQuery = q.Encode()
// 	return url.String(), nil
// }

// func BuildPNBody(key int) ([]PNBody, error) {
// 	ef := func(msg string, err error) ([]PNBody, error) {
// 		slog.Warn("BuildPNBody issue reported", "message", msg, "cuase", err)
// 		return nil, err
// 	}
//
// 	bdys := make([]PNBody, 0)
// 	bdy := PNBody{}
//
// 	var prof prof.IBCProfile
// 	err := LoadIBCProfile(key, &prof)
// 	if err != nil {
// 		return ef(fmt.Sprintf("Profile for key could not be loaded %d", key), err)
// 	}
//
// 	flt := prof.Manifest.GetFlight()
// 	// fmt.Println("FLIGHT DETAILS", flt)
// 	dt1 := dateToFormat(prof.Manifest.Date.String, "20060102")
//
// 	bdy.Type = "86"
// 	bdy.ReferenceQualifier = "AWB"
// 	bdy.Reference = fmt.Sprintf("%d", prof.Profile.Key.Int64)
// 	bdy.TransportMode = "40"
// 	bdy.BillType = "M"
// 	bdy.Master = prof.Manifest.GetMaster()
// 	bdy.House = trim(prof.Profile.BillNum.String)
// 	bdy.ArrivalDate = dt1
// 	bdy.ArrivalTime = "15:00"
// 	bdy.ArrivalPort = toAirportPort(prof.Profile.ExportPort1.String)
// 	bdy.OiDescription = trim(prof.Profile.Description.String)
	// bdy.CarrierName = flt[0:2]
	// bdy.Voyage = trim(flt[3:])
	//
	// counter := 0
	// for _, pks := range prof.Packages {
	// 	for _, gd := range pks.Goods {
	// 		val := gd.GetGoodXValue("fdap")
	// 		if val == "FOO" { // proceed
	// 			var itm PNItem
	// 			counter = counter + 1
	// 			itm.LineNumber = counter
	// 			// itm.ProductCode = strings.ReplaceAll(gd.GetGoodXValue("fdac"), " ", "")
	// 			itm.FDAProductCode = strings.ReplaceAll(gd.GetGoodXValue("fdac"), " ", "")
	// 			// itm.ProductCode = itm.FDAProductCode  // Dont set the ProductCode  -- important
	// 			itm.ProductDescription = max(trim(gd.Description.String), 65)
	// 			itm.ItemDescription = itm.ProductDescription
	// 			itm.AgencyProcessingCode = gd.GetGoodXValue("fdad")
	// 			itm.AgencyProgramCode = val     // gd.GetGoodXValue("fdap") //or val
	// 			itm.IntendedUseCode = "210.000" // gd.GetGoodXValue("fdai")
	// 			itm.CountryOfShipment = trim(prof.Shipper.Country.String)
	// 			itm.TradeName = gd.GetGoodXValue("fdab")
	// 			if itm.TradeName == "" {
	// 				itm.TradeName = max(itm.ItemDescription, 30)
	// 			}
	//
	// 			itm.CountryOfOrigin = trim(gd.OriginCountry.String)
	// 			itm.SourceTypeCode = "39" // 39 countryofProduction, 30 countryOfSource, CSH countryofShipment, 262 DrugFactLabeling FDC food or cosmetic , 294 countryofRefusal
	//
	// 			cons := "CONSIGNEE"
	// 			sell := "SELLER"
				// // submitter
				// itm.SubmitterLoadEntry = &cons
				// // itm.SubmitterEmail = "mia-cs@ibcinc.com"
				// // transmitter
				// itm.TransmitterName = "IBC"
				// itm.TransmitterAddress = "8401 NW 17th St"
				// itm.TransmitterCity = "Miami"
				// itm.TransmitterState = "FL"
				// itm.TransmitterZip = "33126"
				// itm.TransmitterCountry = "US"
				// itm.TransmitterEmail = "it@ibcinc.com"
				// itm.TransmitterPhone = "305-591-8080"
				// itm.TransmitterContactName = "Brokerage Dept"
				//
				// // manufacturer
				// itm.ManufacturerLoadEntry = &sell // either this or add the name address, city and country
				//
				// // importer
				// itm.ImporterLoadEntry = &cons
				// // itm.ImporterEmail = "mia-cs@ibcinc.com"
				// // owner
				// itm.OwnerLoadEntry = &cons
				//
				// // shipper seller
				// itm.ShipperName = trim(prof.Shipper.Name.String)
				// itm.ShipperAddress = max(fmt.Sprintf("%s", trim(prof.Shipper.Address1.String)), 35)
				// itm.ShipperAddress2 = max(fmt.Sprintf("%s", trim(prof.Shipper.Address2.String)), 35)
				// itm.ShipperCountry = trim(prof.Shipper.Country.String) // to 2 us country code
				// itm.ShipperCity = max(trim(prof.Shipper.City.String), 25)
				// itm.ShipperState = max(trim(prof.Shipper.State.String), 9)
				// itm.ShipperZip = trim(prof.Shipper.Zip.String)
				// itm.ShipperPhone = trim(prof.Shipper.Phone.String)
				// itm.ShipperContactName = itm.ShipperName
				//
				// // override to sellert country of origin
				// if itm.ShipperCountry != itm.CountryOfOrigin {
				// 	itm.CountryOfOrigin = itm.ShipperCountry
				// }
				//
				// // consignee
				// cs, sz := extractStateZip(prof.Profile.ConState.String)
				// itm.ConsigneeName = trim(fmt.Sprintf("%s %s", trim(prof.Profile.Consignee1.String), trim(prof.Profile.Consignee2.String)))
				// itm.ConsigneeAddress = max(fmt.Sprintf("%s", trim(prof.Profile.ConStreet1.String)), 35)
				// itm.ConsigneeAddress2 = max(fmt.Sprintf("%s", trim(prof.Profile.ConStreet2.String)), 35)
				// itm.ConsigneeCountry = "US" // trim(prof.Profile.Destination.String), // to 2 us country code
				// itm.ConsigneeCity = max(trim(prof.Profile.ConCity.String), 25)
				// itm.ConsigneeState = max(cs, 9)
				// itm.ConsigneeZip = sz
				// itm.ConsigneePhone = trim(prof.ProfileD.ConsigneePhone.String)
				// if itm.ConsigneePhone == "" {
				// 	itm.ConsigneePhone = "305-591-8080"
				// }
				// itm.ConsigneeEmail = trim(prof.GetProfileXValue("M")) // VERY IMPORTANT
				// if itm.ConsigneeEmail == "" {
				// 	itm.ConsigneeEmail = "mia-cs@ibcinc.com"
				// }
				// itm.ConsigneeContactName = itm.ConsigneeName
				//
				// qt1 := int(gd.Quantity.Int32)
				// if qt1 == 0 {
	// 				qt1 = 1
	// 			}
	// 			itm.Packaging = []PNPackaging{{UOM: "G", Quantity: qt1}}
	//
	// 			ccodes := LimitCompCodes(gd.ToCompCodeMap("FDA"))
	//
	// 			for _, cd1 := range ccodes {
	// 				itm.Compliance = append(itm.Compliance, ComplianceCode{cd1.Code, cd1.Value})
	// 			}
	//
	// 			// https://www.fda.gov/media/186988/download?attachment
	// 			fd1 := gd.GetGoodXValue("fdat")
	// 			fd2 := gd.GetGoodXValue("fdan")
	//
	// 			if fd1 != "" && fd2 != "" {
	// 				itm.Compliance = append(itm.Compliance, ComplianceCode{fd1, fd2})
	// 			}
	//
	// 			seenFME := slices.ContainsFunc(itm.Compliance, func(c ComplianceCode) bool {
	// 				return c.Code == "FME"
	// 			})
	//
	// 			if !seenFME {
	// 				itm.Compliance = append(itm.Compliance, ComplianceCode{"FME", "K"})
	// 			}
	//
	// 			bdy.Items = append(bdy.Items, itm)
	// 		}
	// 	}
	// }

	// test here if any goods were retrieved otherwise return an error
// 	if len(bdy.Items) == 0 {
// 		ms := fmt.Sprintf("no pga content found for key %d", key)
// 		return ef(ms, fmt.Errorf(ms))
// 	}
// 	bdy.Items = Dedup(bdy.Items)
// 	bdys = append(bdys, bdy)
//
// 	return bdys, nil
// }

func Dedup(list []PNItem) []PNItem {
	seen := make(map[string]int)
	unique := []PNItem{}

	for _, item := range list {
		if _, ok := seen[item.FDAProductCode]; !ok {
			seen[item.FDAProductCode] = 1
			unique = append(unique, item)
		} else {
			seen[item.FDAProductCode] = seen[item.FDAProductCode] + 1
		}
	}
	for x, it := range unique {
		unique[x].LineNumber = x + 1
		unique[x].Packaging[0].Quantity = seen[it.FDAProductCode]
	}
	return unique
}

// func LimitCompCodes(input []prof.CompCode) []prof.CompCode {
// 	itms := []string{"VFT", "VES", "RNO", "FME", "PFR"} //, "CAN", "FCE", "SID", "VOL", "FSX", "RNE", "IFE", "PKC", "AIN" }
// 	keepIndx := 0
//
// 	for _, val := range input {
// 		if slices.Contains(itms, val.Code) {
// 			input[keepIndx] = val
// 			keepIndx++
// 		}
// 	}
// 	input = input[:keepIndx]
// 	return input
// }

// func LoadIBCProfile(key int, pr *prof.IBCProfile) error {
// 	err := prof.LoadFullProfile(key, pr, conf.ToDSNDatabaseStringTargetDb("pt_aams"))
// 	if err != nil {
// 		return err
// 	}
// 	// fmt.Printf("p0 %+v\n", pr)
// 	// fmt.Printf("p0 %+v\n", pr.Packages[0].Goods[0].GoodExtra)
// 	// fmt.Printf("p0 %+v\n", pr.Packages[0].Goods[0].GoodExtra)
// 	return nil
// }
//
// func LoadSimpleProfile(key int, pr *prof.IBCProfile) error {
// 	err := prof.LoadProfile(key, pr, conf.ToDSNDatabaseStringTargetDb("pt_aams"))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func StorePN(key int, items []PNCItem) []error {
// 	var errs []error
// 	// ef := func(msg string, err error) []error {
// 	// 	slog.Warn("StorePN issue reported", "message", msg, "cuase", err)
// 	// 	errs = append(errs, err)
// 	// 	return errs
// 	// }
//
// 	qry := "execute procedure add_prior_notice(?, ?, ?, ?)"
// 	dsn := conf.ToDSNDatabaseStringTargetDb("pt_aams")
//
// 	// fmt.Println(dsn)
//
// 	// db, err := sql.Open("odbc", dsn)
// 	// if err != nil {
// 	// 	return ef("db could not be opened!", err)
// 	// }
// 	// defer db.Close()
//
// 	for k, itm := range items {
// 		pc := itm.ProductCodeNumber
// 		pn := itm.PNCNumber
// 		val := "f"
// 		if k == 0 {
// 			val = "t"
// 		}
//
// 		// res, err := db.Exec(qry, key, pc, pn, val)
// 		res, err := db.ExecuteInsertOrUpdate(qry, []any{key, pc, pn, val}, dsn)
// 		if err != nil {
// 			slog.Warn("issue storing pn detail", "key", key, "Product", pc, "PriorNotice", pn)
// 			errs = append(errs, err)
// 		} else {
// 			slog.Info("PN Stored", "key", key, "priorNotice", pn, "result", res)
// 		}
// 	}
// 	return errs
// }
