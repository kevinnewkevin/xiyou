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

	local btnList = self.contentPane:GetChild("n82").asList;
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
	rightInfo.visible = true;
	rightLevelUp.visible = false;
	btnList.selectedIndex = 0;
	level= self.contentPane:GetChild("n59");
	fee = self.contentPane:GetChild("n58");
	name = self.contentPane:GetChild("n60");
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

	levelUpHp = rightLevelUp:GetChild("n271");
	levelUpAtk = rightLevelUp:GetChild("n267");
	levelUpDef = rightLevelUp:GetChild("n268");
	levelUpAgility = rightLevelUp:GetChild("n272");
	levelUpMatk = rightLevelUp:GetChild("n269"); 
	levelUpMdef = rightLevelUp:GetChild("n270");
	needMoneyLab = rightLevelUp:GetChild("n304");
	needItemNumLab = rightLevelUp:GetChild("n301"); 
	needItem= rightLevelUp:GetChild("n297"); 
	needItemIcon= needItem:GetChild("n1"); 
	needItemIconback= needItem:GetChild("n0"); 
	needItemBar = rightLevelUp:GetChild("n300");
	xiangxiziliao_FlushData();
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

	if GamePlayer.IsGroupMax(crtGroupIdx) then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "卡组已满", true);
		UIManager.SetDirty("paiku");
		return;
	end

	isInGroup = UIManager.GetWindow("paiku").IsInGroup();
	if isInGroup then
		GamePlayer.TakeOffCard(crtCardInstID, crtGroupIdx);
	else
		GamePlayer.PutInCard(crtCardInstID, crtGroupIdx);
	end
	--UIManager.HideMessageBox();
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
	Proxy4Lua.UnloadAsset(modelRes);
	Window:Hide();
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
end


function xiangxiziliao_FlushData()
	local instId = UIParamHolder.Get("qiecuo1");
	local hideBtn = UIParamHolder.Get("qiecuo2");
	local displayData = GamePlayer.GetDisplayDataByInstID(instId);
	if modelRes ~= displayData._AssetPath then
		Proxy4Lua.UnloadAsset(modelRes);
		modelRes = displayData._AssetPath;
		holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes));

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

	levelUpHp.text = levelData._Hp + entityInst.CProperties[1] .. "";
	levelUpAtk.text = levelData._Atk + entityInst.CProperties[3] .. "";
	levelUpDef.text = levelData._Def + entityInst.CProperties[4] .. "";
	levelUpAgility.text = levelData._Agile + entityInst.CProperties[7]  .. "";
	levelUpMatk.text = levelData._MagicAtk + entityInst.CProperties[5]  .. "";
	levelUpMdef.text = levelData._MagicDef + entityInst.CProperties[6]  .. "";
	needMoneyLab.text = levelData._ItemNum .. "";
	local itemNum = BagSystem.GetItemMaxNum(levelData._ItemId);
	needItemBar.value = itemNum/levelData._ItemNum*100;
	needItemNumLab.text = itemNum .. "/" .. levelData._ItemNum ;
	local itemdata = ItemData.GetData(levelData._ItemId);
	needItemIcon.asLoader.url = "ui://" .. itemdata._Icon;
	needItemIconback.asLoader.url = "ui://icon/" .. itemdata._IconBack;
	
	if itemNum >= levelData._ItemNum   then
		levelUpBtn.enabled  = true;
		levelUpRad.visible = true;
	else
		levelUpRad.visible = false;
		levelUpBtn.enabled  = false;
	end
	
	for i=1, 4 do
		local skill = skillList:GetChildAt(i - 1);
		local sData = SkillData.GetData(entityData._Skills[i - 1]);
		local loader = skill:GetChild("n8").asLoader;
		loader.url = "";
		if sData ~= nil then
			loader.url = "ui://" .. sData._Icon;
		end
	end
end