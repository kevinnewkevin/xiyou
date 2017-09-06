require "FairyGUI"

bagtips = fgui.window_class(WindowBase)
local Window;
local nameLab;
local descLab;
local levelLab;
local typeLab;
local icon;
local iconBack;
local stackLab;

function bagtips:OnEntry()
	Window = bagtips.New();
	Window:Show();
end

function bagtips:GetWindow()
	return Window;
end

function bagtips:OnInit()
	self.contentPane = UIPackage.CreateObject("bagtips", "bagtips_com").asCom;
	self:Center();
	self.modal = true;

	self.closeButton = self.contentPane:GetChild("n11");

	nameLab = self.contentPane:GetChild("n9");
	descLab= self.contentPane:GetChild("n24");
	levelLab= self.contentPane:GetChild("n21");
	typeLab= self.contentPane:GetChild("n22");

	local itemIcon = self.contentPane:GetChild("n10");
	icon = itemIcon:GetChild("n1");
	iconBack = itemIcon:GetChild("n0");
	stackLab = itemIcon:GetChild("n2");
	bagtips_FlushData();
end

function bagtips:OnUpdate()
	if UIManager.IsDirty("bagtips") then
		bagtips_FlushData();
		UIManager.ClearDirty("bagtips");
	end
end

function bagtips:OnTick()
	
end

function bagtips:isShow()
	return Window.isShowing;
end

function bagtips:OnDispose()
	Window:Dispose();
end

function bagtips:OnHide()
	Window:Hide();
end

function bagtips_FlushData()
	local bagwin = UIManager.GetWindow("bagui");
	local iteminst = BagSystem.GetItemInstByIndex(bagwin.GetClickItemIdx(), bagwin.GetCrtTab());
	local itemdata;
	if iteminst ~= null then
		itemdata = ItemData.GetData(iteminst.ItemId);
	end
	local iconpath = "";
	if itemdata == nil then
		return;
	end
	icon.asLoader.url = "ui://" .. itemdata._Icon;
	stackLab.text ="" .. iteminst.Stack_;
	iconBack.asLoader.url = "ui://icon/" .. itemdata._IconBack;
	descLab.text = itemdata._Desc;
	nameLab.text = itemdata._Name;
	typeLab.text = "碎片";
end