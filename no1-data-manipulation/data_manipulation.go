package no1datamanipulation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Country struct {
	Name     string `json:"name"`
	DialCode string `json:"dialCode"`
	IsoCode  string `json:"isoCode"`
	Flag     string `json:"flag"`
}

func LoadData() ([]Country, error) {
	var countries []Country
	var data []byte

	// get data form server
	resp, err := http.Get("https://citcall.com/test/countries.json")
	if err != nil {
		log.Println("error when get data form server")
		return nil, err
	}
	defer resp.Body.Close()

	// read response
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error when read data from server")
		return nil, err
	}

	// marshal data to countries
	if err := json.Unmarshal(data, &countries); err != nil {
		log.Printf("Failed to load data: %v", err)
		return nil, err
	}
	return countries, nil
}

func MakeTable() error {
	// crete dummy data
	data := ""

	// loda countries data form server
	countries, err := LoadData()
	if err != nil {
		log.Println("error when load data")
		return err
	}

	// Create table
	for _, country := range countries {
		data += "<tr>"
		data += fmt.Sprintf("<td>%s</td>\n", country.Name)
		data += fmt.Sprintf("<td>%s</td>\n", country.DialCode)
		data += fmt.Sprintf("<td>%s</td>\n", country.IsoCode)
		data += fmt.Sprintf("<td>%s</td>\n", country.Flag)
		data += "</tr>"
	}

	// make html file
	htmlData = strings.ReplaceAll(htmlData, "{{data}}", data)
	ioutil.WriteFile("table.html", []byte(htmlData), 0644)
	return nil
}

var htmlData string = `
<!DOCTYPE html>
<html>
<head>
<style>
table {
  font-family: arial, sans-serif;
  border-collapse: collapse;
  width: 100%;
}

td, th {
  border: 1px solid #dddddd;
  text-align: left;
  padding: 8px;
}

tr:nth-child(even) {
  background-color: #dddddd;
}
</style>
</head>
<body>
<h2>The table is automatically build by run function MakeTable() from no1datamanipulation package</h2>
<table>
  <tr>
    <th>Name</th>
    <th>Dial Code</th>
    <th>ISO Code</th>
	<th>Flag</th>
  </tr>
  {{data}}
</table>

</body>
</html>`
