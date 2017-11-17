require "FairyGUI"

squadList = fgui.window_class(WindowBase)
local Window;

local squads;

local createBtn;
local createCom;  
local createName;
local createGold;
local createLevel;
local createCancel;
local createConfirm;
local createClose;
local findBtn;
local findName;
local searchBtn; 
local searchCom;
local searchCancel;
local searchConfirm;
local searchName;
local searchNum;
local searchPlayerNum;
local searchFen;
local searchTiao;
local searchjuan;
local searchType;
local searchList;



function squadList:OnEntry()
	Window = squadList.New();
	Window:Show();
end

function squadList:GetWindow()
	return Window;
end

function squadList:OnInit()
	self.contentPane = UIPackage.CreateObject("squadList", "squadList_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n2").asButton;

	squads = self.contentPane:GetChild("n4").asList;
	squads:SetVirtual();
	squads.itemRenderer = squadList_RenderListItem;
	self.contentPane:GetChild("n22").visible = false;
	self.contentPane:GetChild("n23").visible = false;
	self.contentPane:GetChild("n25").visible = false;
	self.contentPane:GetChild("n26").visible = false;
	createBtn = self.contentPane:GetChild("n27").asButton; 
	createBtn.onClick:Add(squadList_OnCreate);
	createCom = self.contentPane:GetChild("n28").asCom;
	createCom.visible = false;
	createName = createCom:GetChild("n4");
	createGold = createCom:GetChild("n11");
	createLevel = createCom:GetChild("n10");
	createCancel = createCom:GetChild("n15");
	createClose = createCom:GetChild("n16");
	createConfirm = createCom:GetChild("n14");
	createCancel.onClick:Add(squadList_OnCreateCancel);
	createClose.onClick:Add(squadList_OnCreateCancel);
	createConfirm.onClick:Add(squadList_OnCreateConfirm);
	findName = self.contentPane:GetChild("n20");
	findBtn = self.contentPane:GetChild("n21").asButton;
	searchCom = self.contentPane:GetChild("n29").asCom;
	searchCom.visible = false;
	searchBtn = searchCom:GetChild("n22");
	searchBtn.onClick:Add(squadList_OnSearch);
	searchName = searchCom:GetChild("n13");
	searchCancel = searchCom:GetChild("n2");
	searchConfirm = searchCom:GetChild("n22");
	searchNum = searchCom:GetChild("n14");
	searchPlayerNum = searchCom:GetChild("n15");
	searchFen = searchCom:GetChild("n16");
	searchjuan = searchCom:GetChild("n17");
	searchType = searchCom:GetChild("n19");
	searchTiao =searchCom:GetChild("n18");
	searchList = searchCom:GetChild("n20");
	local needGold = Define.GetInt("CreateGuild");
	createGold.text = ""..needGold;
	createLevel.text = "3";
	searchList.itemRenderer = searchList_RenderListItem;
	searchCancel.onClick:Add(squadList_OnSearchCancel);
	findBtn.onClick:Add(squadList_OnFindGuild);
	searchConfirm.onClick:Add(squadList_OnSearchConfirm);

	squadList_FlushData();
end

function squadList_RenderListItem(index, obj)
	local checkBtn = obj:GetChild("n6").asButton;
	local guild = GuildSystem.guildViewerList[index];
	if guild == nil then
	end
	obj:GetChild("n1").text = guild.GuildId .."";
	obj:GetChild("n3").text = guild.GuildName .."";
	obj:GetChild("n4").text = guild.MasterName .."";
	obj:GetChild("n5").text = guild.MemberNum.."";
	obj:GetChild("n2").text = guild.GuildVal.."";
	checkBtn.data = guild.GuildId;
	checkBtn.onClick:Add(squadList_OnCheckSquad);
end

