package main

import (
	"encoding/xml"
	"fetparser/fet_parse"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Command: %+v\n", os.Args)
	fpath := os.Args[1]
	fet := fet_parse.Parse(fpath)
	fmt.Println("Time Constraints:")
	for _, c := range fet.Time_Constraints.Constraints {
		fmt.Printf("== %+v\n", c)
	}
	fmt.Println("Space Constraints:")
	for _, c := range fet.Space_Constraints.Constraints {
		fmt.Printf("== %+v\n", c.XMLName.Local)
	}
	fmt.Println("TODO:")
	for _, x := range fet.TODO {
		fmt.Printf("== %+v\n", x)
	}

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
