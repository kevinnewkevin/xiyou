package game

import (
	"logic/prpc"
	"time"
	"jimny/logs"
)

type Guild struct {
	GuildData 		*prpc.COM_Guild
	GuildMember 	[]*prpc.COM_GuildMember
}

const (
	MAX_GUILD_NAME_LEN = 12
)

var(
	IdGuildMap 			= map[int32]*Guild{}
	NameGuildMap		= map[string]*Guild{}
	PlayerIdGuildMap	= map[int64]*Guild{}
	TempGuildName		[]string		//占位
	TempGuildPlayerId	[]int64			//占位
)

func InitGuild()  {
	guilds := <-FetchGuild()
	if guilds==nil {
		return
	}
	for _,g := range guilds{
		AddGuild(g)
	}
	members := <-FetchGuildMember()
	if members==nil {
		return
	}
	for _,m := range members{
		AddGuildMember(m)
	}
}

func FindGuildById(guildId int32) *Guild {
	return IdGuildMap[guildId]
}

func FindGuildByName(name string) *Guild {
	return NameGuildMap[name]
}

func FindGuildByPlayerId(instId int64) *Guild {
	return PlayerIdGuildMap[instId]
}

func AddGuildCritical(guildName string,playerInstId int64)  {
	TempGuildName = append(TempGuildName,guildName)
	TempGuildPlayerId = append(TempGuildPlayerId,playerInstId)
}

func DelGuildCritical(guildName string,playerInstId int64)  {
	for i:=0;i<len(TempGuildName) ;i++  {
		if TempGuildName[i] == guildName {
			TempGuildName = append(TempGuildName[:i], TempGuildName[i+1:]...)
			break
		}
	}
	for i:=0;i<len(TempGuildPlayerId) ;i++  {
		if TempGuildPlayerId[i] == playerInstId {
			TempGuildPlayerId = append(TempGuildPlayerId[:i], TempGuildPlayerId[i+1:]...)
			break
		}
	}
}

func CheckGuildCritical(guildName string,playerInstId int64) bool {
	isName,isId := true,true
	for i:=0;i<len(TempGuildName) ;i++  {
		if TempGuildName[i] == guildName {
			isName = false
		}
	}
	for i:=0;i<len(TempGuildPlayerId) ;i++ {
		if TempGuildPlayerId[i] == playerInstId {
			isId = false
		}
	}
	return isName&&isId
}

func GuildEveryMonday()  {

	for _,g := range IdGuildMap{
		if g == nil {
			continue
		}
		g.GuildData.Contribution = 0
		g.UpdateGuild()
	}

	for _,g := range NameGuildMap{
		if g == nil {
			continue
		}
		g.GuildData.Contribution = 0
	}

	for _,g := range PlayerIdGuildMap{
		if g == nil {
			continue
		}
		g.GuildData.Contribution = 0
	}
	
	ResetDBGuildContribution()
}

