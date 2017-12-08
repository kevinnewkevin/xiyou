require "FairyGUI"

xiangxiziliao = fgui.window_class(WindowBase)
local Window;

local holder;

local controller;
local isInGroup;

local fee;
local name;
local level;

local skillList;
local hp;
local agility;
local atk;
local def;
local matk;
local mdef;
local crit;
local critdmg;
local combo;
local split;
local rec;
local reflect;
local suck;
local race;
local raceBack;

local levelHp;
local levelAtk;
local levelDef;
local levelAgility;
local levelMatk;
local levelMdef;

local addLevelHp;
local addLevelAtk;
local addLevelDef;
local addLevelAgility;
local addLevelMatk;
local addLevelMdef;

local levelLab;
local levelUpHp;
local levelUpAtk;
local levelUpDef;
local levelUpAgility;
local levelUpMatk;
local levelUpMdef;

local addGroupBtn;
local modelRes;
local rightInfo;
local rightLevelUp;
local levelUpBtn;
local needMoneyLab;
local needItemNumLab; 
local needItemIcon;
local needItemIconback;
local needItem;
local needItemBar;
local levelUpRad;
local btnList;
local levelRadBtn;
local raceHelp;

function xiangxiziliao:OnEntry()
	Window = xiangxiziliao.New();
	Window:Show();
end

function xiangxiziliao:GetWindow()
	return Window;
end

function xiangxiziliao:OnInit()
	self.contentPane = UIPackage.CreateObject("xiangxiziliao", "xiangxiziliao_com").asCom;
	self:Center();
	self.modal = true;

	self.closeButton = self.contentPane:GetChild("n4");
	holder = self.contentPane:GetChild("n63").asGraph;

	addGroupBtn = self.contentPane:GetChild("n61").asButton;
	addGroupBtn.onClick:Add(xiangxiziliao_OnAddGroup);
	controller = addGroupBtn:GetController("huang");

	rightInfo = self.contentPane:GetChild("n80").asCom;
	rightLevelUp = self.contentPane:GetChild("n81").asCom;

	btnList = self.contentPane:GetChild("n82").asList;
	local btnMax = btnList.numItems;
	local btnItem;
	for i=1, btnMax do
		btnItem = btnList:GetChildAt(i-1);
		btnItem.data = i-1;
		btnItem.onClick:Add(xiangxiziliao_OnBtnItemClick);
		if i == btnMax then 
			levelUpRad = btnItem:GetChild("n5");
		end
	end
	levelRadBtn = btnList:GetChildAt(btnMax-1);
	rightInfo.visible = true;
	rightLevelUp.visible = false;
	local movebg =  self.contentPane:GetChild("n83");
	local moveLab =  self.contentPane:GetChild("n84");
	movebg.visible = false;
	moveLab.visible = false;
	btnList.selectedIndex = 0;
	level= self.contentPane:GetChild("n59");
	fee = self.contentPane:GetChild("n58");
	name = self.contentPane:GetChild("n60");
	race = self.contentPane:GetChild("n89");
	race.onClick:Add(xiangxiziliao_OnRace);
	raceBack = self.contentPane:GetChild("n85");
	raceHelp = self.contentPane:GetChild("n87");
	raceHelpCloseBtn = raceHelp:GetChild("n1").asButton;
	raceHelpCloseBtn.onClick:Add(xiangxiziliao_OnRaceClose);
	raceHelp.visible = false;
	skillList = rightInfo:GetChild("n237").asList;
	hp = rightInfo:GetChild("n242");
	agility = rightInfo:GetChild("n246");
	atk = rightInfo:GetChild("n213");
	def = rightInfo:GetChild("n214");
	matk = rightInfo:GetChild("n215");
	mdef = rightInfo:GetChild("n216");
	crit = rightInfo:GetChild("n217");
	critdmg = rightInfo:GetChild("n218");
	combo = rightInfo:GetChild("n219");
	split = rightInfo:GetChild("n220");
	rec = rightInfo:GetChild("n221");
	reflect = rightInfo:GetChild("n222");
	suck = rightInfo:GetChild("n223");
	levelLab = rightLevelUp:GetChild("n308");
	levelUpBtn= rightLevelUp:GetChild("n299");
	levelUpBtn.onClick:Add(xiangxiziliao_OnLevelUpClick);
	levelHp = rightLevelUp:GetChild("n242");
	levelAtk = rightLevelUp:GetChild("n213");
	levelDef = rightLevelUp:GetChild("n214");
	levelAgility = rightLevelUp:GetChild("n246");
	levelMatk = rightLevelUp:GetChild("n215");
	levelMdef = rightLevelUp:GetChild("n216");

	addLevelHp = rightLevelUp:GetChild("n277");
	addLevelAtk = rightLevelUp:GetChild("n279");
	addLevelDef = rightLevelUp:GetChild("n280");  
	addLevelAgility = rightLevelUp:GetChild("n281");
	addLevelMatk = rightLevelUp:GetChild("n282");
	addLevelMdef = rightLevelUp:GetChild("n283");

	--levelUpHp = rightLevelUp:GetChild("n271");
	--levelUpAtk = rightLevelUp:GetChild("n267");
	--levelUpDef = rightLevelUp:GetChild("n268");
	--levelUpAgility = rightLevelUp:GetChild("n272");
	--levelUpMatk = rightLevelUp:GetChild("n269"); 
	--levelUpMdef = rightLevelUp:GetChild("n270");
	needMoneyLab = rightLevelUp:GetChild("n304");
	needItemNumLab = rightLevelUp:GetChild("n301");
	--needItem= rightLevelUp:GetChild("n297"); 
	needItemIcon= rightLevelUp:GetChild("n309"); 
	--needItemIconback= needItem:GetChild("n0"); 
	needItemBar = rightLevelUp:GetChild("n300");

	for i=2, 4 do
		local skill = skillList:GetChildAt(i - 2);
		skill.onClick:Add(xiangxiziliao_OnSkillBtn);
	end
	--xiangxiziliao_FlushData();
