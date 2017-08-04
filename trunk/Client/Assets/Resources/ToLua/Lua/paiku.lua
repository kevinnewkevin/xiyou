require "FairyGUI"

paiku = fgui.window_class(WindowBase)
local Window;

local allCardList;
local allCardGroupList;

local cardItemUrl = "ui://paiku/touxiangkuang_Label";
local cardGroupUrl = "ui://paiku/paizuanniu_Button";

function paiku:OnEntry()
	Window = paiku.New();
	Window:Show();
end

function paiku:GetWindow()
	return Window;
end

function paiku:OnInit()

	self.contentPane = UIPackage.CreateObject("paiku", "paiku_com").asCom;
	self:Center();

	self.closeButton = self.contentPane:GetChild("n7").asButton;

	local leftPart = self.contentPane:GetChild("n6").asCom;
	allCardList = leftPart:GetChild("n27").asList;
	allCardList:SetVirtual();
	allCardList.itemRenderer = paiku_RenderListItem;
	allCardList.numItems = GamePlayer._Cards.Count;

	local rightPart = self.contentPane:GetChild("n5").asCom;
	local bg = rightPart:GetChild("n3");
	allCardGroupList = bg:GetChild("n5").asList;

	--test
--	local cardNum = GamePlayer._Cards.Count;
--	for i=1, 500 do
--		local itemBtn = allCardList:AddItemFromPool(cardItemUrl);
--		itemBtn.onClick:Add(paiku_OnCardItem);
--	end

--	for i=1, 5 do
--		allCardGroupList:AddItemFromPool(cardGroupUrl);
--	end

	paiku_FlushData();
end

function paiku_RenderListItem(index, obj)
	obj.onClick:Add(paiku_OnCardItem);
	obj.data = index;
end

function paiku:OnUpdate()
	if UIManager.IsDirty("paiku") then
		paiku_FlushData();
		UIManager.ClearDirty("paiku");
	end
end

function paiku:OnTick()
	
end

function paiku:isShow()
	return Window.isShowing;
end

function paiku:OnDispose()
	Window:Dispose();
end

function paiku:OnHide()
	Window:Hide();
end

function paiku_FlushData()
	
end

function paiku_OnCardItem(context)
	UIParamHolder.Set("paiku_OnCardItem", context.sender.data);
	UIManager.Show("xiangxiziliao");
end