func CreatGuild(player *GamePlayer,guildName string)  {
	if player == nil {
		return
	}
	if len(guildName) == 0 {
		return
	}
	//屏蔽字

	if len(guildName) > MAX_GUILD_NAME_LEN {
		return
	}
	if FindGuildByName(guildName) != nil {
		return
	}
	if FindGuildByPlayerId(player.MyUnit.InstId) != nil {
		return
	}
	if player.MyUnit.GetIProperty(prpc.IPT_LEVEL) < int32(GetGlobalInt("C_CreatGuildLevel")) {
		return
	}
	if player.MyUnit.GetIProperty(prpc.IPT_GOLD) < int32(GetGlobalInt("C_CreatGuildGold")){
		return
	}
	if CheckGuildCritical(guildName,player.MyUnit.InstId) == false {
		return
	}
	AddGuildCritical(guildName,player.MyUnit.InstId)

	pGuild := prpc.COM_Guild{}
	pGuild.CreateTime 			= time.Now().Unix()
	pGuild.GuildId				= GenGuildInstId()
	pGuild.GuildName			= guildName
	pGuild.Master				= player.MyUnit.InstId
	pGuild.MasterName			= player.MyUnit.InstName
	pGuild.GuildVal				= player.TianTiVal

	pGuildMember := prpc.COM_GuildMember{}
	pGuildMember.GuildId		= pGuild.GuildId
	pGuildMember.RoleName		= player.MyUnit.InstName
	pGuildMember.RoleId			= player.MyUnit.InstId
	pGuildMember.Level			= player.MyUnit.GetIProperty(prpc.IPT_PROMOTE)
	pGuildMember.Job			= prpc.GJ_Premier
	pGuildMember.UnitId			= player.MyUnit.UnitId

	need := GetGlobalInt("C_CreatGuildGold")
	player.AddGold(-int32(need))

	isOK := <- InsertGuild(pGuild,pGuildMember)

	logs.Info("CreateGuild InsertGuild To DB ", isOK)

	if isOK {
		DelGuildCritical(guildName,player.MyUnit.InstId)
		guild := AddGuild(pGuild)
		if guild==nil {
			return
		}
		// To Client CreateGuildOK
		isAdd := AddGuildMember(pGuildMember)
		if !isAdd {
			return
		}
		PlayerIdGuildMap[pGuildMember.RoleId] = guild
		guild.UpdateMemberList(player)
		guild.UpdateGuild()
		if player.session ==nil {
			return
		}
		player.session.CreateGuildOK()

		logs.Info("CreateGuildOK GuildName = ",pGuild.GuildName," MasterName = ",pGuild.MasterName)
	}
}

func AddGuild(guild prpc.COM_Guild) *Guild {
	pGuild := Guild{}
	pGuild.GuildData = &guild
	IdGuildMap[pGuild.GuildData.GuildId] 		= &pGuild
	NameGuildMap[pGuild.GuildData.GuildName]	= &pGuild

	return &pGuild
}

func DelGuild(guildId int32)  {
	pGuild := FindGuildById(guildId)
	if pGuild==nil {
		return
	}
	delete(IdGuildMap,guildId)
	delete(NameGuildMap,pGuild.GuildData.GuildName)
	for k,g := range PlayerIdGuildMap{
		if g.GuildData.GuildId == guildId {
			DeleteDBGuildMember(k)
			delete(PlayerIdGuildMap,k)
		}
	}
	DeleteDBGuild(guildId)
}

func AddGuildMember(member prpc.COM_GuildMember) bool {
	pGuild := FindGuildById(member.GuildId)
	if pGuild==nil {
		logs.Info("AddGuildMember Can Not Find Guild ", member.GuildId)
		return false
	}

	if pGuild.AddMember(member)==false {
		logs.Info("AddMember guild member id conflict ",member.RoleId)
		return false
	}

	player := FindPlayerByInstId(member.RoleId)
	if player!=nil {
		pGuild.GuildMemberOnLine(player)
	}

	return true
}

func DelGuildMember(roleid int64)  {
	pGuild := FindGuildByPlayerId(roleid)
	if pGuild==nil {
		return
	}
	pGuild.DelMember(roleid)
	DeleteDBGuildMember(roleid)
	if len(pGuild.GuildMember)==0 {
		DelGuild(pGuild.GuildData.GuildId)
	}
}

func RequestGuildList(player *GamePlayer)  {
	if player==nil {
		return 
	}
	var guildViewerList []prpc.COM_GuildViewerData
	for _,guild := range NameGuildMap{
		guildViewerData := prpc.COM_GuildViewerData{}
		guildViewerData.GuildId 		= guild.GuildData.GuildId
		guildViewerData.GuildName		= guild.GuildData.GuildName
		guildViewerData.MasterName		= guild.GuildData.MasterName
		guildViewerData.MemberNum		= int32(len(guild.GuildMember))
		guildViewerData.GuildVal		= guild.GuildData.GuildVal
		guildViewerList = append(guildViewerList,guildViewerData)
	}
	//To Client guildViewerList
	if player.session == nil {
		return
	}
	player.session.QueryGuildListResult(guildViewerList)

	logs.Info("RequestGuildList ",guildViewerList)
}

