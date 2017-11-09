package main

import (
	_ "jimny/sqlite3"

	"jimny/network"
	"fmt"
)

var TestTable = "CREATE TABLE IF NOT EXISTS TEST(" +
	"UserName	VARCHAR(60) NOT NULL," +
	"PlayerName VARCHAR(60) NOT NULL," +
	"PlayerGuid INT NOT NULL," +
	"PlayerLevel INT NOT NULL," +
	"PlayerProfession INT NOT NULL," +
	"PlayerGrade INT NOT NULL," +
	"Money INT NOT NULL," +
	"Diamond INT NOT NULL," +
	"Magic INT NOT NULL," +
	"LogoutTime BIGINT NOT NULL," +
	"BinData	BLOB	 NOT NULL," +
	"Seal	INT	 NOT NULL," +
	"Freeze	INT	 NOT NULL," +
	"InDoorId	INT	 NOT NULL," +
	"VersionNumber	INT	 NOT NULL," +
	"PRIMARY KEY(PlayerGuid)" +
	");"

var InsertTable = "INSERT INTO TEST(UserName,PlayerName,PlayerGuid,PlayerLevel,PlayerProfession,PlayerGrade,Money,Diamond,Magic,LogoutTime,BinData,Seal,Freeze,InDoorId,VersionNumber)" +
	"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

var SelectTable = "SELECT * FROM TEST;"

func main() {


	//serv := network.Server{}
	//serv.Open("tcp4","127.0.0.1:10999")
	//serv.Run()



}
