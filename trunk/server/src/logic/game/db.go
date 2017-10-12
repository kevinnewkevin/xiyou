package game

import (
	"logic/prpc"
	"database/sql"

	"fmt"
	"bytes"

	_ "github.com/go-sql-driver/mysql"
)

const(

)

func ConnectDB() (*sql.DB, error) {
	//dsn := beego.AppConfig.String("dbuser") + ":" + beego.AppConfig.String("dbpass") + "@tcp(" + beego.AppConfig.String("dbhost") + ":" + beego.AppConfig.String("dbport") + ")/" + beego.AppConfig.String("dbname")
	dsn := "xysk:123456@tcp(10.10.10.254:3306)/xygame"
	return sql.Open("mysql", dsn)
}

func QueryPlayer(p *prpc.SGE_DBPlayer) bool{
	c,e := ConnectDB()
	if e != nil{
		fmt.Println(e.Error())
		return false
	}

	r, e := c.Query("SELECT * FROM `Player` WHERE `Username` = ?", p.Username)

	if e != nil{
		fmt.Println(e.Error())
		return false
	}

	if r.Next(){
		a := int64(0)
		b := []byte{}
		c := ""
		e = r.Scan(&a,&c,&b)
		if e != nil{
			fmt.Println(e.Error())
			return false
		}

		bb := bytes.NewBuffer(b)
		e = p.Deserialize(bb)
		if e != nil{
			fmt.Println(e.Error())
			return  false
		}

		p.PlayerId = a

		return  true
	}

	return false
}

func InsertPlayer(p prpc.SGE_DBPlayer){
	c,e := ConnectDB()
	if e != nil{
		fmt.Println(e.Error())
		return
	}

	b := bytes.Buffer{}

	p.Serialize(&b)

	_, e = c.Exec("INSERT INTO `Player`(`Username`, `BinData`)VALUES(?,?)", p.Username, b.Bytes())

	if e != nil{
		fmt.Println(e.Error())
		return
	}
}

func UpdatePlayer(p prpc.SGE_DBPlayer){
	c,e := ConnectDB()
	if e != nil{
		fmt.Println(e.Error())
		return
	}

	b := bytes.Buffer{}

	p.Serialize(&b)

	_, e = c.Exec("UPDATE `Player` SET `BinData` = ? WHERE `PlayerId` = ?", b.Bytes(), p.InstId)
}



