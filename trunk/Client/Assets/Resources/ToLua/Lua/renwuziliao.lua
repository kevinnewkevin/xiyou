require "FairyGUI"

renwuziliao = fgui.window_class(WindowBase)
local Window;

local holder;

local rightInfo;
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


function renwuziliao:OnEntry()
	Window = renwuziliao.New();
	Window:Show();
end

function renwuziliao:GetWindow()
	return Window;
end

function renwuziliao:OnInit()
	self.contentPane = UIPackage.CreateObject("renwuziliao", "renwuziliao_com").asCom;
	self:Center();
	self.modal = true;

	self.closeButton = self.contentPane:GetChild("n4");
	holder = self.contentPane:GetChild("n63").asGraph;
	level= self.contentPane:GetChild("n59");
	name = self.contentPane:GetChild("n60");

	rightInfo = self.contentPane:GetChild("n90").asCom;
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
	for i=1, 3 do
		local skill = skillList:GetChildAt(i - 1);
		skill.onClick:Add(renwuziliao_OnSkillBtn);
	end

	renwuziliao_FlushData();
end

function renwuziliao_OnSkillBtn(context)
	if context.sender.data == nil or context.sender.data == 0 then
		return;
	end

	UIParamHolder.Set("jinengxiangqing", context.sender.data);
	UIManager.Show("jinengxiangqing");
end

function renwuziliao:OnUpdate()
	if UIManager.IsDirty("renwuziliao") then
		renwuziliao_FlushData();
		UIManager.ClearDirty("renwuziliao");
	end
end

function renwuziliao:OnTick()
	
end

function renwuziliao:isShow()
	return Window.isShowing;
end

function renwuziliao:OnDispose()
	Window:Dispose();
end

function renwuziliao:OnHide()
	Proxy4Lua.UnloadAsset(modelRes);
	modelRes = "";
	Window:Hide();
end


function renwuziliao_FlushData()
	local instId =GamePlayer._InstID;
	local displayData =GamePlayer.GetMyDisplayData();
	if modelRes ~= displayData._AssetPath then
		Proxy4Lua.UnloadAsset(modelRes);
		modelRes = displayData._AssetPath;
		local scale = Define.GetFloat("UIModelScale") * 2;
		holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes, scale, 0, true));
	end
	name.text =  GamePlayer._Name;

	local entityInst = GamePlayer.GetCardByInstID(instId);
	level.text =  GamePlayer._Data.IProperties[9] .. "";
	hp.text = GamePlayer._Data.CProperties[1];
	agility.text = GamePlayer._Data.CProperties[7];
	atk.text = GamePlayer._Data.CProperties[3];
	def.text = GamePlayer._Data.CProperties[4];
	matk.text = GamePlayer._Data.CProperties[5];
	mdef.text = GamePlayer._Data.CProperties[6];
	crit.text = GamePlayer._Data.CProperties[9];
	critdmg.text = GamePlayer._Data.CProperties[12];
	combo.text = GamePlayer._Data.CProperties[10]; 
	split.text = GamePlayer._Data.CProperties[11];
	rec.text = GamePlayer._Data.CProperties[13];
	reflect.text = GamePlayer._Data.CProperties[14];
	suck.text = GamePlayer._Data.CProperties[15];   

	local entityData = GamePlayer.GetMyEntityData();
	for i=1, 3 do
		local skill = skillList:GetChildAt(i - 1);
		local sData = Proxy4Lua.GetPlayerSkillData(i-1);
		skill.data = entityData._Skills[i - 1];
		local loader = skill:GetChild("n8").asLoader;
		local lv = skill:GetChild("n7").asTextField;
		loader.url = "";
		if sData ~= nil then
			loader.url = "ui://" .. sData._Icon;
			lv.text = sData._Level;
		else
			loader.url = "";
			lv.text = "";
		end
	end

end