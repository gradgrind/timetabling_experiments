package fet_parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type FetTree struct {
	XMLName           xml.Name         `xml:"fet"`
	Version           string           `xml:"version,attr"`
	TODO              []Other          `xml:",any"`
	Time_Constraints  TimeConstraints  `xml:"Time_Constraints_List"`
	Space_Constraints SpaceConstraints `xml:"Space_Constraints_List"`
}

type TimeConstraints struct {
	//	XMLName     xml.Name     `xml:"Time_Constraints_List"`
	Constraints []Constraint `xml:",any"`
}

type SpaceConstraints struct {
	//	XMLName     xml.Name     `xml:"Space_Constraints_List"`
	Constraints []Constraint `xml:",any"`
}

type Constraint struct {
	XMLName xml.Name
	Weight  float64 `xml:"Weight_Percentage"`
	Data    []Other `xml:",any"`
}

type Other struct {
	XMLName xml.Name
	Data    string `xml:",innerxml"`
}

func Parse(fpath string) *FetTree {
	xmlFile, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	fmt.Println("Successfully opened", fpath)
	defer xmlFile.Close()

	var fet FetTree
	err = xml.Unmarshal(byteValue, &fet)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return nil
	}

	return &fet
}
