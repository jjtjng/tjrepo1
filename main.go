package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*
Output format
{
    "vendorDetails":{
        "oui":"443839",
        "isPrivate":false,
        "companyName":"Cumulus Networks, Inc",
        "companyAddress":"650 Castro Street, suite 120-245 Mountain View  CA  94041 US",
        "countryCode":"US"
    },
    "blockDetails":{
        "blockFound":true,
        "borderLeft":"443839000000",
        "borderRight":"443839FFFFFF",
        "blockSize":16777216,
        "assignmentBlockSize":"MA-L",
        "dateCreated":"2012-04-08",
        "dateUpdated":"2015-09-27"
    },
    "macAddressDetails":{
        "searchTerm":"44:38:39:ff:ef:57",
        "isValid":true,
        "virtualMachine":"Not detected",
        "applications":[
            "Multi-Chassis Link Aggregation (Cumulus Linux)"
        ],
        "transmissionType":"unicast",
        "administrationType":"UAA",
        "wiresharkNotes":"No details",
        "comment":""
    }
}
*/


type VendorDetails struct {
        Oui string  `json:"oui"`
        IsPrivate bool `json:"isPrivate"`
        CompanyName string `json:"companyName"`
        CompanyAddress string `json:"companyAddress"`
	CountryCode string `json:"countryCode"`
}
type BlockDetails struct {
        BlockFound bool `json:"blockFound"`
        BorderLeft string `json:"borderLeft"`
        BorderRight string `json:"borderRight"`
        BlockSize int64 `json:"blockSize"`
        AssignmentBlockSize string `json:"assignmentBlockSize"`
        DateCreated string `json:"dateCreated"`
        DateUpdated string  `json:"dateUpdated"`
}
type MacAddressDetails struct {
        SearchTerm string `json:"searchTerm"`
        IsValid bool `json:"isValid"`
        VirtualMachine string `json:"virtualMachine"`
        Applications []string `json:"applications"`
        TransmissionType string `json:"transmissionType"`
        AdministrationType string `json:"administrationType"`
        WiresharkNotes string `json:"wiresharkNotes"`
        Comment string `json:"comment"`
}

type MacInfo struct {
	VendorDetails VendorDetails `json:"vendorDetails"`
	BlockDetails BlockDetails `json:"blockDetails"`
	MacAddressDetails MacAddressDetails `json:"macAddressDetails"`
}

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("Usage %s <mac address>\n", os.Args[0])
		os.Exit(1)
	}

	movie := args[0]
	lookupMac(movie)
}

func lookupMac(movie string) {

	fmt.Printf("Looking up %s\n", movie)

	key := os.Getenv("MACADDRESS_IO_API_KEY")
	if key == "" {
		log.Fatalf("MACADDRESS_IO_API_KEY env-var not defined")
	}

	url := fmt.Sprintf("https://api.macaddress.io/v1?apiKey=%s&output=json&search=%s", key, movie)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to query %v\n", err)
	}
	defer resp.Body.Close()

	jsondata, _ := ioutil.ReadAll(resp.Body)

	var data MacInfo

	if err := json.Unmarshal(jsondata, &data); err != nil {
		log.Fatalf("Failed to unmarshall %v\n", err)
	}

	//fmt.Printf("MacInfo: %v\n", data)
	fmt.Printf("Company name: %s\n", data.VendorDetails.CompanyName)
}
