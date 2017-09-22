require "FairyGUI"

jineng = fgui.window_class(WindowBase)
local Window;

--info group
local skillName;
local skillLv;
local skillNowLv;
local skillDesc;
local skillIcon;
local skillLock;
local skillUnlockCondition;

--lv up group
local itemQua;
local itemIcon;
local itemNum;
local itemBar;
local itemBarNum;
local needGold;
local upgradeBtn;

--list
local skillItemUrl = "ui://jineng/jineng_Button";
local activeList;
local passiveList;
local criticalList;

--equipment
local playerSkillList;
local skill1;
local skill2;
local skill3;
local skill4;

local crtSelectRoleSkillId;
local crtSelectRoleSkilltype;
local crtSelectIsDown;

local isNewOpen;
function jineng:OnEntry()
	Window = jineng.New();
	Window:Show();
end

function jineng:GetWindow()
	return Window;
end

function jineng:OnInit()
	self.contentPane = UIPackage.CreateObject("jineng", "jineng_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n2").asButton;

	skillName = self.contentPane:GetChild("n21").asTextField;
	skillNowLv = self.contentPane:GetChild("n24").asTextField;
	skillDesc = self.contentPane:GetChild("n29").asTextField;

	local item = self.contentPane:GetChild("n31").asCom;
	itemQua = item:GetChild("n0").asLoader;
	itemIcon = item:GetChild("n1").asLoader;
	itemNum = item:GetChild("n2").asTextField;
	itemBar = self.contentPane:GetChild("n32").asProgress;
	itemBarNum = self.contentPane:GetChild("n38").asTextField;
	needGold = self.contentPane:GetChild("n36").asTextField;
	upgradeBtn = self.contentPane:GetChild("n37").asButton;
	upgradeBtn.onClick:Add(jineng_OnUpgrade);

	local iconCom = self.contentPane:GetChild("n20").asCom;
	iconCom.touchable = false;
	skillIcon = iconCom:GetChild("n8").asLoader;
	skillLock = iconCom:GetChild("n12");
	skillLv = iconCom:GetChild("n7").asTextField;
	skillUnlockCondition = iconCom:GetChild("n13").asTextField;
	iconCom:GetChild("n10").visible = false;

	local skillList = self.contentPane:GetChild("n16").asList;
	local acCom = skillList:GetChildAt(0).asCom;
	local paCom = skillList:GetChildAt(1).asCom;
	local crCom = skillList:GetChildAt(2).asCom;
	activeList = acCom:GetChild("n6").asList;
	passiveList = paCom:GetChild("n6").asList;
	criticalList = crCom:GetChild("n6").asList;

	playerSkillList = self.contentPane:GetChild("n40").asList;
	skill1 = playerSkillList:GetChildAt(0).asCom;
	skill2 = playerSkillList:GetChildAt(1).asCom;
	skill3 = playerSkillList:GetChildAt(2).asCom;
	skill4 = playerSkillList:GetChildAt(3).asCom;
	skill1:GetChild("n12").visible = false;
	skill1:GetChild("n13").visible = false;
	skill2:GetChild("n12").visible = false;
	skill2:GetChild("n13").visible = false;
	skill3:GetChild("n12").visible = false;
	skill3:GetChild("n13").visible = false;
	skill4:GetChild("n12").visible = false;
	skill4:GetChild("n13").visible = false;

	skill1.data = {};
	skill1.data.rsId = 0;
	skill1.data.down = 1;
	skill2.data = {};
	skill2.data.rsId = 0;
	skill2.data.down = 1;
	skill3.data = {};
	skill3.data.rsId = 0;
	skill3.data.down = 1;
	skill4.data = {};
	skill4.data.rsId = 0;
	skill4.data.down = 1;
	skill1.onDrop:Add(jineng_OnDropSkill);
	skill2.onDrop:Add(jineng_OnDropSkill);
	skill3.onDrop:Add(jineng_OnDropSkill);
	skill4.onDrop:Add(jineng_OnDropSkill);

	skill1.onClick:Add(jineng_OnSkillSelected);
	skill2.onClick:Add(jineng_OnSkillSelected);
	skill3.onClick:Add(jineng_OnSkillSelected);
	skill4.onClick:Add(jineng_OnSkillSelected);

	jineng_Clear();
	jineng_FlushData();
end