func RequestGuildDetails(player *GamePlayer,guildId int32)  {
	if player==nil {
		return
	}
	pGuild := FindGuildById(guildId)
	if pGuild==nil {
		return
	}
	data := prpc.COM_GuildDetails{}
	data.GuildId 		= pGuild.GuildData.GuildId
	data.GuildName		= pGuild.GuildData.GuildName
	data.MasterName		= pGuild.GuildData.MasterName
	data.MemberNum		= int32(len(pGuild.GuildMember))
	data.GuildVal		= pGuild.GuildData.GuildVal
	data.IsRatify		= pGuild.GuildData.IsRatify
	data.Require		= pGuild.GuildData.Require
	data.Contribution	= pGuild.GuildData.Contribution

	temp1 := pGuild.FindMemberByJob(prpc.GJ_Premier)
	for i:=0;i<len(temp1) ;i++  {
		data.Member = append(data.Member,*temp1[i])
	}
	temp2 := pGuild.FindMemberByJob(prpc.GJ_VicePremier)
	for i:=0;i<len(temp2) ;i++  {
		data.Member = append(data.Member,*temp2[i])
	}

	//To Client data
	if player.session == nil {
		return
	}
	player.session.QueryGuildDetailsResult(data)

	logs.Info("RequestGuildDetails ",data)
}

func RequestjoinGuild(player *GamePlayer,guildId int32)  {
	if player==nil {
		return
	}
	if player.MyUnit.GetIProperty(prpc.IPT_PROMOTE) < int32(GetGlobalInt("C_CreatGuildLevel")) {
		return
	}
	pGuild := FindGuildById(guildId)
	if pGuild==nil {
		return
	}
	if FindGuildByPlayerId(player.MyUnit.InstId) != nil {
		return
	}

	if len(pGuild.GuildMember) >= GetGlobalInt("C_CreatGuildMemberMax") {
		return
	}

	if pGuild.GuildData.IsRatify {
		myVal := GetTianTiIdByVal(player.TianTiVal)			//天梯段位
		if myVal < pGuild.GuildData.Require {
			return
		}
		if len(pGuild.GuildData.RequestList) >= GetGlobalInt("C_CreatGuildMemberMax") {
			return
		}
		data := prpc.COM_GuildRequestData{}
		data.RoleId 		= player.MyUnit.InstId
		data.RoleName		= player.MyUnit.InstName
		data.Level			= player.MyUnit.GetIProperty(prpc.IPT_PROMOTE)
		data.Time			= time.Now().Unix()
		data.UnitId			= player.MyUnit.UnitId
		data.TianTiVal		= player.TianTiVal
		pGuild.AddGuildRequestList(data)
	}else {
		member := prpc.COM_GuildMember{}
		member.GuildId 		= pGuild.GuildData.GuildId
		member.RoleId		= player.MyUnit.InstId
		member.RoleName 	= player.MyUnit.InstName
		member.Level		= player.MyUnit.GetIProperty(prpc.IPT_PROMOTE)
		member.TianTiVal	= player.TianTiVal
		member.Job			= prpc.GJ_People
		member.UnitId		= player.MyUnit.UnitId
		isOK := AddGuildMember(member)
		if isOK {
			InsertGuildMember(member)
			pGuild.UpdateGuildVal()
		}
	}
}

