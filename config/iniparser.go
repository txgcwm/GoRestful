package main

import (
	"fmt"
	"bufio"
	"bytes"
	"io"
	"os"
	"unicode"
)

const (
	stat_none = iota
	stat_group
	stat_key
	stat_value
	stat_comment
)

type Attr struct {
	Name    string
	Value   string
	Comment string
	next    *Attr
}

type Element struct {
	Element string
	Attr    *Attr
	next    *Element
}

type Decoder struct {
	state int
	b     byte
	t     bytes.Buffer
	r     io.ByteReader
	err   error
	m     *Element
	n     string
}

func (d *Decoder) getAttr(m *Element, e string) string {
	for n := m.Attr; nil != n; n = n.next {
		if e == n.Name {
			return n.Value
		}
	}

	return ""
}

func (d *Decoder) GetElement(e string, a string) string {
	for n := d.m; nil != n; n = n.next {
		if e == n.Element {
			return d.getAttr(n, a)
		}
	}

	return ""
}

func (d *Decoder) newAttrNextComment(value string) {
	d.m.Attr.Comment = value
	println(value)
}

func (d *Decoder) newAttrNext(name string, value string) {
	attr := new(Attr)
	attr.Name = name
	attr.Value = value

	if nil == d.m.Attr {
		attr.next = nil
	} else {
		attr.next = d.m.Attr
	}

	d.m.Attr = attr
}

func (d *Decoder) newElement(name string) {
	element := new(Element)
	element.Element = name
	element.Attr = nil

	if nil == d.m {
		element.next = nil
	} else {
		element.next = d.m
	}

	d.m = element
}

func (d *Decoder) switchToMap() {
	for {
		d.b, d.err = d.r.ReadByte()
		if d.err != nil {
			fmt.Println("error: ", d.err)
			break
		}

		switch d.state {
		case stat_none:
			if d.b == '[' {
				d.state = stat_group
			} else if d.b == ';' {
				d.state = stat_comment
			} else if !unicode.IsSpace(rune(d.b)) {
				d.state = stat_key
				d.t.WriteByte(byte(d.b))
			}
			break

		case stat_group:
			if d.b == ']' {
				d.state = stat_none
				d.newElement(d.t.String())
				d.t.Reset()
			} else if !unicode.IsSpace(rune(d.b)) {
				d.t.WriteByte(byte(d.b))
			}
			break

		case stat_key:
			if d.b == '=' {
				d.state = stat_value
				d.n = d.t.String()
				d.t.Reset()
			} else if !unicode.IsSpace(rune(d.b)) {
				d.t.WriteByte(byte(d.b))
			}
			break

		case stat_value:
			if !unicode.IsSpace(rune(d.b)) {
				d.t.WriteByte(byte(d.b))
			} else {
				d.state = stat_none
				d.newAttrNext(d.n, d.t.String())
				d.t.Reset()
			}
			break

		case stat_comment:
			if !unicode.IsSpace(rune(d.b)) {
				d.t.WriteByte(byte(d.b))
			} else {
				d.state = stat_none
				d.newAttrNextComment(d.t.String())
				d.t.Reset()
			}
			break

		default :
			d.state = stat_none
			break
		}
	}
}

func (d *Decoder) switchToReader(r io.Reader) {
	if rb, ok := r.(io.ByteReader); ok {
		d.r = rb
	} else {
		d.r = bufio.NewReader(r)
	}

	d.switchToMap()
}

func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{}
	d.switchToReader(r)

	return d
}

func main() {
	xmlFile, err := os.Open("test.ini")
	if nil != err {
		return
	}

	defer xmlFile.Close()
	d := NewDecoder(xmlFile)
	println(d.GetElement("test", "ac"))
}

