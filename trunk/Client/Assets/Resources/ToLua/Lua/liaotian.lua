require "FairyGUI"

liaotian = fgui.window_class(WindowBase)
local Window;

local sysCom = "ui://liaotian/xitong_com";
local otherCom = "ui://liaotian/duifang_com";
local selfCom = "ui://liaotian/wofang_com";

local contentList;
local typeList;
local crtType;
local crtList;

local sendBtn;
local emojiBtn;
local yyBtn;
local content;

function liaotian:OnEntry()
	Window = liaotian.New();
	Window:Show();
end

function liaotian:GetWindow()
	return Window;
end

function liaotian:OnInit()
	self.contentPane = UIPackage.CreateObject("liaotian", "liaotian_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n1");

	sendBtn = self.contentPane:GetChild("n8").asButton;
	emojiBtn = self.contentPane:GetChild("n10").asButton;
	yyBtn = self.contentPane:GetChild("n11").asButton;
	content = self.contentPane:GetChild("n12");
	sendBtn.onClick:Add(liaotian_OnSend);
	yyBtn.onTouchBegin:Add(liaotian_OnYYBegin);
	yyBtn.onTouchEnd:Add(liaotian_OnYYEnd);

	typeList = self.contentPane:GetChild("n6").asList;

	contentList = self.contentPane:GetChild("n13").asList;
	contentList:SetVirtual();
	contentList.itemProvider = liaotian_GetListItemResource;
	contentList.itemRenderer = liaotian_OnRenderListItem;

	crtType = 0;
	liaotian_FlushData();
end

function liaotian_GetListItemResource(index)
	if crtList == nil then
		return;
	end

	if GamePlayer.IsMe(crtList[index].PlayerInstId) then
		return selfCom;
	elseif crtList[index].Type == 0 then
		return sysCom;
	else
		return otherCom;
	end
end

function liaotian_OnRenderListItem(index, obj)
	if crtList == nil then
		return;
	end

	for i=0, crtList.Count - 1 do
		if crtList[index].Type == 0 then --系统
			local content = obj:GetChild("n4").asTextField;
			content.text = "系统:" .. crtList[index].Content;
		else
			local yybtn = obj:GetChild("n8").asCom;
			local yybg = obj:GetChild("n9");
			local icon = obj:GetChild("n1").asLoader;
			local name = obj:GetChild("n5").asTextField;
			local content = obj:GetChild("n7");
			local lv = obj:GetChild("n3");

			yybtn.onClick.Add(liaotian_OnPlayRecord);

			if crtList[index].AudioId ~= 0 then
				icon.visible = false;
				name.visible = false;
				content.visible = false;
				lv.visible = false;
				yybtn.visible = true;
				yybg.visible = true;
				yybtn.data = crtList[index].AudioId;
				yybtn:GetChild("n3").visible = crtList[index].Audio == nil;
			else
				icon.visible = true;
				name.visible = true;
				content.visible = true;
				lv.visible = true;
				yybtn.visible = false;
				yybg.visible = false;

				icon.url = "ui://" .. crtList[index].HeadIcon;
				name.text = crtList[index].PlayerName;
				content.text = crtList[index].Content;
				lv.text = crtList[index].Level;
			end
		end
	end
end

function liaotian_OnYYBegin()
	YYSystem.StartRecord();
end

function liaotian_OnYYEnd()
	YYSystem.StopRecord();
end

function liaotian_OnPlayRecord(context)
	local record = ChatSystem.GetRecord(context.sender.data);
	if record == nil then
		Proxy4Lua.PlayAudio(context.sender.data);
	else
		YYSystem.PlayRecord(record);
	end
end

function liaotian_OnSend()
	if content.text == "" then
		return;
	end

	local chat = COM_Chat.New();
	chat.Type = 1;
	chat.PlayerInstId = GamePlayer._InstID;
	chat.PlayerName = GamePlayer._Name;
	chat.HeadIcon = GamePlayer.GetMyDisplayData()._HeadIcon;
	chat.Level = GamePlayer._Data.IProperties[9];
	chat.Content = content.text;
	Proxy4Lua.SendChat(chat);
	content.text = "";
end

function liaotian:OnUpdate()
	if UIManager.IsDirty("liaotian") then
		liaotian_FlushData();
		UIManager.ClearDirty("liaotian");
	end
end

function liaotian_OnTypeSelect()
	crtType = typeList.selectedIndex;
	UIManager.SetDirty("liaotian");
end

function liaotian:OnTick()
	
end

function liaotian:isShow()
	return Window.isShowing;
end

function liaotian:OnDispose()
	Window:Dispose();
end

function liaotian:OnHide()
	crtType = 0;
	Window:Hide();
end

function liaotian_FlushData(context)
	local isScrollBottom = contentList.scrollPane.isBottomMost;
	local type = crtType;
	if type == 0 then
		type = -1;
	elseif type == 1 then
		type = 1;
	elseif type == 2 then
		type = 0;
	end
	crtList = ChatSystem.MsgByType(type);
	contentList.numItems = crtList.Count;

	if isScrollBottom then
		contentList.scrollPane:ScrollBottom();
	end
end