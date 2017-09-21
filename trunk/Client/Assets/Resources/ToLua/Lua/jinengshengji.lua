require "FairyGUI"

jinengshengji = fgui.window_class(WindowBase)
local Window;

local skillIcon;
local skillLv;
local skillPreLv;
local skillNowLv;
local skillDesc;
local holder;
local effRes;

function jinengshengji:OnEntry()
	Window = jinengshengji.New();
	Window:Show();
end

function jinengshengji:GetWindow()
	return Window;
end

function jinengshengji:OnInit()
	self.contentPane = UIPackage.CreateObject("qianghuachenggong", "jinengshengji_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n52");

	holder = self.contentPane:GetChild("n56").asGraph;
	local skill = self.contentPane:GetChild("n57");
	skillIcon = skill:GetChild("n8").asLoader;
	skillLv = skill:GetChild("n7");
	fee = skill:GetChild("n7");
	inGroup = skill:GetChild("n9");
	inGroup.visible = false;
	skillPreLv = self.contentPane:GetChild("n48");
	skillNowLv = self.contentPane:GetChild("n49");
	skillDesc = self.contentPane:GetChild("n61");

	jinengshengji_FlushData();
end

function jinengshengji:OnUpdate()
	if UIManager.IsDirty("jinengshengji") then
		jinengshengji_FlushData();
		UIManager.ClearDirty("jinengshengji");
	end
end


function jinengshengji:OnTick()
	
end

function jinengshengji:isShow()
	return Window.isShowing;
end

function jinengshengji:OnDispose()
	Window:Dispose();
end

function jinengshengji:OnHide()
	Proxy4Lua.UnloadAsset(effRes);
	effRes = "";
	Window:Hide();
end

function jinengshengji_FlushData()
	effRes = "Effect/dengjitishen";
	holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(effRes));
	local preSkillId = UIParamHolder.Get("jineng");
	local nowSkillId = UIParamHolder.Get("jineng2");
	local sData = SkillData.GetData(preSkillId);
	local sData2 = SkillData.GetData(nowSkillId);
	if sData ~= nil then
		skillLv.text = sData._Level;
		skillPreLv.text = "等级" .. sData._Level;
	else
		skillLv.text = "";
		skillPreLv.text = "";
	end
	if sData2 ~= nil then
		skillIcon.url = "ui://" .. sData2._Icon;
		skillNowLv.text = "等级" .. sData2._Level;
		skillDesc.text = sData2._Desc;
	else
		skillIcon.url = "";
		skillNowLv.text = "";
		skillDesc.text = "";
	end
end