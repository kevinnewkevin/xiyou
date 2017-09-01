require "FairyGUI"

zhujiemian = fgui.window_class(WindowBase)
local Window;

local playerName;
local playerLevel;
local playerExp;
local gold;
local chopper;
local stamaPoint;
local expBar;
local headIcon;

local stateIcon;
local stateBar;

function zhujiemian:OnEntry()
	Window = zhujiemian.New();
	Window:Show();
end

function zhujiemian:OnInit()
	self.contentPane = UIPackage.CreateObject("zhujiemian", "zhujiemian_com").asCom;
	self:Center();

	stateIcon = self.contentPane:GetChild("n17").asButton;
	stateIcon.onClick:Add(zhujiemian_OnFolder);

	local bottomGroup = self.contentPane:GetChild("n18");
	stateBar = bottomGroup:GetController("c1");
	local bottomBtnList = bottomGroup:GetChild("n15").asList;
	local cardCargo = bottomBtnList:GetChildAt(3);
	cardCargo.onClick:Add(zhujiemian_OnCardCargo);
	local taskBtn = bottomBtnList:GetChildAt(4);
	taskBtn.onClick:Add(zhujiemian_OnTaskBtn);
	local qiecuoBtn = bottomBtnList:GetChildAt(6);
	qiecuoBtn.onClick:Add(zhujiemian_OnQieCuoBtn);
	local infoGroup = self.contentPane:GetChild("n15").asCom;
	playerName = infoGroup:GetChild("n9");
	playerExp = infoGroup:GetChild("n10");
	playerLevel = infoGroup:GetChild("n11");
	expBar = infoGroup:GetChild("n8").asProgress;
	headIcon = infoGroup:GetChild("n13").asLoader;

	local moneyGroup = self.contentPane:GetChild("n16").asCom;
	gold = moneyGroup:GetChild("n5");
	chopper = moneyGroup:GetChild("n7");
	stamaPoint = moneyGroup:GetChild("n12");

	zhujiemian_FlushData();
end

function zhujiemian:GetWindow()
	return Window;
end

function zhujiemian:OnUpdate()
	if UIManager.IsDirty("zhujiemian") then
		zhujiemian_FlushData();
		UIManager.ClearDirty("zhujiemian");
	end
end

function zhujiemian:OnTick()
	
end

function zhujiemian:isShow()
	return Window.isShowing;
end

function zhujiemian:OnDispose()
	Window:Dispose();
end

function zhujiemian:OnHide()
	Window:Hide();
end

function zhujiemian_FlushData()
	local displayData = GamePlayer.GetMyDisplayData();
	headIcon.url = "ui://" .. displayData._HeadIcon;
	playerName.text = GamePlayer._Name;
	playerExp.text = GamePlayer._Data.IProperties[4] .. "/" .. 99999;
	playerLevel.text = GamePlayer._Data.IProperties[5];

	gold.text = GamePlayer._Data.IProperties[8];
	chopper.text = GamePlayer._Data.IProperties[6];
	stamaPoint.text = GamePlayer._Data.IProperties[2];

	expBar.value = GamePlayer._Data.IProperties[4] / 99999.0 * 100;
end

function zhujiemian_OnFolder()
	stateIcon:GetController("icon").selectedIndex = (stateIcon:GetController("icon").selectedIndex + 1) % 2;
	stateBar.selectedIndex = stateIcon:GetController("icon").selectedIndex;
end

function zhujiemian_OnCardCargo()
	UIManager.Show("paiku");
end

function zhujiemian_OnTaskBtn()
	UIManager.Show("jiehun");
end

function zhujiemian_OnQieCuoBtn()
	UIManager.Show("qiecuo");
end
