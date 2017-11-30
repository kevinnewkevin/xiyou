package game

import (
	"logic/prpc"
	"jimny/logs"
	"time"
)

var(
	IsFatchMail				bool  = false
	FatchMailId				int32 = 0
	FatchMailTimeout		float64 = 0
	Mails			[]prpc.COM_Mail
)

const(
	FatchMailTimer = 30
)

func (player *GamePlayer)MailTick(dt float64)  {
	FatchMailTimeout -= dt
	if(FatchMailTimeout <= 0){
		player.FatchMail()
		FatchMailTimeout += FatchMailTimer
	}
}

func (player *GamePlayer)FatchMail()  {
	if IsFatchMail {
		return
	}
	IsFatchMail = true
	//Fatch db mail
	mails := <-FatchDBMail(player.MyUnit.InstName,FatchMailId)
	player.AppendMail(mails)
}

func (player *GamePlayer)AppendMail(mails []prpc.COM_Mail)  {
	IsFatchMail = false
	if len(mails) == 0 {
		return
	}
	for _,m := range mails{
		Mails = append(Mails,m)
		FatchMailId = m.MailId
	}
	//TO Client
	logs.Info("AppendMail = ", mails)
	if player.session != nil {
		player.session.AppendMail(mails)
	}
}

func (player *GamePlayer)InitMail()  {
	//TO Client
	if len(Mails) != 0 && player.session != nil{
		player.session.AppendMail(Mails)
	}
}

func SendMailByDrop(sendName string,recvName string,title string,content string,dropId int32)  {
	mail := prpc.COM_Mail{}
	drop := GetDropById(dropId)
	if drop == nil {
		logs.Info("SendMailByDrop Can Not Find Drop By DropId=", dropId)
		return
	}

	mail.Mailtype 		= prpc.MT_System
	mail.SendPlayerName = sendName
	mail.RecvPlayerName = recvName
	mail.Content		= content
	mail.Title			= title
	mail.Copper 		= drop.Money
	mail.Hero			= drop.Hero
	mail.MailTimestamp	= time.Now().Unix()
	item := prpc.COM_MailItem{}
	for _,i := range drop.Items{
		item.ItemId 	= i.ItemId
		item.ItemStack	= i.ItemNum
		mail.Items = append(mail.Items,item)
	}
	//insert to db
	isOK := <-InsertMail(mail)
	if isOK {
		logs.Info("SendMailByDrop",mail)
	}
}

func (player *GamePlayer)DeleteMail(mailId int32)  {
	for i:=0;i<len(Mails) ;i++  {
		if Mails[i].MailId == mailId {
			Mails = append(Mails[:i], Mails[i+1:]...)
			//del db mail
			isOK := <-EraseMail(mailId)
			if isOK && player.session != nil {
				player.session.DelMailOK(mailId)
			}
			return
		}
	}
}

func (player *GamePlayer)ReadMail(mailId int32)  {
	for i:=0;i<len(Mails) ;i++  {
		if Mails[i].MailId == mailId {
			Mails[i].IsRead = true
			//updata db mail
			isOK := <-UpdateMail(Mails[i])
			//TO Client
			if isOK && player.session != nil {
				player.session.UpdateMailOk(Mails[i])
			}
			return
		}
	}
}

func (player *GamePlayer)GetMailItem(mailId int32)  {
	for i:=0;i<len(Mails) ;i++  {
		if Mails[i].MailId == mailId {
			if len(Mails[i].Items) != 0 {
				for _,item := range Mails[i].Items{
					player.AddBagItemByItemId(item.ItemId,item.ItemStack)
				}
				Mails[i].Items = nil
			}
			if Mails[i].Copper != 0 {
				player.AddCopper(Mails[i].Copper)
				Mails[i].Copper = 0
			}
			if Mails[i].Gold != 0 {
				player.AddGold(Mails[i].Gold)
				Mails[i].Gold = 0
			}
			if Mails[i].Hero != 0 {
				if player.HasUnitByTableId(Mails[i].Hero) {
					//有这个卡就不给了
					logs.Info("PlayerName=", player.MyUnit.InstName, "GiveDrop AddUnit Have not to UnitId=", Mails[i].Hero)
				} else {
					unit := player.NewGameUnit(Mails[i].Hero)
					if unit != nil {
						logs.Info("PlayerName=", player.MyUnit.InstName, "GiveDrop AddUnit OK UnitId=", Mails[i].Hero)
						temp := unit.GetUnitCOM()
						if player.session != nil {
							player.session.AddNewUnit(temp)
						}
					}
				}
				Mails[i].Hero = 0
			}
			//updata db and to client
			isOK := <-UpdateMail(Mails[i])
			//TO Client
			if isOK && player.session != nil {
				player.session.UpdateMailOk(Mails[i])
			}

			return
		}
	}
}