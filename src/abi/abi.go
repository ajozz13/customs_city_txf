package abi

import "fmt"

type ManifestDoc struct {
	Type        string    `json:"type"`                  // abi    air-document
	Action      string    `json:"action,omitempty"`      // add
	SendAs      string    `json:"sendAs,omitempty"`      // add
	Application string    `json:"application,omitempty"` // entry-summary-cargo-release, entry-summary  cargo-release
	SendType    string    `json:"sendType,omitempty"`    // air|acas|air-acas
	Mawb        string    `json:"MBOLNumber,omitempty"`
	Entries     *[]string `json:"entryNumber,omitempty"`
	Houses      *[]string `json:"HBOLNumber,omitempty"`
	SendAll     bool      `json:"sendAllHBOLS,omitempty"`
}

type ABIDoc struct { // ABI+Air v2
	Type    string    `json:"type"`    // abi
	Version string    `json:"version"` // 2.1
	Body    []ABIBody `json:"body"`
}

func (d *ABIDoc) Defaults() {
	d.Type = "abi"
	d.Version = "2.1"
}

type ABIReference struct {
	Code   string `json:"code"`
	Number string `json:"number"`
}

type ABIEntity struct {
	Type               string `json:"type,omitempty"`
	MID                string `json:"mid,omitempty"`
	LoadFrom           string `json:"loadFrom,omitempty"`
	TaxID              string `json:"taxId,omitempty"`
	Name               string `json:"name,omitempty"`
	Address            string `json:"address,omitempty"`
	UnitNumber         string `json:"unitNumber,omitempty"`
	City               string `json:"city,omitempty"`
	State              string `json:"state,omitempty"`
	PostalCode         string `json:"postalCode,omitempty"`
	Country            string `json:"country,omitempty"`
	Phone              string `json:"telephone,omitempty"`
	Email              string `json:"email,omitempty"`
	FEIDuns            string `json:"feiDuns,omitempty"`
	POC                string `json:"pointOfContact,omitempty"`
	RegistrationType   string `json:"registrationNumberType,omitempty"`
	RegistrationNumber string `json:"registrationNumber,omitempty"`
}

type ABIPayment struct {
	Type           int    `json:"typeCode,omitempty"`
	StatementDate  string `json:"preliminaryStatementDate,omitempty"` //
	StatementMonth string `json:"periodicStatementMonth,omitempty"`   //
}

type ABIBond struct {
	Type   string `json:"type,omitempty"`
	Code   string `json:"suretyCode,omitempty"` //
	Amount int    `json:"amount,omitempty"`     //
}

type IOR struct {
	Number string `json:"number"`
	Name   string `json:"name"` //
}

type ABILocation struct {
	EntryPort        string `json:"portOfEntry"`
	DestinationState string `json:"destinationStateUS"` //
}

type ABIDates struct {
	Entry       string `json:"entryDate,omitempty"`   // YYYYMMDD
	Import      string `json:"importDate,omitempty"`  // YYYYMMDD
	Arrival     string `json:"arrivalDate,omitempty"` // YYYYMMDD
	ArrivalTime string `json:"arrivalTime,omitempty"` // HHMM
}

type ABIBody struct {
	EntryType            string         `json:"entryType"`               // 11
	CustomerID           string         `json:"customerId,omitempty"`    //
	FileNumber           string         `json:"fileNumber,omitempty"`    //
	Version              string         `json:"version,omitempty"`       //
	Reference            string         `json:"internalReferenceNumber"` // serialKey
	TransportMode        string         `json:"modeOfTransport"`         // 40
	ABIDates             ABIDates       `json:"dates"`
	ABILocation          ABILocation    `json:"location"`
	IOR                  IOR            `json:"ior"`
	Bond                 ABIBond        `json:"bond"`
	Payment              ABIPayment     `json:"payment"`
	Firms                string         `json:"firms"`                                // when fda
	ExamSite             string         `json:"electedExamSite,omitempty"`            //
	IsKnownImporter      string         `json:"knownImporter,omitempty"`              // Y\N  when fda
	IsPerishable         string         `json:"perishableGoods,omitempty"`            // Y\N  when fda
	IsNonAMS             string         `json:"nonAMSIndicator,omitempty"`            // Y\N  when fda
	IsExpress            string         `json:"expressConsignmentShipment,omitempty"` // Y\N  when fda
	EntryConsignee       *ABIEntity     `json:"entryConsignee,omitempty"`             // when fda
	AdditionalReferences []ABIReference `json:"additionalReferences,omitempty"`
	Manifests            []ABIManifest  `json:"manifest"`
}

