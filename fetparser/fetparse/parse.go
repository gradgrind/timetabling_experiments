package fetparse

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type FetTree struct {
	XMLName         xml.Name `xml:"fet"`
	Version         string   `xml:"version,attr"`
	Mode            string
	InstitutionName string `xml:"Institution_Name"`
	Comments        string
	DaysList        daysList  `xml:"Days_List"`
	HoursList       hoursList `xml:"Hours_List"`

	TODO             []Other          `xml:",any"`
	TimeConstraints  timeConstraints  `xml:"Time_Constraints_List"`
	SpaceConstraints spaceConstraints `xml:"Space_Constraints_List"`
}

type daysList struct {
	NumberOfDays int   `xml:"Number_of_Days"`
	Days         []Day `xml:"Day"`
}

type Day struct {
	Tag  string `xml:"Name"`
	Name string `xml:"Long_Name"`
}

type hoursList struct {
	NumberOfHours int    `xml:"Number_of_Hours"`
	Hours         []Hour `xml:"Hour"`
}

type Hour struct {
	Tag  string `xml:"Name"`
	Name string `xml:"Long_Name"`
}

type timeConstraints struct {
	//	XMLName     xml.Name     `xml:"Time_Constraints_List"`
	Constraints []Constraint `xml:",any"`
}

type spaceConstraints struct {
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

	fet := &FetTree{}
	err = xml.Unmarshal(byteValue, fet)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return nil
	}

	return fet
}
