require "FairyGUI"

paiku = fgui.window_class(WindowBase)
local Window;

local allCardList;
local allCardGroupList;
local total;

local cardItemUrl = "ui://paiku/touxiangkuang_Label";
local cardItemGapUrl = "ui://paiku/kong_com";
local cardGroupUrl = "ui://paiku/paizuanniu_Button";

local cardGroupList;
local crtGroupIdx = 0;
local crtCardInstID = 0;
local crtCardsFee = 0;
local crtCardsType = 0;
local crtCardName;

local crtCardIdx = 0; --当前所有卡的起始索引
local countPerPage = 9; --所有卡组列表 每页的最大数量

local nextBtn;
local prevBtn;
local pageText;

local doubleClickTicker;

local isInGroup;

function paiku:OnEntry()
	Define.LaunchUIBundle("icon");
	Window = paiku.New();
	Window:Show();
end

function paiku:GetWindow()
	return Window;
end

function paiku:OnInit()

	self.contentPane = UIPackage.CreateObject("paiku", "paiku_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n7").asButton;

	local leftPart = self.contentPane:GetChild("n6").asCom;
	allCardList = leftPart:GetChild("n27").asList;
--	allCardList:SetVirtual();
--	allCardList.itemRenderer = paiku_RenderListItem;
	allCardList.onDrop:Add(paiku_OnDropCard);
	allCardList.data = 0;
	allCardList.scrollItemToViewOnClick = false;
	total = leftPart:GetChild("n33").asTextField;

	prevBtn = leftPart:GetChild("n36");
	nextBtn = leftPart:GetChild("n37");
	pageText = leftPart:GetChild("n39");
	prevBtn.onClick:Add(paiku_OnPrevPage);
	nextBtn.onClick:Add(paiku_OnNextPage);

	local feeList = leftPart:GetChild("n34").asList;
	local feeMax = feeList.numItems;
	local feeItem;
	for i=1, feeMax do
		feeItem = feeList:GetChildAt(i-1);
		feeItem.data = i-1;
		feeItem.onClick:Add(paiku_OnFeeItemClick);
	end
	feeList.selectedIndex = crtCardsFee;

	local typeList = leftPart:GetChild("n35").asList;
	local typeMax = typeList.numItems;
	local typeItem;
	for i=1, typeMax do
		typeItem = typeList:GetChildAt(i-1);
		typeItem.data = i-1;
		typeItem.onClick:Add(paiku_OnTypeItemClick);
	end
	typeList.selectedIndex = crtCardsType;

	local rightPart = self.contentPane:GetChild("n5").asCom;
	local bg = rightPart:GetChild("n3");
	allCardGroupList = bg:GetChild("n5").asList;

	local cardGroup = rightPart:GetChild("n6");
	local deleteBtn = cardGroup:GetChild("n24").asButton;
--	crtCardName = cardGroup:GetChild("n23");
	deleteBtn.onClick:Add(paiku_OnDeleteGroup);
	cardGroupList = cardGroup:GetChild("n27").asList;
	cardGroupList.onDrop:Add(paiku_OnDropCard);
	cardGroupList.data = 1;
	local setBattleBtn = cardGroup:GetChild("n29").asButton;
	setBattleBtn.onClick:Add(paiku_OnSetBattle);

--	local changeNameBtn = cardGroup:GetChild("n25").asButton;
--	changeNameBtn.onClick:Add(paiku_OnChangeGroupName);

	--test

	for i=1, 5 do
		local groupItem = allCardGroupList:AddItemFromPool(cardGroupUrl);
		groupItem.onClick:Add(paiku_OnSelectGroup);
		groupItem.data = i - 1;
	end
	allCardGroupList.selectedIndex = crtGroupIdx;

	UIManager.SetDirty("paiku");
end

function paiku_OnFeeItemClick(context)
	crtCardsFee = context.sender.data;
	crtCardIdx = 0;
	UIManager.SetDirty("paiku");
end

function paiku_OnTypeItemClick(context)
	crtCardsType = context.sender.data;
	crtCardIdx = 0;
	UIManager.SetDirty("paiku");
end

