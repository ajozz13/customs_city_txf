package abi

import (
	"fmt"
	// "log/slog"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	ibc "ibcinc.com/customs_city/ibc"
	// "ibcinc.com/ibcdb/conf"
	// prof "ibcinc.com/ibcdb/pactrak/profiles"
)

// func BuildABIBody(man_num int) ([]ABIBody, error) {
func BuildABIBody(man ibc.Manifest) []ABIBody {
	// ef := func(msg string, err error) ([]ABIBody, error) {
	// slog.Warn("BuildABIBody issue reported", "message", msg, "cuase", err)
	// return nil, err
	// }

	output := make([]ABIBody, 1)
	bdy := ABIBody{EntryType: "11", TransportMode: "40", Reference: man.Mawb}

	// px, err := loadIBCManifest(man_num)
	// if err != nil {
	// 	return ef("loading IBCManifest", err)
	// }
	// p1 := px[0]

	dt1 := man.Date // dateToFormat(p1.Manifest.Date.String, "20060102")
	bdy.ABIDates = ABIDates{Import: dt1, Entry: dt1, Arrival: dt1}
	// bdy.ABILocation = stationToLocation(p1.Manifest.Destination.String, false) // Manifest destination to state and port
	bdy.ABILocation = stationToLocation(man.Destination, false) // Manifest destination to state and port
	bdy.IOR = IOR{Number: "59-216679300", Name: "International Bonded Couriers, Inc."}
	bdy.Bond = ABIBond{Type: "0"}
	// bdy.Payment = ABIPayment{StatementDate: tomorrow.Format("20060102"), Type: 2}
	bdy.Payment = ABIPayment{StatementDate: GetPaymentDate(dt1), Type: 2}

	// isExpress := "N"
	// if is_express {
	// 	isExpress = "Y"
	// }
	bdy.Firms = ibc.ToFirmsCode(man.Destination, man.IsExpress)
	bdy.IsKnownImporter = "Y"
	bdy.IsPerishable = "N"
	bdy.IsNonAMS = "N"
	// if p1.Manifest.ClearMethod.String == "E" {
	// 	isExpress = "Y"
	// }
	bdy.IsExpress = man.IsExpress // CFS N - ECCF Y
	// bdy.EntryConsignee = stationToEntity(p1.Manifest.Destination.String)
	bdy.EntryConsignee = stationToEntity(man.Destination)

	// TODO HERE
	// bdy.Manifests, err = BuildPactrakManifest(px)
	// if err != nil {
	// return output, err
	// }

	output[0] = bdy
	return output
}

func GetPaymentDate(entrydt string) string {
	holidays := []string{"0101", "0525", "0704", "0907", "1111", "1126", "1225"}
	now, _ := time.Parse("20060102", entrydt)

	today := time.Now()
	// fmt.Println( now )
	// fmt.Println( today )
	// fmt.Println( now.Before(today))

	if now.Before(today) {
		now = today
	}

	tomorrow := now.AddDate(0, 0, 1)
	dt := tomorrow.Format("20060102")

	// fmt.Println(slices.Contains(holidays, dt[4:]))

	if slices.Contains(holidays, dt[4:]) {
		tomorrow = now.AddDate(0, 0, 2)
		dt = tomorrow.Format("20060102")
	}

	return dt
}

func stationToLocation(input string, use_station_code bool) ABILocation {
	switch input {
	case "MIA":
		if use_station_code {
			return ABILocation{DestinationState: "FL", EntryPort: "MIA"}
		}
		return ABILocation{DestinationState: "FL", EntryPort: "5206"}
	case "ORD":
		if use_station_code {
			return ABILocation{DestinationState: "IL", EntryPort: "ORD"}
		}
		return ABILocation{DestinationState: "IL", EntryPort: "3901"}
	case "LAX":
		if use_station_code {
			return ABILocation{DestinationState: "CA", EntryPort: "LAX"}
		}
		return ABILocation{DestinationState: "CA", EntryPort: "2720"}
	case "DFW":
		if use_station_code {
			return ABILocation{DestinationState: "TX", EntryPort: "DFW"}
		}
		return ABILocation{DestinationState: "TX", EntryPort: "5501"}
	default: // NYC
		if use_station_code {
			return ABILocation{DestinationState: "NY", EntryPort: "JFK"}
		}
		return ABILocation{DestinationState: "NY", EntryPort: "4701"}
	}
}

