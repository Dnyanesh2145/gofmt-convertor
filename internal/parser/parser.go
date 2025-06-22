package parser

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Parser struct {
}

var Record = map[string]interface{}{}

func (p *Parser) ParseFile(path, format string) ([]map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	switch format {
	case "csv":
		return p.ParseCSV(file)
	case "json":
		return p.ParseJSON(file)
	case "yaml":
		return p.ParseYAML(file)
	case "xml":
		return p.ParseXML(file)
	default:
		return nil, errors.New("unsupported input format")
	}
}

func (p *Parser) ParseCSV(file io.Reader) ([]map[string]interface{}, error) {
	var records []map[string]interface{}
	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		record := make(map[string]interface{})
		for i, col := range header {
			record[col] = row[i]
		}
		records = append(records, record)
	}

	return records, nil
}

func (p *Parser) ParseJSON(file io.Reader) ([]map[string]interface{}, error) {
	var records []map[string]interface{}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}
func (p *Parser) ParseYAML(file io.Reader) ([]map[string]interface{}, error) {
	var records []map[string]interface{}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (p *Parser) ParseXML(file io.Reader) ([]map[string]interface{}, error) {
	var wrapper struct {
		Records []map[string]interface{} `xml:"record"`
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}
	return wrapper.Records, nil
}