function paiku_RenderListItem(index, obj)
	if obj == nil then
		return;
	end

	local feeImg = obj:GetChild("n3");
	if feeImg == nil then
		return;
	end

	local img = obj:GetChild("n5");
	if img == nil then
		return;
	end

	local level = obj:GetChild("n6");
	if level == nil then
		return;
	end

	local fee = obj:GetChild("n7");
	if fee == nil then
		return;
	end

	local inGroup = obj:GetChild("n9");
	if inGroup == nil then
		return;
	end

	local radImg = obj:GetChild("n10");
	if radImg == nil then
		return;
	end

	local iconloader = obj:GetChild("n11");
	if iconloader == nil then
		return;
	end

	local displayData = GamePlayer.GetDisplayDataByIndex(crtCardsFee, crtCardsType, index);
	local entityData = GamePlayer.GetEntityDataByIndex(crtCardsFee, crtCardsType, index);
	img.asLoader.url = "ui://" .. displayData._HeadIcon;
	iconloader.asLoader.url = "ui://" .. displayData._Quality;
	obj.onClick:Add(paiku_OnCardItem);
	local instId = GamePlayer.GetInstID(crtCardsFee, crtCardsType, index);
	obj.data = instId;
	obj.onDragStart:Add(paiku_OnDragCard);
	fee.text = entityData._Cost;
	local isIn = GamePlayer.IsInGroup(instId, crtGroupIdx);
	inGroup.visible = isIn;
	img.enabled = not isIn;
	fee.enabled = not isIn;
	feeImg.enabled = not isIn;
	obj.draggable = not isIn;
	local entityInst = GamePlayer.GetCardByInstID(instId);
	local  levelData =  StrengthenData.GetData( entityInst.UnitId,  entityInst.IProperties[9]+1);
	level.text = entityInst.IProperties[9] .. "";
	local itemNum = BagSystem.GetItemMaxNum(levelData._ItemId);
	if itemNum >= levelData._ItemNum   then
		radImg.visible = true;
	else
		radImg.visible = false;
	end
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
	UIManager.SetDirty("paiku");
end

function paiku:OnUpdate()
	if doubleClickTicker ~= nil then
		doubleClickTicker.crt = doubleClickTicker.crt + Proxy4Lua.DeltaTime;
		if doubleClickTicker.crt >= doubleClickTicker.max then
			doubleClickTicker = nil;
			paiku_ExcuteOnCardItem();
		end
	end
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

function paiku_OnNextPage()
	local cardList = GamePlayer.CardsByFeeAndType(crtCardsFee, crtCardsType);
	if cardList ~= nil then
		if crtCardIdx + countPerPage <= cardList.Count then
			crtCardIdx = crtCardIdx + countPerPage;
		end
	end
	UIManager.SetDirty("paiku");
end

function paiku_OnPrevPage()
	crtCardIdx = crtCardIdx - countPerPage;
	if crtCardIdx < 0 then
		crtCardIdx = 0;
	end
	UIManager.SetDirty("paiku");
end

function paiku_FlushData()
	total.text = "(数量:" .. GamePlayer.CardsByFeeAndType(0, 0).Count .. ")";
	local page = crtCardIdx / countPerPage + 1;
	pageText.text = "第" .. page .. "页";
--	local cards = 0;
--	if GamePlayer.CardsByFeeAndType(crtCardsFee, crtCardsType) ~= nil then
--		cards = GamePlayer.CardsByFeeAndType(crtCardsFee, crtCardsType).Count;
--	end
--	allCardList.numItems = cards;
	local cardList = GamePlayer.CardsByFeeAndType(crtCardsFee, crtCardsType);
	allCardList:RemoveChildrenToPool();

	if cardList ~= nil then
		local max = crtCardIdx + countPerPage;
		if max > cardList.Count then
			max = cardList.Count;
		end
		for i=crtCardIdx, max - 1 do
			paiku_RenderListItem(i, allCardList:AddItemFromPool(cardItemUrl));
		end
		nextBtn.enabled = crtCardIdx + countPerPage < cardList.Count;
		prevBtn.enabled = crtCardIdx ~= 0;
	else
		nextBtn.enabled = false;
		prevBtn.enabled = false;
	end

	cardGroupList:RemoveChildrenToPool(); 
	local groupCards = GamePlayer.GetGroupCards(crtGroupIdx);
	if groupCards == nil then
		return;
	end
	local displayData;
	local entityData;
	for i=1, groupCards.Count do
		if i == 1 or i == 4 or i == 8 then
			cardGroupList:AddItemFromPool(cardItemGapUrl);
		end

		displayData = GamePlayer.GetDisplayDataByIndexFromGroup(crtGroupIdx, i - 1);
		entityData = GamePlayer.GetEntityDataByIndexFromGroup(crtGroupIdx, i - 1);
		local itemBtn = cardGroupList:AddItemFromPool(cardItemUrl);
		itemBtn:GetChild("n5").asLoader.url = "ui://" .. displayData._HeadIcon;
		itemBtn:GetChild("n11").asLoader.url = "ui://" .. displayData._Quality;
		local fee = itemBtn:GetChild("n7");
		fee.text = entityData._Cost
		local radImg = itemBtn:GetChild("n10");
		local level = itemBtn:GetChild("n6");
		local instId = GamePlayer.GetInstIDFromGroup(crtGroupIdx,  i - 1);
		local entityInst = GamePlayer.GetCardByInstID(instId);
		local  levelData =  StrengthenData.GetData( entityInst.UnitId,  entityInst.IProperties[9]+1);
		level.text = entityInst.IProperties[9] .. "";
		local itemNum = BagSystem.GetItemMaxNum(levelData._ItemId);
		if itemNum >= levelData._ItemNum then
			radImg.visible = true;
		else
			radImg.visible = false;
		end

		itemBtn.onClick:Add(paiku_OnCardItem); 
		itemBtn.data = GamePlayer.GetInstIDFromGroup(crtGroupIdx, i - 1);
		itemBtn.draggable = true;
		itemBtn.onDragStart:Add(paiku_OnDragCard);
	end

	local groupItem;
	local groupName;
	for i=1, 5 do
		groupItem = allCardGroupList:GetChildAt(i-1);
