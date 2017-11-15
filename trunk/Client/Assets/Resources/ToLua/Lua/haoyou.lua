require "FairyGUI"

haoyou = fgui.window_class(WindowBase)
local Window;
local sysCom = "ui://haoyou/xitong_com";
local otherCom = "ui://haoyou/duifang_com";
local selfCom = "ui://haoyou/wofang_com";
local friendPanel;
local findFriendPanel;
local friendList;
local friendBtns; 
local funInfoBtn;
local funAddBtn;
local funBlackBtn;
local chatPanel;
local contentList;
local sendBtn;
local content;
local findFriendList;
local applyFriendList;
local findNameLab;
local findNameBtn;
local facePanel;
local emojiBtn;
local emojiCom;
local changeBtn;

local yyAnim;
local yyBtn;

local crtTab = 0;
local fCrtTab = 0;
local friendInstId;
local crtList;
local addFriendRad;
local chatFriendRad;
function haoyou:OnEntry()
	Window = haoyou.New();
	Window:Show();
end

function haoyou:GetWindow()
	return Window;
end

function haoyou:OnInit()
	Define.LaunchUIBundle("liaotian");
	self.contentPane = UIPackage.CreateObject("haoyou", "haoyou_com").asCom;
	self:Center();
	self.modal = true;
    self.closeButton = self.contentPane:GetChild("n7").asButton;
    yyAnim = self.contentPane:GetChild("n15");
    yyAnim.visible = false;
    local feeList = self.contentPane:GetChild("n3").asList;
	local feeMax = feeList.numItems;
	local feeItem;
	for i=1, feeMax do
		feeItem = feeList:GetChildAt(i-1);
		feeItem.data = i-1;
		feeItem.onClick:Add(haoyu_OnMainBtnsClick);
	end
	feeList.selectedIndex = 0;
    friendPanel = self.contentPane:GetChild("n2");
    chatPanel = friendPanel:GetChild("n3");
    chatPanel.visible = false;
    contentList = chatPanel:GetChild("n3").asList;
	--contentList:SetVirtual();
	contentList.itemProvider = haoyou_GetListItemResource;
	contentList.itemRenderer = haoyou_OnRenderListItem;
	local gestureMoveUp = Proxy4Lua.SwipeGesture(self.contentPane);
	gestureMoveUp.onMove:Add(liaotian_OnSwipeMoveEnd);

    sendBtn = chatPanel:GetChild("n6");
    sendBtn.onClick:Add(haoyu_OnSendClick);
    content = chatPanel:GetChild("n11");
   	emojiBtn = chatPanel:GetChild("n8");
	emojiBtn.onClick:Add(haoyou_OnEmoji);
	yyBtn = chatPanel:GetChild("n7");
	yyBtn.visible  = false;
	yyBtn.onTouchBegin:Add(liaotian_OnYYBegin);
	yyBtn.onTouchEnd:Add(liaotian_OnYYEnd);
    local fbBtns = friendPanel:GetChild("n2").asList;
	local fbMax = fbBtns.numItems;
	local fbItem;
	for i=1, fbMax do
		fbItem = fbBtns:GetChildAt(i-1);
		fbItem.data = i-1;
		fbItem.onClick:Add(haoyu_OnFBBtnsClick);
	end
	fbBtns.selectedIndex = 0;

	emojiCom = self.contentPane:GetChild("n13").asCom;
	emojiCom.fairyBatching = true;
	emojiCom:GetChild("n1").asList.onClickItem:Add(haoyou_OnEmojiItem);
	emojiCom:RemoveFromParent();


	findFriendPanel = self.contentPane:GetChild("n12");
	findFriendList = findFriendPanel:GetChild("n7");
	applyFriendList = findFriendPanel:GetChild("n11");
	findNameLab = findFriendPanel:GetChild("n5");
	findNameBtn = findFriendPanel:GetChild("n2");
	changeBtn = findFriendPanel:GetChild("n12");
	changeBtn.onClick:Add(haoyu_OnChangeBtnClick);
	findNameBtn.onClick:Add(haoyu_OnFindBtnClick);
	friendBtns = self.contentPane:GetChild("n14");
	friendBtns.visible = false;--:RemoveFromParent();
	findFriendPanel.visible = false;
	friendList = friendPanel:GetChild("n4");
	friendList.itemRenderer = haoyu_RenderListItem;
	findFriendList.itemRenderer = haoyu_RenderFindListItem;
	applyFriendList.itemRenderer = haoyu_RenderApplyListItem;
	funInfoBtn = friendBtns:GetChild("n5");
	funBlackBtn = friendBtns:GetChild("n8");
	funInfoBtn.onClick:Add(haoyu_OnFunInfoClick);
	funBlackBtn.onClick:Add(haoyu_OnFunBlackClick);
	friendBtns.onClick:Add(haoyu_OnfunBtnsClick);
	friendInstId = 0;
	addFriendRad = self.contentPane:GetChild("n17");
	chatFriendRad = self.contentPane:GetChild("n16");
	haoyou_FlushData();
