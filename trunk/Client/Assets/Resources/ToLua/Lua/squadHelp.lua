require "FairyGUI"

squadHelp = fgui.window_class(WindowBase)
local Window;

local cards;
local cardList;

function squadHelp:OnEntry()
	Window = squadHelp.New();
	Window:Show();
end

function squadHelp:GetWindow()
	return Window;
end

function squadHelp:OnInit()
	self.contentPane = UIPackage.CreateObject("bangpai", "juanzeng_com").asCom;
	self:Center();

	cards = self.contentPane:GetChild("n6").asList;
	cards:SetVirtual();
	cards.itemRenderer = squadHelp_RenderListItem;

	squadHelp_FlushData();
end

function squadHelp_RenderListItem(index, obj)
	if obj == nil then
		return;
	end

	local instid = GamePlayer.GetInstID(0, 0, index);
	local cardInst = GamePlayer.GetCardByInstID(instid);
	if cardInst == nil then
		return;
	end
	local entityData = EntityData.GetData(cardInst.UnitId);
	if entityData == nil then
		return;
	end
	local displayData = DisplayData.GetData(entityData._DisplayId);
	if displayData == nil then
		return;
	end
	obj:GetChild("n15").asLoader.url = "ui://" .. displayData._HeadIcon;
	obj:GetChild("n16").asLoader.url = "ui://" .. displayData._Quality;
	local levelData =  StrengthenData.GetData(cardInst.UnitId, cardInst.IProperties[9]+1);
	if levelData == nil then
		return;
	end
	local itemNum = BagSystem.GetItemMaxNum(levelData._ItemId);
	obj:GetChild("n18").text = itemNum .. "/" .. levelData._ItemNum;
	obj.data = levelData._ItemId;
	obj.onClick:Add(squadHelp_OnHelp);
end

function squadHelp_OnHelp(context)
	if context.sender.data ~= nil then
		Proxy4Lua.NeedAssistantItem(context.sender.data);
	end
	UIManager.Hide("squadHelp");
end

function squadHelp:OnUpdate()
	if UIManager.IsDirty("squadHelp") then
		squadHelp_FlushData();
		UIManager.ClearDirty("squadHelp");
	end
end

function squadHelp:OnTick()
	
end

function squadHelp:isShow()
	return Window.isShowing;
end

function squadHelp:OnDispose()
	Window:Dispose();
end

function squad:OnHide()
	Window:Hide();
end

function squadHelp_FlushData()
	cardList = GamePlayer.CardsByFeeAndType(0, 0);
	cards.numItems = cardList.Count();
end