function jineng_OnSkillSelected(context)
	local rsData = RoleSkillData.GetData(context.sender.data.rsId);
	if rsData == nil then
		return;
	end

	crtSelectRoleSkillId = context.sender.data.rsId;
	crtSelectRoleSkilltype = RoleSkillData.GetData(crtSelectRoleSkillId)._Type;
	if crtSelectRoleSkilltype == 0 then
		activeList:SelectNone();
		criticalList:SelectNone();
	end
	if crtSelectRoleSkilltype == 1 then
		passiveList:SelectNone();
		criticalList:SelectNone();
	end
	if crtSelectRoleSkilltype == 2 then
		activeList:SelectNone();
		passiveList:SelectNone();
	end

	if context.sender.data.down == 0 then
		playerSkillList:SelectNone();
	else
		activeList:SelectNone();
		passiveList:SelectNone();
		criticalList:SelectNone();
	end
	crtSelectIsDown = context.sender.data.down;
	jineng_OnlyLeftFlushData();
end

function jineng_OnDragSkill(context)
	jineng_OnSkillSelected(context);
	context:PreventDefault();
	DragDrop.inst:StartDrag(context.sender, context.sender:GetChild("n8").asLoader.url, context.sender.data, context.data);
end

function jineng_OnDropSkill(context)
	local rsData = RoleSkillData.GetData(context.sender.data.rsId);
	if rsData == nil then
		return;
	end
	if rsData._Type == crtSelectRoleSkilltype then
		local idx = Proxy4Lua.GetIndexBySkillID(rsData._SkillId);
		if idx ~= -1 then
			Proxy4Lua.EquipSkill(idx, crtSelectRoleSkillId);
		end
	end
end

function jineng:OnUpdate()
	if UIManager.IsDirty("jineng") then
		jineng_FlushData();
		UIManager.ClearDirty("jineng");
	end

	if isNewOpen and activeList.numItems > 0 then
		activeList.selectedIndex = 0;
		crtSelectRoleSkillId = activeList:GetChildAt(0).data.rsId;
		jineng_OnlyLeftFlushData();
		isNewOpen = false;
	end
end

function jineng:OnTick()
	
end

function jineng:isShow()
	return Window.isShowing;
end

function jineng:OnDispose()
	Window:Dispose();
end

function jineng:OnHide()
	jineng_Clear();
	Window:Hide();
end

function jineng_OnUpgrade()
	local rsData = RoleSkillData.GetData(crtSelectRoleSkillId);
	if rsData ~= nil then
		local rsuData = RoleSkillUpdateData.GetData(rsData._SkillId);
		if rsuData ~= nil then
			UIParamHolder.Set("jineng", rsuData._SkillId);
			UIParamHolder.Set("jineng2", rsuData._NextId);
			Proxy4Lua.UpdateSkill(rsData._ID, rsData._SkillId);
		end
	end
end

function jineng_Clear()
	crtSelectRoleSkilltype = -1;
	crtSelectRoleSkillId = 0;
	playerSkillList:SelectNone();
	criticalList:SelectNone();
	passiveList:SelectNone();
	activeList:SelectNone();
	isNewOpen = true;
end

