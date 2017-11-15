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
	cards.itemRenderer = squad_RenderListItem;

	squadHelp_FlushData();
end

function squadHelp_RenderListItem(index, obj)
	if obj == nil then
		return;
	end

	local displayData = GamePlayer.GetDisplayDataByIndex(0, 0, index);
	obj:GetChild("n15").asLoader.url = "ui://" .. displayData._HeadIcon;
	obj:GetChild("n16").asLoader.url = "ui://" .. displayData._Quality;
	obj:GetChild("n18").text = "";
	obj.data = index;
	obj.onClick:Add(squadHelp_OnHelp);
end

function squadHelp_OnHelp(context)
--	local chat = COM_Chat.New();
--	chat.Type = 3;
--	chat.PlayerInstId = GamePlayer._InstID;
--	chat.PlayerName = GamePlayer._Name;
--	chat.HeadIcon = GamePlayer.GetMyDisplayData()._HeadIcon;
--	chat.Level = GamePlayer._Data.IProperties[9];
--	chat.Content = "";
--	Proxy4Lua.SendChat(chat);
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