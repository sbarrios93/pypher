package license

import (
	"encoding/json"
	"log"
	"os"
)

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id,omitempty"`
}

type Licenses []License

func (l Licenses) NameList() []string {
	var list []string
	for _, license := range l {
		list = append(list, license.Name)
	}
	return list
}

func GetLicenses() *Licenses {
	var licenses Licenses
	licenseFile, errReadFile := os.ReadFile("./api/github/licenses.json")
	if errReadFile != nil {
		log.Fatalf("could not read file: %s", errReadFile)
	}

	errJsonUnmarshal := json.Unmarshal(licenseFile, &licenses)
	if errJsonUnmarshal != nil {
		log.Fatalf("could not unmarshal json file: %s", errJsonUnmarshal)
	}
	return &licenses
}