func stationToEntity(input string) *ABIEntity {
	ent := ABIEntity{Name: "IBC Inc.", Country: "US"}

	switch input {
	case "MIA":
		ent.Address = "8401 NW 17th St"
		ent.City = "Miami"
		ent.State = "FL"
		ent.PostalCode = "33126"
	case "ORD":
		ent.Address = "9742 W Foster Ave"
		ent.City = "Chicago"
		ent.State = "IL"
		ent.PostalCode = "60018"
	case "LAX":
		ent.Address = "11034 S La Cienega Blvd"
		ent.City = "Inglewood"
		ent.State = "CA"
		ent.PostalCode = "90304"
	case "DFW":
		ent.Address = "601 Hanover #500"
		ent.City = "Grapevine"
		ent.State = "TX"
		ent.PostalCode = "76051"
	default: // NYC
		ent.Address = "152-01 Rockaway Blvd"
		ent.City = "Jamaica"
		ent.State = "NY"
		ent.PostalCode = "11434"
	}

	return &ent
}

func dateToFormat(input, output_format string) string {
	dt, err := time.Parse(time.RFC3339, input)
	if err != nil {
		dt = time.Now()
	}
	return dt.Format(output_format)
}

func extractFlightDetails(input string) (string, string, string) {
	// returns carrier flight, extra, err

	carrier := ""
	flight := ""
	extra := ""
	ft := strings.Split(input, "/")

	carrier = ft[0][0:2]
	flight = strings.TrimSpace(ft[0][2:])
	if len(ft) > 1 {
		extra = strings.TrimSpace(ft[1])
	}

	return carrier, flight, extra
}

func extractStateZip(input string) (string, string) {
	state := ""
	zip := ""

	st1 := strings.Split(input, ",")

	state = strings.TrimSpace(st1[0])
	if len(st1) == 2 {
		zip = strings.TrimSpace(strings.ReplaceAll(st1[1], "-", ""))
	}

	return state, zip
}

func max(input string, max int) string {
	op := cleanup(strings.TrimSpace(input))
	if len(op) > max {
		return op[0:max]
	}
	return op
}

func jmax(input string, max int) string {
	// op := cleanup(strings.TrimSpace(input))
	op := strings.TrimSpace(input)
	if len(op) > max {
		return op[0:max]
	}
	return op
}

func BuildManifest(out *[]ABIManifest, mn ibc.Manifest, hwb *ibc.Hawb) {
	// var output []ABIManifest
	var location ABILocation
	var man ABIManifest
	man.Bill.Marks = mn.Mawb

	on := strings.TrimSpace(hwb.House)
	ln := len(on)
	if ln > 12 {
		on = on[ln-12:]
	}

	// man.Bill.House = strings.TrimSpace(hwb.House)
	man.Bill.House = on
	man.Bill.Master = mn.Mawb
	man.Bill.Type = "M"
	// when PGA
	man.Bill.Group = "N"

	var inv Invoice
	location = stationToLocation(mn.Destination, false)
	inv.ExportDate = dateToFormat(mn.Date, "20060102")

	man.Carrier.Code = mn.Flight[0:2]
	man.Carrier.Vessel = mn.Flight[2:]
	man.Ports.Origin = mn.Origin
	man.Ports.Operator = man.Carrier.Code
	// man.Ports.Flight = mn.Flight
	man.Ports.Flight = man.Carrier.Vessel
	man.Ports.Unlading = location.EntryPort // already a numeric port code

	inv.ExportCountry = hwb.GoodsOrigin
	if inv.ExportCountry == "" {
		inv.ExportCountry = hwb.ShipperCountry
	}
	inv.Currency = hwb.Currency
	if inv.Currency == "" {
		inv.Currency = "USD"
	}
	inv.PurchaseOrder = hwb.House
	lpo := len(inv.PurchaseOrder)
	if lpo > 20 {
		inv.PurchaseOrder = hwb.House[lpo-20:]
	}

	inv.InvoiceNumber = hwb.House // inv.PurchaseOrder                         // when fda
	inv.TrackNumber = hwb.House
	inv.MarksNumber = hwb.House
	inv.RelatedParties = "N"
	inv.Package.Description = max(hwb.Description, 45)
	inv.Package.UOM = "PCS"
	inv.Package.Quantity = hwb.Pieces
	inv.Package.OriginCountry = inv.ExportCountry
	inv.Package.IsFDA = false

	// load good items inside the items

	inv.Items = make([]Item, 0)
	LoadItem(hwb, &inv.Items)
	man.Invoices = append(man.Invoices, inv)
	*out = append(*out, man)

	// return output
}

