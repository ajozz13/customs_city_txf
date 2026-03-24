package ibc

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

type Manifest struct {
	RecordType    string `json:"record_type"`
	RecordVersion int    `json:"record_version"`
	Code          string `json:"manifest_code"`
	Date          string `json:"manifest_date"`
	Origin        string `json:"manifest_origin"`
	Destination   string `json:"manifest_destination"`
	Flight        string `json:"manifest_flight"`
	Mawb          string `json:"master_number"`
	IsExpress     string `json:"-"`
}

func (e *Manifest) Defaults() {
	e.RecordType = "mawb"
	e.RecordVersion = 1
}

func (e *Manifest) ReadLine(input string) error {
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

func (e *Manifest) Load(rc []string) error {
	e.RecordType = rc[0]
	if x, err := strconv.Atoi(rc[1]); err != nil {
		return err
	} else {
		e.RecordVersion = x
	}
	e.Code = rc[2]
	e.Date = rc[3]
	e.Origin = rc[4]
	e.Destination = rc[5]
	e.Flight = rc[6]
	e.Mawb = rc[7]
	return nil
}

func (e *Manifest) ToString() string {
	return fmt.Sprintf("%s [%d] master: '%s' flight: %s [%s to %s] on %s",
		e.RecordType, e.RecordVersion, e.Mawb, e.Flight, e.Origin, e.Destination, e.Date)
}
