package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "net/url"
	"os"
	// "sync"
	"time"

	abi "ibcinc.com/customs_city/abi"
	// ibc "ibcinc.com/customs_city/ibc"
	"ibcinc.com/customs_city/utils"
)

type CCService struct {
	Name       string
	Version    string
	URL        string
	AMSURL     string
	AMSSENDURL string
	ABIURL     string
	ABISENDURL string
	PNURL      string
}

func (c *CCService) Defaults() {
	c.Name = "CustomsCity Submission Service"
	c.Version = fmt.Sprintf("0.%s", "1749669293")
	env, ok := os.LookupEnv("GO_ENV")
	if !ok {
		os.Setenv("GO_ENV", "development")
	}
	switch env {
	case "production": // correct this one in the future
		c.URL = "https://api.customscity.com"
		// c.URL = "https://api-cert.customscity.com"

	default: // development
		c.URL = "https://api-cert.customscity.com"
	}
	c.AMSURL = fmt.Sprintf("%s/api/documents", c.URL)
	c.AMSSENDURL = fmt.Sprintf("%s/api/send", c.URL)
	c.ABIURL = fmt.Sprintf("%s/api/abi/documents", c.URL)
	c.ABISENDURL = fmt.Sprintf("%s/api/abi/send", c.URL)
	// c.PNURL = fmt.Sprintf("%s/api/products", c.URL)
}

func (c *CCService) Info() {
	nv := os.Getenv("GO_ENV")
	log.Printf("Running app %s - %s\ntarget: %s\nenvironment: %s\n", c.Name, c.Version, c.ABIURL, nv)
	// fmt.Println("------------------------------------------------------------------")
}

func (c *CCService) ProcessManifest(bdy []abi.ABIBody, submit bool, dbg bool) error {
	ts := time.Now()
	doc := abi.ABIDoc{Type: "abi", Version: "2.1"}
	var err error

	doc.Body = bdy
	// doc.Body, err = abi.BuildABIBody(man)
	// if err != nil {
	// 	return err
	// }

	jsonstr, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	if dbg {
		fmt.Println(string(jsonstr))
	}
	if submit {
		client := &http.Client{}
		var resp abi.CCResponse
		err = utils.ProcessRequest(client, c.ABIURL, jsonstr,
			&resp, utils.DebugDefaults())
		if err != nil {
			resp.PresentErrorMessages()
			return err
		}
		fmt.Println(resp.Message)
		fmt.Printf("%+v\n", resp.Summary)

		mwb := doc.Body[0].Manifests[0].Bill.Master
		fmt.Println("sleep 5 secs...")
		time.Sleep(5 * time.Second)
		fmt.Println("Send manifest")
		err = c.SendABIManifest(client, false, mwb, dbg) // Firts attempt is for AMS
		// err = c.SendABIManifest(client, true, mwb, dbg) // second is for entry summary
		if err != nil {
			return err
		}
	}
	log.Printf("process manifest completed in: %v seconds", time.Since(ts).Seconds())
	return nil
}

func (c *CCService) SendABIManifest(client *http.Client, entry_summary bool, mawb string, dbg bool) error {
	mn := abi.ManifestDoc{Mawb: mawb, SendAll: true}

	url := ""
	if entry_summary {
		mn.Type = "abi"
		mn.Action = "add"
		mn.Application = "entry-summary-cargo-release"
		// mn.Application = "entry-summary"
		url = c.ABISENDURL
		// Above is for cargo release
	} else {
		mn.Type = "air-document"
		mn.SendAs = "add"
		mn.SendType = "air-acas"
		url = c.AMSSENDURL
		// Above is FOR AMS
	}

	json, err := json.Marshal(mn)
	if err != nil {
		return err
	}

	if dbg {
		fmt.Println(string(json))
	}
	var resp abi.CCResponse

	// opts := utils.NewRequestOptions(utils.WithDebug(dbg))

	err = utils.ProcessRequest(client, url, json,
		// err = utils.ProcessRequest(client, c.AMSSENDURL, json,
		&resp, utils.DebugDefaults())
	if err != nil {
		resp.PresentErrorMessages()
		return err
	}
	fmt.Println(resp.Message)
	fmt.Printf("%+v\n", resp.Summary)

	return nil
}

