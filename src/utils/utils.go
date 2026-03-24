package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type ReqOpts func(*RequestOptions)

type RequestOptions struct {
	Method string
	Debug  bool
}

func WithDebug(dbg bool) ReqOpts {
	return func(opts *RequestOptions) {
		opts.Debug = dbg
	}
}

func WithMethod(mtd string) ReqOpts {
	return func(opts *RequestOptions) {
		opts.Method = mtd
	}
}

func Defaults() RequestOptions {
	return RequestOptions{"POST", false}
}

func DebugDefaults() RequestOptions {
	return RequestOptions{"POST", true}
}

func NewRequestOptions(ofs ...ReqOpts) RequestOptions {
	opts := Defaults()
	for _, fn := range ofs {
		fn(&opts)
	}
	return opts
}

func GetKey() (string, error) {
	tpe := "dev"
	env := os.Getenv("GO_ENV")
	if env == "production" {
		tpe = "prod"
		// tpe = "dev"
	}

	// pth := fmt.Sprintf("/go/dist/secret/key.%s", tpe)
	pth := fmt.Sprintf("secret/key.%s", tpe)

	if content, err := os.ReadFile(pth); err != nil {
		log.Println("issue reading key from file:", pth, err)
		return "", err
	} else {
		return strings.TrimSpace(string(content)), nil
	}
}

func ProcessRequest(c *http.Client, target_url string, uploadJson []byte, target any, opts RequestOptions) error {
	key, err := GetKey()
	if err != nil {
		log.Fatal("could not get key to ProcessRequest", err)
	}
	// data := strings.NewReader(uploadJson)
	data := bytes.NewReader(uploadJson)

	req, err := http.NewRequest(opts.Method, target_url, data)
	if err != nil {
		log.Printf("issue creating request %s cause: %+v\n", string(uploadJson), err)
		return err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", key))
	req.Header.Set("Content-Length", fmt.Sprintf("%d", data.Size()))

	if opts.Debug {
		fmt.Printf("HEADERS: %+v\n", req.Header)
		log.Printf("%s - to %s - [%s]\n", opts.Method, target_url, uploadJson)
	} else {
		log.Printf("%s - to %s", opts.Method, target_url)
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Println("issue processing response", err)
		return err
	}
	defer resp.Body.Close()

	log.Printf("status received: %s", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("could not read body", err)
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		if target != nil {
			err = json.Unmarshal(body, target)
			if err != nil {
				log.Println("could not create error Json body", err)
			}
			if opts.Debug {
				fmt.Printf("----ERROR BODY----\n%s\n",
					string(body))
			}
		}
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)

	} else {
		if target != nil {
			err = json.Unmarshal(body, target)
			if err != nil {
				log.Println("could not create Json body", err)
			}
		}
		if opts.Debug {
			fmt.Printf("Response: %+v\n", string(body))
		}
	}
	return nil
}

func ProcessSimpleRequest(c *http.Client, target_url string, target any, opts RequestOptions) error {
	key, err := GetKey()
	if err != nil {
		log.Fatal("could not get key to ProcessRequest", err)
	}

	req, err := http.NewRequest(opts.Method, target_url, nil)
	if err != nil {
		log.Printf("issue creating request '%s' cause: %+v\n", target_url, err)
		return err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", key))

	if opts.Debug {
		fmt.Printf("HEADERS: %+v\n", req.Header)
	}
	log.Printf("%s - to %s", opts.Method, target_url)

	resp, err := c.Do(req)
	if err != nil {
		log.Println("issue processing response", err)
		return err
	}
	defer resp.Body.Close()

	log.Printf("status received: %s", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("could not read body", err)
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		if target != nil {
			err = json.Unmarshal(body, target)
			if err != nil {
				log.Println("could not create error Json body", err)
			}
			if opts.Debug {
				fmt.Printf("----ERROR BODY----\n%s\n",
					string(body))
			}
		}
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)

	} else {
		if target != nil {
			err = json.Unmarshal(body, target)
			if err != nil {
				log.Println("could not create Json body", err)
			}
		}
		if opts.Debug {
			fmt.Printf("Response: %+v\n", string(body))
		}
	}

	return nil
}
