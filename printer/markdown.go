package printer

import (
	"bicep-doc/parser"
	"fmt"
	"math"
	"strings"
)

const (
	NAME_HEADER        = "Name"
	TYPE_HEADER        = "Type"
	REQUIRED_HEADER    = "Required"
	DESCRIPTION_HEADER = "Description"
)

func PrintMarkdown(p *parser.Parser) {
	nameLen := int(math.Max(float64(p.InputColumns.MaxNameLen), float64(len(NAME_HEADER))))
	typeLen := int(math.Max(float64(p.InputColumns.MaxTypeLen), float64(len(TYPE_HEADER))))
	requiredLen := int(math.Max(float64(p.InputColumns.MaxRequiredLen), float64(len(REQUIRED_HEADER))))
	descriptionLen := int(math.Max(float64(p.InputColumns.MaxDescriptionLen), float64(len(DESCRIPTION_HEADER))))
	rowFmt := fmt.Sprintf("| %%-%ds | %%-%ds | %%-%ds | %%-%ds |\n", nameLen, typeLen, requiredLen, descriptionLen)

	fmt.Println("## Inputs")
	fmt.Println("")
	fmt.Printf(rowFmt, "Name", "Type", "Required", "Description")
	fmt.Printf(
		"| %s | %s | %s | %s |\n",
		strings.Repeat("-", nameLen),
		strings.Repeat("-", typeLen),
		strings.Repeat("-", requiredLen),
		strings.Repeat("-", descriptionLen))
	for _, i := range(p.Inputs) {
		fmt.Printf(rowFmt, i.Name, i.Type, i.Required, i.Description)
	}

	nameLen = int(math.Max(float64(p.OutputColumns.MaxNameLen), float64(len(NAME_HEADER))))
	typeLen = int(math.Max(float64(p.OutputColumns.MaxTypeLen), float64(len(TYPE_HEADER))))
	descriptionLen = int(math.Max(float64(p.OutputColumns.MaxDescriptionLen), float64(len(DESCRIPTION_HEADER))))
	rowFmt = fmt.Sprintf("| %%-%ds | %%-%ds | %%-%ds |\n", nameLen, typeLen, descriptionLen)

	fmt.Println("## Outputs")
	fmt.Println("")
	fmt.Printf(rowFmt, "Name", "Type", "Description")
	fmt.Printf(
		"| %s | %s | %s |\n",
		strings.Repeat("-", nameLen),
		strings.Repeat("-", typeLen),
		strings.Repeat("-", descriptionLen))
	for _, o := range(p.Outputs) {
		fmt.Printf(rowFmt, o.Name, o.Type, o.Description)
	}
}