// ABIManifest is part of the ABIBody sent to CC
// func BuildPactrakManifest(ps []prof.IBCProfile) ([]ABIManifest, error) {
// ef := func(msg string, err error) ([]ABIManifest, error) {
// 	slog.Warn("BuildPactrakManifest issue reported", "message", msg, "cuase", err)
// 	return nil, err
// }

// var output []ABIManifest
//
// airline := ""
// flight := ""
// xtra := ""
// // fmt.Println(xtra)
//
// var location ABILocation
//
// for indx, prof := range ps {
//
// 	var mn ABIManifest
// 	mn.Bill.Marks = fmt.Sprintf("%d", prof.Manifest.Number.Int64)
// 	mn.Bill.House = strings.TrimSpace(prof.Profile.BillNum.String)
// 	mn.Bill.Master = prof.Manifest.GetMaster()
// 	mn.Bill.Type = "M"
// 	// when PGA
// 	mn.Bill.Group = "N"
//
// 	var inv Invoice
// 	if indx == 0 {
// 		airline, flight, xtra = extractFlightDetails(prof.Manifest.GetFlight())
// 		location = stationToLocation(prof.Manifest.Destination.String, false)
// 		inv.ExportDate = dateToFormat(prof.Manifest.Date.String, "20060102")
// 	}
// 	mn.Carrier.Code = airline
// 	mn.Carrier.Vessel = flight
//
// 	mn.Ports.Origin = prof.Manifest.Origin.String
// 	mn.Ports.Operator = airline
// mn.Ports.Flight = flight
// mn.Ports.Unlading = location.EntryPort
//
// inv.ExportCountry = trim(prof.GetProfileXValue("O"))
// inv.Currency = "USD"
// inv.PurchaseOrder = fmt.Sprintf("%d", prof.Profile.Key.Int64) // probably order id
// inv.InvoiceNumber = mn.Bill.House                             // inv.PurchaseOrder                         // when fda
// inv.TrackNumber = prof.ProfileD.SecondReference.String
// inv.MarksNumber = fmt.Sprintf("%d", prof.Profile.Key.Int64)
// inv.RelatedParties = "N"
// inv.Package.Description = max(prof.Profile.Description.String, 45)
// inv.Package.UOM = "PCS"
// inv.Package.Quantity = fmt.Sprintf("%d", prof.Profile.TotalPieces.Int32)
// inv.Package.OriginCountry = trim(prof.GetProfileXValue("O"))
// inv.Package.IsFDA = false
//
// if inv.ExportCountry == "" {
// 	inv.ExportCountry = trim(prof.Shipper.Country.String)
// }
// if inv.Package.OriginCountry == "" {
// 	inv.Package.OriginCountry = inv.ExportCountry
// }
//
// // fmt.Println("PACKAGES", len(prof.Packages))
// // fmt.Printf("PACKAGES %+v\n", prof.Packages)
// // fmt.Println("GOODS", len(prof.Packages[0].Goods))
//
// for i, pk := range prof.Packages { // i
// 	for _, gd := range pk.Goods { // j
// 		var it Item
// it.Row = i + 1
// it.Reference = fmt.Sprintf("%d-%d", gd.PackageKey.Int32, gd.GoodKey.Int32)
// it.CountryOrigin = gd.OriginCountry.String
// it.OriginDetail = &Origin{Country: gd.OriginCountry.String}
// // it.OriginDetail.Country = it.CountryOrigin
// it.Description = max(gd.Description.String, 70)
// it.HTS = gd.HtsCode.String
//
// it.Quantity1 = fmt.Sprintf("%d", gd.Quantity.Int32)
// it.Values.Currency = "USD"
// it.Values.ValueOfGoods = gd.Value.Float64
// // fixes
// if it.Values.ValueOfGoods == 0 {
// 	it.Values.ValueOfGoods = prof.Profile.TotalValue.Float64 / float64(len(pk.Goods))
// }
//
// it.Weight.UOM = "L" // or K
// if gd.Weight.Float64 == 0 {
// 	it.Weight.Gross = fmt.Sprintf("%.2f", prof.Profile.TotalWeight.Float64/float64(len(pk.Goods)))
// } else {
// 	it.Weight.Gross = fmt.Sprintf("%.2f", gd.Weight.Float64)
// }
//
// it.Parties = append(it.Parties, ABIEntity{
// 	Type:       "seller",
// 	Name:       trim(prof.Shipper.Name.String),
// 	Address:    jmax(fmt.Sprintf("%s %s", trim(prof.Shipper.Address1.String), trim(prof.Shipper.Address2.String)), 35),
// 	Country:    trim(prof.Shipper.Country.String), // to 2 us country code
// 	City:       trim(prof.Shipper.City.String),
// 	State:      jmax(trim(prof.Shipper.State.String), 9),
// 	PostalCode: jmax(trim(prof.Shipper.Zip.String), 9),
// })
//
// midc := gd.GetGoodXValue("midc")
// manu := ABIEntity{
// 	Type:       "manufacturer",
// 	Name:       trim(prof.Shipper.Name.String),
// 	Address:    jmax(fmt.Sprintf("%s %s", trim(prof.Shipper.Address1.String), trim(prof.Shipper.Address2.String)), 70),
// 	Country:    trim(prof.Shipper.Country.String), // to 2 us country code
// 	City:       trim(prof.Shipper.City.String),
// 	State:      jmax(trim(prof.Shipper.State.String), 9),
// 	PostalCode: jmax(trim(prof.Shipper.Zip.String), 9),
// }
// if midc != "" {
// 	manu.MID = midc
// }
// // else {
// // 	manu.Name = trim(prof.Shipper.Name.String)
// // 	manu.Address = fmt.Sprintf("%s %s", trim(prof.Shipper.Address1.String), trim(prof.Shipper.Address2.String))
// // 	manu.Country = trim(prof.Shipper.Country.String) // to 2 us country code
// // 	manu.City = trim(prof.Shipper.City.String)
// // 	manu.State = trim(prof.Shipper.State.String)
// // 	manu.PostalCode = trim(prof.Shipper.Zip.String)
// // }
// it.Parties = append(it.Parties, manu)
// // it.Parties = append(it.Parties, ABIEntity{
// // 	Type: "manufacturer",
// // 	// MID:  "",
// // 	Name:       trim(prof.Shipper.Name.String),
// // 	Address:    fmt.Sprintf("%s %s", trim(prof.Shipper.Address1.String), trim(prof.Shipper.Address2.String)),
// // 	Country:    trim(prof.Shipper.Country.String), // to 2 us country code
// // 	City:       trim(prof.Shipper.City.String),
// // 	State:      trim(prof.Shipper.State.String),
// // 	PostalCode: trim(prof.Shipper.Zip.String),
// // })
// cs, sz := extractStateZip(prof.Profile.ConState.String)
//
// it.Parties = append(it.Parties, ABIEntity{
// 	Type:       "consignee",
// 	Name:       cleanup(fmt.Sprintf("%s %s", trim(prof.Profile.Consignee1.String), trim(prof.Profile.Consignee2.String))),
// 	Address:    jmax(fmt.Sprintf("%s %s", trim(prof.Profile.ConStreet1.String), trim(prof.Profile.ConStreet2.String)), 35),
// 	Country:    "US", // trim(prof.Profile.Destination.String), // to 2 us country code
// 	City:       trim(prof.Profile.ConCity.String),
// 	State:      jmax(cs, 9),
// 	PostalCode: jmax(sz, 9),
// })
//
// it.Parties = append(it.Parties, ABIEntity{
// 	Type:     "buyer",
// 	LoadFrom: "consignee",
// 	// Name:       trim(fmt.Sprintf("%s %s", trim(prof.Profile.Consignee1.String), trim(prof.Profile.Consignee2.String))),
// 	// Address:    fmt.Sprintf("%s %s", trim(prof.Profile.ConStreet1.String), trim(prof.Profile.ConStreet2.String)),
// 	// Country:    "US", // to 2 us country code
// 	// City:       trim(prof.Profile.ConCity.String),
// 	// State:      cs,
// 	// PostalCode: sz,
// })
//
// it.Parties = append(it.Parties, ABIEntity{
// 	Type:     "shipTo",
// 	LoadFrom: "consignee",
// })
//
// if len(gd.GoodExtra) > 1 {
// 	inv.Package.IsFDA = true
//
// 	pct := 0
// 	one := 1
// 	it.AluminumPct = &pct
// 	it.SteelPct = &pct
// it.CopperPct = &pct
// it.HasCottonExemption = "N"
// it.HasAutoPartsExemption = "N"
// it.HasCompletedKitchenParts = "N"
// it.HasInformationalExemption = "N"
// it.ForReligiousPurpose = "N"
// it.HasAgriculturalExemption = "N"
// it.SemiConductorExemption = &one
//
// // when fda
// it.Parties = append(it.Parties, ABIEntity{
// 	Type:     "pgaShipper",
// 	LoadFrom: "seller",
// })
//
// it.Parties = append(it.Parties, ABIEntity{
// 	Type:     "pgaImporter",
// 	LoadFrom: "consignee",
// })
//
// it.Parties = append(it.Parties, ABIEntity{
// 	Type:     "pgaOwner",
// 	LoadFrom: "seller",
// })
//
// it.Parties = append(it.Parties, ABIEntity{
// 	Type:     "pgaDeliverToParty",
// 	LoadFrom: "consignee",
// })
//
// 			it.Parties = append(it.Parties, ABIEntity{
// 				Type:     "pgaForeignSupplierVerificationImporter",
// 				LoadFrom: "manufacturer",
// 			})
//
// 			// when fda
// 			it.PGAPackaging = []PGAPackaging{{Quantity: "1", UOM: "PCS"}}
// 			it.PGAData = &PGAData{FDA: &FDA{
// 				ProgramCode:     gd.GetGoodXValue("fdap"),
// 				ProcessingCode:  gd.GetGoodXValue("fdad"),
// 				IntendedUseCode: gd.GetGoodXValue("fdai"),
// 				ProductCode:     gd.GetGoodXValue("fdac"),
// 				ItemType:        "P",  // don't know what this is
// 				SourceType:      "39", // don't know what this is
// 				Compliance:      []ComplianceCode{},
// 			}}
//
// 			c1 := gd.GetGoodXValue("fdat")
// 			c2 := gd.GetGoodXValue("fdan")
//
// 			if c1 != "" {
// 				it.PGAData.FDA.Compliance = append(it.PGAData.FDA.Compliance, ComplianceCode{Code: c1, Value: c2})
// 			}
//
// 		}
// 		// fmt.Printf("Pack: %d %d GOOD Extra: %d - %d len: %d\n ", i, pk.PackageKey.Int32, j, gd.GoodKey.Int32, len(gd.GoodExtra))
//
// 		inv.Items = append(inv.Items, it)
// 	}
// }

