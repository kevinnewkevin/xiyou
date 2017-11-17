require "FairyGUI"

squadSetting = fgui.window_class(WindowBase)
local Window;

local typeList;
local setCom;
local requestCom;
local crtType;

local guildName;
local guildId;
local guildCount;
local guildScore;
local guildGiftWeek;
local guildNeedCheck;
local guildNeededLoader;
local guildNeededBtn;

local requestList;
local guildNeededCom;

function squadSetting:OnEntry()
	Window = squadSetting.New();
	Window:Show();
end

function squadSetting:GetWindow()
	return Window;
end

function squadSetting:OnInit()
	self.contentPane = UIPackage.CreateObject("bangpai", "bangpaishezhi_com").asCom;
	self:Center();
	self.closeButton = self.contentPane:GetChild("n37");

	typeList = self.contentPane:GetChild("n4");
	setCom = self.contentPane:GetChild("n5");
	requestCom = self.contentPane:GetChild("n39");
	typeList.onClickItem:Add(squadSetting_OnTypeSelect);

	guildName = setCom:GetChild("n12");
	guildId = setCom:GetChild("n13");
	guildCount = setCom:GetChild("n14");
	guildScore = setCom:GetChild("n15");
	guildGiftWeek = setCom:GetChild("n16");
	guildNeedCheck = setCom:GetChild("n21");
	guildNeedCheck.onClick:Add(squadSetting_OnNeedCheck);
	guildNeededLoader = setCom:GetChild("n33").asLoader;
	guildNeededBtn = setCom:GetChild("n22");
	guildNeededBtn.onClick:Add(squadSetting_OnNeededChange);

	requestList = requestCom:GetChild("n39").asList;
	requestList:SetVirtual();
	requestList.itemRenderer = squadSetting_RenderListItem;

	guildNeededCom = UIPackage.CreateObject("bangpai", "xiugaitiaojian_com").asCom;
	guildNeededCom.fairyBatching = true;
	guildNeededCom:GetChild("n24").asList.onClickItem:Add(squadSetting_OnChangeNeed);
	guildNeededCom:RemoveFromParent();
	guildNeededCom:GetChild("n2").onClick:Add(squadSetting_OnChangeNeedClose);

	crtType = 0;
	typeList.selectedIndex = crtType;

	squadSetting_FlushData();
end

function squadSetting_OnTypeSelect()
	crtType = typeList.selectedIndex;
	UIManager.SetDirty("squadSetting");
end

function squadSetting:OnUpdate()
	if UIManager.IsDirty("squadSetting") then
		squadSetting_FlushData();
		UIManager.ClearDirty("squadSetting");
	end
end

function squadSetting:OnTick()
	
end

function squadSetting:isShow()
	return Window.isShowing;
end

function squadSetting:OnDispose()
	Window:Dispose();
end

function squadSetting:OnHide()
	Window:Hide();
end

function squadSetting_OnNeedCheck()
	GuildSystem.myGuild.IsRatify = not GuildSystem.myGuild.IsRatify;
	UIManager.SetDirty("squadSetting");
	Proxy4Lua.ChangeJoinGuildFlag(GuildSystem.myGuild.IsRatify, GuildSystem.myGuild.Require);
end

function squadSetting_OnNeededChange(context)
	GRoot.inst:ShowPopup(guildNeededCom, context.sender, false);
end

function squadSetting_OnChangeNeed(context)
	GuildSystem.myGuild.Require = guildNeededCom:GetChild("n24").selectedIndex;
	Proxy4Lua.ChangeJoinGuildFlag(GuildSystem.myGuild.IsRatify, GuildSystem.myGuild.Require);
	guildNeededCom:RemoveFromParent();
	UIManager.SetDirty("squadSetting");
end

function squadSetting_OnChangeNeedClose(context)
	guildNeededCom:RemoveFromParent();
end

function squadSetting_RenderListItem(index, obj)
	if obj == nil then
		return;
	end

	local data = GuildSystem.requestList[index];
	if data == nil then
		return;
	end
	local playerName = obj:GetChild("n40");
	local headCom = obj:GetChild("n39");
	local headIcon = headCom:GetChild("n5").asLoader;
	local lv = headCom:GetChild("n3").asTextField;
	local dayBefore = obj:GetChild("n42");
	local tiantiLv = obj:GetChild("n41").asLoader;
	local add = obj:GetChild("n43").asButton;
	local ignore = obj:GetChild("n44").asButton;
	add.onClick:Add(squadSetting_OnAdd);
	ignore.onClick:Add(squadSetting_OnIgnore);
	obj.data = data.RoleId;

	playerName.text = data.RoleName;
	local eData = EntityData.GetData(data.UnitId);
	local dData;
	if eData ~= nil then
		dData = DisplayData.GetData(eData._DisplayId);
	end
	if dData ~= nil then
		headIcon.url = "ui://" .. dData._HeadIcon;
	else
		headIcon.url = "";
	end
	lv.text = data.Level;
	dayBefore.text = TimerManager.TimeAgo(data.Time);
	tiantiLv.url = "ui://bangpai/xiao_duanwei" .. GamePlayer.RankLevel(data.TianTiVal);
end

function squadSetting_OnAdd(context)
	if context.sender.data == nil then
		return;
	end

	Proxy4Lua.AcceptRequestGuild(context.sender.data);
	GuildSystem.DeleteGuildRequest(context.sender.data);
	UIManager.SetDirty("squadSetting");
end

function squadSetting_OnIgnore(context)
	if context.sender.data == nil then
		return;
	end

	Proxy4Lua.RefuseRequestGuild(context.sender.data);
	GuildSystem.DeleteGuildRequest(context.sender.data);
	UIManager.SetDirty("squadSetting");
end
	
function squadSetting_FlushData()
	requestList.numItems = GuildSystem.requestList.Count;

	if crtType == 0 then
		setCom.visible = true;
		requestCom.visible = false;
	else
		setCom.visible = false;
		requestCom.visible = true;
	end
	--TODO
	guildName.text = "帮派名称: " .. GuildSystem.myGuild.GuildName;
	guildId = "帮派编号: " .. GuildSystem.myGuild.GuildId;
	guildCount = "帮派人数: " .. GuildSystem.myGuild.GuildName;
	guildScore = "帮派积分: " .. GuildSystem.myGuild.GuildVal;
	guildGiftWeek = "每周捐献: " .. GuildSystem.myGuild.Contribution;
	if GuildSystem.myGuild.IsRatify then
		guildNeedCheck:GetController("button").selectedIndex = 0;
	else
		guildNeedCheck:GetController("button").selectedIndex = 1;
	end
	guildNeededLoader.url = "ui://bangpai/xiao_duanwei" .. GamePlayer.RankLevel(GuildSystem.myGuild.Require);
end