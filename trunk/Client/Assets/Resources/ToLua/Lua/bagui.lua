require "FairyGUI"

bagui = fgui.window_class(WindowBase)
local Window;

local crtTab;
local crtClickItemIdx;

function bagui:OnEntry()
	Window = bagui.New();
	Window:Show();
end

function bagui:GetWindow()
	return Window;
end

function bagui:OnInit()
	self.contentPane = UIPackage.CreateObject("bagui", "bagui_com").asCom;
	self:Center();
	self.modal = true;

	self.closeButton = self.contentPane:GetChild("n4");

	bagList = self.contentPane:GetChild("n27").asList;
	bagList:SetVirtual();
	bagList.itemRenderer = bagui_RenderListItem;

	crtTab = 0;

	bagui_FlushData();
end

function bagui_RenderListItem(index, obj)
	local stack = obj:GetChild("n1");
	local icon = obj:GetChild("n2").asLoader;
	obj.onClick:Add(bag_OnItem);
	obj.data = index;
	local iteminst;
	iteminst = BagSystem.GetItemInstByIndex(index, crtTab);
	local itemdata;
	local stackNum = "";
	if iteminst ~= null then
		itemdata = ItemData.GetData(iteminst.tableid);
		stackNum = iteminst.stack;
	end
	local iconpath = "";
	if itemdata ~= nil then
		iconpath = itemdata._Icon;
	end
	stack.text = stackNum;
	icon.url = "ui://" + iconpath;
end

function bag_OnItem(context)
	crtClickItemIdx = context.sender.data;
	UIManager.Show("bagtips");
end

function bagui:OnUpdate()
	if UIManager.IsDirty("bagui") then
		xiangxiziliao_FlushData();
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
	allCardList.numItems = BagSystem.GetCount(crtTab);
end