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
if minTimer == nil then
	return;
end
--	if 结构不是空 then
		rollResultName.text = "";
		rollResultIcon.url = "ui://" .. "";
		rollResultIconShadow.url = "ui://" .. "";

		rollCom:GetTransition("t1"):Stop();
		rollCom:GetTransition("t0"):Play();
--	end
end

function shihun_OnRoll()
	GRoot.inst:AddChild(rollCom);
	minTimer = 3;
end

function shihun_OnRollResult()
	rollCom:RemoveFromParent();
	UIManager.Hide("shihun");
end