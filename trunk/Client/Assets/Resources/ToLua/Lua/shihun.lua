require "FairyGUI"

shihun = fgui.window_class(WindowBase)
local Window;

local rollBtn;

local rollCom;
local rollResultName;
local rollResultIcon;
local rollResultIconShadow;
local rollResultBtn;

local test;

local minTimer = nil;

function shihun:OnEntry()
	Window = shihun.New();
	Window:Show();
end

function shihun:GetWindow()
	return Window;
end

function shihun:OnInit()
	self.contentPane = UIPackage.CreateObject("shihun", "shihun_com").asCom;
	self:Center();

	rollBtn = self.contentPane:GetChild("n5").asButton;
	rollBtn.onClick:Add(shihun_OnRoll);

	rollCom = UIPackage.CreateObject("shihun", "huode_com").asCom;
	rollCom:RemoveFromParent();
	rollResultName = rollCom:GetChild("n10").asTextField;
	rollResultIcon = rollCom:GetChild("n12").asLoader;
	rollResultIconShadow = rollCom:GetChild("n15").asLoader;
	rollResultBtn = rollCom:GetChild("n13").asButton;
	rollResultBtn.onClick:Add(shihun_OnRollResult);

	test = self.contentPane:GetChild("n7");
	test:RemoveFromParent();

	shihun_FlushData();
end

function shihun:OnUpdate()
	if UIManager.IsDirty("shihun") then
		shihun_FlushData();
		UIManager.ClearDirty("shihun");
	end
end

function shihun:OnTick()
	if minTimer ~= nil then
		minTimer = minTimer - 1;
		if minTimer <= 0 then
			shihun_CheckResult();
			minTimer = nil;
		end
	end
end

function shihun:isShow()
	return Window.isShowing;
end

function shihun:OnDispose()
	Window:Dispose();
end

function shihun:OnHide()
	minTimer = nil;
	Window:Hide();
end

function shihun_FlushData()
	shihun_CheckResult();
end

function shihun_CheckResult()
	if JieHunSystem._LastestChapter ~= nil then
		local hsData = HeroStoryData.GetData(JieHunSystem._LastestChapter.ChapterId);
		if hsData ~= nil then
			local eData = EntityData.GetData(hsData.EntityID_);
			if eData ~= nil then
				local dData = DisplayData.GetData(eData._DisplayId);
				if dData ~= nil then
					rollResultName.text = eData._Name;
					rollResultIcon.url = "ui://" .. dData._HeadIcon;
					rollResultIconShadow.url = "ui://" .. dData._HeadIcon;

					rollCom:GetTransition("t1"):Stop();
					rollCom:GetTransition("t0"):Play();
				end
			end
		end
		JieHunSystem._LastestChapter = nil;
	end
end

function shihun_OnRoll()
	GRoot.inst:AddChild(rollCom);
	minTimer = 3;
	Proxy4Lua.RandChapter();
end

function shihun_OnRollResult()
	rollCom:RemoveFromParent();
	UIManager.Hide("shihun");
end