end

function xiangxiziliao_OnAddGroup()
	xiangxiziliao_OnMessageConfirm();
	--[[local MessageBox = UIManager.ShowMessageBox();
	isInGroup = UIManager.GetWindow("paiku").IsInGroup();
	if isInGroup then
		MessageBox:SetData("提示", "是否取出卡组？", false, xiangxiziliao_OnMessageConfirm, xiangxiziliao_OnMessageCancel);
	else
		MessageBox:SetData("提示", "是否加入卡组？", false, xiangxiziliao_OnMessageConfirm, xiangxiziliao_OnMessageCancel);
	end--]]
end

function xiangxiziliao_OnMessageConfirm()
	local crtCardInstID = UIManager.GetWindow("paiku").GetCrtCard();
	local crtGroupIdx = UIManager.GetWindow("paiku").GetCrtGroup();

	isInGroup = UIManager.GetWindow("paiku").IsInGroup();
	if isInGroup then
		GamePlayer.TakeOffCard(crtCardInstID, crtGroupIdx);
	else 
		if GamePlayer.IsGroupMax(crtGroupIdx) then
			local MessageBox = UIManager.ShowMessageBox();
			MessageBox:SetData("提示", "卡组已满", true);
		else
			GamePlayer.PutInCard(crtCardInstID, crtGroupIdx);
		end
	end
	UIManager.Hide("xiangxiziliao");
	UIManager.SetDirty("paiku"); 
end

function xiangxiziliao_OnMessageCancel()
	UIManager.HideMessageBox();
	UIManager.SetDirty("paiku");
end

function xiangxiziliao:OnUpdate()
	if UIManager.IsDirty("xiangxiziliao") then
		xiangxiziliao_FlushData();
		UIManager.ClearDirty("xiangxiziliao");
	end
end

function xiangxiziliao:OnTick()
	
end

function xiangxiziliao:isShow()
	return Window.isShowing;
end

