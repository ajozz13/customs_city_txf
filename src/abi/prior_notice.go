package abi

import (
	// "fmt"
	"strings"
)

type PNDoc struct {
	Type   string   `json:"type"` // product-fda-pn
	Send   bool     `json:"send"`
	SendAs string   `json:"sendAs"` // add|replace|cancel
	Body   []PNBody `json:"body"`
}

type PNBody struct {
	PNC                *string  `json:"pncNumber"`
	Type               string   `json:"entryType"`
	ReferenceQualifier string   `json:"referenceQualifier"` // AWB
	TransportMode      string   `json:"modeOfTransport"`    // 40
	Reference          string   `json:"referenceNumber"`    // AWB
	EntryNumber        *string  `json:"entryNumber"`
	FTZAdmission       *string  `json:"ftzAdmission"`
	InbondNumber       *string  `json:"inbondNumber"`
	BillType           string   `json:"billType"` // M
	Master             string   `json:"MBOLNumber"`
	House              string   `json:"HBOLNumber"`
	Trip               *string  `json:"trip"`
	BOL                *string  `json:"scnBol"`
	ConsolId           *string  `json:"consolidationId"`
	ExpressTrack       *string  `json:"expressCarrierTrackNumber"`
	ImportingCarrier   *string  `json:"importingCarrier"`
	ArrivalDate        string   `json:"dateOfArrival"`
	ArrivalTime        string   `json:"timeOfArrival"`
	ArrivalPort        string   `json:"portOfArrival"`
	EquipmentNumber    *string  `json:"equipmentNumber"`
	OiDescription      string   `json:"oiDescription"`
	CarrierName        string   `json:"carrierName"`
	VesselName         *string  `json:"vesselName"`
	Voyage             string   `json:"voyageNumber"`
	RailcarNumber      *string  `json:"railCarNumber"`
	Items              []PNItem `json:"items"`
}