end

function haoyou:OnUpdate()
	if UIManager.IsDirty("haoyou") then
		haoyou_FlushData();
		UIManager.ClearDirty("haoyou");
	end
end

function haoyou:OnTick()
	
end

function haoyou:isShow()
	return Window.isShowing;
end

function haoyou:OnDispose()
	Window:Dispose();
end

function haoyou:OnHide()
	yyAnim.visible = false;
	Window:Hide();
end

function haoyou_FlushData()

	addFriendRad.visible = FriendSystem.isApplyFriend;
	local num = FriendSystem.GetNewCahtListNum();	
	if num > 0 then 
  		chatFriendRad.visible = true;
  	else
  		chatFriendRad.visible = false;
  	end

	if findFriendPanel.visible == true then
		if FriendSystem.findFriend ~= nil then
			findFriendList.numItems = 1;
		elseif FriendSystem.randomFriends ~= nil then
			findFriendList.numItems = FriendSystem.randomFriends.Length;
		end
		applyFriendList.numItems = FriendSystem.GetApplyNum();
		FriendSystem.isApplyFriend = false;
		addFriendRad.visible = FriendSystem.isApplyFriend;
		return;
	end
	if fCrtTab == 0 then 
   		friendList.numItems = FriendSystem.GetFriendNum();
   	elseif fCrtTab == 1 then 
   		friendList.numItems = FriendSystem.GetBalckNum();
   	elseif fCrtTab == 2 then
   		friendList.numItems = FriendSystem.GetLatelyListNum();
  	end

  	haoyu_UpdataChat();
end

function haoyu_UpdataChat()
		--crtList =FriendSystem.GetFriendChat(friendInstId);
		local friend = FriendSystem.GetFriend(friendInstId);
		if friend ~= nil then
			local name = friend.Name;
			crtList =FriendSystem.GetFriendChatStr(name);
			contentList.numItems = crtList.Count;	
			FriendSystem.DelNewCahtList(name);
			local num = FriendSystem.GetNewCahtListNum();
		  	if num > 0 then 
		  		chatFriendRad.visible = true;
		  	else
		  		chatFriendRad.visible = false;
		  	end
		end
end

function haoyu_RenderApplyListItem(indx, obj)
	local palyer = FriendSystem.applyFriendList[indx];
	--local onLine= obj:GetChild("n8");
	--onLine.visible =false;
	local nameLab = obj:GetChild("n7");
	local levelLab = obj:GetChild("n6");
	nameLab.text = palyer.Name;
	levelLab.text = palyer.Level .. "";
	local icon = obj:GetChild("n3");
	local displayData = DisplayData.GetData(palyer.DisplayID);
	icon.asLoader.url = "ui://" .. displayData._HeadIcon;
	local addBtn = obj:GetChild("n9");
	addBtn.data = palyer.Name;
	addBtn.onClick:Set(haoyou_OnAddApplyFriendClick);
	local delBtn = obj:GetChild("n10");
	delBtn.data = palyer.InstId;
	delBtn.onClick:Set(haoyou_OnDelApplyFriendClick);
end

