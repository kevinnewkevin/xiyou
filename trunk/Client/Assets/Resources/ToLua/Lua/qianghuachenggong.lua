require "FairyGUI"

qianghuachenggong = fgui.window_class(WindowBase)
local Window;


local hp;
local agility;
local atk;
local def;
local matk;
local mdef;

local levelUpHp;
local levelUpAtk;
local levelUpDef;
local levelUpAgility;
local levelUpMatk;
local levelUpMdef;

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


	qianghuachenggong_FlushData();
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
	Proxy4Lua.UnloadAsset(modelRes);
	Window:Hide();
end





function qianghuachenggong_FlushData()
	local instId = UIParamHolder.Get("qiecuo1");
	local displayData = GamePlayer.GetDisplayDataByInstID(instId);

	local entityData = GamePlayer.GetEntityDataByInstID(instId);

	local entityInst = GamePlayer.GetCardByInstID(instId);
	hp.text = entityInst.CProperties[1];
	agility.text = entityInst.CProperties[7];
	atk.text = entityInst.CProperties[3];
	def.text = entityInst.CProperties[4];
	matk.text = entityInst.CProperties[5];
	mdef.text = entityInst.CProperties[6];



	local  levelData =  StrengthenData.GetData( entityInst.UnitId, entityInst.Level );
	
	levelUpHp.text = levelData._Hp + entityInst.CProperties[1] .. "";
	levelUpAtk.text = levelData._Atk + entityInst.CProperties[3] .. "";
	levelUpDef.text = levelData._Def + entityInst.CProperties[4] .. "";
	levelUpAgility.text = levelData._Agile + entityInst.CProperties[7]  .. "";
	levelUpMatk.text = levelData._MagicAtk + entityInst.CProperties[5]  .. "";
	levelUpMdef.text = levelData._MagicDef + entityInst.CProperties[6]  .. "";


end