// all fei codes null, 16 or 47  16=duns number 47=fei number
// LoadEntries should be null, SELLER or CONSIGNEE
type PNItem struct {
	LineNumber              int              `json:"pgaLineNumber"`
	ProductCode             *string          `json:"productCode"`
	ProductDescription      string           `json:"productDescription"`
	AgencyProcessingCode    string           `json:"agencyProcessingCode"`
	AgencyProgramCode       string           `json:"agencyProgramCode"`
	IntendedUseCode         string           `json:"intendedUseCode"`
	TradeName               string           `json:"tradeNameBrandName"`
	CountryOfShipment       string           `json:"countryOfShipment"`
	FDAProductCode          string           `json:"fdaProductCodeNumber"`
	SourceTypeCode          string           `json:"sourceTypeCode"`
	CountryOfOrigin         string           `json:"countryOfOrigin"`
	Compliance              []ComplianceCode `json:"affirmationOfCompliance,omitempty"`
	Packaging               []PNPackaging    `json:"packaging,omitempty"` // when fda
	ItemDescription         string           `json:"itemDescription"`
	ConsigneeLoadEntry      *string          `json:"ultimateConsigneeLoadEntry,omitempty"`
	ConsigneeName           string           `json:"ultimateConsigneeName"`
	ConsigneeFeiCode        *string          `json:"ultimateConsigneeFeiOrDunsCode"`
	ConsigneeFei            *string          `json:"ultimateConsigneeFeiOrDuns"`
	ConsigneeAddress        string           `json:"ultimateConsigneeAddress"`
	ConsigneeAddress2       string           `json:"ultimateConsigneeAddress2"`
	ConsigneeUnitNo         string           `json:"ultimateConsigneeUnitNumber"`
	ConsigneeCountry        string           `json:"ultimateConsigneeCountry"`
	ConsigneeState          string           `json:"ultimateConsigneeStateOrProvince"`
	ConsigneeCity           string           `json:"ultimateConsigneeCity"`
	ConsigneeZip            string           `json:"ultimateConsigneeZipPostalCode"`
	ConsigneeContactName    string           `json:"ultimateConsigneePointOfContactName"`
	ConsigneePhone          string           `json:"ultimateConsigneePointOfContactPhone"`
	ConsigneeEmail          string           `json:"ultimateConsigneePointOfContactEmail"`
	ShipperLoadEntry        *string          `json:"shipperLoadEntry,omitempty"`
	ShipperName             string           `json:"shipperName"`
	ShipperFeiCode          *string          `json:"shipperFeiOrDunsCode"`
	ShipperFei              *string          `json:"shipperFeiOrDuns"`
	ShipperAddress          string           `json:"shipperAddress"`
	ShipperAddress2         string           `json:"shipperAddress2"`
	ShipperUnitNo           string           `json:"shipperUnitNumber"`
	ShipperCountry          string           `json:"shipperCountry"`
	ShipperState            string           `json:"shipperStateOrProvince"`
	ShipperCity             string           `json:"shipperCity"`
	ShipperZip              string           `json:"shipperZipPostalCode"`
	ShipperContactName      string           `json:"shipperPointOfContactName"`
	ShipperPhone            string           `json:"shipperPointOfContactPhone"`
	ShipperEmail            string           `json:"shipperPointOfContactEmail"`
	SubmitterLoadEntry      *string          `json:"pnSubmitterLoadEntry,omitempty"`
	SubmitterName           string           `json:"pnSubmitterName"`
	SubmitterFeiCode        *string          `json:"pnSubmitterFeiOrDunsCode"`
	SubmitterFei            *string          `json:"pnSubmitterFeiOrDuns"`
	SubmitterAddress        string           `json:"pnSubmitterAddress"`
	SubmitterAddress2       string           `json:"pnSubmitterAddress2"`
	SubmitterUnitNo         string           `json:"pnSubmitterUnitNumber"`
	SubmitterCountry        string           `json:"pnSubmitterCountry"`
	SubmitterState          string           `json:"pnSubmitterStateOrProvince"`
	SubmitterCity           string           `json:"pnSubmitterCity"`
	SubmitterZip            string           `json:"pnSubmitterZipPostalCode"`
	SubmitterContactName    string           `json:"pnSubmitterPointOfContactName"`
	SubmitterPhone          string           `json:"pnSubmitterPointOfContactPhone"`
	SubmitterEmail          *string          `json:"pnSubmitterPointOfContactEmail"`
	TransmitterLoadEntry    *string          `json:"pnTransmitterLoadEntry,omitempty"`
	TransmitterName         string           `json:"pnTransmitterName"`
	TransmitterFeiCode      *string          `json:"pnTransmitterFeiOrDunsCode"`
	TransmitterFei          *string          `json:"pnTransmitterFeiOrDuns"`
	TransmitterAddress      string           `json:"pnTransmitterAddress"`
	TransmitterAddress2     string           `json:"pnTransmitterAddress2"`
	TransmitterUnitNo       string           `json:"pnTransmitterUnitNumber"`
	TransmitterCountry      string           `json:"pnTransmitterCountry"`
	TransmitterState        string           `json:"pnTransmitterStateOrProvince"`
	TransmitterCity         string           `json:"pnTransmitterCity"`
	TransmitterZip          string           `json:"pnTransmitterZipPostalCode"`
	TransmitterContactName  string           `json:"pnTransmitterPointOfContactName"`
	TransmitterPhone        string           `json:"pnTransmitterPointOfContactPhone"`
	TransmitterEmail        string           `json:"pnTransmitterPointOfContactEmail"`
	ManufacturerLoadEntry   *string          `json:"manufacturerOfGoodsLoadEntry,omitempty"`
	ManufacturerName        string           `json:"manufacturerOfGoodsName"`
	ManufacturerFeiCode     *string          `json:"manufacturerOfGoodsFeiOrDunsCode"`
	ManufacturerFei         *string          `json:"manufacturerOfGoodsFeiOrDuns"`
	ManufacturerAddress     string           `json:"manufacturerOfGoodsAddress"`
	ManufacturerAddress2    string           `json:"manufacturerOfGoodsAddress2"`
	ManufacturerUnitNo      string           `json:"manufacturerOfGoodsUnitNumber"`
	ManufacturerCountry     string           `json:"manufacturerOfGoodsCountry"`
	ManufacturerState       string           `json:"manufacturerOfGoodsStateOrProvince"`
	ManufacturerCity        string           `json:"manufacturerOfGoodsCity"`
	ManufacturerZip         string           `json:"manufacturerOfGoodsZipPostalCode"`
	ManufacturerContactName string           `json:"manufacturerOfGoodsPointOfContactName"`
	ManufacturerPhone       string           `json:"manufacturerOfGoodsPointOfContactPhone"`
	ManufacturerEmail       string           `json:"manufacturerOfGoodsPointOfContactEmail"`
	ImporterLoadEntry       *string          `json:"fdaImporterLoadEntry,omitempty"`
	ImporterName            string           `json:"fdaImporterName"`
	ImporterFeiCode         *string          `json:"fdaImporterFeiOrDunsCode"`
	ImporterFei             *string          `json:"fdaImporterFeiOrDuns"`
	ImporterAddress         string           `json:"fdaImporterAddress"`
	ImporterAddress2        string           `json:"fdaImporterAddress2"`
	ImporterUnitNo          string           `json:"fdaImporterUnitNumber"`
	ImporterCountry         string           `json:"fdaImporterCountry"`
	ImporterState           string           `json:"fdaImporterStateOrProvince"`
	ImporterCity            string           `json:"fdaImporterCity"`
	ImporterZip             string           `json:"fdaImporterZipPostalCode"`
	ImporterContactName     string           `json:"fdaImporterPointOfContactName"`
	ImporterPhone           string           `json:"fdaImporterPointOfContactPhone"`
	ImporterEmail           *string          `json:"fdaImporterPointOfContactEmail"`
	OwnerLoadEntry          *string          `json:"ownerLoadEntry,omitempty"`
	OwnerName               string           `json:"ownerName"`
	OwnerFeiCode            *string          `json:"ownerFeiOrDunsCode"`
	OwnerFei                *string          `json:"ownerFeiOrDuns"`
	OwnerAddress            string           `json:"ownerAddress"`
	OwnerAddress2           string           `json:"ownerAddress2"`
	OwnerUnitNo             string           `json:"ownerUnitNumber"`
	OwnerCountry            string           `json:"ownerCountry"`
	OwnerState              string           `json:"ownerStateOrProvince"`
	OwnerCity               string           `json:"ownerCity"`
	OwnerZip                string           `json:"ownerZipPostalCode"`
	OwnerContactName        string           `json:"ownerPointOfContactName"`
	OwnerPhone              string           `json:"ownerPointOfContactPhone"`
	OwnerEmail              string           `json:"ownerPointOfContactEmail"`
	LocationOfGoodsName     string           `json:"locationOfGoodsName"`
	LocationOfGoodsFeiCode  *string          `json:"locationOfGoodsFeiOrDunsCode"`
	LocationOfGoodsFei      *string          `json:"locationOfGoodsFeiOrDuns"`
	LocationOfGoodsAddress  string           `json:"locationOfGoodsAddress"`
	LocationOfGoodsAddress2 string           `json:"locationOfGoodsAddress2"`
	LocationOfGoodsUnitNo   string           `json:"locationOfGoodsUnitNumber"`
	LocationOfGoodsCountry  string           `json:"locationOfGoodsCountry"`
	LocationOfGoodsState    string           `json:"locationOfGoodsStateOrProvince"`
	LocationOfGoodsCity     string           `json:"locationOfGoodsCity"`
	LocationOfGoodsZip      string           `json:"locationOfGoodsZipPostalCode"`
}