function searchList_RenderListItem(index, obj)

	local info = GuildSystem.searchData.Member[index];
	local icon = obj:GetChild("n1");
	icon:GetChild("n3").text = info.Level .."";
	local edata = EntityData.GetData(info.UnitId);
	local ddata = DisplayData.GetData(edata._DisplayId);
	icon:GetChild("n5").asLoader.url = "ui://" .. ddata._HeadIcon;
	obj:GetChild("n4").text = info.RoleName .."";
	obj:GetChild("n5").text = info.TianTiVal .."";
	obj:GetChild("n3").asLoader.url = "ui://squadList/xiao_duanwei" .. GamePlayer.RankLevel(info.TianTiVal);
	if info.Job == 3 then
		obj:GetChild("n2").asLoader.url = "ui://squadList/cy_07";
	elseif info.Job == 2 then
		obj:GetChild("n2").asLoader.url = "ui://squadList/cu_06"
	end
end

function squadList_OnCheckSquad(context)
	Proxy4Lua.QueryGuildDetails(context.sender.data);
	GuildSystem.IsShowSearch = true;
end

function squadList_OnCreate()
	createCom.visible = true;
end

function squadList_OnCreateCancel()
	createCom.visible = false;  
end                                                                             

function squadList_OnCreateConfirm()
	if createName.text == "" then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "请输入帮派名", true);
		return;
	end
	local needGold = Define.GetInt("CreateGuild");
	if GamePlayer._Data.IProperties[8] < needGold then
		local MessageBox1 = UIManager.ShowMessageBox();
		MessageBox1:SetData("提示", "金币不够", true);
		return;
	end
	Proxy4Lua.CreateGuild(createName.text);
end

function squadList_OnSearch(context)
	local guild = GuildSystem.searchData;
	if guild.Require > GamePlayer.GetTianTiLevel() then
		local MessageBox1 = UIManager.ShowMessageBox();
		MessageBox1:SetData("提示", "不满足条件", true);
		return;
	end
	Proxy4Lua.RequestJoinGuild(context.sender.data);
	Proxy4Lua.PopMsg("申请发送成功");
	searchCom.visible = false;
end

function squadList_OnSearchCancel()
	searchCom.visible = false;
end

function squadList_OnSearchConfirm()
	searchCom.visible = false;
end

function squadList_OnFindGuild()
	if findName.text == "" then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "请输入家族名", true);
		return;
	end

	local guild = GuildSystem.findGuildViewer(findName.text);
	if guild == nil then
		Proxy4Lua.PopMsg("没有此家族");
		return;
	end
	findName.text = "";
	Proxy4Lua.QueryGuildDetails(guild.GuildId);
	GuildSystem.IsShowSearch = true;
end



function squadList:OnUpdate()
	if UIManager.IsDirty("squadList") then
		squadList_FlushData();
		UIManager.ClearDirty("squadList");
	end
end

function squadList:OnTick()
	
end

function squadList:isShow()
	return Window.isShowing;
end

function squadList:OnDispose()
	Window:Dispose();
end

function squadList:OnHide()
	findName.text = "";
	searchCom.visible = false;
	createName.text = "";
	createCom.visible = false;
	Window:Hide();
end

function squadList_FlushData()
	if GuildSystem.IsShowSearch == true then
		searchCom.visible = true;
		squadList_UpdateSearch();
		GuildSystem.IsShowSearch = false;
	end

	squads.numItems = GuildSystem.guildViewerList.Count;
end


function squadList_UpdateSearch()
	local info = GuildSystem.searchData;
	if info == nil then
		return;
	end

	searchName.text = "帮派名称："..info.GuildName;
	searchNum.text = "帮派编号："..info.GuildId.."";
	searchPlayerNum.text = "帮派人数："..info.MemberNum.."/50";
	searchFen.text = "帮派积分："..info.GuildVal.."";
	searchjuan.text = "帮派捐献："..info.Contribution;
	searchTiao.text = "加入条件：天梯等级"..info.Require;
	if info.IsRatify == true then
		searchType.text = "类型：需申请";
	else
		searchType.text = "类型:不需申请";
	end
	if info.Member == nil then
		searchList.numItems = 0;
	else
		searchList.numItems = info.Member.Length;
	end
	searchBtn.data = info.GuildId;
end