type Bill struct {
	Type    string `json:"type"`                              // M
	Master  string `json:"mBOL"`                              //
	House   string `json:"hBOL"`                              //
	Marks   string `json:"manifestMarksAndNumbers,omitempty"` //
	OnlyAMS string `json:"itemConsigneeOnlyForAMS,omitempty"` // Y|N
	Group   string `json:"groupBOL,omitempty"`                // Y|N  when fda
}

type Carrier struct {
	Code   string `json:"code"`                 //
	Vessel string `json:"vesselName,omitempty"` //
	Flag   string `json:"vesselFlag,omitempty"` //
}

type Ports struct {
	Unlading         string `json:"portOfUnlading,omitempty"`          //
	Origin           string `json:"airportOfOrigin,omitempty"`         //
	Flight           string `json:"airFlightNumber,omitempty"`         //
	Operator         string `json:"airTerminalOperator,omitempty"`     //
	Foreign          string `json:"lastForeignPort,omitempty"`         //
	ReceiptByCarrier string `json:"placeOfReceiptByCarrier,omitempty"` //
}

type SplitDetail struct {
	Flight      string `json:"voyageFlightTrip,omitempty"`
	ArrivalDate string `json:"dateOfArrival,omitempty"`
	Quantity    string `json:"quantity,omitempty"`
	ReleaseCode string `json:"releaseCode,omitempty"`
}

type AccountDetail struct {
	Holder      string `json:"holder,omitempty"`       //
	Name        string `json:"name,omitempty"`         //
	Issuer      string `json:"issuer,omitempty"`       //
	Number      string `json:"number,omitempty"`       //
	Frequency   string `json:"shippingFreq,omitempty"` //
	Email       string `json:"email,omitempty"`        //
	CreatedDate string `json:"creationDate,omitempty"` //
	BillingType string `json:"billingType,omitempty"`  //
	IP          string `json:"ipAddress,omitempty"`    //
	URL         string `json:"url,omitempty"`          //
}

type ACASDetail struct {
	Account         *AccountDetail `json:"customerAccount,omitempty"`
	IsKnown         string         `json:"verifiedKnownConsignor,omitempty"` // YES|NO
	BioData         string         `json:"biographicData,omitempty"`         //
	TransactionType string         `json:"transactionType,omitempty"`        //
	FilingType      string         `json:"filingType,omitempty"`             //
	PartyURL        string         `json:"partyUrl,omitempty"`               //
}

type Package struct {
	Quantity      string `json:"quantity,omitempty"`
	UOM           string `json:"quantityUOM,omitempty"` // PCS
	Description   string `json:"cargoDescription,omitempty"`
	OriginCountry string `json:"countryOfOrigin,omitempty"`
	IsFDA         bool   `json:"fdaIndicator,omitempty"`
}

type Charges struct {
	Deductible int `json:"totalDeductible,omitempty"`
	Additional int `json:"totalAdditional,omitempty"`
}
type Origin struct {
	Country string `json:"country"`         //
	State   string `json:"state,omitempty"` //
}

