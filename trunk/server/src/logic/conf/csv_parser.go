package conf

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func (this *CSV) parseFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)

	line, err := buf.ReadString('\n')

	err2 := this.parseHeader(line)
	if err2 != nil {
		return err2
	}

	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		err = this.parseSource(line)
		if err != nil {
			return err
		}

	}

	return nil
}

func (this *CSV) parseString(s string) error {
	ss := strings.Split(s, "\n")
	if len(ss) == 0 {
		return errors.New("CSV strings is empty")
	}
	if err := this.parseHeader(ss[0]); err != nil {
		this.ErrorLine = 1
		return err
	}
	for i := 1; i < len(ss); i++ {
		if err := this.parseSource(ss[i]); err != nil {
			this.ErrorLine = i + 1
			return err
		}

	}
	return nil
}

func (this *CSV) parseHeader(s string) error {
	ss := strings.Split(s, ",")
	for _, v := range ss {
		v = strings.Trim(v, "\r\n\t\" ")
		if this.index(v) != kInvalideIndex {
			this.ErrorColum = strings.Index(s, v)
			return errors.New("CSV : same column name")
		}

		this.names = append(this.names, v)
	}
	return nil
}

func (this *CSV) parseSource(s string) error {
	if len(s) == 0 {
		return nil
	}
	s = strings.Trim(s, "\n\r\t ")
	r := []string{}
	q := 0
	ii := 0
	for i, v := range s {
		if v == '\'' || v == rune("\""[0]) {
			q++
		} else if v == ',' {
			if q&1 == 0 {
				r = append(r, strings.Trim(s[ii:i], "\"\n\r\t"))
				ii = i + 1
			}
		}
	}
	r = append(r, s[ii:])
	if len(r) != len(this.names) {
		return errors.New("CSV : name record length not match")
	}
	this.data = append(this.data, r)
	return nil
}

func (this *CSV) parseJson(j []map[string]interface{}) error {
	this.data = make([][]string, len(j))
	for i, l := range j {

		for k, _ := range l {
			if this.index(k) == kInvalideIndex {
				this.names = append(this.names, k)
			}
		}
		this.data[i] = make([]string, len(this.names))
		for k, v := range l {
			this.data[i][this.index(k)] = fmt.Sprint(v)
		}
	}
	return nil
}
