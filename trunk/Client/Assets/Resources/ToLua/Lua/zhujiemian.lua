require "FairyGUI"

zhujiemian = fgui.window_class(WindowBase)
local Window;

local playerName;
local playerLevel;
local playerExp;

local stateIcon;
local stateBar;

function zhujiemian:OnEntry()
	Window = zhujiemian.New();
	Window:Show();
end

function zhujiemian:OnInit()
	UIPackage.AddPackage("UI/UI_Fairy/export/zhujiemian");
	self.contentPane = UIPackage.CreateObject("zhujiemian", "Component1").asCom;
	self:Center();

	stateIcon = self.contentPane:GetChild("n17").asButton;
	stateIcon.onClick:Add(zhujiemian_OnFolder);

	local bottomGroup = self.contentPane:GetChild("n18");
	stateBar = bottomGroup:GetController("c1");
	local bottomBtnList = bottomGroup:GetChild("n15").asList;
	local cardCargo = bottomBtnList:GetChildAt(3);
	cardCargo.onClick:Add(zhujiemian_OnCardCargo);

	local infoGroup = self.contentPane:GetChild("n15").asCom;
	playerName = infoGroup:GetChild("n9");
	playerExp = infoGroup:GetChild("n10");
	playerLevel = infoGroup:GetChild("n11");

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
	playerName.text = GamePlayer._Name;
	playerExp.text = "754782";
	playerLevel.text = "89";
end

function zhujiemian_OnFolder()
	stateIcon:GetController("icon").selectedIndex = (stateIcon:GetController("icon").selectedIndex + 1) % 2;
	stateBar.selectedIndex = stateIcon:GetController("icon").selectedIndex;
end

function zhujiemian_OnCardCargo()
	UIManager.Show("paiku_Component");
end