type Values struct {
	Currency             string  `json:"currency,omitempty"`                  //
	ExchangeRate         int     `json:"exchangeRate,omitempty"`              //
	InvoiceCurrency      int     `json:"totalValueInvoiceCurrency,omitempty"` //
	ValueInvoice         int     `json:"totalValueInvoice,omitempty"`         //
	DDP                  int     `json:"ddp,omitempty"`                       //
	ValueOfGoods         float64 `json:"totalValueOfGoods,omitempty"`         //
	ValueOfGoodsCurrency string  `json:"totalValueOfGoodsCurrency,omitempty"` //
	Charges              int     `json:"chargesAmount,omitempty"`             //
	ChargesCurrency      string  `json:"chargesAmountCurrency,omitempty"`     //
}

type Weight struct {
	Gross string `json:"gross,omitempty"` //
	UOM   string `json:"uom,omitempty"`   //
}

type PGAPackaging struct {
	Quantity string `json:"quantity,omitempty"` //
	UOM      string `json:"uom,omitempty"`      //
}

type ComplianceCode struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}

type LCPO struct {
	Type   int    `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

type FDA struct {
	ProgramCode       string           `json:"programCode,omitempty"`
	ProcessingCode    string           `json:"processingCode,omitempty"`
	DisclaimCode      string           `json:"disclaimCode,omitempty"`
	IntendedUseCode   string           `json:"intendedUseCode,omitempty"`
	ProductCode       string           `json:"productCodeNumber,omitempty"`
	BrandName         string           `json:"tradeNameBrandName,omitempty"`
	ItemType          string           `json:"itemType,omitempty"`
	SourceType        string           `json:"sourceTypeCode,omitempty"`
	Remarks           string           `json:"remarks,omitempty"`
	Compliance        []ComplianceCode `json:"affirmationOfCompliance,omitempty"`
	PriorNoticeNumber string           `json:"priorNoticeConfirmationNumber,omitempty"`
	LCPO              *LCPO            `json:"lcpo,omitempty"`
}

type EPA struct {
	ProgramCode            string        `json:"programCode,omitempty"`
	ProcessingCode         string        `json:"processingCode,omitempty"`
	DisclaimCode           string        `json:"disclaimCode,omitempty"`
	IntendedUseCode        string        `json:"intendedUseCode,omitempty"`
	IntendedUseDescription string        `json:"intendedUseDescription,omitempty"`
	ImageSubmitted         string        `json:"electronicImageSubmitted,omitempty"` // Y|N
	Producer               string        `json:"producerEstablishmentNumber,omitempty"`
	ItemType               string        `json:"itemType,omitempty"`
	Model                  string        `json:"model,omitempty"`
	ManufactureDate        string        `json:"manufactureDate,omitempty"`
	Identity               Identity      `json:"identity,omitempty"`
	Commodity              Commodity     `json:"commodity,omitempty"`
	Remarks                Remarks       `json:"remarks,omitempty"`
	Measures               Measures      `json:"measures,omitempty"`
	Documentation          Documentation `json:"documentation,omitempty"`
}

type AMS struct {
	ProgramCode            string    `json:"programCode,omitempty"`
	ProcessingCode         string    `json:"processingCode,omitempty"`
	DisclaimCode           string    `json:"disclaimCode,omitempty"`
	ImageSubmitted         string    `json:"electronicImageSubmitted,omitempty"` // Y|N
	IntendedUseCode        string    `json:"intendedUseCode,omitempty"`
	IntendedUseDescription string    `json:"intendedUseDescription,omitempty"`
	ProductCode            string    `json:"productCodeQualifier,omitempty"`
	CodeNumber             string    `json:"productCodeNumber,omitempty"`
	Value                  string    `json:"lineValue,omitempty"`
	UnitValue              string    `json:"unitValue,omitempty"`
	TestStatus             string    `json:"inspectionLaboratoryTestStatus,omitempty"`
	TempQualifier          string    `json:"temperatureQualifier,omitempty"`
	Identity               Identity  `json:"identity,omitempty"`
	Container              Container `json:"container,omitempty"`
}

type Container struct {
	Number string `json:"number,omitempty"` //
	Type   string `json:"type,omitempty"`   //
	Length string `json:"length,omitempty"` //
}

type Identity struct {
	Qualifier string   `json:"numberQualifier,omitempty"` //
	Numbers   []string `json:"numbers"`                   //
}

type Commodity struct {
	Code        string `json:"qualifierCode,omitempty"`             //
	Qualifier   string `json:"characteristicsQualifier,omitempty"`  //
	Description string `json:"characteristicDescription,omitempty"` //
}

type CommodityGroup struct {
	Code        []string `json:"qualifierCode,omitempty"`             //
	Qualifier   []string `json:"characteristicsQualifier,omitempty"`  //
	Description []string `json:"characteristicDescription,omitempty"` //
}

type Remarks struct {
	Type string `json:"type,omitempty"` //
	Code string `json:"code,omitempty"` //
	Text string `json:"text,omitempty"` //
}

type Measures struct {
	UOM      string `json:"grossNetUOM,omitempty"`      //
	Quantity string `json:"grossNetQuantity,omitempty"` //
}

type Documentation struct {
	Type               string `json:"lpcoType,omitempty"`           //
	Number             string `json:"lpcoNumber,omitempty"`         //
	RegistrationNumber string `json:"registrationNumber,omitempty"` //
	DocID              string `json:"documentId,omitempty"`         //
	DeclarationCode    string `json:"declarationCode,omitempty"`    //
	Date               string `json:"dateOfSignature,omitempty"`    //
}

type Declaration struct {
	Code          string `json:"code,omitempty"`          //
	Certification string `json:"certification,omitempty"` //
}

type POC struct {
	Qualifier string `json:"individualQualifier,omitempty"` //
	Name      string `json:"individualName,omitempty"`      //
	Phone     string `json:"telephone,omitempty"`           //
	Email     string `json:"email,omitempty"`               //
}

type TSCA struct {
	ItemType       string      `json:"itemType,omitempty"`       //
	Code           string      `json:"programCode,omitempty"`    //
	ProcessingCode string      `json:"processingCode,omitempty"` //
	DisclaimCode   string      `json:"disclaimCode,omitempty"`   //
	Declaration    Declaration `json:"declaration,omitempty"`
	EntityRole     string      `json:"entityRoleCode,omitempty"` //
	POC            POC         `json:"poc,omitempty"`
}

type APHIS struct {
	ProgramCode            string         `json:"programCode,omitempty"`
	ProcessingCode         string         `json:"processingCode,omitempty"`
	DisclaimCode           string         `json:"disclaimCode,omitempty"`
	ImageSubmitted         string         `json:"electronicImageSubmitted,omitempty"` // Y|N
	IntendedUseCode        string         `json:"intendedUseCode,omitempty"`
	IntendedUseDescription string         `json:"intendedUseDescription,omitempty"`
	ProductCode            string         `json:"productCodeQualifier,omitempty"`
	CodeNumber             string         `json:"productCodeNumber,omitempty"`
	ItemType               string         `json:"itemType,omitempty"`
	GenusName              string         `json:"genusName,omitempty"`
	SpeciesName            string         `json:"speciesName,omitempty"`
	SubSpeciesName         string         `json:"subSpeciesName,omitempty"`
	SourceType             string         `json:"sourceTypeCode,omitempty"`
	CountryCode            string         `json:"countryCode,omitempty"`
	GeoLocation            string         `json:"geographicLocation,omitempty"`
	Processing             Processing     `json:"processing,omitempty"`
	RoutingName            string         `json:"politicalSubmitOfRoutingName,omitempty"`
	PG07                   PGDetail       `json:"pg07,omitempty"`
	PG08                   PGDetail       `json:"pg08,omitempty"`
	Category               CategoryDetail `json:"category,omitempty"`
	Commodity              CommodityGroup `json:"commdotiy,omitempty"`
	LPCO                   LPCO           `json:"lpco,omitempty"`
	CommonName             []string       `json:"commonName,omitempty"`
	Remarks                RemarksGroup   `json:"remarks,omitempty"`
}

type LPCO struct {
	Code                []string `json:"govermentGeographicCodeQualifier,omitempty"` //
	Location            []string `json:"locationOfIssuer,omitempty"`                 //
	LocationDescription []string `json:"regionalDescriptionOfLocation,omitempty"`    //
	Type                []string `json:"type,omitempty"`                             //
	Number              []string `json:"number,omitempty"`                           //
	DateQualifier       []string `json:"dateQualifier,omitempty"`                    //
	Date                []string `json:"date,omitempty"`                             //
	Quantity            []string `json:"quantity,omitempty"`                         //
	UOM                 []string `json:"uom,omitempty"`                              //
}

type RemarksGroup struct {
	Type []string `json:"typeCode,omitempty"` //
	Code []string `json:"code,omitempty"`     //
	Text []string `json:"text,omitempty"`     //
}

type PGDetail struct {
	Qualifier []string `json:"numberQualifier,omitempty"`
	Number    []string `json:"number"`
}

type CategoryDetail struct {
	Code []string `json:"typeCode,omitempty"`
	Type []string `json:"type"`
}

type Processing struct {
	StartDate   string `json:"startDate,omitempty"`   // YYYYMMDD
	EndDate     string `json:"endDate,omitempty"`     // YYYYMMDD
	Code        string `json:"typeCode,omitempty"`    //
	Description string `json:"description,omitempty"` //
}

type PGAData struct {
	FDA   *FDA   `json:"fda,omitempty"`
	EPA   *EPA   `json:"epa,omitempty"`
	TSCA  *TSCA  `json:"epaTsca,omitempty"`
	APHIS *APHIS `json:"aphis,omitempty"`
	AMS   *AMS   `json:"ams,omitempty"`
}

type Item struct {
	Row                       int            `json:"row,omitempty"`                           //
	ProductCode               string         `json:"productCode,omitempty"`                   //
	SKU                       string         `json:"sku,omitempty"`                           //
	Reference                 string         `json:"productReference,omitempty"`              //
	HTS                       string         `json:"htsNumber,omitempty"`                     //
	Description               string         `json:"description"`                             //
	CountryOrigin             string         `json:"countryOfOrigin,omitempty"`               //
	OriginDetail              *Origin        `json:"origin,omitempty"`                        //
	ProgramCode               string         `json:"specialProgramCode,omitempty"`            //
	Values                    Values         `json:"values,omitempty"`                        //
	Quantity1                 string         `json:"quantity1,omitempty"`                     //
	Quantity2                 string         `json:"quantity2,omitempty"`                     //
	Quantity3                 string         `json:"quantity3,omitempty"`                     //
	LCPO                      string         `json:"lcpoNumber,omitempty"`                    //
	PrimarySmeltCountry       string         `json:"aluminumPrimarySmeltCountry,omitempty"`   //
	SecondarySmeltCountry     string         `json:"aluminumSecondarySmeltCountry,omitempty"` //
	CastCountry               string         `json:"aluminumCastCountry,omitempty"`           //
	SteelMeltPourCountry      string         `json:"steelMeltPourContry,omitempty"`           //
	SteelMeltAppCode          string         `json:"steelMeltPourApplicationCode,omitempty"`  //
	Weight                    Weight         `json:"weight,omitempty"`
	ProductURL                string         `json:"productUrl,omitempty"`
	AluminumPct               *int           `json:"aluminumPercentage,omitzero"`
	SteelPct                  *int           `json:"steelPercentage,omitzero"`
	CopperPct                 *int           `json:"copperPercentage,omitzero"`
	HasCottonExemption        string         `json:"cottonFeeExemption,omitempty"`
	HasAutoPartsExemption     string         `json:"autoPartsExemption,omitempty"`
	HasCompletedKitchenParts  string         `json:"otherThanCompletedKitchenParts,omitempty"`
	HasInformationalExemption string         `json:"informationalMaterialsExemption,omitempty"`
	ForReligiousPurpose       string         `json:"religiousPurposes,omitempty"`
	HasAgriculturalExemption  string         `json:"agriculturalExemption,omitempty"`
	SemiConductorExemption    *int           `json:"semiConductorExemption,omitempty"`
	Parties                   []ABIEntity    `json:"parties"`
	PGAPackaging              []PGAPackaging `json:"pgaPackaging,omitempty"` // when fda
	PGAData                   *PGAData       `json:"pgaData,omitempty"`      // when fda
}

type Invoice struct {
	PurchaseOrder  string   `json:"purchaseOrder"`
	InvoiceNumber  string   `json:"invoiceNumber,omitempty"` // when fda
	Package        Package  `json:"package"`
	ExportDate     string   `json:"exportDate"`
	LadingPort     string   `json:"portOfLading,omitempty"`
	RelatedParties string   `json:"relatedParties"`            // Y\N
	ExportCountry  string   `json:"countryOfExport"`           //
	Currency       string   `json:"currency"`                  //
	ExchangeRate   int      `json:"exchangeRate,omitempty"`    // 1
	Charges        *Charges `json:"totalCharges,omitempty"`    //
	TrackNumber    string   `json:"trackingNumber,omitempty"`  //
	MarksNumber    string   `json:"MarksAndNumbers,omitempty"` //
	Items          []Item   `json:"items"`                     //
}

type ABIManifest struct {
	Bill       Bill          `json:"bill"`
	Carrier    Carrier       `json:"carrier"`
	Ports      Ports         `json:"ports"`
	SNP        string        `json:"snp,omitempty"`         //
	Quantity   string        `json:"quantity,omitempty"`    //
	UOM        string        `json:"quantityUOM,omitempty"` //
	Splits     []SplitDetail `json:"split,omitempty"`
	AcasDetail *ACASDetail   `json:"acas,omitempty"`
	Invoices   []Invoice     `json:"invoices,omitempty"`
}

// Responses
type CCResponse struct {
	Name      string      `json:"name"`
	Code      int         `json:"code,omitempty,string"`
	ClassName string      `json:"className"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Summary   struct {
		DocumentsDeleted int               `json:"documentsDeleted"`
		TotalDeleted     int               `json:"totalItemsDeleted"`
		Documents        []DocumentDetails `json:"documents"`
		Totalrecords     int               `json:"total_records"`
		Saved            int               `json:"saved"`
		Failed           int               `json:"failed"`
	} `json:"summary"`
	Errors struct {
		Entry1 struct {
			Entry     []string            `json:"entry"`
			Manifests map[string][]string `json:"manifests"`
		} `json:"Entry: 1"`
	} `json:"errors"`
}

type DocumentDetails struct {
	EntryNumber  string `json:"entryNumber"`
	Master       string `json:"mbolNumber"`
	ItemsDeleted int    `json:"itemsDeleted"`
}

func (resp *CCResponse) PresentErrorMessages() {
	fmt.Println(resp.Message)
	fmt.Printf("%+v\n", resp.Summary)

	if len(resp.Errors.Entry1.Entry) > 0 {
		reportArray("Entry", resp.Errors.Entry1.Entry)
	}

	if len(resp.Errors.Entry1.Manifests) > 0 {
		for nm, msgs := range resp.Errors.Entry1.Manifests {
			reportArray(nm, msgs)
		}
	}
}

func reportArray(name string, items []string) {
	fmt.Println("-", name)
	for j, mx := range items {
		fmt.Printf("--%d - %v\n", j, mx)
	}
	fmt.Println("--------------------------------------------")
}