function jineng_OnlyLeftFlushData()

	upgradeBtn.enabled = false;
	local playerlv = GamePlayer._Data.IProperties[9];
	if crtSelectRoleSkillId ~= 0 then
		local rsData = RoleSkillData.GetData(crtSelectRoleSkillId);
		if rsData ~= nil then
			local skilldata = SkillData.GetData(rsData._SkillId);
			if skilldata ~= nil then
				skillName.text = skilldata._Name;
				skillNowLv.text = skilldata._Level;
				skillLv.text = skilldata._Level;
				skillIcon.url = "ui://" .. skilldata._Icon;
				skillDesc.text = skilldata._Desc;
				skillLock.visible = playerlv < rsData._OpenLv;
				skillUnlockCondition.visible = playerlv < rsData._OpenLv;
				skillUnlockCondition.text = rsData._OpenLv .. "级解锁";

				local rsuData = RoleSkillUpdateData.GetData(skilldata._Id);
				if rsuData ~= nil then
					local iData = ItemData.GetData(rsuData._NeedItem);
					if iData ~= nil then
						itemQua.url = "ui://" .. iData._IconBack;
						itemIcon.url = "ui://" .. iData._Icon;
						print("ui://" .. iData._IconBack);
						print("ui://" .. iData._Icon);
						local has = BagSystem.GetItemMaxNum(iData._Id);
						itemNum.text = has;
						itemBar.value = has / rsuData._NeedNum * 100;
						itemBarNum.text = has .. "/" .. rsuData._NeedNum;
						if has < rsuData._NeedNum then
							itemBarNum.color = Color.red;
						else
							itemBarNum.color = Color.white;
						end
						if has >= rsuData._NeedNum and GamePlayer._Data.IProperties[6] >= rsuData._NeedGold and playerlv >= rsData._OpenLv then
							upgradeBtn.enabled = true;
						end
					else
						itemQua.url = "";
						itemIcon.url = "";
						itemNum.text = "";
						itemBar.value = 0;
						itemBarNum.text = "0/" ..rsuData._NeedNum;
					end
					needGold.text = rsuData._NeedGold;
					if GamePlayer._Data.IProperties[6] < rsuData._NeedGold then
						needGold.color = Color.red;
					else
						needGold.color = Color.white;
					end
				end
			end
		end
	else
		skillName.text = "";
		skillNowLv.text = "";
		skillLv.text = "";
		skillIcon.url = "";
		skillDesc.text = "";
		skillLock.visible = false;
		skillUnlockCondition.visible = false;

		itemQua.url = "";
		itemIcon.url = "";
		itemNum.text = "";
		itemBar.value = 0;
		itemBarNum.text = "0/0";
		needGold.text = "0";
	end

	if crtSelectIsDown == 0 then
		if crtSelectRoleSkilltype == 0 then
			skill1:GetChild("n10").visible = true;
			skill2:GetChild("n10").visible = false;
			skill3:GetChild("n10").visible = false;
			skill4:GetChild("n10").visible = false;
		elseif crtSelectRoleSkilltype == 1 then
			skill1:GetChild("n10").visible = false;
			skill2:GetChild("n10").visible = true;
			skill3:GetChild("n10").visible = true;
			skill4:GetChild("n10").visible = false;
		elseif crtSelectRoleSkilltype == 2 then
			skill1:GetChild("n10").visible = false;
			skill2:GetChild("n10").visible = false;
			skill3:GetChild("n10").visible = false;
			skill4:GetChild("n10").visible = true;
		else
			skill1:GetChild("n10").visible = false;
			skill2:GetChild("n10").visible = false;
			skill3:GetChild("n10").visible = false;
			skill4:GetChild("n10").visible = false;
		end
	else
		skill1:GetChild("n10").visible = false;
		skill2:GetChild("n10").visible = false;
		skill3:GetChild("n10").visible = false;
		skill4:GetChild("n10").visible = false;
	end
end

function jineng_FlushData()

	local playerlv = GamePlayer._Data.IProperties[9];

	jineng_OnlyLeftFlushData();

	passiveList:RemoveChildrenToPool();
	activeList:RemoveChildrenToPool();
	criticalList:RemoveChildrenToPool();

	
	local allData = RoleSkillData.metaData;
	local skilldata;
	local skillitem;
	local skillicon;
	local skilllv;
	local skilllock;
	local skillunlockcond;
	for i=1, allData.Count do
		if allData[i-1]._Type == 0 then	--passive
			skilldata = SkillData.GetData(allData[i-1]._SkillId);
			skillitem = passiveList:AddItemFromPool(skillItemUrl);
			skillitem:GetChild("n10").visible = false;
			skillicon = skillitem:GetChild("n8").asLoader;
			skillicon.url = "ui://" .. skilldata._Icon;
			skilllv = skillitem:GetChild("n7").asTextField;
			skilllv.text = skilldata._Level;
			skilllock = skillitem:GetChild("n12");
			skillunlockcond = skillitem:GetChild("n13").asTextField;
			skilllock.visible = playerlv < allData[i-1]._OpenLv;
			skillunlockcond.visible = playerlv < allData[i-1]._OpenLv;
			skillunlockcond.text = allData[i-1]._OpenLv .. "级解锁";

			skillitem.data = {};
			skillitem.data.rsId = allData[i-1]._ID;
			skillitem.data.down = 0;
			skillitem.onClick:Add(jineng_OnSkillSelected);
			skillitem.onDragStart:Add(jineng_OnDragSkill);
			skillitem.draggable = playerlv >= allData[i-1]._OpenLv;

			skillitem.enabled = Proxy4Lua.GetIndexBySkillID(allData[i-1]._SkillId) == -1;
		end

		if allData[i-1]._Type == 1 then	--active
			skilldata = SkillData.GetData(allData[i-1]._SkillId);
			skillitem = activeList:AddItemFromPool(skillItemUrl);
			skillitem:GetChild("n10").visible = false;
			skillicon = skillitem:GetChild("n8").asLoader;
			skillicon.url = "ui://" .. skilldata._Icon;
			skilllv = skillitem:GetChild("n7").asTextField;
			skilllv.text = skilldata._Level;
			skilllock = skillitem:GetChild("n12");
			skillunlockcond = skillitem:GetChild("n13").asTextField;
			skilllock.visible = playerlv < allData[i-1]._OpenLv;
			skillunlockcond.visible = playerlv < allData[i-1]._OpenLv;
			skillunlockcond.text = allData[i-1]._OpenLv .. "级解锁";

			skillitem.data = {};
			skillitem.data.rsId = allData[i-1]._ID;
			skillitem.data.down = 0;
			skillitem.onClick:Add(jineng_OnSkillSelected);
			skillitem.onDragStart:Add(jineng_OnDragSkill);
			skillitem.draggable = playerlv >= allData[i-1]._OpenLv;

			skillitem.enabled = Proxy4Lua.GetIndexBySkillID(allData[i-1]._SkillId) == -1;
		end

		if allData[i-1]._Type == 2 then	--critical
			skilldata = SkillData.GetData(allData[i-1]._SkillId);
			skillitem = criticalList:AddItemFromPool(skillItemUrl);
			skillitem:GetChild("n10").visible = false;
			skillicon = skillitem:GetChild("n8").asLoader;
			skillicon.url = "ui://" .. skilldata._Icon;
			skilllv = skillitem:GetChild("n7").asTextField;
			skilllv.text = skilldata._Level;
			skilllock = skillitem:GetChild("n12");
			skillunlockcond = skillitem:GetChild("n13").asTextField;
			skilllock.visible = playerlv < allData[i-1]._OpenLv;
			skillunlockcond.visible = playerlv < allData[i-1]._OpenLv;
			skillunlockcond.text = allData[i-1]._OpenLv .. "级解锁";

			skillitem.data = {};
			skillitem.data.rsId = allData[i-1]._ID;
			skillitem.data.down = 0;
			skillitem.onClick:Add(jineng_OnSkillSelected);
			skillitem.onDragStart:Add(jineng_OnDragSkill);
			skillitem.draggable = playerlv >= allData[i-1]._OpenLv;

			skillitem.enabled = Proxy4Lua.GetIndexBySkillID(allData[i-1]._SkillId) == -1;
		end
	end
	skilldata = Proxy4Lua.GetPlayerSkillData(0);
	skilllv = skill1:GetChild("n7").asTextField;
	skillicon = skill1:GetChild("n8").asLoader;
	if skilldata ~= nil then
		skillicon.url = "ui://" .. skilldata._Icon;
		skilllv.text = skilldata._Level;
		local rsData = RoleSkillData.GetDataBySkillID(skilldata._Id);
		if rsData ~= nil then
			skill1.data.rsId = rsData._ID;
		end