function haoyu_RenderFindListItem(indx, obj)

	local palyer;
	if FriendSystem.findFriend ~= nil then
		palyer	= FriendSystem.findFriend;
	else
		palyer	= FriendSystem.randomFriends[indx];
	end

	local nameLab = obj:GetChild("n5");
	local levelLab = obj:GetChild("n4");
	nameLab.text = palyer.Name;
	levelLab.text = palyer.Level .. "";
	local addBtn = obj:GetChild("n9");
	local delBtn = obj:GetChild("n8");
	local funBtn = obj:GetChild("n7");
	local icon = obj:GetChild("n2");
	local rad = obj:GetChild("n3");
	rad.visible = false;
	local displayData = DisplayData.GetData(palyer.DisplayID);
	icon.asLoader.url = "ui://" .. displayData._HeadIcon;
	addBtn.visible = true;
	delBtn.visible = false;
	funBtn.visible = false;
	addBtn.data = palyer.Name;
	addBtn.onClick:Set(haoyou_OnAddFriendClick);

end

function haoyu_RenderListItem(indx, obj)

	local palyer;
	if fCrtTab == 0 then
		palyer =  FriendSystem.friendList[indx];
	elseif fCrtTab == 1 then 
	 	palyer =  FriendSystem.blackList[indx];
	 elseif fCrtTab == 2 then 
	 	palyer =  FriendSystem.GetFriend(FriendSystem.latelyList[indx]);
	end
	if palyer == nil then
		return;
	end
	local panel = obj:GetChild("n1");
	local onLine= panel:GetChild("n6");
	onLine.visible =false;
	local nameLab = panel:GetChild("n5");
	local levelLab = panel:GetChild("n4");
	local icon = panel:GetChild("n2");
	local displayData = DisplayData.GetData(palyer.DisplayID);
	icon.asLoader.url = "ui://" .. displayData._HeadIcon;
	nameLab.text = palyer.Name;
	local newChat = FriendSystem.IsNewCaht(palyer.Name);
	local rad = obj:GetChild("n3");
	if newChat then
		rad.visible = true;
	else
		rad.visible = false;
	end
	levelLab.text = palyer.Level .. "";
	local addBtn = panel:GetChild("n9");
	local delBtn = panel:GetChild("n8");
	local funBtn = panel:GetChild("n7");
	addBtn.visible = false;
	delBtn.visible = false;
	funBtn.visible = false;
	if fCrtTab == 0 then
		funBtn.visible = true;
		funBtn.data = palyer.InstId;
		obj.data = palyer.InstId;
		funBtn.onClick:Add(haoyou_OnFunClick);
		obj.onClick:Add(haoyou_SelectFriendClick);
	elseif fCrtTab == 1 then 
	 	delBtn.visible = true;
	 	delBtn.data = palyer.InstId;
		obj.data = palyer.InstId;
		obj.onClick:Remove(haoyou_SelectFriendClick);
		delBtn.onClick:Add(haoyou_OnFunRemoveClick);
 	elseif fCrtTab == 2 then 
 		funBtn.visible = true;
		funBtn.data = palyer.InstId;
		obj.data = palyer.InstId;
		funBtn.onClick:Add(haoyou_OnFunClick);
		obj.onClick:Add(haoyou_SelectFriendClick);
	end
end

function haoyu_OnMainBtnsClick(context)
	crtTab = context.sender.data;
	if crtTab == 0 then
		friendPanel.visible = true;
		findFriendPanel.visible = false;
	else
		friendPanel.visible = false;
		findFriendPanel.visible = true;
	end
	haoyou_FlushData();
end

function haoyu_OnFBBtnsClick(context)
	fCrtTab = context.sender.data;
	chatPanel.visible = false;
	haoyou_FlushData();
end

function haoyou_OnFunClick(context)
	--GRoot.inst:ShowPopup(friendBtns, context.sender, false);
	friendBtns.visible = true;
end

function haoyou_OnFunRemoveClick(context)
	local MessageBox = UIManager.ShowMessageBox();
	MessageBox:SetData("提示", "将从黑名单中删除？", false, haoyou_OnDel);
end
function haoyou_OnDel(context)
	UIManager.HideMessageBox();
	Proxy4Lua.DeleteEnemy(friendInstId);
end


