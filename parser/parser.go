package parser

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type parserState int64
const (
	processLine      parserState = iota
	processMultiLine
)

type attributes struct {
	Description string
}

type inputColumns struct {
	MaxNameLen        int
	MaxDescriptionLen int
	MaxTypeLen        int
	MaxRequiredLen    int
}

type input struct {
	Name        string
	Description string
	Type        string
	Required    string
}

type outputColumns struct {
	MaxNameLen        int
	MaxDescriptionLen int
	MaxTypeLen        int
}

type output struct {
	Name        string
	Description string
	Type        string
}

type Parser struct {
	file          string
	attrs         attributes
	Inputs        []input
	Outputs       []output
	InputColumns  inputColumns
	OutputColumns outputColumns
	state         parserState
}

func (p *Parser) updateInputColumns(c *inputColumns) {
	inputCount := len(p.Inputs)
	if inputCount > 0 {
		lastInput := p.Inputs[inputCount-1]
		c.MaxDescriptionLen = int(math.Max(float64(len(lastInput.Description)), float64(c.MaxDescriptionLen)))
		c.MaxNameLen = int(math.Max(float64(len(lastInput.Name)), float64(c.MaxNameLen)))
		c.MaxRequiredLen = int(math.Max(float64(len(lastInput.Required)), float64(c.MaxRequiredLen)))
		c.MaxTypeLen = int(math.Max(float64(len(lastInput.Type)), float64(c.MaxTypeLen)))	
	}
}

func (p *Parser) updateOutputColumns(c *outputColumns) {
	outputCount := len(p.Outputs)
	if outputCount > 0 {
		lastOutput := p.Outputs[outputCount-1]
		c.MaxDescriptionLen = int(math.Max(float64(len(lastOutput.Description)), float64(c.MaxDescriptionLen)))
		c.MaxNameLen = int(math.Max(float64(len(lastOutput.Name)), float64(c.MaxNameLen)))
		c.MaxTypeLen = int(math.Max(float64(len(lastOutput.Type)), float64(c.MaxTypeLen)))	
	}
}

func (p *Parser) processLine(line string) {
	if strings.HasPrefix(line, "@description('''") {
		p.attrs.Description = ""
		p.state = processMultiLine
	} else if strings.HasPrefix(line, "@description(") {
		p.attrs.Description = strings.TrimSuffix(strings.TrimPrefix(line, "@description('"), "')")
	} else if strings.HasPrefix(line, "param ") {
		e := strings.Split(line, " ")
		required := "no"
		if len(e) > 3 {
			required = "yes"
		}
		p.Inputs = append(p.Inputs, input{
			Name:        e[1],
			Description: p.attrs.Description,
			Type:        e[2],
			Required:    required,
		})
		p.updateInputColumns(&p.InputColumns)
		p.attrs.Description = ""
	}
	
	if strings.HasPrefix(line, "output ") {
		e := strings.Split(line, " ")
		p.Outputs = append(p.Outputs, output{
			Name:        e[1],
			Description: p.attrs.Description,
			Type:        e[2],
		})
		p.updateOutputColumns(&p.OutputColumns)
		p.attrs.Description = ""
	}
}

func (p *Parser) processMultiLine(line string) {
	if strings.HasPrefix(line, "''')") {
		p.state = processLine
	} else {
		if len(p.attrs.Description) == 0 {
			p.attrs.Description = line
		} else {
			p.attrs.Description = fmt.Sprintf("%s<br>%s", p.attrs.Description, line)
		}
	}
}

func (p *Parser) ProcessFile() {
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		switch p.state {
		case processLine:
			p.processLine(line)
		case processMultiLine:
			p.processMultiLine(line)
		}
	}

	sort.Slice(p.Inputs, func(i, j int) bool {
		return strings.Compare(p.Inputs[i].Name, p.Inputs[j].Name) < 1
	})
	sort.Slice(p.Outputs, func(i, j int) bool {
		return strings.Compare(p.Outputs[i].Name, p.Outputs[j].Name) < 1
	})
}

func NewParser(file string) (* Parser) {
	return &Parser{
		file:    file,
		Inputs:  make([]input, 0),
		Outputs: make([]output, 0),
		state:   processLine,
	}
}