// func (c *CCService) SendPNRequest(client *http.Client, key int, dbg bool) error {
// 	var err error
// 	mn := abi.PNDoc{Type: "fda-pn", Send: true, SendAs: "add"}
//
// 	mn.Body, err = abi.BuildPNBody(key)
// 	if err != nil {
// 		return err
// 	}
//
// 	json, err := json.Marshal(mn)
// 	if err != nil {
// 		return err
// 	}
//
// 	if dbg {
// 		fmt.Println(string(json))
// 	}
//
// 	var resp abi.CCResponse
//
// 	// opts := utils.NewRequestOptions(utils.WithDebug(dbg))
// 	err = utils.ProcessRequest(client, c.AMSURL, json,
// 		&resp, utils.DebugDefaults())
// 	if err != nil {
// 		resp.PresentErrorMessages()
// 		return err
// 	}
// 	fmt.Println(resp.Message)
// 	//
// 	if !mn.Send {
// 		time.Sleep(5 * time.Second)
//
// 		fmt.Println("Send Notice")
// 		err = c.SendPN(client, dbg)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	//
// 	return nil
// }

// func (c *CCService) SendPN(client *http.Client, dbg bool) error {
// 	mn := abi.ManifestDoc{ // for aams
// 		Type:   "fda-pn",
// 		SendAs: "add",
// 		// Mawb:    "93277311264",
// 		// Houses:  &[]string{"11509874"},
// 		SendAll: true,
// 	}
// 	json, err := json.Marshal(mn)
// 	if err != nil {
// 		return err
// 	}
//
// 	if dbg {
// 		fmt.Println(string(json))
// 	}
// 	var resp abi.CCResponse
//
// 	// opts := utils.NewRequestOptions(utils.WithDebug(dbg))
// 	err = utils.ProcessRequest(client, c.AMSSENDURL, json,
// 		&resp, utils.DebugDefaults())
// 	if err != nil {
// 		resp.PresentErrorMessages()
// 		return err
// 	}
// 	fmt.Println(resp.Message)
//
// 	return nil
// }

// func (c *CCService) GetPNNumbers(client *http.Client, key int, dbg bool) error {
// 	// now := time.Now()
// 	// dateString := now.Format("2006-01-02")
// 	dateString := "2025-12-24"
// 	u, err := url.Parse(c.AMSURL)
// 	if err != nil {
// 		return err
// 	}
// 	url, err := abi.BuildPNRequest(key, dateString, u)
// 	if err != nil {
// 		return err
// 	}
// 	var resp abi.PNCResponse
// 	err = utils.ProcessSimpleRequest(client, url, &resp,
// 		utils.NewRequestOptions(utils.WithMethod("GET"), utils.WithDebug(dbg)))
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Println(resp.GetPNCNumbers())
// if len(resp.GetPNCNumbers()) > 0 {
// 	errs := resp.StorePNCNumbers(key)
// 	if len(errs) != 0 {
// 		for _, err := range errs {
// 			fmt.Println("error", err)
// 		}
// 	} else {
// 		fmt.Println("Notices stored!")
// 	}
// }
//
// return nil
// }

// func submitList(input []int, service CCService, s, d bool) {
// 	var wg sync.WaitGroup
// 	for indx, man := range input {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Printf("process item #%d\n", indx)
// 			err := service.ProcessManifest(man, s, d) // typical process for AMS And aCAS only
// 			if err != nil {
// 				fmt.Println("completed %d with error", indx, err)
// 			}
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("list completed.")
// }

func (c *CCService) CancelMawb(client *http.Client, mawb string, dbg bool) error {
	url := fmt.Sprintf("%s?mbol-number=%s", c.ABIURL, mawb)

	var resp abi.CCResponse
	err := utils.ProcessSimpleRequest(client, url, &resp,
		utils.NewRequestOptions(utils.WithMethod("DELETE"), utils.WithDebug(dbg)))
	if err != nil {
		return err
	}

	fmt.Sprintf("Manfiest has been deleted: %+v", resp)

	return nil
}

//func main() {
//	submit := true
//	dbg := true
//	is_entry := true
//	service := CCService{}
//	service.Defaults()
//	service.Info()
//
//	man := 652309          // 648482          // 647607          // 648180
//	mawb := "618-52272802" //"205-84732045" // "020-03005240" //"001-17462664"
//	// below is for pn test
//	key := 250236072
//
//	fmt.Println(man, mawb, key, is_entry, submit)
//
//	err := service.ProcessManifest(man, submit, dbg) // typical process for AMS And aCAS only
//	if err != nil {
//		log.Panic(err)
//	}
//
//	// list := []int{647755, 648608, 648614, 648479}
//	// submitList(list, service, submit, dbg)
//
//	// client := &http.Client{}
//	// err := service.CancelMawb(client, mawb, dbg)
//	// err := service.SendABIManifest(client, is_entry, mawb, dbg)
//	// err := service.SendPNRequest(client, key, dbg)
//	// err := service.GetPNNumbers(client, key, dbg)
//	// if err != nil {
//	// 	log.Panic(err)
//	// }
//}
