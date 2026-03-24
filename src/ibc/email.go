package ibc

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

type Email struct {
	RecordType    string `json:"record_type"`
	RecordVersion int    `json:"record_version"`
	NoticeType    string `json:"notice_type"`
	Address       string `json:"email_address"`
}

func (e *Email) Defaults() {
	e.RecordType = "email"
	e.RecordVersion = 1
	e.NoticeType = "all"
}

func (e *Email) ReadLine(input string) error {
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

func (e *Email) Load(rc []string) error {
	e.RecordType = rc[0]
	if x, err := strconv.Atoi(rc[1]); err != nil {
		return err
	} else {
		e.RecordVersion = x
	}
	e.NoticeType = rc[2]
	e.Address = rc[3]
	return nil
}

func (e *Email) ToString() string {
	return fmt.Sprintf("%s [%d] type: '%s' to: %s", e.RecordType, e.RecordVersion, e.NoticeType, e.Address)
}
