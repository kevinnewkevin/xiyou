require "FairyGUI"

squad = fgui.window_class(WindowBase)
local Window;

local member;

local helpBtn;
local helpTime;

local squadChatCom;
local sysCom = "ui://bangpai/xitong_com";
local otherCom = "ui://bangpai/duifang_com";
local selfCom = "ui://bangpai/wofang_com";

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

function squad:OnEntry()
	Window = squad.New();
	Window:Show();
end

function squad:GetWindow()
	return Window;
end

function squad:OnInit()
	self.contentPane = UIPackage.CreateObject("bangpai", "bangpai_com").asCom;
	self:Center();

	member = self.contentPane:GetChild("").asList;
	member:SetVirtual();
	member.itemRenderer = squad_RenderListItem;

	helpBtn = self.contentPane:GetChild("").asButton;
	helpBtn.onClick:Add(squad_OnHelp);
	helpTime = self.contentPane:GetChild("");

	-------------------------------chat----------------------------------------
	squadChatCom = self.contentPane:GetChild("n3").asCom;
	sendBtn = squadChatCom:GetChild("n8").asButton;
	emojiBtn = squadChatCom:GetChild("n10").asButton;
	yyBtn = squadChatCom:GetChild("n11").asButton;
	content = squadChatCom:GetChild("n12");
	sendBtn.onClick:Add(squadliaotian_OnSend);
	emojiBtn.onClick:Add(squadliaotian_OnEmoji);
	yyBtn.onTouchBegin:Add(squadliaotian_OnYYBegin);
	yyBtn.onTouchEnd:Add(squadliaotian_OnYYEnd);
	yyAnim = squadChatCom:GetChild("n15").asCom;
	yyAnim.visible = false;

	local gestureMoveUp = Proxy4Lua.SwipeGesture(squadChatCom);
	gestureMoveUp.onMove:Add(squadliaotian_OnSwipeMoveEnd);

	typeList = squadChatCom:GetChild("n6").asList;
	typeList.onClickItem:Add(squadliaotian_OnTypeSelect);

	emojiCom = squadChatCom:GetChild("n14").asCom;
	emojiCom.fairyBatching = true;
	emojiCom:GetChild("n1").asList.onClickItem:Add(squadliaotian_OnEmojiItem);
	emojiCom:RemoveFromParent();

	contentList = squadChatCom:GetChild("n13").asList;
	contentList:SetVirtual();
	contentList.itemProvider = squadliaotian_GetListItemResource;
	contentList.itemRenderer = squadliaotian_OnRenderListItem;

	crtType = 0;
	typeList.selectedIndex = crtType;
	-------------------------------endchat-------------------------------------

	squad_FlushData();
end

function squad_RenderListItem(index, obj)
	
end

function squad_OnHelp()
	
end

function squad:OnUpdate()
	if UIManager.IsDirty("squad") then
		squad_FlushData();
		UIManager.ClearDirty("squad");
	end
end

function squad:OnTick()
	
end

function squad:isShow()
	return Window.isShowing;
end

function squad:OnDispose()
	Window:Dispose();
end

function squad:OnHide()
	yyAnim.visible = false;
	crtType = 0;
	typeList.selectedIndex = crtType;
	Window:Hide();
end

function squad_FlushData()

	local isScrollBottom = contentList.scrollPane.isBottomMost;
	local type = squadliaotian_GetChatType(crtType);
	crtList = ChatSystem.MsgByType(type);
	contentList.numItems = crtList.Count;

	if isScrollBottom then
		contentList.scrollPane:ScrollBottom();
	end
end


function squadliaotian_OnEmoji(context)
	GRoot.inst:ShowPopup(emojiCom, context.sender, false);
end

function squadliaotian_OnSwipeMoveEnd(context)
	if context.sender.delta.y > -100 then
		return;
	end
	if yyAnim.visible == true then
		yyAnim.visible = false;
		YYSystem.StopRecord(true, 0);
		Proxy4Lua.PopMsg("语音发送取消");
	end
end

function squadliaotian_GetListItemResource(index)
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

function squadliaotian_OnEmojiItem(context)
	content:ReplaceSelection("[:" .. context.data.gameObjectName .. "]");
end

function squadliaotian_OnRenderListItem(index, obj)
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
				yybtn.onClick:Add(squadliaotian_OnPlayRecord);
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

function squadliaotian_OnYYBegin()
	yyAnim.visible = true;
	YYSystem.StartRecord();
end

function squadliaotian_OnYYEnd()
	yyAnim.visible = false;
	YYSystem.StopRecord(false, liaotian_GetChatType(crtType));
end

function squadliaotian_OnPlayRecord(context)
	local record = ChatSystem.GetRecord(context.sender.data);
	if record == nil then
		return;
	end
	YYSystem.PlayRecord(record.AudioPath, record.AudioUrl);
	ChatSystem.SetRecord(record.AudioId);
end

function squadliaotian_OnSend()
	if content.text == "" then
		return;
	end

	local chat = COM_Chat.New();
	local chatType = squadliaotian_GetChatType(crtType);
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

function squadliaotian_OnTypeSelect()
	crtType = typeList.selectedIndex;
	UIManager.SetDirty("squad");
end

function squadliaotian_GetChatType(uitype)
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