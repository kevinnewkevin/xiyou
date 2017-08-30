require "FairyGUI"

xiangxiziliao = fgui.window_class(WindowBase)
local Window;

local holder;

local controller;
local isInGroup;

local fee;
local name;

local skillList;

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

local atkplus;
local defplus;
local matkplus;
local mdefplus;
local critplus;
local critdmgplus;
local comboplus;
local splitplus;
local recplus;
local reflectplus;
local suckplus;

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

	local addGroupBtn = self.contentPane:GetChild("n61").asButton;
	addGroupBtn.onClick:Add(xiangxiziliao_OnAddGroup);
	controller = addGroupBtn:GetController("huang");

	fee = self.contentPane:GetChild("n58");
	name = self.contentPane:GetChild("n60");

	skillList = self.contentPane:GetChild("n49").asList;

	atk = self.contentPane:GetChild("n31");
	def = self.contentPane:GetChild("n32");
	matk = self.contentPane:GetChild("n33");
	mdef = self.contentPane:GetChild("n34");
	crit = self.contentPane:GetChild("n35");
	critdmg = self.contentPane:GetChild("n36");
	combo = self.contentPane:GetChild("n37");
	split = self.contentPane:GetChild("n38");
	rec = self.contentPane:GetChild("n67");
	reflect = self.contentPane:GetChild("n68");
	suck = self.contentPane:GetChild("n69");

	atkplus = self.contentPane:GetChild("n39");
	defplus = self.contentPane:GetChild("n40");
	matkplus = self.contentPane:GetChild("n41");
	mdefplus = self.contentPane:GetChild("n42");
	critplus = self.contentPane:GetChild("n43");
	critdmgplus = self.contentPane:GetChild("n44");
	comboplus = self.contentPane:GetChild("n45");
	splitplus = self.contentPane:GetChild("n46");
	recplus = self.contentPane:GetChild("n70");
	reflectplus = self.contentPane:GetChild("n71");
	suckplus = self.contentPane:GetChild("n72");

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
	Window:Hide();
end

function xiangxiziliao_FlushData()
	local instId = UIManager.GetWindow("paiku").GetCrtCard();
	local displayData = GamePlayer.GetDisplayDataByInstID(instId);
	local modelRes = displayData._AssetPath;
	holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes));

	isInGroup = UIManager.GetWindow("paiku").IsInGroup();
	if isInGroup then
		controller.selectedIndex = 1;
	else
		controller.selectedIndex = 0;
	end

	local entityData = GamePlayer.GetEntityDataByInstID(instId);
	fee.text = entityData._Cost;
	name.text = entityData._Name;

	local entityInst = GamePlayer.GetCardByInstID(instId);
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

	atkplus.text = "";
	defplus.text = "";
	matkplus.text = "";
	mdefplus.text = "";
	critplus.text = "";
	critdmgplus.text = "";
	comboplus.text = "";
	splitplus.text = "";
	recplus.text = "";
	reflectplus.text = "";
	suckplus.text = "";

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