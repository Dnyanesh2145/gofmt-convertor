# 📊 gofmt-convertor

A fast and flexible Go-based CLI tool to convert between **CSV**, **JSON**, **YAML**, **XML**, and **Excel** formats.  
Supports **filtering**, **sorting**, and **bi-directional** data transformation.  
Built using **clean architecture**, **SOLID principles**, and **extensible design patterns**.

---

## 🚀 Features

- ✅ Convert between CSV, JSON, YAML, XML, Excel
- ♻️ Bi-directional support (any-to-any format)
- 🧱 Modular architecture for extension

---

## 📦 Installation

Make sure you have [Go](https://golang.org/dl/) installed (Go 1.18+).

```bash
go install github.com/Dnyanesh2145/gofmt-convertor@latest
```

This installs the `gofmt-convertor` binary to your `$GOBIN`.

---

## 🧪 Usage

```bash
gofmt-convertor --input <input-file> --output <output-file> \
  --format <input-format> --out-format <output-format> \
  [--filter <condition>] [--sort <column:asc|desc>]
```

### ✅ Supported Formats:
- `csv`
- `json`
- `yaml`
- `xml`
- `excel`

---

### 📄 Examples

#### Convert CSV to JSON:
```bash
gofmt-convertor --input data.csv --output data.json --format csv --out-format json
```

#### Convert JSON to Excel:
```bash
gofmt-convertor --input data.json --output result.xlsx --format json --out-format excel
```

## 🛠 Flags

| Flag         | Description                              |
|--------------|------------------------------------------|
| `--input`     | Input file path                          |
| `--output`    | Output file path                         |
| `--format`    | Input format (`csv`, `json`, etc.)       |
| `--out-format`| Output format (`csv`, `json`, etc.)      |
| `--filter`    | Filter condition (`field=value`)         |
| `--sort`      | Sort by column (`field:asc` or `desc`)   |

---

## 📁 Project Structure

```
gofmt-convertor/
├── cmd/            # CLI entrypoint (Cobra)
├── internal/
│   ├── parser/     # File parsers
│   ├── writer/     # File writers
│   ├── filter/     # Filtering logic
│   ├── sort/       # Sorting logic
│   └── model/      # Shared data model
├── go.mod
└── main.go
```

---

## 💡 Contributing

Pull requests, feature ideas, and bug reports are welcome!  
Feel free to fork the project and submit a PR 🚀

---

## 📜 License

MIT © [Dnyanesh Rode](https://github.com/Dnyanesh2145)
