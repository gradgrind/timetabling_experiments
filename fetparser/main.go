package main

import (
	"encoding/xml"
	"fetparser/fetparse"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Command: %+v\n", os.Args)
	fpath := os.Args[1]
	fet := fetparse.Parse(fpath)

	fmt.Println("Days:")
	for _, d := range fet.DaysList.Days {
		fmt.Printf("== %+v\n", d)
	}
	fmt.Println("Hours:")
	for _, h := range fet.HoursList.Hours {
		fmt.Printf("== %+v\n", h)
	}
	/*
		fmt.Println("Time Constraints:")
		for _, c := range fet.TimeConstraints.Constraints {
			fmt.Printf("== %+v\n", c)
		}
		fmt.Println("Space Constraints:")
		for _, c := range fet.SpaceConstraints.Constraints {
			fmt.Printf("== %+v\n", c.XMLName.Local)
		}
		fmt.Println("TODO:")
		for _, x := range fet.TODO {
			fmt.Printf("== %+v\n", x)
		}
	*/

	out, err := xml.MarshalIndent(fet, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Create(fpath + ".output")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString(xml.Header)
	_, err = file.Write(out)
	if err != nil {
		fmt.Println(err)
	}
}
