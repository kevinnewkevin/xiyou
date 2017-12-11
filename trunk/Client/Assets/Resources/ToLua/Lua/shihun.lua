require "FairyGUI"

shihun = fgui.window_class(WindowBase)
local Window;

local rollBtn;
local itemIcon;
local itemNameAndNum;
local levelIcon;

local rollCom;
local rollResultName;
local rollResultIcon;
local rollResultIconShadow;
local rollResultBtn;

local test;

local minTimer = nil;
local enableNextTimer = nil;

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
	self.modal = true;
	self.bringToFontOnClick = false;
	self.closeButton = self.contentPane:GetChild("n10");

	rollBtn = self.contentPane:GetChild("n5").asButton;
	rollBtn.onClick:Add(shihun_OnRoll);

	itemIcon = self.contentPane:GetChild("n8").asLoader
	itemNameAndNum = self.contentPane:GetChild("n9").asTextField;
	levelIcon = self.contentPane:GetChild("n12").asLoader

	rollCom = UIPackage.CreateObject("shihun", "huode_com").asCom;
	rollCom:RemoveFromParent();
	rollResultName = rollCom:GetChild("n10").asTextField;
	rollResultIcon = rollCom:GetChild("n12").asLoader;
	rollResultIconShadow = rollCom:GetChild("n15").asLoader;
	rollResultBtn = rollCom:GetChild("n13").asButton;
	--rollResultBtn.onClick:Add(shihun_OnRollResult);

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
			enableNextTimer = 0.6;
		end
	end

	if enableNextTimer ~= nil then
		enableNextTimer = enableNextTimer - 1;
		if enableNextTimer < 0 then
			enableNextTimer = nil;
			rollResultBtn.onClick:Add(shihun_OnRollResult);
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
	enableNextTimer = nil;
	Window:Hide();
end

function shihun_FlushData()
	rollResultBtn.onClick:Remove(shihun_OnRollResult);
	rollBtn.enabled = false;
	shihun_CheckResult();
	if JieHunSystem.instance._NextDrawData ~= nil then
		local iData = ItemData.GetData(JieHunSystem.instance._NextDrawData._ItemId);
		if iData ~= nil then
			local has = BagSystem.GetItemMaxNum(iData._Id);
			local need = JieHunSystem.instance._NextDrawData._ItemNum;
			local color = "white";
			if has < need then
				color = "red";
			end
			itemIcon.url = "ui://" .. iData._Icon;
			itemNameAndNum.text = iData._Name .. "(" .. Proxy4Lua.ChangeColor(has, color) .. "/" .. need .. ")";
			levelIcon.url = "ui://shihun/nandu_" .. JieHunSystem.instance._NextDrawData._ID;
			rollBtn.enabled = has >= need and JieHunSystem.instance._NextDrawData ~= nil;
		else
			itemIcon.url = "";
			itemNameAndNum.text = "";
			levelIcon.url = "";
		end
	end
end

function shihun_CheckResult()
 	if minTimer ~= nil and minTimer > 0 then
 		return;
 	end
	if JieHunSystem.instance._LastestChapter ~= nil then
		local hsData = HeroStroyData.GetData(JieHunSystem.instance._LastestChapter.ChapterId);
		if hsData ~= nil then
			local eData = EntityData.GetData(hsData.EntityID_);
			if eData ~= nil then
				local dData = DisplayData.GetData(eData._DisplayId);
				if dData ~= nil then
					rollResultName.text = eData._Name;
					rollResultIcon.url = "ui://" .. dData._HeadIcon;
					rollResultIconShadow.url = "ui://" .. dData._HeadIcon;

					rollCom:GetTransition("t1"):Stop();
					rollCom:GetTransition("t2"):Play();
				end 
			end
		end
		JieHunSystem.instance._LastestChapter = nil;
	end
end

function shihun_OnRoll()
	GRoot.inst:AddChild(rollCom);
	minTimer = 3;
	Proxy4Lua.RandChapter();
end

function shihun_OnRollResult()
	rollCom:RemoveFromParent();
	UIManager.SetDirty("shihun");
end