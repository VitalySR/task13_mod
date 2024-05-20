package entity

import "encoding/xml"

type Patient struct {
	Name  string `json:"name" xml:"Name"`
	Age   int    `json:"age" xml:"Age"`
	Email string `json:"email" xml:"Email"`
}

type Patients struct {
	List    []Patient `xml:"Patient"`
	XMLName xml.Name  `xml:"patients"`
}