function xiangxiziliao:OnDispose()
	Window:Dispose();
end

function xiangxiziliao:OnHide()
	holder:SetNativeObject(Proxy4Lua.GetAssetGameObject("", false));
	Proxy4Lua.UnloadAsset(modelRes);
	modelRes = "";
	raceHelp.visible = false;
	rightLevelUp.visible = false;
	rightInfo.visible = true;
	btnList.selectedIndex = 0;
	Window:Hide();
	GuideSystem.CloseUI("xiangxiziliao");
end


function xiangxiziliao_OnBtnItemClick(context)
	if  context.sender.data == 0 then
		rightInfo.visible = true;
		rightLevelUp.visible = false;
	else
		rightInfo.visible = false;
		rightLevelUp.visible = true;
	end
end

function xiangxiziliao_OnLevelUpClick(context)
	local instId = UIParamHolder.Get("qiecuo1");
	Proxy4Lua.PromoteUnit(instId);
	Window:ShowModalWait();
end

function xiangxiziliao_OnSkillBtn(context)
	if context.sender.data == nil or context.sender.data == 0 then
		return;
	end

	UIParamHolder.Set("jinengxiangqing", context.sender.data);
	UIManager.Show("jinengxiangqing");
end

function xiangxiziliao_FlushData()
	local showBoos = UIParamHolder.Get("showBoos");
	if showBoos == true then
		xiangxiziliao_BoosInfo();
		return;
	end

	local instId = UIParamHolder.Get("qiecuo1");
	local hideBtn = UIParamHolder.Get("qiecuo2");
	local displayData = GamePlayer.GetDisplayDataByInstID(instId);
	modelRes = displayData._AssetPath;
	holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes, true, 1200, 1.5));

	race.asLoader.url = "ui://" .. displayData._Race;
	raceBack.asLoader.url = "ui://" .. displayData._Race.. "_bj";
	if hideBtn == false then
		isInGroup = UIManager.GetWindow("paiku").IsInGroup();
		if isInGroup then
			controller.selectedIndex = 1;
		else
			controller.selectedIndex = 0;
		end
		addGroupBtn.visible = true;
	else
		addGroupBtn.visible = false;
	end


	local entityData = GamePlayer.GetEntityDataByInstID(instId);
	fee.text = entityData._Cost;
	name.text = entityData._Name;
	local entityInst = GamePlayer.GetCardByInstID(instId);
	level.text =  entityInst.IProperties[9] .. "";
	hp.text = entityInst.CProperties[1];
	agility.text = entityInst.CProperties[7];
	atk.text = entityInst.CProperties[3];
	def.text = entityInst.CProperties[4];
	matk.text = entityInst.CProperties[5];
	mdef.text = entityInst.CProperties[6];
	crit.text = entityInst.CProperties[9];
	critdmg.text = entityInst.CProperties[12];
	combo.text = entityInst.CProperties[10];
	split.text = entityInst.CProperties[11];
	rec.text = entityInst.CProperties[13];
	reflect.text = entityInst.CProperties[14];
	suck.text = entityInst.CProperties[15];
	levelLab.text = entityInst.IProperties[9].."";
	levelHp.text = entityInst.CProperties[1];
	levelAtk.text = entityInst.CProperties[3];
	levelDef.text = entityInst.CProperties[4];
	levelAgility.text = entityInst.CProperties[7];
	levelMatk.text = entityInst.CProperties[5];
	levelMdef.text = entityInst.CProperties[6];
	
	local  levelData =  StrengthenData.GetData( entityInst.UnitId,  entityInst.IProperties[9]+1);
	addLevelHp.text = levelData._Hp .. "";
	addLevelAtk.text = levelData._Atk .. "";
	addLevelDef.text = levelData._Def .. "";
	addLevelAgility.text = levelData._Agile .. "";
	addLevelMatk.text = levelData._MagicAtk .. "";
	addLevelMdef.text = levelData._MagicDef .. "";

	--levelUpHp.text = levelData._Hp + entityInst.CProperties[1] .. "";
	--levelUpAtk.text = levelData._Atk + entityInst.CProperties[3] .. "";
	--levelUpDef.text = levelData._Def + entityInst.CProperties[4] .. "";
	--levelUpAgility.text = levelData._Agile + entityInst.CProperties[7]  .. "";
	--levelUpMatk.text = levelData._MagicAtk + entityInst.CProperties[5]  .. "";
	--levelUpMdef.text = levelData._MagicDef + entityInst.CProperties[6]  .. "";
	needMoneyLab.text = levelData._ItemNum .. "";
	local itemNum = BagSystem.GetItemMaxNum(levelData._ItemId);
	needItemBar.value = itemNum/levelData._ItemNum*100;
	needItemNumLab.text = itemNum .. "/" .. levelData._ItemNum ;
	local itemdata = ItemData.GetData(levelData._ItemId);
	needItemIcon.asLoader.url = "ui://" .. itemdata._Icon;
	--needItemIconback.asLoader.url = "ui://" .. itemdata._IconBack;
	levelRadBtn.visible = true;
	if itemNum >= levelData._ItemNum   then
		levelUpBtn.enabled  = true;
		levelUpRad.visible = true;
	else
		levelUpRad.visible = false;
		levelUpBtn.enabled  = false;
	end

	for i=2, 4 do
		local skill = skillList:GetChildAt(i - 2);
		local sData = Proxy4Lua.GetCardInstSkillData(instId, i-1);
		local loader = skill:GetChild("n8").asLoader;
		local lv = skill:GetChild("n7").asTextField;
		skill:GetChild("n5").visible = false;
		lv.visible = false;
		if sData ~= nil then
			loader.url = "ui://" .. sData._Icon;
			--lv.text = sData._Level;
			skill.data = sData._Id;
		else
			loader.url = "";
			--lv.text = "";
		end
	end
