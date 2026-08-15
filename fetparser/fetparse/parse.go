package fetparse

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type FetTree struct {
	XMLName          xml.Name `xml:"fet"`
	Version          string   `xml:"version,attr"`
	Mode             string
	InstitutionName  string `xml:"Institution_Name"`
	Comments         string
	DaysList         daysList         `xml:"Days_List"`
	HoursList        hoursList        `xml:"Hours_List"`
	SubjectsList     subjectsList     `xml:"Subjects_List"`
	ActivityTagsList activityTagsList `xml:"Activity_Tags_List"`
	TeachersList     teachersList     `xml:"Teachers_List"`
	ClassesList      classesList      `xml:"Students_List"`

	TODO []Other `xml:",any"`

	TimeConstraints  timeConstraints  `xml:"Time_Constraints_List"`
	SpaceConstraints spaceConstraints `xml:"Space_Constraints_List"`
}

type daysList struct {
	NumberOfDays int   `xml:"Number_of_Days"`
	Days         []Day `xml:"Day"`
}

type Item struct {
	Tag  string `xml:"Name"`
	Name string `xml:"Long_Name"`
}

type Day struct {
	Item
}

type hoursList struct {
	NumberOfHours int    `xml:"Number_of_Hours"`
	Hours         []Hour `xml:"Hour"`
}

type Hour struct {
	Item
}

type subjectsList struct {
	Subjects []Subject `xml:"Subject"`
}

type Subject struct {
	Item
	Code     string
	Comments string
}

type activityTagsList struct {
	ActivityTags []ActivityTag `xml:"Activity_Tag"`
}

type ActivityTag struct {
	Item
	Code      string
	Printable bool
	Comments  string
}

type teachersList struct {
	Teachers []Teacher `xml:"Teacher"`
}

type Teacher struct {
	Item
	Code                   string
	Target_Number_of_Hours int
	Qualified_Subjects     []string
	Comments               string
}

type classesList struct {
	Classes []FetClass `xml:"Year"`
}

type FetClass struct {
	Item
	Code             string
	NumberOfStudents int `xml:"Number_of_Students"`
	Comments         string
	// The information regarding categories, divisions of each category, and separator
	// is only used in the divide year automatically by FET categories dialog.
	// NOTE that the FET "Categories" are here called "Divisions", so be careful:
	// These have the field "NumberOfGroups", which corresponds to the FET "Number_of_Divisions"!
	NumberOfDivisions        int        `xml:"Number_of_Categories"`
	Divisions                []division `xml:"Category"`
	FirstDivisionIsPermanent bool       `xml:"First_Category_Is_Permanent"`
	Separator                string
	Groups                   []FetClassGroup `xml:"Group"`
}

type division struct {
	NumberOfGroups int      `xml:"Number_of_Divisions"`
	Groups         []string `xml:"Division"`
}

type FetClassGroup struct {
	Item
	Code             string
	NumberOfStudents int `xml:"Number_of_Students"`
	Comments         string
	Subgroups        []FetClassSubgroup `xml:"Subgroup"`
}

type FetClassSubgroup struct {
	Item
	Code             string
	NumberOfStudents int `xml:"Number_of_Students"`
	Comments         string
}

/*
<!-- The information regarding categories, divisions of each category, and separator is only used in the divide year automatically by categories dialog. -->
      <Number_of_Categories>1</Number_of_Categories>
      <Category>
        <Number_of_Divisions>2</Number_of_Divisions>
        <Division>A</Division>
        <Division>B</Division>
      </Category>
      <First_Category_Is_Permanent>false</First_Category_Is_Permanent>
      <Separator>.</Separator>
      <Group>
        <Name>1.A</Name>
        <Long_Name></Long_Name>
        <Code></Code>
        <Number_of_Students>0</Number_of_Students>
        <Comments></Comments>
        <Subgroup>
          <Name>1#A</Name>
          <Long_Name></Long_Name>
          <Code></Code>
          <Number_of_Students>0</Number_of_Students>
          <Comments></Comments>
        </Subgroup>
      </Group>
      <Group>
        <Name>1.B</Name>
        <Long_Name></Long_Name>
        <Code></Code>
        <Number_of_Students>0</Number_of_Students>
        <Comments></Comments>
        <Subgroup>
          <Name>1#B</Name>
          <Long_Name></Long_Name>
          <Code></Code>
          <Number_of_Students>0</Number_of_Students>
          <Comments></Comments>
        </Subgroup>
      </Group>
    </Year>
*/

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
