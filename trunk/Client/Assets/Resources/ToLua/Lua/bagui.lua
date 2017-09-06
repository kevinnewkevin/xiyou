require "FairyGUI"

bagui = fgui.window_class(WindowBase)
local Window;

local crtTab = 0;
local crtClickItemIdx = 0;
local fullImg;

function bagui:OnEntry()
	Window = bagui.New();
	Window:Show();
end

function bagui:GetWindow()
	return Window;
end

function bagui:OnInit()
	self.contentPane = UIPackage.CreateObject("bagui", "beibao_com").asCom;
	self:Center();
	self.modal = true;

	self.closeButton = self.contentPane:GetChild("n6");

	bagList = self.contentPane:GetChild("n10").asList;
	bagList:SetVirtual();
	bagList.itemRenderer = bagui_RenderListItem;
	fullImg = self.contentPane:GetChild("n23");
	fullImg.visible = false;
	crtTab = 0;

	bagui_FlushData();
end

function bagui_RenderListItem(index, obj)
	local stackLab = obj:GetChild("n2");
	local icon = obj:GetChild("n1");
	local iconBack = obj:GetChild("n0");
	obj.onClick:Add(bag_OnItem);
	obj.data = index;
	local iteminst;
	iteminst = BagSystem.GetItemInstByIndex(index, crtTab);
	local itemdata;
	local stackNum;
	if iteminst ~= null then
		itemdata = ItemData.GetData(iteminst.ItemId);
		stackNum = iteminst.Stack_;
	end
	local iconpath = "";
	if itemdata ~= nil then
		iconpath = itemdata._Icon;
	end
	icon.asLoader.url = "ui://" .. iconpath;
	stackLab.text ="" .. stackNum;
	iconBack.asLoader.url = "ui://icon/" .. itemdata._IconBack;
end

function bag_OnItem(context)
	crtClickItemIdx = context.sender.data;
	UIManager.Show("bagtips");
end

function bagui:OnUpdate()
	if UIManager.IsDirty("bagui") then
		bagui_FlushData();
		UIManager.ClearDirty("bagui");
	end
end

function bagui:OnTick()
	
end

function bagui:isShow()
	return Window.isShowing;
end

function bagui:OnDispose()
	Window:Dispose();
end

function bagui:OnHide()
	Window:Hide();
end

function bagui_FlushData()
	bagList.numItems = BagSystem.GetItemCount(crtTab);
end

function bagui:GetClickItemIdx()
	return crtClickItemIdx;
end

function bagui:GetCrtTab()
	return crtTab;
end