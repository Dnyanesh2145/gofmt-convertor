package writer

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"errors"
	"os"

	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v3"
)

type Writer struct {
}

func (w *Writer) WriteFile(records []map[string]interface{}, path, format string) error {
	switch format {
	case "csv":
		return w.writeCSV(records, path)
	case "json":
		return w.writeJSON(records, path)
	case "yaml":
		return w.writeYAML(records, path)
	case "xml":
		return w.writeXML(records, path)
	case "excel":
		return w.writeExcel(records, path)
	default:
		return errors.New("unsupported output format")
	}
}

func (w *Writer) writeCSV(records []map[string]interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if len(records) == 0 {
		return nil
	}

	// Write header
	headers := make([]string, 0, len(records[0]))
	for k := range records[0] {
		headers = append(headers, k)
	}
	writer.Write(headers)

	// Write rows
	for _, rec := range records {
		row := make([]string, len(headers))
		for i, h := range headers {
			row[i] = rec[h].(string)
		}
		writer.Write(row)
	}
	return nil
}

func (w *Writer)  writeJSON(records []map[string]interface{}, path string) error {
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (w *Writer)  writeYAML(records []map[string]interface{}, path string) error {
	data, err := yaml.Marshal(records)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (w *Writer) writeXML(records []map[string]interface{}, path string) error {
	type Wrapper struct {
		Records []map[string]interface{} `xml:"record"`
	}
	data, err := xml.MarshalIndent(Wrapper{Records: records}, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func  (w *Writer)  writeExcel(records []map[string]interface{}, path string) error {
	f := excelize.NewFile()
	sheet := f.GetSheetName(f.GetActiveSheetIndex())

	if len(records) == 0 {
		return f.SaveAs(path)
	}

	headers := make([]string, 0, len(records[0]))
	for k := range records[0] {
		headers = append(headers, k)
	}

	// Write headers
	for i, h := range headers {
		col, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, col, h)
	}

	// Write data
	for rIdx, rec := range records {
		for cIdx, h := range headers {
			col, _ := excelize.CoordinatesToCellName(cIdx+1, rIdx+2)
			f.SetCellValue(sheet, col, rec[h])
		}
	}

	return f.SaveAs(path)
}
