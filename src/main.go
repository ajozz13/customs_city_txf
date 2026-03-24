package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	abi "ibcinc.com/customs_city/abi"
	ibc "ibcinc.com/customs_city/ibc"
)

func Demo() {
	fmt.Println("Ok go to work")

	el := `email,1,all,ajozz13@gmail.com`
	em := ibc.Email{}

	err := em.ReadLine(el)
	if err != nil {
		panic(err)
	}
	fmt.Println(em.ToString())

	ml := `mawb,1,1232dr,20260301,HKG,JFK,AA20264,235-31567815`
	mx := ibc.Manifest{}

	err = mx.ReadLine(ml)
	if err != nil {
		panic(err)
	}
	fmt.Println(mx.ToString())
}

func WriteDemo() {
	file, err := prepFile("out/20646564.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wrtr := csv.NewWriter(file)

	rcrd := make([]string, 650)
	rcrd[0] = "11"
	rcrd[4] = "20260319"
	rcrd[5] = "20260319"
	rcrd[6] = "20260319"
	rcrd[10] = "20656464"
	rcrd[30] = "XXXXXX"
	rcrd[59] = "YYYYYY"

	if err = wrtr.Write(rcrd); err != nil {
		log.Fatalf("could not write to file")
	}

	wrtr.Flush()

	if err = wrtr.Error(); err != nil {
		panic(err)
	}

	fmt.Println("done writing file")
}

func ReadDemo(inputf string) error {
	fmt.Println("read input file: ", inputf)
	file, err := os.Open(inputf)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // disable consistent field check
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("read record ", record[0])
	}
	return nil
}

func prepFile(input string) (*os.File, error) {
	err := Copy("template/header.csv", input)
	if err != nil {
		return nil, err
	}

	of, err := os.OpenFile(input, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return of, nil
}

func clean(fn string) error {
	if FileExists(fn) {
		if err := os.Remove(fn); err != nil {
			log.Printf("file could not be removed! %s ", fn)
			return err
		}
	}
	return nil
}

func ProcessManifestFile(inputf string, is_express bool) error {
	fmt.Println("read input file: ", inputf)
	file, err := os.Open(inputf)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // disable consistent field check
	total_lines := 0
	last_ln := 0
	// var of *os.File
	var bll ibc.Hawb
	var mwb ibc.Manifest
	var obdy []abi.ABIBody
	// var ofn string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if last_ln == 0 {
			// remvove Byte order mark
			record[0] = strings.TrimPrefix(record[0], "\ufeff")
		}

		switch record[0] {
		case "email":
			fmt.Println("ignore email")
		case "mawb":
			mwb = ibc.Manifest{}
			if err := mwb.Load(record); err != nil {
				log.Fatalf("manifest could not be loaded correctly %+v", err)
			}
			mwb.IsExpress = "N"
			if is_express {
				mwb.IsExpress = "Y"
			}

			//
			fmt.Println("load manifest", record[7])
			// ofn = fmt.Sprintf("out/%s.csv", mwb.Mawb)
			// fmt.Print("Prepare of....")
			// of, err = prepFile(ofn)
			// if err != nil {
			// panic(err)
			// }
			// defer of.Close()
			// fmt.Println("OF HAS BEEN CREATED")
			// fmt.Println("OF", of)
			obdy = abi.BuildABIBody(mwb)
			obdy[0].Manifests = make([]abi.ABIManifest, 0)

		case "hawb":
			// if of == nil {
			// panic("No mawb line detected")
			// }
			if last_ln < total_lines {
				last_ln = total_lines
				fmt.Println("CALL DO OUTPUT NOW\n", "#", total_lines)
				fmt.Printf("with total goods %d\n", len(bll.Goods))
				// err := bll.Write(of, mwb)
				// if err != nil {
				// log.Printf("issue writing ln: %d track: %s cause: %+v", last_ln, bll.House, err)
				// if ferr := clean(ofn); ferr != nil {
				// log.Printf("output file could not be removed %s", ofn)
				// }

				// log.Fatalf("cause: %+v", err)
				// }

				abi.BuildManifest(&obdy[0].Manifests, mwb, &bll)

				bll = ibc.Hawb{}
			}

			// fmt.Println("process ", record[0], record[3])
			if err := bll.Load(record); err != nil {
				fmt.Println(err)
			}
			fmt.Println("loaded", bll.RecordType, bll.House)
			total_lines = total_lines + 1

		case "good", "goods", "commodity":
			var gd ibc.Commodity
			if err := gd.Load(record); err != nil {
				fmt.Println(err)
			}
			fmt.Println("process ", gd.RecordType)
			bll.AddGood(gd)

		default:
			fmt.Println("skip record ", record[0])

		}
	}
	fmt.Println(last_ln, total_lines)
	if last_ln < total_lines {
		fmt.Println("CALL DO OUTPUT NOW\n", "#", last_ln+1)
		// err := bll.Write(of, mwb)
		// if err != nil {
		// fmt.Println(err)
		// }
		abi.BuildManifest(&obdy[0].Manifests, mwb, &bll)
	}
	// fmt.Printf( "%+v\n", obdy )

	// send manifest
	service := CCService{}
	service.Defaults()
	service.Info()

	dbg := true
	auto_submit := false

	err = service.ProcessManifest(obdy, auto_submit, dbg)
	if err != nil {
		log.Fatalf("could not submit cause: %+v", err)
	}
	return nil
}

func main() {
	start := time.Now()
	// Demo()
	// WriteDemo()
	//
	//
	//ONLIVE MODE TURN ON auto_submit above

	fmt.Println(len(os.Args))
	if len(os.Args) != 3 {
		log.Fatalf("usage: cmd path_to_input_file is_express")
	}
	fn := os.Args[1]

	is_express := strings.ToLower(os.Args[2]) == "t"

	fmt.Printf("process file %s is_express? %v\n", fn, is_express)

	if err := ProcessManifestFile(fn, is_express); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("completed in %.3f seconds\n", time.Since(start).Seconds())
	// time.Sleep(3 * time.Second)
}