--		groupName = GamePlayer.GetGroupName(i - 1);
--		if groupName == "" then
			groupName = "" .. i;
--		end
		groupItem:GetChild("n3").text = groupName;
--		if crtGroupIdx == i - 1 then
--			crtCardName.text = groupName;
--		end
		if i - 1 == GamePlayer._CrtBattleGroupIdx then
			groupItem:GetChild("n4").visible = true;
		else
			groupItem:GetChild("n4").visible = false;
		end
	end
end

function paiku_OnCardItem(context)
	crtCardInstID = context.sender.data;
	if doubleClickTicker == nil then
		doubleClickTicker = {};
		doubleClickTicker.crt = 0;
		doubleClickTicker.max = 0.2;
	else
		if doubleClickTicker.crt <= doubleClickTicker.max then
			paiku_OnMessageConfirm();
			doubleClickTicker = nil;
		else
			doubleClickTicker = {};
			doubleClickTicker.crt = 0;
			doubleClickTicker.max = 0.2;
		end
	end
end

function paiku_ExcuteOnCardItem()
	UIParamHolder.Set("qiecuo1", crtCardInstID);
	UIParamHolder.Set("qiecuo2", false);
	UIParamHolder.Set("showBoos",false);
	UIManager.Show("xiangxiziliao");
end

--function paiku_OnChangeGroupName(context)
--	GamePlayer.ChangeGroupName(crtGroupIdx, crtCardName.text);
--	UIManager.SetDirty("paiku");
--end

function paiku_OnDragCard(context)
	context:PreventDefault();
	DragDrop.inst:StartDrag(context.sender, context.sender:GetChild("n5").asLoader.url, context.sender, context.data);
end

function paiku_OnDropCard(context)
	local onGroupArea = context.sender.data == 1;
	crtCardInstID = context.data.data;
	isInGroup = GamePlayer.IsInGroup(crtCardInstID, crtGroupIdx);
	if onGroupArea and isInGroup then
		DragDrop.inst:Cancel();
		UIManager.SetDirty("paiku");
		return;
	end

	if onGroupArea == false and isInGroup == false then
		DragDrop.inst:Cancel();
		UIManager.SetDirty("paiku");
		return;
	end

--	if onGroupArea and GamePlayer.IsGroupMax(crtGroupIdx) then
--		local MessageBox = UIManager.ShowMessageBox();
--		MessageBox:SetData("提示", "卡组已满", true);
--		UIManager.SetDirty("paiku");
--		return;
--	end

	paiku_OnMessageConfirm();
--	--[[local MessageBox = UIManager.ShowMessageBox();
--	if isInGroup then
--		MessageBox:SetData("提示", "是否取出卡组？", false, paiku_OnMessageConfirm, paiku_OnMessageCancel);
--	else
--		MessageBox:SetData("提示", "是否加入卡组？", false, paiku_OnMessageConfirm, paiku_OnMessageCancel);
--	end--]]
end

function paiku_OnMessageConfirm()
	isInGroup = GamePlayer.IsInGroup(crtCardInstID, crtGroupIdx);
	if isInGroup then
		GamePlayer.TakeOffCard(crtCardInstID, crtGroupIdx);
	else
		local max = GamePlayer.IsGroupMax(crtGroupIdx);
		if max then
			local MessageBox = UIManager.ShowMessageBox();
			MessageBox:SetData("提示", "卡组已满", true);
		else
			GamePlayer.PutInCard(crtCardInstID, crtGroupIdx);
		end
	end
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