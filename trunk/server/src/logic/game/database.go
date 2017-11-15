package game

import (
	"bytes"
	"database/sql"
	"jimny/logs"

	_ "github.com/go-sql-driver/mysql"

	"logic/prpc"
	_ "jimny/sqlite3"

	//"log"
	"sync/atomic"
)

var (
	MaxPlayerInstId int64 =  1
	MaxUnitInstId int64  = 1
)

func GenPlayerInstId() int64{
	return atomic.AddInt64(&MaxPlayerInstId,1)
}

func GenUnitInstId() int64 {
	return atomic.AddInt64(&MaxUnitInstId,1)
}

func InitDB() {
	c, e := ConnectDB()
	if e != nil {
		logs.Debug(e.Error())
		return
	}
	defer c.Close()
	r, e := c.Query("SELECT MAX(`PlayerId`) AS MaxID FROM `Player`")

	if e != nil {
		logs.Debug(e.Error())
		return
	}


	if r.Next() {
		e = r.Scan(&MaxPlayerInstId)

		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}

	r, e = c.Query("SELECT MAX(`UnitId`) AS MaxID FROM `Unit`")

	if e != nil {
		logs.Debug(e.Error())
		return
	}


	if r.Next() {
		e = r.Scan(&MaxUnitInstId)
		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}

	logs.Infof("MAX PLAYER ID %d MAX UNIT ID %d",MaxPlayerInstId,MaxUnitInstId)
}



func ConnectDB() (*sql.DB, error) {
	//dsn := beego.AppConfig.String("dbuser") + ":" + beego.AppConfig.String("dbpass") + "@tcp(" + beego.AppConfig.String("dbhost") + ":" + beego.AppConfig.String("dbport") + ")/" + beego.AppConfig.String("dbname")
	dsn := GetEnvString("V_MySqlData")
	return sql.Open("mysql", dsn)
}

func QueryPlayer(username string) <- chan *prpc.SGE_DBPlayer {
	logs.Debug("Query player")
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
			d := int64(0)
			e = r.Scan(&a, &c, &b, &d)
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


			p.COM_Player.Employees = <- QueryUnit(a)

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


func QueryPlayerById(InstId int64) <- chan *prpc.SGE_DBPlayer {
	logs.Debug("Query player")
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
		r, e := c.Query("SELECT * FROM `Player` WHERE `InstId` = ?", InstId)

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
			d := int64(0)
			e = r.Scan(&a, &c, &b, &d)
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


			p.COM_Player.Employees = <- QueryUnit(a)

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

func InsertPlayer(p prpc.SGE_DBPlayer) <- chan int64 {

	rChan := make (chan int64)

	go func () {
		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}
		defer c.Close()
		b :=  bytes.NewBuffer(nil)

		p.Serialize(b)

		r , e := c.Exec("INSERT INTO `Player`(`PlayerId`, `Username`, `InstId`, `BinData`)VALUES(?,?,?,?)", p.PlayerId , p.Username, p.COM_Player.InstId, b.Bytes())

		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		i, e := r.LastInsertId()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		rChan <-  (i + 1)
		close(rChan)
	}()
	return rChan
}

func QueryUnit(ownerId int64) <- chan []prpc.COM_Unit {

	rChan := make(chan []prpc.COM_Unit)
	go func() {

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()
		r, e := c.Query("SELECT * FROM `Unit` WHERE `OwnerId` = ?",ownerId )

		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		arr := []prpc.COM_Unit{}

		for r.Next() {
			a := int64(0)
			b := []byte{}
			c := int64(0)
			e = r.Scan(&a, &c, &b)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p := prpc.COM_Unit{}

			bb := bytes.NewBuffer(b)
			e = p.Deserialize(bb)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p.InstId = a


			arr = append(arr, p)

		}

		rChan <- arr
		close(rChan)

	}()

	return rChan
}

func InsertUnit(ownerId int64, p prpc.COM_Unit) <- chan int64 {
	rChan := make (chan int64)
	go func () {
		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}
		defer c.Close()
		b := bytes.NewBuffer(nil)

		p.Serialize(b)

		r , e := c.Exec("INSERT INTO `Unit`(`UnitId`, `OwnerId`, `BinData`)VALUES(?,?,?)", p.InstId, ownerId, b.Bytes())

		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		i, e := r.LastInsertId()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		rChan <- (i + 1)
		close(rChan)
	}()
	return  rChan
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

		for _, u := range p.Employees{
			UpdateUnit(u)
		}

		p.Employees = nil

		e = p.Serialize(&b)

		if e != nil {
			logs.Debug(e.Error())
			return
		}

		logs.Debug("GamePlayerSave", p.Friends)
		_, e = c.Exec("UPDATE `Player` SET `BinData` = ? WHERE `PlayerId` = ?", b.Bytes(), p.PlayerId)

		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}()
}


func UpdateUnit(p prpc.COM_Unit) {
	go func () {
		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			return
		}
		defer c.Close()
		b := bytes.Buffer{}

		p.Serialize(&b)

		_, e = c.Exec("UPDATE `Unit` SET `BinData` = ? WHERE `UnitId` = ?", b.Bytes(), p.InstId)

		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}()
}


func QueryAllTopList()  <- chan []prpc.COM_TopUnit {		//取出来整张表的数据
	rChan := make(chan []prpc.COM_TopUnit)
	go func() {

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()
		r, e := c.Query("SELECT * FROM `TopList`" )

		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		arr := []prpc.COM_TopUnit{}

		for r.Next() {
			a := int64(0)
			b := []byte{}
			e = r.Scan(&a, &b)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p := prpc.COM_TopUnit{}

			bb := bytes.NewBuffer(b)
			e = p.Deserialize(bb)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			arr = append(arr, p)

		}

		rChan <- arr
		close(rChan)

	}()

	return rChan
}

func UpdateTopList(InstId int64, t prpc.SGE_DBTopUnit) {
	go func () {
		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			return
		}
		defer c.Close()
		b := bytes.Buffer{}

		t.Serialize(&b)

		_, e = c.Exec("UPDATE `TopList` SET `BinData` = ? WHERE `InstId` = ?", b.Bytes(), InstId)

		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}()
}


func InsertTopList (InstId int64, t prpc.SGE_DBTopUnit) <- chan int64 {

	rChan := make (chan int64)

	go func () {
		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}
		defer c.Close()
		b :=  bytes.NewBuffer(nil)

		t.Serialize(b)

		r , e := c.Exec("INSERT INTO `TopList`(`InstId`, `BinData`)VALUES(?,?)", InstId, b.Bytes())

		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		i, e := r.LastInsertId()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		rChan <-  (i + 1)
		close(rChan)
	}()
	return rChan
}