function haoyou_SelectFriendClick(context)
	chatPanel.visible = true;
	friendInstId = context.sender.data;
	local friend = FriendSystem.GetFriend(friendInstId);
	if friend ~= nil then
		local name = friend.Name;
		FriendSystem.DelNewCahtList(name);
	end
	local cnt = friendList.numChildren;

	for  i = 1, cnt do
		local obj = friendList:GetChildAt(i-1);
		local id = obj.data;
		local frd = FriendSystem.GetFriend(id);
		local newChat = FriendSystem.IsNewCaht(frd.Name);
		local rad = obj:GetChild("n3");
		if newChat then
			rad.visible = true;
		else
			rad.visible = false;
		end
	end
	haoyu_UpdataChat();
end

function haoyu_OnFunInfoClick(context)
	--friendBtns:RemoveFromParent();
	friendBtns.visible = false;
	Proxy4Lua.QueryPlayerInfo(friendInstId);
end

function haoyu_OnFunAddClick(context)
	--friendBtns:RemoveFromParent();
	friendBtns.visible = false;
end

function haoyu_OnFunBlackClick(context)
	--friendBtns:RemoveFromParent();
	friendBtns.visible = false;
	Proxy4Lua.AddEnemy(friendInstId);
end

function haoyu_OnSendClick(context)
	if content.text == "" then
		return;
	end
	local chat = COM_Chat.New();
	chat.Type = 3;
	chat.PlayerInstId = friendInstId;
	chat.Content = content.text;
	chat.PlayerName = GamePlayer._Name;
	chat.HeadIcon = GamePlayer.GetMyDisplayData()._HeadIcon;
	chat.Level = GamePlayer._Data.IProperties[9];
	Proxy4Lua.SendChat(chat);
	content.text = "";

	chat.PlayerInstId =  GamePlayer._InstID;
	FriendSystem.chatFriend (friendInstId,chat);
	FriendSystem.AddLatelyFriend (friendInstId);
	FriendSystem.chatFriendStr (FriendSystem.GetFriend(friendInstId).Name,chat);
	haoyou_FlushData();
end

function haoyu_OnFindBtnClick(context)
	if findNameLab.text == "" then
		return;
	end	
	Proxy4Lua.SerchFriendByName(findNameLab.text);
	findNameLab.text = "";
end

function haoyu_OnChangeBtnClick(context)
	FriendSystem.findFriend = nil;
	Proxy4Lua.SerchFriendRandom();
end

function haoyu_OnfunBtnsClick(context)
	friendBtns.visible = false;
end

function haoyou_OnAddFriendClick(context)
	local name = context.sender.data;
	Proxy4Lua.ApplicationFriend(name);
end

function haoyou_OnAddApplyFriendClick(context)
	local name = context.sender.data;
	Proxy4Lua.ProcessingFriend(name);
end
function haoyou_OnDelApplyFriendClick(context)
	local id = context.sender.data;
	FriendSystem.DelApplyFriend(id);
	haoyou_FlushData();
end



function haoyou_GetListItemResource(index)
	if crtList == nil then
		return;
	end

	if GamePlayer.IsMe(crtList[index].PlayerInstId) then
		return selfCom;
	else 
		print("otherCom otherCom otherCom otherCom otherCom");
		return otherCom;
	end
end

function haoyou_OnRenderListItem(index, obj)
	if crtList == nil then
		return;
	end

	for i=0, crtList.Count - 1 do
	
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


function haoyou_OnEmojiItem(context)
	content:ReplaceSelection("[:" .. context.data.gameObjectName .. "]");
end

function haoyou_OnEmoji(context)
	GRoot.inst:ShowPopup(emojiCom, context.sender, false);
end

function liaotian_OnYYBegin()
	yyAnim.visible = true;
	YYSystem.StartRecord();
end

function liaotian_OnYYEnd()
	yyAnim.visible = false;
	YYSystem.StopRecord(false);
end

function haoyou_OnPlayRecord(context)
	local record = ChatSystem.GetRecord(context.sender.data);
	if record == nil then
		return;
	end
	YYSystem.PlayRecord(record.AudioPath, record.AudioUrl);
	ChatSystem.SetRecord(record.AudioId);
end

function liaotian_OnSwipeMoveEnd(context)
	if context.sender.delta.y > -100 then
		return;
	end
	if yyAnim.visible == true then
		yyAnim.visible = false;
		YYSystem.StopRecord(true);
		Proxy4Lua.PopMsg("语音发送取消");
	end
end