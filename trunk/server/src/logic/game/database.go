package game

import (
	"bytes"
	"database/sql"
	"jimny/logs"

	_ "github.com/go-sql-driver/mysql"

	"logic/prpc"
	_ "jimny/sqlite3"

	//"log"
)

func InitDB() {
	ptable := "CREATE TABLE IF NOT EXISTS `Player` ( `PlayerId` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, `Username` VARCHAR(255) ,`BinData` BLOB NOT NULL ) ;"

	c, e := ConnectDB()
	if e != nil {
		logs.Debug(e.Error())
		return
	}

	_, e = c.Exec(ptable)

	if e != nil {
		//log.Error(e.Error())
	}
}

type Database struct {

}


func ConnectDB() (*sql.DB, error) {
	//dsn := beego.AppConfig.String("dbuser") + ":" + beego.AppConfig.String("dbpass") + "@tcp(" + beego.AppConfig.String("dbhost") + ":" + beego.AppConfig.String("dbport") + ")/" + beego.AppConfig.String("dbname")
	dsn := GetEnvString("V_MySqlData")
	return sql.Open("mysql", dsn)
}

func QueryPlayer(username string) <- chan *prpc.SGE_DBPlayer {

	rChan := make(chan *prpc.SGE_DBPlayer)
	go func() {

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()
		r, e := c.Query("SELECT * FROM `Player` WHERE `Username` = ?", username)

		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		if r.Next() {
			a := int64(0)
			b := []byte{}
			c := ""
			e = r.Scan(&a, &c, &b)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p := &prpc.SGE_DBPlayer{}

			bb := bytes.NewBuffer(b)
			e = p.Deserialize(bb)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p.PlayerId = a

			rChan <- p

			close(rChan)
			return
		}

		rChan <- nil
		close(rChan)
		return
	}()

	return rChan
}

func InsertPlayer(p prpc.SGE_DBPlayer) {

	go func () {
		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			return
		}
		defer c.Close()
		b := bytes.Buffer{}

		p.Serialize(&b)

		_, e = c.Exec("INSERT INTO `Player`(`Username`, `BinData`)VALUES(?,?)", p.Username, b.Bytes())

		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}()
}

func UpdatePlayer(p prpc.SGE_DBPlayer) {

	//logs.Debug(p.UnitGroup)
	go func () {
		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			return
		}
		defer c.Close()
		b := bytes.Buffer{}

		p.Serialize(&b)

		_, e = c.Exec("UPDATE `Player` SET `BinData` = ? WHERE `Username` = ?", b.Bytes(), p.Username)

		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}()
}
