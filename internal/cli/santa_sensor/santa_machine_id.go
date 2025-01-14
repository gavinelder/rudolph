package santa_sensor

import (
	"encoding/xml"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

const (
	// This should match the MachineOwnerPlist file, specified in the Santa configuration profile
	defaultSantaMachinePlistFile = "/Library/Preferences/com.google.santa.machine-mapping.plist"
)

type santaPlistXMLDocument struct {
	XMLName xml.Name            `xml:"plist"`
	Data    santaPlistKeyValues `xml:"dict"`
}

type santaPlistKeyValues struct {
	// Annoyingly, because the plist files use <key> and <string> as the respective values in order, =
	// you have to infer which string maps to which key based upon their array positions.
	Keys    []string `xml:"key"`
	Strings []string `xml:"string"`
}

func GetMyMachineUUID() (machineUUID string, err error) {
	// Open our xmlFile
	xmlFile, err := os.Open(defaultSantaMachinePlistFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		err = errors.Wrap(err, "could not open santa machine mapping plist")
		return
	}

	// fmt.Println("Successfully opened santa machine mapping xml")

	// defer the closing of our xmlFile so that we can parse it later on
	defer func() {
		if err := xmlFile.Close(); err != nil {
			log.Fatal("failed to close file")
		}
	}()

	byteValue, _ := io.ReadAll(xmlFile)

	// we initialize our Users array
	var doc santaPlistXMLDocument
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	err = xml.Unmarshal(byteValue, &doc)
	if err != nil {
		return "Unable to unmarshal into users", err
	}

	if len(doc.Data.Keys) != 2 || len(doc.Data.Strings) != 2 {
		err = errors.New("Invalid XML parsed or something")
		return
	}

	for i, thing := range doc.Data.Keys {
		if thing == "MachineUUID" {
			machineUUID = doc.Data.Strings[i]
		}
	}

	if machineUUID == "" {
		err = errors.New("Could not find machineUUID out of XML doc or something")
	}

	return
}