// if len(inv.Items) == 0 { // no good data
// 	inv.ExportDate = dateToFormat(prof.Manifest.Date.String, "20060102")
// 	var it Item
// 	it.Row = 1
// 	it.Reference = fmt.Sprintf("%d", prof.Profile.Key.Int64)
// 	it.CountryOrigin = inv.ExportCountry
// 	it.OriginDetail = &Origin{Country: inv.ExportCountry}
// 	it.Description = max(prof.Profile.Description.String, 70)
// 	it.HTS = trim(prof.GetProfileXValue("h"))
// 	if len(it.HTS) == 0 {
// 		it.HTS = trim(prof.GetProfileXValue("H"))
// 	}
// 	// if len(it.HTS) == 0 {
// 	// 	it.HTS = "9999999999"
// 	// }
//
// 	it.Quantity1 = fmt.Sprintf("%d", prof.Profile.TotalPieces.Int32)
// 	it.Values.Currency = "USD"
// 	it.Values.ValueOfGoods = prof.Profile.TotalValue.Float64
// 	if it.Values.ValueOfGoods == 0 {
// 		it.Values.ValueOfGoods = 1
// 	}
// 	it.Weight.UOM = "L" // or K
// 	it.Weight.Gross = fmt.Sprintf("%.2f", prof.Profile.TotalWeight.Float64)
// 	it.Parties = append(it.Parties, ABIEntity{
// 		Type:       "seller",
// 		Name:       trim(prof.Shipper.Name.String),
// 		Address:    jmax(fmt.Sprintf("%s %s", trim(prof.Shipper.Address1.String), trim(prof.Shipper.Address2.String)), 35),
// 		Country:    trim(prof.Shipper.Country.String), // to 2 us country code
// 		City:       jmax(trim(prof.Shipper.City.String), 25),
// 	State:      jmax(trim(prof.Shipper.State.String), 9),
// 	PostalCode: jmax(trim(prof.Shipper.Zip.String), 9),
// })
// cs, sz := extractStateZip(prof.Profile.ConState.String)
// it.Parties = append(it.Parties, ABIEntity{
// 	Type:       "consignee",
// 	Name:       cleanup(fmt.Sprintf("%s %s", trim(prof.Profile.Consignee1.String), trim(prof.Profile.Consignee2.String))),
// 	Address:    jmax(fmt.Sprintf("%s %s", trim(prof.Profile.ConStreet1.String), trim(prof.Profile.ConStreet2.String)), 35),
// 	Country:    "US", // trim(prof.Profile.Destination.String), // to 2 us country code
// 	City:       jmax(trim(prof.Profile.ConCity.String), 25),
// 	State:      jmax(cs, 9),
// 	PostalCode: jmax(sz, 9),
// })
//
// it.Parties = append(it.Parties, ABIEntity{
// 	Type:     "buyer",
// 	LoadFrom: "consignee",
// })
// it.Parties = append(it.Parties, ABIEntity{
// 	Type:     "shipTo",
// 				LoadFrom: "consignee",
// 			})
// 			inv.Items = append(inv.Items, it)
// 		}
//
// 		mn.Invoices = append(mn.Invoices, inv)
// 		output = append(output, mn)
// 	}
//
// 	return output, nil
// }