func AcceptrequestGuild(player *GamePlayer,proposerId int64)  {
	if player==nil {
		return
	}

	if player.GuildId == 0 {
		return
	}
	
	pGuild := FindGuildById(player.GuildId)
	if pGuild==nil {
		return
	}
	if len(pGuild.GuildMember) >= GetGlobalInt("C_CreatGuildMemberMax") {
		return
	}

	member := pGuild.FindMember(player.MyUnit.InstId)
	if member==nil {
		return
	}
	if member.Job < prpc.GJ_VicePremier {
		return
	}

	proposerRequest := pGuild.FindRequest(proposerId)
	if proposerRequest==nil {
		return
	}

	guild := FindGuildByPlayerId(proposerId)
	if guild != nil {
		//申请人有帮派
		pGuild.DeleteGuildRequestList(proposerId)
		return
	}
	newMember := prpc.COM_GuildMember{}
	proposer := FindPlayerByInstId(proposerId)
	if proposer == nil {
		var p *prpc.SGE_DBPlayer
		if p = <- QueryPlayerById(proposerId); p!=nil {		//数据库里有这个人
			newMember.RoleName 	= p.Name
			newMember.RoleId	= p.InstId
			newMember.Level		= p.Unit.Level
			newMember.TianTiVal = p.TianTiVal
			newMember.GuildId	= pGuild.GuildData.GuildId
			newMember.Job		= prpc.GJ_People
			newMember.UnitId	= p.Unit.UnitId
		} else {
			logs.Info("AcceptRequestGuild Fetch DB Data Nil ",proposerId)
			return
		}
	} else {
		newMember.RoleName 	= proposer.MyUnit.InstName
		newMember.RoleId	= proposer.MyUnit.InstId
		newMember.Level		= proposer.MyUnit.GetIProperty(prpc.IPT_PROMOTE)
		newMember.TianTiVal = proposer.TianTiVal
		newMember.GuildId	= pGuild.GuildData.GuildId
		newMember.Job		= prpc.GJ_People
		newMember.UnitId	= proposer.MyUnit.UnitId
	}

	if !AddGuildMember(newMember) {
		return
	}
	InsertGuildMember(newMember)

	pGuild.DeleteGuildRequestList(proposerId)
}

func RefuserequestGuild(player *GamePlayer,proposerId int64)  {
	if player==nil {
		return
	}
	if player.GuildId == 0 {
		return
	}
	pGuild := FindGuildById(player.GuildId)
	if pGuild==nil {
		return
	}
	member := pGuild.FindMember(player.MyUnit.InstId)
	if member==nil {
		return
	}
	if member.Job < prpc.GJ_VicePremier {
		return
	}
	proposerRequest := pGuild.FindRequest(proposerId)
	if proposerRequest==nil {
		return
	}
	pGuild.DeleteGuildRequestList(proposerId)
}

func SycnGuildData(player *GamePlayer)  {
	if player == nil {
		return
	}
	if player.GuildId == 0 {
		return
	}
	pGuild := FindGuildById(player.GuildId)
	if pGuild == nil {
		return
	}
	if player.session == nil {
		return
	}
	pGuild.UpdateMemberList(player)
	player.session.InitGuildData(*pGuild.GuildData)
}

///-----------------------------------------------------------------------------------------

func (this *Guild)AddMember(member prpc.COM_GuildMember) bool {
	if this.FindMember(member.RoleId) != nil {
		return false
	}

	this.GuildMember = append(this.GuildMember,&member)

	this.UpdateMember(member,prpc.MLF_Add)

	return true
}

func (this *Guild)FindMember(playerInstId int64) *prpc.COM_GuildMember {
	for i:=0;i<len(this.GuildMember) ;i++  {
		if this.GuildMember[i].RoleId == playerInstId {
			return this.GuildMember[i]
		}
	}
	return nil
}

func (this *Guild)FindMemberByJob(job int) []*prpc.COM_GuildMember {
	var members []*prpc.COM_GuildMember
	for i := 0 ; i < len(this.GuildMember) ;i++  {
		if this.GuildMember[i].Job == job {
			members = append(members,this.GuildMember[i])
		}
	}
	return members
}

func (this *Guild)AddGuildRequestList(request prpc.COM_GuildRequestData)  {
	this.GuildData.RequestList = append(this.GuildData.RequestList,request)
	UpdateGuildRequestList(this.GuildData.GuildId,this.GuildData.RequestList)

	logs.Info("AddGuildRequestList", request)

	this.UpdateGuild()
}

