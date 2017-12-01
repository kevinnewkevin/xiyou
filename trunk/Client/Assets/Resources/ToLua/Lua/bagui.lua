require "FairyGUI"

bagui = fgui.window_class(WindowBase)
local Window;

local crtTab = 0;
local crtClickItemIdx = 0;
local fullImg;
local gold;
local hunBi;
local chopper;
local stamaPoint;
local bagList;

function bagui:OnEntry()
	Define.LaunchUIBundle("icon");
	Window = bagui.New();
	Window:Show();
end

function bagui:GetWindow()
	return Window;
end

function bagui:OnInit()
	self.contentPane = UIPackage.CreateObject("bagui", "beibao_com").asCom;
	self:Center();

	self.closeButton = self.contentPane:GetChild("n6");

	bagList = self.contentPane:GetChild("n10").asList;
	bagList:SetVirtual();
	bagList.itemRenderer = bagui_RenderListItem;
	fullImg = self.contentPane:GetChild("n23");
	fullImg.visible = false;
	crtTab = 0;
	local moneyGroup = self.contentPane:GetChild("n29").asCom;
	gold = moneyGroup:GetChild("n5");
	chopper = moneyGroup:GetChild("n7");
	stamaPoint = moneyGroup:GetChild("n12");
	hunBi = moneyGroup:GetChild("n16");
	local feeList = self.contentPane:GetChild("n9").asList;
	local feeMax = feeList.numItems;
	local feeItem;
	for i=1, feeMax do
		feeItem = feeList:GetChildAt(i-1);
		feeItem.data = i-1;
		feeItem.onClick:Add(bagui_OnFeeItemClick);
	end
	feeList.selectedIndex = crtTab;
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
		stackNum = iteminst.Stack;
	end
	local iconpath = "";
	if itemdata ~= nil then
		iconpath = itemdata._Icon;
	end
	icon.asLoader.url = "ui://" .. iconpath;
	stackLab.text ="" .. stackNum;
	iconBack.asLoader.url = "ui://" .. itemdata._IconBack;
end

function bag_OnItem(context)
	crtClickItemIdx = context.sender.data;
	UIParamHolder.Set("tipsItem", 0);
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
	gold.text = GamePlayer._Data.IProperties[8];
	chopper.text = GamePlayer._Data.IProperties[6];
	stamaPoint.text = GamePlayer._Data.IProperties[10];
	hunBi.text = GamePlayer._Data.IProperties[11];
end

function bagui:GetClickItemIdx()
	return crtClickItemIdx;
end

function bagui:GetCrtTab()
	return crtTab;
end

function bagui_OnFeeItemClick(context)
	crtTab =  context.sender.data;
	bagui_FlushData();
end