require "FairyGUI"

qianghuachenggong = fgui.window_class(WindowBase)
local Window;


local hp;
local agility;
local atk;
local def;
local matk;
local mdef;
local holder;

local levelUpHp;
local levelUpAtk;
local levelUpDef;
local levelUpAgility;
local levelUpMatk;
local levelUpMdef;
local oldLevel;
local nowLevel;
local head;
local headIcon;
local headLevel;
local nameLab;
local fee;
local inGroup;

local effRes;

function qianghuachenggong:OnEntry()
	Window = qianghuachenggong.New();
	Window:Show();
end

function qianghuachenggong:GetWindow()
	return Window;
end

function qianghuachenggong:OnInit()
	self.contentPane = UIPackage.CreateObject("qianghuachenggong", "qianghuachenggong_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n52");
	head = self.contentPane:GetChild("n47");
	headIcon = head:GetChild("n5").asLoader;
	headLevel = head:GetChild("n6");
	fee = head:GetChild("n7");
	inGroup = head:GetChild("n9");
	inGroup.visible = false;
	oldLevel = self.contentPane:GetChild("n48");
	nowLevel = self.contentPane:GetChild("n49");
	nameLab =  self.contentPane:GetChild("n59");
	hp = self.contentPane:GetChild("n20");
	agility = self.contentPane:GetChild("n24");
	atk = self.contentPane:GetChild("n15");
	def = self.contentPane:GetChild("n16");
	matk = self.contentPane:GetChild("n17");
	mdef = self.contentPane:GetChild("n18");

	levelUpHp = self.contentPane:GetChild("n29");
	levelUpAtk = self.contentPane:GetChild("n25");
	levelUpDef = self.contentPane:GetChild("n26");
	levelUpAgility = self.contentPane:GetChild("n30");
	levelUpMatk = self.contentPane:GetChild("n27"); 
	levelUpMdef = self.contentPane:GetChild("n28");
	holder = self.contentPane:GetChild("n56").asGraph;


	qianghuachenggong_FlushData();
end

function qianghuachenggong:OnUpdate()
	if UIManager.IsDirty("qianghuachenggong") then
		qianghuachenggong_FlushData();
		UIManager.ClearDirty("qianghuachenggong");
	end
end


function qianghuachenggong:OnTick()
	
end

function qianghuachenggong:isShow()
	return Window.isShowing;
end

function qianghuachenggong:OnDispose()
	Window:Dispose();
end

function qianghuachenggong:OnHide()
	Proxy4Lua.UnloadAsset(effRes);
	effRes = "";
	Window:Hide();
end

function qianghuachenggong_FlushData()
	UIManager.GetWindow("xiangxiziliao"):CloseModalWait();
	effRes = "Effect/dengjitishen";
	holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(effRes, false));
	local instId = UIParamHolder.Get("qiecuo1");
	local displayData = GamePlayer.GetDisplayDataByInstID(instId);
	local entityInst = GamePlayer.GetCardByInstID(instId);
	local entityData = GamePlayer.GetEntityDataByInstID(instId);
	local  levelData =  StrengthenData.GetData( entityInst.UnitId,  entityInst.IProperties[9] );
	headIcon.url = "ui://" .. displayData._HeadIcon;
	fee.text = entityData._Cost;
	nameLab.text = entityData._Name;
	--oldLevel.text =  entityInst.IProperties[9] -1 .. "";
	nowLevel.text =  "等级:".. entityInst.IProperties[9];
	headLevel.text =  entityInst.IProperties[9] .. "";
	hp.text =  entityInst.CProperties[1] - levelData._Hp .. "";
	agility.text =  entityInst.CProperties[7] - levelData._Agile   .. "";
	atk.text = entityInst.CProperties[3] - levelData._Atk .. "";
	def.text = entityInst.CProperties[4] - levelData._Def .. "";
	matk.text = entityInst.CProperties[5] - levelData._MagicAtk  .. "";
	mdef.text = entityInst.CProperties[6] - levelData._MagicDef .. "";

	--levelUpHp.text = entityInst.CProperties[1] .. "";
	--levelUpAtk.text = entityInst.CProperties[3].. "";
	--levelUpDef.text = entityInst.CProperties[4]  .. "";
	--levelUpAgility.text = entityInst.CProperties[7]   .. "";
	--levelUpMatk.text = entityInst.CProperties[5]  .. "";
	--levelUpMdef.text = entityInst.CProperties[6] .. "";

	levelUpHp.text = levelData._Hp .. "";
	levelUpAtk.text = levelData._Atk .. "";
	levelUpDef.text = levelData._Def .. "";
	levelUpAgility.text = levelData._Agile .. "";
	levelUpMatk.text = levelData._MagicAtk .. "";
	levelUpMdef.text = levelData._MagicDef .. "";

end