func (this *Guild)DeleteGuildRequestList(player int64)  {
	for i:=0;i<len(this.GuildData.RequestList) ;i++  {
		if this.GuildData.RequestList[i].RoleId == player {
			this.GuildData.RequestList = append(this.GuildData.RequestList[:i], this.GuildData.RequestList[i+1:]...)
			break
		}
	}
	UpdateGuildRequestList(this.GuildData.GuildId,this.GuildData.RequestList)

	logs.Info("DeleteGuildRequestList ",player)

	this.UpdateGuild()
}

func (this *Guild)UpdateGuild()  {
	for _,m := range this.GuildMember{
		player := FindPlayerByInstId(m.RoleId)
		if player==nil {
			continue
		}
		//To Client this.GuildData
		if player.session == nil {
			continue
		}
		player.session.InitGuildData(*this.GuildData)
	}
}

func (this *Guild)UpdateGuildMemberVal(player int64,val int32)  {
	member := this.FindMember(player)
	if member==nil {
		return
	}
	member.TianTiVal = val
	UpdateDBGuildMemberVal(player,val)
	this.UpdateGuildVal()

	logs.Info("UpdateGuildMemberVal player = ",player," TianTiVal = " ,member.TianTiVal ," UpdateGuildVal =",this.GuildData.GuildVal)
}

func (this *Guild)UpdateGuildVal()  {
	this.GuildData.GuildVal = 0
	for _,m := range this.GuildMember{
		this.GuildData.GuildVal += m.TianTiVal
	}
	UpdateDBGuildVal(this.GuildData.GuildId,this.GuildData.GuildVal)
}

func (this *Guild)FindRequest(player int64) *prpc.COM_GuildRequestData {
	for i:=0;i<len(this.GuildData.RequestList) ;i++  {
		if this.GuildData.RequestList[i].RoleId == player {
			return &this.GuildData.RequestList[i]
		}
	}
	return nil
}

func (this *Guild)SetGuildRequestFlag(isFlag bool,require int32)  {
	this.GuildData.IsRatify = isFlag
	if this.GuildData.IsRatify {
		this.GuildData.Require = require
	}
	UpdateDBGuildRatify(this.GuildData.GuildId,this.GuildData.IsRatify,this.GuildData.Require)
	//To Client
	this.UpdateGuild()
	logs.Info("SetGuildRequestFlag Guild=",this.GuildData.GuildId,isFlag,require)
}

func (this *Guild)DelMember(roleId int64)  {
	tempMem := prpc.COM_GuildMember{}
	for i:=0;i<len(this.GuildMember) ;i++  {
		if this.GuildMember[i].RoleId == roleId {
			tempMem = *this.GuildMember[i]
			this.GuildMember = append( this.GuildMember[:i],  this.GuildMember[i+1:]...)
			break
		}
	}
	this.UpdateMember(tempMem,prpc.MLF_Delete)

	logs.Info("Guild DelMember = ",roleId)
}

func (this *Guild)GuildChangeJob(myId int64,roleId int64,job int)  {

	if myId == roleId {
		return
	}

	sender := this.FindMember(myId)
	if sender==nil {
		return
	}
	if sender.Job != prpc.GJ_Premier {
		return
	}
	roleMember := this.FindMember(roleId)
	if roleMember==nil {
		return
	}
	if sender.Job <= job {
		return
	}
	if job==prpc.GJ_VicePremier {
		num := len(this.FindMemberByJob(prpc.GJ_VicePremier))
		if num >= GetGlobalInt("C_VicePremierMaxNum") {
			return
		}
	}
	isOK := <- UpdateDBGuildMemberJob(roleId,job)
	if !isOK {
		return
	}
	roleMember.Job = job
	//To Client
	this.UpdateMember(*roleMember,prpc.MLF_ChangePosition)

	logs.Info("GuildChangeJob ",roleMember)
}

