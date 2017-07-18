package conf

import (
	"strings"
	"errors"
	"io/ioutil"
)

func(this *CSV)parseFile(f string) error{
	s, e := ioutil.ReadFile(f)
	if e != nil {
		return e
	}

	return this.parseString(string(s))
}

func(this *CSV)parseString(s string) error {
	ss := strings.Split(s,"\n")
	if len(ss) == 0{
		return errors.New("CSV strings is empty")
	}
	if err := this.parseHeader(ss[0]); err != nil{
		this.ErrorLine = 1
		return  err
	}
	for i:=1; i<len(ss); i++{
		if err := this.parseSource(ss[i]); err != nil{
			this.ErrorLine = i + 1
			return  err;
		}

	}
	return nil
}

func(this *CSV)parseHeader(s string) error{
	ss := strings.Split(s,",")
	for i, v := range ss{
		v = strings.Trim(v,"\r\n\t\" ")
		if _ , ok := this.names[v]; ok {
			this.ErrorColum = strings.Index(s,v)
			return  errors.New("CSV : same column name")
		}
		this.names[v] = i
	}
	return  nil
}

func(this* CSV)parseSource(s string)error{
	s = strings.Trim(s, "\n\r\t ");
	r := []string{}
	q := 0
	ii := 0
	for i, v := range s {
		if v == '\'' || v == rune("\""[0]){
			q++;
		}else if v == ','{
			if q&1 == 0{
				r = append(r,s[ii:i])
				ii = i + 1
			}
		}
	}
	r = append(r, s[ii:])
	if len(r)!=len(this.names){
		return errors.New("CSV : name record length not match")
	}
	this.data = append(this.data,r)
	return  nil
}