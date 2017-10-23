package conf

import (
	"fmt"
	"testing"
)

func Test_IniFile(t *testing.T) {
	c, e := NewConfigerFile("test/main.ini")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(c.GetBool("bool"))
	fmt.Println(c.GetInt("int"))
	fmt.Println(c.GetInt("int32"))
	fmt.Println(c.GetInt64("int64"))
	fmt.Println(c.GetString("string"))
	fmt.Println(c.GetStrings("strings"))

	fmt.Println(c.GetBool("test::bool"))
	fmt.Println(c.GetInt("test::int"))
	fmt.Println(c.GetFloat32("test::int32"))
	fmt.Println(c.GetInt64("test::int64"))
	fmt.Println(c.GetString("test::string"))
	fmt.Println(c.GetStrings("test::strings"))

	fmt.Println(c.GetBool("a::bool"))
	fmt.Println(c.GetInt("a::int"))
	fmt.Println(c.GetInt("a::int32"))
	fmt.Println(c.GetInt64("a::int64"))
	fmt.Println(c.GetString("a::string"))
	fmt.Println(c.GetStrings("a::strings"))

	fmt.Println(c.GetInt("b::int"))
	fmt.Println(c.GetInt("b::int32"))
	fmt.Println(c.GetInt("b::int64"))
	fmt.Println(c.GetString("b::string"))
	fmt.Println(c.GetStrings("b::strings"))

}

func Test_IniString(t *testing.T) {
	s := fmt.Sprintln(";Test comment")
	s = fmt.Sprintln("#Test comment")
	s = fmt.Sprintln("include 'test\\main.ini'")

	c, e := NewConfigerString(s)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(c.GetInt("int"))
	fmt.Println(c.GetInt("int32"))
	fmt.Println(c.GetInt("int64"))
	fmt.Println(c.GetString("string"))
	fmt.Println(c.GetStrings("strings"))

	fmt.Println(c.GetInt("test::int"))
	fmt.Println(c.GetInt("test::int32"))
	fmt.Println(c.GetInt("test::int64"))
	fmt.Println(c.GetString("test::string"))
	fmt.Println(c.GetStrings("test::strings"))

	fmt.Println(c.GetInt("a::int"))
	fmt.Println(c.GetInt("a::int32"))
	fmt.Println(c.GetInt("a::int64"))
	fmt.Println(c.GetString("a::string"))
	fmt.Println(c.GetStrings("a::strings"))

	fmt.Println(c.GetInt("b::int"))
	fmt.Println(c.GetInt("b::int32"))
	fmt.Println(c.GetInt("b::int64"))
	fmt.Println(c.GetString("b::string"))
	fmt.Println(c.GetStrings("b::strings"))
	fmt.Println(c.GetString("b::path"))
}
