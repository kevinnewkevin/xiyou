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

local emojiCom;
local yyAnim;

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
	emojiBtn.onClick:Add(liaotian_OnEmoji);
	yyBtn.onTouchBegin:Add(liaotian_OnYYBegin);
	yyBtn.onTouchEnd:Add(liaotian_OnYYEnd);
	yyAnim = self.contentPane:GetChild("n15").asCom;
	yyAnim.visible = false;

	local gestureMoveUp = Proxy4Lua.SwipeGesture(self.contentPane);
	gestureMoveUp.onMove:Add(liaotian_OnSwipeMoveEnd);

	typeList = self.contentPane:GetChild("n6").asList;
	typeList.onClickItem:Add(liaotian_OnTypeSelect);

	emojiCom = self.contentPane:GetChild("n14").asCom;
	emojiCom.fairyBatching = true;
	emojiCom:GetChild("n1").asList.onClickItem:Add(liaotian_OnEmojiItem);
	emojiCom:RemoveFromParent();

	contentList = self.contentPane:GetChild("n13").asList;
	contentList:SetVirtual();
	contentList.itemProvider = liaotian_GetListItemResource;
	contentList.itemRenderer = liaotian_OnRenderListItem;

	crtType = 0;
	typeList.selectedIndex = crtType;
	liaotian_FlushData();
end

function liaotian_OnEmoji(context)
	GRoot.inst:ShowPopup(emojiCom, context.sender, false);
end

function liaotian_OnSwipeMoveEnd(context)
	if context.sender.delta.y > -100 then
		return;
	end
	if yyAnim.visible == true then
		yyAnim.visible = false;
		YYSystem.StopRecord(true, 0);
		Proxy4Lua.PopMsg("语音发送取消");
	end
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

function liaotian_OnEmojiItem(context)
	content:ReplaceSelection("[:" .. context.data.gameObjectName .. "]");
end

function liaotian_OnRenderListItem(index, obj)
	if crtList == nil then
		return;
	end

	for i=0, crtList.Count - 1 do
		if crtList[index].Type == 0 then --系统
			local content = obj:GetChild("n4").asTextField;
			content.text = Proxy4Lua.ChangeColor("系统:", "yellow") .. EmojiParser.inst:Parse(crtList[index].Content);
		else
			local yybtn = obj:GetChild("n8").asCom;
			local yybg = obj:GetChild("n9");
			local icon = obj:GetChild("n1").asLoader;
			local name = obj:GetChild("n5").asTextField;
			local content = obj:GetChild("n7");
			local contentBg = obj:GetChild("n6");
			local lv = obj:GetChild("n3");

			if crtList[index].AudioUrl ~= nil then
				content.visible = false;
				contentBg.visible = false;
				yybtn.visible = true;
				yybg.visible = true;
				yybtn.onClick:Add(liaotian_OnPlayRecord);
				yybtn.data = crtList[index].AudioId;
				yybtn:GetChild("n3").visible = not crtList[index].AudioOld;
				yybtn:GetChild("n2").text = crtList[index].AudioLen .. "\"";
			else
				content.visible = true;
				contentBg.visible = true;
				yybtn.visible = false;
				yybg.visible = false;
				content.width = content.initWidth;
				content.text = EmojiParser.inst:Parse(crtList[index].Content);
				content.width = content.textWidth;
			end
			icon.url = "ui://" .. crtList[index].HeadIcon;
			name.text = Proxy4Lua.ChangeColor(crtList[index].PlayerName, "blue");
			lv.text = crtList[index].Level;
		end
	end
end

function liaotian_OnYYBegin()
	yyAnim.visible = true;
	YYSystem.StartRecord();
end

function liaotian_OnYYEnd()
	yyAnim.visible = false;
	YYSystem.StopRecord(false, liaotian_GetChatType(crtType));
end

function liaotian_OnPlayRecord(context)
	local record = ChatSystem.GetRecord(context.sender.data);
	if record == nil then
		return;
	end
	YYSystem.PlayRecord(record.AudioPath, record.AudioUrl);
	ChatSystem.SetRecord(record.AudioId);
end

function liaotian_OnSend()
	if content.text == "" then
		return;
	end

	local chat = COM_Chat.New();
	local chatType = liaotian_GetChatType(crtType);
	if chatType == -1 then
		chatType = 1;
	end
	chat.Type = chatType;
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
	yyAnim.visible = false;
	crtType = 0;
	typeList.selectedIndex = crtType;
	Window:Hide();
end

function liaotian_FlushData(context)
	local isScrollBottom = contentList.scrollPane.isBottomMost;
	local type = liaotian_GetChatType(crtType);
	crtList = ChatSystem.MsgByType(type);
	contentList.numItems = crtList.Count;

	if isScrollBottom then
		contentList.scrollPane:ScrollBottom();
	end

	if type == 0 or type == 4 then
		sendBtn.visible = false;
		emojiBtn.visible = false;
		yyBtn.visible = false;
		content.visible = false;
	else
		sendBtn.visible = true;
		emojiBtn.visible = true;
		yyBtn.visible = true;
		content.visible = true;
	end
end

function liaotian_GetChatType(uitype)
	if uitype == 0 then
		return -1;--全部
	elseif uitype == 1 then
		return 1;--世界
	elseif uitype == 2 then
		return 0;--系统
	elseif uitype == 3 then
		return 3;--帮派
	elseif uitype == 4 then
		return 4;--帮派求助
	end
end