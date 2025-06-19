package print

import (
	"fmt"
	"log"
	"regexp"
)

type Printer struct {
	indent int8
}

func (p Printer) printIndentation() {
	for i := 0; i < int(p.indent)+1; i++ {
		fmt.Printf("\t")
	}
}

func (p Printer) PrintHead() {
	fmt.Println("<?php")
	fmt.Println("")
	fmt.Println("$CONFIG = [")
}

func (p Printer) PrintFoot() {
	fmt.Println("];")
}

func (p *Printer) StartArray(key string) {
	p.printIndentation()
	fmt.Printf("'%s' => [\n", key)
	p.indent += 1
}

func (p *Printer) EndArray() {
	p.indent -= 1

	p.printIndentation()
	fmt.Printf("],\n")
}

func (p Printer) PrintBool(data bool, key string) {
	p.printIndentation()
	fmt.Printf("'%s' => %t,\n", key, data)
}

func (p Printer) PrintString(data string, key string) {
	if len(data) == 0 {
		return
	}

	p.printIndentation()
	fmt.Printf("'%s' => '%s',\n", key, data)
}

func (p Printer) PrintInt16(data int16, key string) {
	p.printIndentation()
	fmt.Printf("'%s' => %d,\n", key, data)
}

func (p Printer) PrintStringSlice(data []string, key string) {
	if len(data) == 0 {
		return
	}

	p.printIndentation()
	fmt.Printf("'%s' => [", key)
	for _, d := range data {
		fmt.Printf("'%s',", d)
	}
	fmt.Print("],\n")
}

func (p Printer) PrintStringMap(data []string, key string) {
	if len(data) == 0 {
		return
	}

	pattern, err := regexp.Compile("(?<key>[a-zA-Z]+)=(?<value>[a-zA-Z0-9-_]+)")
	if err != nil {
		log.Fatal(err)
	}

	p.printIndentation()
	fmt.Printf("'%s' => [\n", key)
	for _, d := range data {
		matches := pattern.FindAllStringSubmatch(d, -1)
		p.printIndentation()
		fmt.Print("\t[")
		for _, match := range matches {
			fmt.Printf("'%s' => '%s',", match[1], match[2])
		}
		fmt.Print("]")
	}
	fmt.Print("\n\t],\n")
}