func (this *Guild)GuildChat(chat prpc.COM_Chat)  {
	for _,m:=range this.GuildMember{
		player := FindPlayerByInstId(m.RoleId)
		if player==nil {
			continue
		}
		if player.session == nil {
			continue
		}
		player.session.ReceiveChat(chat)
	}
}

func (this *Guild)GuildMemberOnLine(player *GamePlayer)  {
	if player==nil {
		return
	}
	member := this.FindMember(player.MyUnit.InstId)
	if member == nil {
		return
	}
	member.IsOnline = true
	player.GuildId = this.GuildData.GuildId

	player.SycnGuildAssistants()
}

func (this *Guild)GuildMemberOffOnLine(player int64)  {
	member := this.FindMember(player)
	if member == nil {
		return
	}
	member.IsOnline = false
}

func (this *Guild)GuildMemberLevelUp(player int64,level int32)  {
	member := this.FindMember(player)
	if member == nil {
		return
	}
	member.Level = level
	UpdateDBGuildMemberLevel(player,level)
}

func (this *Guild)UpdateMemberList(player *GamePlayer)  {
	if player==nil {
		return
	}
	members := []prpc.COM_GuildMember{}
	for _,m := range this.GuildMember{
		if m == nil {
			logs.Info("UpdateMemberList ERROR UpdateMemberList ERROR")
			continue
		}
		members = append(members,*m)
	}
	if player.session == nil {
		return
	}
	player.session.InitGuildMemberList(members)
}

func (this *Guild)UpdateMember(member prpc.COM_GuildMember,mlf int)  {
	for _,m:=range this.GuildMember{
		player := FindPlayerByInstId(m.RoleId)
		if player==nil {
			continue
		}
		if player.session == nil {
			continue
		}
		player.session.ModifyGuildMemberList(member, mlf)
	}
}

func (this *Guild)Kickout(sender int64,roleId int64)  {
	if sender == roleId {
		return
	}

	senderm := this.FindMember(sender)
	if senderm==nil {
		return
	}
	if senderm.Job < prpc.GJ_VicePremier {
		return
	}
	roleMember := this.FindMember(roleId)
	if roleMember==nil {
		return
	}
	DelGuildMember(roleId)
}

func (this *Guild)Leave(roleId int64)  {
	roleMember := this.FindMember(roleId)
	if roleMember==nil {
		return
	}
	if roleMember.Job == prpc.GJ_Premier {
		return
	}
	DelGuildMember(roleId)
}

/////////////////////////////////////求助AND捐赠/////////////////////////////////////

func GetAssistantItemNum(itemId int32) int32 {
	itemData := GetItemTableDataById(itemId)
	if itemData==nil {
		return 0
	}
	if itemData.ItemQuality == 1 {
		return 5
	} else if itemData.ItemQuality == 2 {
		return 4
	} else if itemData.ItemQuality == 3 {
		return 3
	} else if itemData.ItemQuality == 4 {
		return 2
	} else if itemData.ItemQuality == 5 || itemData.ItemQuality == 6 {
		return 1
	}
	return 0
}

func (player *GamePlayer)CheckGuildAssistantTime()  {
	if player==nil {
		return
	}
	if player.GuildId == 0 {
		return
	}
	if player.AssistantCreateTime == 0 {
		return
	}
	if player.AssistantId == 0 {
		return
	}
	nowTime := time.Now()
	requestTime := time.Unix(player.AssistantCreateTime,0)

	if nowTime.Day() != requestTime.Day() || nowTime.Month() != requestTime.Month() || nowTime.Year() != requestTime.Year() {
		DeleteGuildAssistant(player.AssistantId)
		player.AssistantCreateTime = 0
		player.AssistantId = 0
	}
}

