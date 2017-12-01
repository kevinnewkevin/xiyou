package game

import (
	"logic/prpc"
	"jimny/logs"
	"time"
)

const(
	FatchMailTimer = 30
)

func (player *GamePlayer)MailTick(dt float64)  {
	player.FatchMailTimeout -= dt
	if(player.FatchMailTimeout <= 0){
		player.FatchMail()
		player.FatchMailTimeout += FatchMailTimer
	}
}

func (player *GamePlayer)FatchMail()  {
	if player.IsFatchMail {
		return
	}
	player.IsFatchMail = true
	//Fatch db mail
	mails := <-FatchDBMail(player.MyUnit.InstName,player.FatchMailId)
	player.AppendMail(mails)
}

func (player *GamePlayer)AppendMail(mails []prpc.COM_Mail)  {
	player.IsFatchMail = false
	if len(mails) == 0 {
		return
	}
	for _,m := range mails{
		player.Mails = append(player.Mails,m)
		player.FatchMailId = m.MailId
	}
	//TO Client
	logs.Info("AppendMail = ", mails)
	if player.session != nil {
		player.session.AppendMail(mails)
	}
}

func (player *GamePlayer)InitMail()  {
	//TO Client
	if len(player.Mails) != 0 && player.session != nil{
		player.session.AppendMail(player.Mails)
	}
}

func SendMailByDrop(sendName string,recvName string,title string,content string,dropId int32)  {
	mail := prpc.COM_Mail{}
	drop := GetDropById(dropId)
	if drop == nil {
		logs.Info("SendMailByDrop Can Not Find Drop By DropId=", dropId)
		return
	}
	logs.Info("SendMailByDrop Drop ", drop)
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
	for i:=0;i<len(player.Mails) ;i++  {
		if player.Mails[i].MailId == mailId {
			player.Mails = append(player.Mails[:i], player.Mails[i+1:]...)
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
	for i:=0;i<len(player.Mails) ;i++  {
		if player.Mails[i].MailId == mailId {
			player.Mails[i].IsRead = true
			//updata db mail
			isOK := <-UpdateMail(player.Mails[i])
			//TO Client
			if isOK && player.session != nil {
				player.session.UpdateMailOk(player.Mails[i])
			}
			return
		}
	}
}

func (player *GamePlayer)GetMailItem(mailId int32)  {
	for i:=0;i<len(player.Mails) ;i++  {
		if player.Mails[i].MailId == mailId {
			if len(player.Mails[i].Items) != 0 {
				for _,item := range player.Mails[i].Items{
					player.AddBagItemByItemId(item.ItemId,item.ItemStack)
				}
				player.Mails[i].Items = nil
			}
			if player.Mails[i].Copper != 0 {
				player.AddCopper(player.Mails[i].Copper)
				player.Mails[i].Copper = 0
			}
			if player.Mails[i].Gold != 0 {
				player.AddGold(player.Mails[i].Gold)
				player.Mails[i].Gold = 0
			}
			if player.Mails[i].Hero != 0 {
				if player.HasUnitByTableId(player.Mails[i].Hero) {
					//有这个卡就不给了
					logs.Info("PlayerName=", player.MyUnit.InstName, "GetMailItem AddUnit Have not to UnitId=",player. Mails[i].Hero)
				} else {
					unit := player.NewGameUnit(player.Mails[i].Hero)
					if unit != nil {
						logs.Info("PlayerName=", player.MyUnit.InstName, "GetMailItem AddUnit OK UnitId=", player.Mails[i].Hero)
						temp := unit.GetUnitCOM()
						if player.session != nil {
							player.session.AddNewUnit(temp)
						}
					}
				}
				player.Mails[i].Hero = 0
			}
			//updata db and to client
			logs.Info("PlayerName=", player.MyUnit.InstName, "GetMailItem OK")
			isOK := <-UpdateMail(player.Mails[i])
			//TO Client
			if isOK && player.session != nil {
				player.session.UpdateMailOk(player.Mails[i])
			}

			return
		}
	}
}