package utils

import (
	"encoding/xml"
)

type ContextObjects struct {
	XMLName        xml.Name `xml:"context-objects"`
	Text           string   `xml:",chardata"`
	Ctx            string   `xml:"ctx,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Xsi            string   `xml:"xsi,attr"`
	ContextObject  struct {
		Text       string `xml:",chardata"`
		Identifier string `xml:"identifier,attr"`
		Timestamp  string `xml:"timestamp,attr"`
		Version    string `xml:"version,attr"`
		Referent   struct {
			Text          string `xml:",chardata"`
			MetadataByVal struct {
				Text   string `xml:",chardata"`
				Format struct {
					Text string `xml:",chardata"`
				} `xml:"format"`
				Metadata struct {
					Text    string `xml:",chardata"`
					Journal struct {
						Text           string `xml:",chardata"`
						Rft            string `xml:"rft,attr"`
						SchemaLocation string `xml:"schemaLocation,attr"`
						Btitle         struct {
							Text string `xml:",chardata"`
						} `xml:"btitle"`
						EISSN struct {
							Text string `xml:",chardata"`
						} `xml:"eissn"`
						Genre struct {
							Text string `xml:",chardata"`
						} `xml:"genre"`
						ISSN struct {
							Text string `xml:",chardata"`
						} `xml:"issn"`
						Jtitle struct {
							Text string `xml:",chardata"`
						} `xml:"jtitle"`
						Language struct {
							Text string `xml:",chardata"`
						} `xml:"language"`
						LCCN struct {
							Text string `xml:",chardata"`
						} `xml:"lccn"`
						ObjectID struct {
							Text string `xml:",chardata"`
						} `xml:"object_id"`
						Oclcnum struct {
							Text string `xml:",chardata"`
						} `xml:"oclcnum"`
						Place struct {
							Text string `xml:",chardata"`
						} `xml:"place"`
						Pub struct {
							Text string `xml:",chardata"`
						} `xml:"pub"`
						Stitle struct {
							Text string `xml:",chardata"`
						} `xml:"stitle"`
						Title struct {
							Text string `xml:",chardata"`
						} `xml:"title"`
						Authors struct {
							Text string `xml:",chardata"`
						} `xml:"authors"`
					} `xml:"journal"`
				} `xml:"metadata"`
			} `xml:"metadata-by-val"`
		} `xml:"referent"`
	} `xml:"context-object"`
}