func (player *GamePlayer)AddGuildAssistant(item int32)  {
	if player == nil {
		return
	}
	if player.GuildId == 0 {
		return
	}

	if player.AssistantCreateTime != 0 {
		return
	}

	if player.AssistantId != 0 {
		return
	}

	pGuild := FindGuildById(player.GuildId)
	if pGuild == nil {
		return
	}
	sge := prpc.SGE_DBGuildAssistant{}
	sge.GuildId 		= pGuild.GuildData.GuildId
	sge.ItemId			= item
	sge.RoleName		= player.MyUnit.InstName
	sge.Id				= GenGuildAssistantInstId()
	sge.MaxCount		= GetAssistantItemNum(item)
	InsertGuildAssistant(sge)

	player.AssistantCreateTime = time.Now().Unix()
	player.AssistantId = sge.Id
	//广播PLAYER的请求
	com := prpc.COM_Assistant{}
	com.Id 			= sge.Id
	com.CrtCount	= sge.CrtCount
	com.MaxCount	= sge.MaxCount
	com.ItemId		= sge.ItemId
	com.PlayerName  = sge.RoleName
	BroadcastAssistant(player.GuildId,com,"")

	logs.Info("AddGuildAssistant ",sge)
}

func (player *GamePlayer)GuildAssistantItem(ass int32)  {

	if player.GuildId == 0 {
		return
	}

	pGuild := FindGuildById(player.GuildId)
	if pGuild==nil {
		return
	}

	data := <- FindGuildAssistantById(ass)
	if data==nil {
		return
	}

	if player.MyUnit.InstName == data.RoleName {
		return
	}

	for _,id := range data.Donator{
		if id == player.MyUnit.InstId{
			return
		}
	}
	items := player.GetBagItemByTableId(data.ItemId)
	if items == nil {
		return
	}
	quality := GetItemQualityById(data.ItemId)

	dropid := GetGiveGiftDropIdByQuality(quality)

	if dropid == 0 {
		return
	}
	
	player.DelItemByTableId(data.ItemId,1)

	receiver := FindPlayerByInstName(data.RoleName)
	if receiver != nil {
		receiver.GiveDrop(dropid)
	}else {
		data.CatchNum++
	}

	pGuild.GuildData.Contribution++

	UpdateDBGuildContribution(pGuild.GuildData.GuildId,pGuild.GuildData.Contribution)

	data.CrtCount++
	data.Donator = append(data.Donator,player.MyUnit.InstId)

	UpdateGuildAssistant(*data)
	//广播PLAYER的请求
	com := prpc.COM_Assistant{}
	com.Id 			= data.Id
	com.CrtCount	= data.CrtCount
	com.MaxCount	= data.MaxCount
	com.ItemId		= data.ItemId
	com.PlayerName  = data.RoleName
	BroadcastAssistant(player.GuildId,com,player.MyUnit.InstName)

	logs.Info("GuildAssistantItem ",data)
}

func DeleteGuildAssistant(assistantId int32)()  {
	DelGuildAssistant(assistantId)
	logs.Info("DeleteGuildAssistant ",assistantId)
}

func (player *GamePlayer)SycnGuildAssistants()  {
	if player==nil {
		return
	}
	if player.GuildId == 0 {
		return
	}
	//将所有现有请求同步给PLAYER
	data := <- FindGuildAssistantByGuildId(player.GuildId)
	if data == nil {
		return
	}

	infos := []prpc.COM_Assistant{}

	for _,m := range data{
		com := prpc.COM_Assistant{}
		com.Id 			= m.Id
		com.CrtCount	= m.CrtCount
		com.MaxCount	= m.MaxCount
		com.ItemId		= m.ItemId
		com.PlayerName  = m.RoleName
		infos = append(infos,com)
	}

	player.session.SycnGuildAssistant(infos)
}

func BroadcastAssistant(guildid int32,info prpc.COM_Assistant,donator string)  {
	pGuild := FindGuildById(guildid)
	if pGuild==nil {
		return
	}
	for _,m:=range pGuild.GuildMember{
		player := FindPlayerByInstId(m.RoleId)
		if player==nil {
			continue
		}
		if player.session == nil {
			continue
		}
		player.session.UpdateGuildAssistant(info,donator)
	}
}