end

function xiangxiziliao_OnRace()
	raceHelp.visible = true;
end

function xiangxiziliao_OnRaceClose()
	raceHelp.visible = false;
end

function xiangxiziliao_BoosInfo()
	local showBoos = UIParamHolder.Get("showBoos");
	local boosId = UIParamHolder.Get("showBoosId");
	local entityData = EntityData.GetData(boosId);
	local displayData =  DisplayData.GetData(entityData._DisplayId);
	modelRes = displayData._AssetPath;
	holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes, true, 1200, 1.5));
	race.asLoader.url = "ui://" .. displayData._Race;
	raceBack.asLoader.url = "ui://" .. displayData._Race.. "_bj";

	addGroupBtn.visible = false;
	levelRadBtn.visible = false;

	fee.text = entityData._Cost;
	name.text = entityData._Name;

	level.text =  entityData.IPT_LEVEL .. "";
	hp.text = entityData.CPT_HP .. "";
	agility.text = entityData.CPT_AGILE .. "";
	atk.text = entityData.CPT_ATK .. "";
	def.text = entityData.CPT_DEF .. "";
	matk.text = entityData.CPT_MAGIC_ATK .. "";
	mdef.text = entityData.CPT_MAGIC_DEF .. "";
	crit.text = entityData.CPT_CRIT .. "";
	critdmg.text =  "0";--entityData.;
	combo.text = "0";--entityData.;
	split.text = "0";--entityData.;
	rec.text = "0";--entityData.;
	reflect.text = "0";--entityData.;
	suck.text = "0";--entityData.;


	for i=2, 4 do
		local skill = skillList:GetChildAt(i - 2);
		local sData = SkillData.GetData(entityData._Skills[i-1]);
		local loader = skill:GetChild("n8").asLoader;
		local lv = skill:GetChild("n7").asTextField;
		skill:GetChild("n5").visible = false;
		lv.visible = false;
		if sData ~= nil then
			loader.url = "ui://" .. sData._Icon;
			--lv.text = sData._Level;
			skill.data = sData._Id;
		else
			loader.url = "";
			--lv.text = "";
		end
	end


end