func LoadItem(record *ibc.Hawb, itms *[]Item) {
	// Basically the part below will be added if the hawb line has not goods
	// for now just do hawb line

	var it Item
	if len(record.Goods) == 0 {
		// inv.ExportDate = dateToFormat(prof.Manifest.Date.String, "20060102")
		it.Row = 1
		it.Reference = record.Reference
		// it.CountryOrigin = inv.ExportCountry
		it.CountryOrigin = record.GoodsOrigin
		if it.CountryOrigin == "" {
			it.CountryOrigin = trim(record.ShipperCountry)
		}
		// it.OriginDetail = &Origin{Country: inv.ExportCountry}
		it.OriginDetail = &Origin{Country: it.CountryOrigin}
		it.Description = max(record.Description, 70)
		it.HTS = trim(record.HTS)

		it.Quantity1 = record.Pieces
		it.Values.Currency = record.Currency

		if fvl, err := strconv.ParseFloat(record.Value, 64); err != nil {
			it.Values.ValueOfGoods = 1
		} else {
			it.Values.ValueOfGoods = fvl
			if it.Values.ValueOfGoods == 0 {
				it.Values.ValueOfGoods = 1
			}
		}

		it.Weight.UOM = record.WeightUnit[0:1]
		it.Weight.Gross = record.Weight
		it.Parties = append(it.Parties, ABIEntity{
			Type:       "seller",
			Name:       trim(record.ShipperName),
			Address:    jmax(fmt.Sprintf("%s %s", trim(record.ShipperAddress), trim(record.ShipperAddress2)), 35),
			Country:    trim(record.ShipperCountry), // to 2 us country code
			City:       jmax(trim(record.ShipperCity), 25),
			State:      jmax(trim(record.ShipperState), 9),
			PostalCode: jmax(trim(record.ShipperZip), 9),
		})
		// cs, sz := extractStateZip(prof.Profile.ConState.String)
		it.Parties = append(it.Parties, ABIEntity{
			Type:       "consignee",
			Name:       cleanup(fmt.Sprintf("%s %s", trim(record.ConsigneeName), trim(record.ConsigneeName2))),
			Address:    jmax(fmt.Sprintf("%s %s", trim(record.ConsigneeAddress), trim(record.ConsigneeAddress2)), 35),
			Country:    record.ConsigneeCountry, // to 2 us country code
			City:       jmax(trim(record.ConsigneeCity), 25),
			State:      jmax(record.ConsigneeState, 9),
			PostalCode: jmax(record.ConsigneeZip, 9),
		})

		it.Parties = append(it.Parties, ABIEntity{
			Type:     "buyer",
			LoadFrom: "consignee",
		})
		it.Parties = append(it.Parties, ABIEntity{
			Type:     "shipTo",
			LoadFrom: "consignee",
		})

		*itms = append(*itms, it)
	}

	for i := range record.Goods {

		it.Row = i + 1
		it.Reference = record.Reference
		// it.CountryOrigin = inv.ExportCountry
		it.CountryOrigin = record.Goods[i].OriginCountry
		if it.CountryOrigin == "" {
			it.CountryOrigin = trim(record.ShipperCountry)
		}
		// it.OriginDetail = &Origin{Country: inv.ExportCountry}
		it.OriginDetail = &Origin{Country: it.CountryOrigin}
		it.Description = max(record.Goods[i].Description, 70)
		it.HTS = trim(record.Goods[i].HTS)

		it.Quantity1 = record.Goods[i].Quantity
		it.Values.Currency = record.Goods[i].ItemCurrency
		if it.Values.Currency == "" {
			it.Values.Currency = record.Currency
		}

		if fvl, err := strconv.ParseFloat(record.Goods[i].ItemValue, 64); err != nil {
			it.Values.ValueOfGoods = 1
		} else {
			it.Values.ValueOfGoods = fvl
			if it.Values.ValueOfGoods == 0 {
				it.Values.ValueOfGoods = 1
			}
		}

		if record.Goods[i].ItemWeightUnit != "" {
			it.Weight.UOM = record.Goods[i].ItemWeightUnit[0:1]
		} else {
			it.Weight.UOM = record.WeightUnit[0:1]
		}
		if record.Goods[i].ItemWeight != "" {
			it.Weight.Gross = record.Goods[i].ItemWeight
		} else {
			p1, _ := strconv.ParseFloat(record.Weight, 32)
			pw := p1 / float64(len(record.Goods))

			it.Weight.Gross = fmt.Sprintf("%.2f", pw)
		}

		it.Parties = append(it.Parties, ABIEntity{
			Type:       "seller",
			Name:       trim(record.ShipperName),
			Address:    jmax(fmt.Sprintf("%s %s", trim(record.ShipperAddress), trim(record.ShipperAddress2)), 35),
			Country:    trim(record.ShipperCountry), // to 2 us country code
			City:       jmax(trim(record.ShipperCity), 25),
			State:      jmax(trim(record.ShipperState), 9),
			PostalCode: jmax(trim(record.ShipperZip), 9),
		})
		// cs, sz := extractStateZip(prof.Profile.ConState.String)
		it.Parties = append(it.Parties, ABIEntity{
			Type:       "consignee",
			Name:       cleanup(fmt.Sprintf("%s %s", trim(record.ConsigneeName), trim(record.ConsigneeName2))),
			Address:    jmax(fmt.Sprintf("%s %s", trim(record.ConsigneeAddress), trim(record.ConsigneeAddress2)), 35),
			Country:    record.ConsigneeCountry, // to 2 us country code
			City:       jmax(trim(record.ConsigneeCity), 25),
			State:      jmax(record.ConsigneeState, 9),
			PostalCode: jmax(record.ConsigneeZip, 9),
		})

		it.Parties = append(it.Parties, ABIEntity{
			Type:     "buyer",
			LoadFrom: "consignee",
		})
		it.Parties = append(it.Parties, ABIEntity{
			Type:     "shipTo",
			LoadFrom: "consignee",
		})

		*itms = append(*itms, it)
	}
}

func trim(input string) string {
	if input != "" {
		return strings.TrimSpace(input)
	}
	return input
}

func cleanup(input string) string {
	regex := regexp.MustCompile(`[^a-zA-Z ]+`)
	return trim(regex.ReplaceAllString(input, ""))
}
