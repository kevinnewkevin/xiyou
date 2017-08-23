require "FairyGUI"

paiku = fgui.window_class(WindowBase)
local Window;

local allCardList;
local allCardGroupList;

local cardItemUrl = "ui://paiku/touxiangkuang_Label";
local cardGroupUrl = "ui://paiku/paizuanniu_Button";

local cardGroupList;
local crtGroupIdx = 0;
local crtCardInstID = 0;
local crtCardsFee = 0;

local isInGroup;

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

	local feeList = leftPart:GetChild("n34").asList;
	local feeMax = feeList.numItems;
	local feeItem;
	for i=1, feeMax do
		feeItem = feeList:GetChildAt(i-1);
		feeItem.data = i-1;
		feeItem.onClick:Add(paiku_OnFeeItemClick);
	end
	feeList.selectedIndex = crtCardsFee;

	local rightPart = self.contentPane:GetChild("n5").asCom;
	local bg = rightPart:GetChild("n3");
	allCardGroupList = bg:GetChild("n5").asList;

	local cardGroup = rightPart:GetChild("n6");
	local deleteBtn = cardGroup:GetChild("n24").asButton;
	deleteBtn.onClick:Add(paiku_OnDeleteGroup);
	cardGroupList = cardGroup:GetChild("n27").asList;
	local setBattleBtn = cardGroup:GetChild("n29").asButton;
	setBattleBtn.onClick:Add(paiku_OnSetBattle);

	--test

	for i=1, 5 do
		local groupItem = allCardGroupList:AddItemFromPool(cardGroupUrl);
		groupItem.onClick:Add(paiku_OnSelectGroup);
		groupItem.data = i - 1;
	end
	allCardGroupList.selectedIndex = crtGroupIdx;

	paiku_FlushData();
end

function paiku_OnFeeItemClick(context)
	crtCardsFee = context.sender.data;
	UIManager.SetDirty("paiku");
end

function paiku_RenderListItem(index, obj)
	obj:GetChild("n5").asLoader.url = "ui://" .. GamePlayer.GetCardHeadIcon(index);
	obj.onClick:Add(paiku_OnCardItem);
	obj.data = GamePlayer.GetInstIDInMyCards(crtCardsFee, index);
	obj.draggable = true;
	obj.onDragEnd:Add(paiku_OnDropCard);
	local fee = obj:GetChild("n7");
	fee.text = GamePlayer.GetFeeInMyCards(crtCardsFee, index);
end

function paiku_OnDeleteGroup(context)
	local MessageBox = UIManager.ShowMessageBox();
	MessageBox:SetData("提示", "是否删除卡组？", false, paiku_OnDelete);
end

function paiku_OnSetBattle(context)
	local MessageBox = UIManager.ShowMessageBox();
	MessageBox:SetData("提示", "是否设置当前卡组为出战卡组？", false, paiku_OnSet);
end

function paiku_OnSet(context)
	GamePlayer._CrtBattleGroupIdx = crtGroupIdx;
	UIManager.HideMessageBox();
	UIManager.SetDirty("paiku");
end

function paiku_OnDelete()
	GamePlayer.DeleteGroup(crtGroupIdx);
	UIManager.HideMessageBox();
	UIManager.SetDirty("paiku");
end

function paiku_OnSelectGroup(context)
	crtGroupIdx = context.sender.data;
	paiku_FlushData();
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
	local cards = 0;
	if GamePlayer.CardsByFee(crtCardsFee) ~= nil then
		cards = GamePlayer.CardsByFee(crtCardsFee).Count;
	end
	allCardList.numItems = cards;
	cardGroupList:RemoveChildrenToPool();
	local groupCards = GamePlayer.GetGroupCards(crtGroupIdx);
	if groupCards == nil then
		return;
	end
	for i=1, groupCards.Count do
		local itemBtn = cardGroupList:AddItemFromPool(cardItemUrl);
		itemBtn:GetChild("n5").asLoader.url = "ui://" .. GamePlayer.GetGroupCardHeadIcon(crtGroupIdx, i - 1);
		local fee = itemBtn:GetChild("n7");
		fee.text = GamePlayer.GetFeeInGroupCards(crtGroupIdx, i - 1);
		itemBtn.onClick:Add(paiku_OnCardInGroup);
		itemBtn.data = GamePlayer.GetInstIDInMyGroup(crtGroupIdx, i - 1);
		itemBtn.draggable = true;
		itemBtn.onDragEnd:Add(paiku_OnDropCard);
	end

	local groupItem;
	for i=1, 5 do
		groupItem = allCardGroupList:GetChildAt(i-1);
		if i - 1 == GamePlayer._CrtBattleGroupIdx then
			groupItem:GetChild("n4").visible = true;
		else
			groupItem:GetChild("n4").visible = false;
		end
	end
end

function paiku_OnCardInGroup(context)
	crtCardInstID = context.sender.data;
	UIManager.Show("xiangxiziliao");
end

function paiku_OnCardItem(context)
	crtCardInstID = context.sender.data;
	UIManager.Show("xiangxiziliao");
end

function paiku_OnDropCard(context)
	crtCardInstID = context.sender.data;
	isInGroup = GamePlayer.IsInGroup(crtCardInstID, crtGroupIdx);
	local MessageBox = UIManager.ShowMessageBox();
	if isInGroup then
		MessageBox:SetData("提示", "是否取出卡组？", false, paiku_OnMessageConfirm, paiku_OnMessageCancel);
	else
		MessageBox:SetData("提示", "是否加入卡组？", false, paiku_OnMessageConfirm, paiku_OnMessageCancel);
	end
end

function paiku_OnMessageConfirm()
	isInGroup = GamePlayer.IsInGroup(crtCardInstID, crtGroupIdx);
	if isInGroup then
		GamePlayer.TakeOffCard(crtCardInstID, crtGroupIdx);
		print("TakeOffCard");
	else
		GamePlayer.PutInCard(crtCardInstID, crtGroupIdx);
		print("PutInCard");
	end
	UIManager.HideMessageBox();
	UIManager.SetDirty("paiku");
end

function paiku_OnMessageCancel()
	UIManager.HideMessageBox();
	UIManager.SetDirty("paiku");
end

function paiku:GetCrtGroup()
	return crtGroupIdx;
end

function paiku:GetCrtCard()
	return crtCardInstID;
end

function paiku:IsInGroup()
	return GamePlayer.IsInGroup(crtCardInstID, crtGroupIdx);
end