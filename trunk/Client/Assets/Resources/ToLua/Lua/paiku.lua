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
local crtCardName;

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
	crtCardName = cardGroup:GetChild("n23");
	deleteBtn.onClick:Add(paiku_OnDeleteGroup);
	cardGroupList = cardGroup:GetChild("n27").asList;
	local setBattleBtn = cardGroup:GetChild("n29").asButton;
	setBattleBtn.onClick:Add(paiku_OnSetBattle);
	local changeNameBtn = cardGroup:GetChild("n25").asButton;
	changeNameBtn.onClick:Add(paiku_OnChangeGroupName);

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
	local displayData = GamePlayer.GetDisplayDataByIndex(crtCardsFee, index);
	local entityData = GamePlayer.GetEntityDataByIndex(crtCardsFee, index);
	local img = obj:GetChild("n5");
	local feeImg = obj:GetChild("n3");
	img.asLoader.url = "ui://" .. displayData._HeadIcon;
	obj.onClick:Add(paiku_OnCardItem);
	local instId = GamePlayer.GetInstID(crtCardsFee, index);
	obj.data = instId;
	obj.onDragEnd:Add(paiku_OnDropCard);
	local fee = obj:GetChild("n7");
	local inGroup = obj:GetChild("n9");
	fee.text = entityData._Cost;
	local isIn = GamePlayer.IsInGroup(instId, crtGroupIdx);
	inGroup.visible = isIn;
	img.enabled = not isIn;
	fee.enabled = not isIn;
	feeImg.enabled = not isIn;
	obj.draggable = not isIn;
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
	local displayData;
	local entityData;
	for i=1, groupCards.Count do
		displayData = GamePlayer.GetDisplayDataByIndexFromGroup(crtGroupIdx, i - 1);
		entityData = GamePlayer.GetEntityDataByIndexFromGroup(crtGroupIdx, i - 1);
		local itemBtn = cardGroupList:AddItemFromPool(cardItemUrl);
		itemBtn:GetChild("n5").asLoader.url = "ui://" .. displayData._HeadIcon;
		local fee = itemBtn:GetChild("n7");
		fee.text = entityData._Cost
		itemBtn.onClick:Add(paiku_OnCardInGroup);
		itemBtn.data = GamePlayer.GetInstIDFromGroup(crtGroupIdx, i - 1);
		itemBtn.draggable = true;
		itemBtn.onDragEnd:Add(paiku_OnDropCard);
	end

	local groupItem;
	local groupName;
	for i=1, 5 do
		groupItem = allCardGroupList:GetChildAt(i-1);
		groupName = GamePlayer.GetGroupName(i - 1);
		if groupName == "" then
			groupName = "卡组" .. i;
		end
		groupItem:GetChild("n3").text = groupName;
		if crtGroupIdx == i - 1 then
			crtCardName.text = groupName;
		end
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

function paiku_OnChangeGroupName(context)
	GamePlayer.ChangeGroupName(crtGroupIdx, crtCardName.text);
	UIManager.SetDirty("paiku");
end

function paiku_OnDropCard(context)
	crtCardInstID = context.sender.data;
	isInGroup = GamePlayer.IsInGroup(crtCardInstID, crtGroupIdx);
	paiku_OnMessageConfirm();
	--[[local MessageBox = UIManager.ShowMessageBox();
	if isInGroup then
		MessageBox:SetData("提示", "是否取出卡组？", false, paiku_OnMessageConfirm, paiku_OnMessageCancel);
	else
		MessageBox:SetData("提示", "是否加入卡组？", false, paiku_OnMessageConfirm, paiku_OnMessageCancel);
	end--]]
end

function paiku_OnMessageConfirm()
	if GamePlayer.IsGroupMax(crtGroupIdx) then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "卡组已满", true);
		UIManager.SetDirty("paiku");
		return;
	end

	isInGroup = GamePlayer.IsInGroup(crtCardInstID, crtGroupIdx);
	if isInGroup then
		GamePlayer.TakeOffCard(crtCardInstID, crtGroupIdx);
		print("TakeOffCard");
	else
		GamePlayer.PutInCard(crtCardInstID, crtGroupIdx);
		print("PutInCard");
	end
	--UIManager.HideMessageBox();
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