package main

import (
	"fmt"
	"github.com/Station-Manager/config"
	"github.com/Station-Manager/database"
	"github.com/Station-Manager/types"
	"path/filepath"
)

func main() {
	fp, err := filepath.Abs("../build/db/data.db")
	if err != nil {
		panic(err)
	}

	cfg := types.AppConfig{
		DatastoreConfig: types.DatastoreConfig{
			Driver: database.SqliteDriver,
			Path:   fp,
		},
	}
	cfgSvc := &config.Service{AppConfig: cfg}
	if err := cfgSvc.Initialize(); err != nil {
		panic(err)
	}

	svc := &database.Service{ConfigService: cfgSvc}
	if err := svc.Initialize(); err != nil {
		panic(err)
	}
	if err := svc.Open(); err != nil {
		panic(err)
	}
	defer svc.Close()

	typeQso := types.Qso{
		QsoDetails: types.QsoDetails{
			Band:    "20m",
			Freq:    "14.320",
			Mode:    "SSB",
			QsoDate: "20251108",
			TimeOn:  "1140",
			TimeOff: "1146",
			RstRcvd: "59",
			RstSent: "56",
		},
		ContactedStation: types.ContactedStation{Call: "M0CMC", Country: "England"},
		LoggingStation:   types.LoggingStation{StationCallsign: "7Q5MLV", MyCountry: "Mzuzu", MyAntenna: "VHQ Hex Beam"},
	}

	qso, err := svc.InsertQso(typeQso)
	if err != nil {
		fmt.Println("InsertQso returned error:", err)
		// try to get cause via .Cause() if available
		if de, ok := err.(interface{ Cause() error }); ok {
			fmt.Println("Cause:", de.Cause())
		}
		return
	}
	fmt.Println("Inserted qso ID:", qso.ID)
}