type PNPackaging struct {
	Quantity int    `json:"quantity,omitempty"`      //
	UOM      string `json:"unitOfMeasure,omitempty"` //
}

type PNCResponse struct {
	Total int               `json:"total"`
	Skip  int               `json:"skip"`
	Limit int               `json:"limit"`
	Data  []PNCResponseData `json:"data"`
}

type PNCResponseData struct {
	Type            string    `json:"type"`
	Status          string    `json:"status"`
	LastEvent       string    `json:"lastEvent"`
	CreatedAt       string    `json:"createdAt"`
	UpdatedAt       string    `json:"updatedAt"`
	LastEventDt     string    `json:"lastEventDate"`
	PNCNumbers      string    `json:"pncNumber"` // comma separated list
	MBOLNumber      string    `json:"MBOLNumber"`
	HBOLNumber      string    `json:"HBOLNumber"`
	Entry           PNCEntry  `json:"entryType"`
	ReferenceType   string    `json:"referenceQualifier"`
	Transport       PNCEntry  `json:"modeOfTransport"`
	Reference       string    `json:"referenceNumber"`
	BillType        PNCEntry  `json:"billType"`
	ImportCarrier   string    `json:"importingCarrier"`
	ArrivalDate     string    `json:"dateOfArrival"`
	ArrivalTime     string    `json:"timeOfArrival"`
	ArrivalPort     string    `json:"portOfArrival"`
	EquipmentNumber string    `json:"equipmentNumber"`
	Description     string    `json:"oiDescription"`
	CarrierName     string    `json:"carrierName"`
	VesselName      string    `json:"vesselName"`
	VoyageNumber    string    `json:"voyageNumber"`
	RailcarNumber   string    `json:"railCarNumber"`
	Items           []PNCItem `json:"items"`
	IORType         string    `json:"iorType"`
	BondType        string    `json:"bondType"`
}

type PNCEntry struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}

type PNCItem struct {
	PNCNumber         string `json:"pncNumber"`
	LineNumber        int    `json:"pgaLineNumber"`
	ProductCode       string `json:"productCode"`
	Description       string `json:"productDescription"`
	OriginCountry     string `json:"countryOfOrigin"`
	ProductCodeNumber string `json:"productCodeNumber"`
}

func (r *PNCResponseData) GetPNCNumbers() []string {
	return strings.Split(r.PNCNumbers, ",")
}

func (r *PNCResponse) GetPNCNumbers() []string {
	var out []string

	for _, dta := range r.Data {
		out = append(out, dta.GetPNCNumbers()...)
	}
	return out
}

// func (r *PNCResponse) StorePNCNumbers(key int) []error {
// 	var errs []error
// 	if len(r.Data) == 0 {
// 		errs = append(errs, fmt.Errorf("There are no items to store!"))
// 		return errs
// 	}
// 	for _, dta := range r.Data {
// 		ers := StorePN(key, dta.Items)
// 		if len(ers) > 0 {
// 			errs = append(errs, ers...)
// 		}
// 	}
// 	return errs
// }
