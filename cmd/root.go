package cmd

import (
	"fmt"
	"os"

	"github.com/Dnyanesh2145/gofmt-convertor/internal/parser"
	"github.com/Dnyanesh2145/gofmt-convertor/internal/writer"
	"github.com/spf13/cobra"
)

var (
	input      string
	output     string
	inFormat   string
	outFormat  string
	// filterCond string
	// sortOption string
)
func init() {
	rootCmd.Flags().StringVar(&input, "input", "", "Input file path")
	rootCmd.Flags().StringVar(&output, "output", "", "Output file path")
	rootCmd.Flags().StringVar(&inFormat, "format", "csv", "Input format (csv|json|yaml|xml)")
	rootCmd.Flags().StringVar(&outFormat, "out-format", "json", "Output format (csv|json|yaml|xml|excel)")
	// rootCmd.Flags().StringVar(&filterCond, "filter", "", "Filter condition (e.g., age>25)")
	// rootCmd.Flags().StringVar(&sortOption, "sort", "", "Sort column (e.g., name:asc)")
}


var rootCmd = &cobra.Command{
	Use:   "csv-converter-cli",
	Short: "Convert between CSV, JSON, YAML, XML, Excel with filtering and sorting",
	Long: `Data Converter CLI is a tool to convert between structured file formats.

	Examples:
	  data-converter --input file.csv --output out.json --format csv --out-format json
	  data-converter --input data.yaml --output result.xlsx --format yaml --out-format excel`,
	Run: func(cmd *cobra.Command, args []string) {
		var parser parser.Parser
		records, err := parser.ParseFile(input, inFormat)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			return
		}

		// if filterCond != "" {
		// 	records = filter.ApplyFilter(records, filterCond)
		// }
		// if sortOption != "" {
		// 	records = sort.ApplySort(records, sortOption)
		// }
		var writer writer.Writer
		err = writer.WriteFile(records, output, outFormat)
		if err != nil {
			fmt.Println("Error writing output:", err)
			return
		}

		fmt.Println("✅ Conversion successful")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
