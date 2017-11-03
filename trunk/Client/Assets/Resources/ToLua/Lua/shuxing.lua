require "FairyGUI"

shuxing = fgui.window_class(WindowBase)
local Window;
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

function shuxing:OnEntry()
	Window = shuxing.New();
	Window:Show();
end

function shuxing:GetWindow()
	return Window;
end

function shuxing:OnInit()
	self.contentPane = UIPackage.CreateObject("shuxing", "shuxing_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n268");

	skillList = self.contentPane:GetChild("n237").asList;
	hp = self.contentPane:GetChild("n242");
	agility = self.contentPane:GetChild("n246");
	atk = self.contentPane:GetChild("n213");
	def = self.contentPane:GetChild("n214");
	matk = self.contentPane:GetChild("n215");
	mdef = self.contentPane:GetChild("n216");
	crit = self.contentPane:GetChild("n217");
	critdmg = self.contentPane:GetChild("n218");
	combo = self.contentPane:GetChild("n219");
	split = self.contentPane:GetChild("n220");
	rec = self.contentPane:GetChild("n221");
	reflect = self.contentPane:GetChild("n222");
	suck = self.contentPane:GetChild("n223");

	for i=2, 4 do
		local skill = skillList:GetChildAt(i - 2);
		skill.onClick:Add(shuxing_OnSkillBtn);
	end

	shuxing_FlushData();
end

function shuxing_OnSkillBtn(context)
	if context.sender.data == nil or context.sender.data == 0 then
		return;
	end

	UIParamHolder.Set("jinengxiangqing", context.sender.data);
	UIManager.Show("jinengxiangqing");
end

function shuxing:OnUpdate()
	if UIManager.IsDirty("shuxing") then
		shuxing_FlushData();
		UIManager.ClearDirty("shuxing");
	end
end

function shuxing:OnTick()
	
end

function shuxing:isShow()
	return Window.isShowing;
end

function shuxing:OnDispose()
	Window:Dispose();
end

function shuxing:OnHide()
	Window:Hide();
end

function shuxing_FlushData()

	local entityInst = GamePlayer.GetCardByInstID(Battle._SelectedHandCardInstID);
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

	for i=2, 4 do
		local skill = skillList:GetChildAt(i - 2);
		local sData = Proxy4Lua.GetCardInstSkillData(Battle._SelectedHandCardInstID, i-1);
		local loader = skill:GetChild("n8").asLoader;
		local lv = skill:GetChild("n7").asTextField;
		if sData ~= nil then
			loader.url = "ui://" .. sData._Icon;
			lv.text = sData._Level;
			skill.data = sData._Id;
		else
			loader.url = "";
			lv.text = "";
		end
	end
end