--		skill1.touchable = true;
	else
		skillicon.url = "";
		skilllv.text = "";
		skill1.data.rsId = 0;
--		skill1.touchable = false;
	end

	skilldata = Proxy4Lua.GetPlayerSkillData(1);
	skilllv = skill2:GetChild("n7").asTextField;
	skillicon = skill2:GetChild("n8").asLoader;
	if skilldata ~= nil then
		skillicon.url = "ui://" .. skilldata._Icon;
		skilllv.text = skilldata._Level;
		local rsData = RoleSkillData.GetDataBySkillID(skilldata._Id);
		if rsData ~= nil then
			skill2.data.rsId = rsData._ID;
		end
--		skill2.touchable = true;
	else
		skillicon.url = "";
		skilllv.text = "";
		skill2.data.rsId = 0;
--		skill2.touchable = false;
	end

	skilldata = Proxy4Lua.GetPlayerSkillData(2);
	skilllv = skill3:GetChild("n7").asTextField;
	skillicon = skill3:GetChild("n8").asLoader;
	if skilldata ~= nil then
		skillicon.url = "ui://" .. skilldata._Icon;
		skilllv.text = skilldata._Level;
		local rsData = RoleSkillData.GetDataBySkillID(skilldata._Id);
		if rsData ~= nil then
			skill3.data.rsId = rsData._ID;
		end
--		skill3.touchable = true;
	else
		skillicon.url = "";
		skilllv.text = "";
		skill3.data.rsId = 0;
--		skill3.touchable = false;
	end

	skilldata = Proxy4Lua.GetPlayerSkillData(3);
	skilllv = skill4:GetChild("n7").asTextField;
	skillicon = skill4:GetChild("n8").asLoader;
	if skilldata ~= nil then
		skillicon.url = "ui://" .. skilldata._Icon;
		skilllv.text = skilldata._Level;
		local rsData = RoleSkillData.GetDataBySkillID(skilldata._Id);
		if rsData ~= nil then
			skill4.data.rsId = rsData._ID;
		end
--		skill4.touchable = true;
	else
		skillicon.url = "";
		skilllv.text = "";
		skill4.data.rsId = 0;
--		skill4.touchable = false;
	end
end