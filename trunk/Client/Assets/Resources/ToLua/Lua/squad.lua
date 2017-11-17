require "FairyGUI"

squad = fgui.window_class(WindowBase)
local Window;

local member;

local helpBtn;
local helpTime;

local squadChatCom;
local sysCom = "ui://liaotian/xitong_com";
local otherCom = "ui://liaotian/duifang_com";
local selfCom = "ui://liaotian/wofang_com";
local assisCom = "ui://bangpai/zhiyuan_com";

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
local contentBg;

local settingBtn;
local quitBtn;

local willKickPlayer; --计划踢出的玩家id

function squad:OnEntry()
	Window = squad.New();
	Window:Show();
end

function squad:GetWindow()
	return Window;
end

function squad:OnInit()
	Define.LaunchUIBundle("liaotian");
	self.contentPane = UIPackage.CreateObject("bangpai", "bangpai_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n6");

	member = self.contentPane:GetChild("n8").asList;
	member:SetVirtual();
	member.itemRenderer = squad_RenderListItem;

	quitBtn = self.contentPane:GetChild("n10").asButton;
	quitBtn.onClick:Add(squad_OnQuit);

	settingBtn = self.contentPane:GetChild("n9").asButton;
	settingBtn.onClick:Add(squad_OnSetting);

	-------------------------------chat----------------------------------------
	squadChatCom = self.contentPane:GetChild("n3").asCom;
	--unique
	helpBtn = squadChatCom:GetChild("n17").asButton;
	helpBtn.onClick:Add(squad_OnHelp);
	helpTime = squadChatCom:GetChild("n18");
	--unique
	sendBtn = squadChatCom:GetChild("n8").asButton;
	emojiBtn = squadChatCom:GetChild("n10").asButton;
	yyBtn = squadChatCom:GetChild("n11").asButton;
	content = squadChatCom:GetChild("n12");
	contentBg = squadChatCom:GetChild("n7");
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
	if obj == nil then
		return;
	end

	local data = GuildSystem.guildMemberList[index];
	local name = obj:GetChild("n6").asTextField;
	local online = obj:GetChild("n9").asTextField;
	local pos = obj:GetChild("n8").asLoader;
	local score = obj:GetChild("n7").asLoader;
	local headCom = obj:GetChild("n5").asCom;
	local headIcon = headCom:GetChild("n5").asLoader;
	local lv = headCom:GetChild("n3").asTextField;
	local rkBtnBg = obj:GetChild("n12");
--	local renmingBtn = obj:GetChild("n11").asButton;
--	local kickBtn = obj:GetChild("n10").asButton;
--	renmingBtn.data = data.RoleId;
--	kickBtn.data = data.RoleId;
--	renmingBtn.onClick:Add(squad_OnRenming);
--	kickBtn.onClick:Add(squad_OnKick);
--
--	renmingBtn.visible = GuildSystem.MyJob() == 3 and data.RoleName ~= GamePlayer._Name;
--	kickBtn.visible = (GuildSystem.MyJob() == 2 or GuildSystem.MyJob() == 3) and data.RoleName ~= GamePlayer._Name;
--
--	rkBtnBg.visible = renmingBtn.visible or kickBtn.visible;
--	headIcon.onClick:Add(squad_OnGuildPlayer);
--	headIcon.data =  data.RoleId;

	local allPop = obj:GetChild("n14");
	allPop.visible = false;
	obj.data = data;
	obj.onClick:Add(squad_OnOperateList);
	name.text = data.RoleName;
	if data.IsOnline then
		online.text = "在线";
	else
		online.text = "离线";
	end
	pos.url = "ui://bangpai/cu_" .. data.Job;
	score.url = "ui://bangpai/xiao_duanwei" .. GamePlayer.RankLevel(data.TianTiVal);
	local eData = EntityData.GetData(data.UnitId);
	local dData;
	if eData ~= nil then
		dData = DisplayData.GetData(eData._DisplayId);
	end
	if dData ~= nil then
		headIcon.url = "ui://" .. dData._HeadIcon;
	else
		headIcon.url = "";
	end
	lv.text = data.Level;
end

function squad_OnOperateList(memberData)
	local allPop = memberData.sender:GetChild("n14");
	local popMenu = allPop:GetChild("n15").asList;
	if allPop.visible == true then
		allPop.visible = false;
		return;
	end
	local item = popMenu:GetChildAt(2).asButton;
	local count = 0;
	if memberData.sender.data.RoleName ~= GamePlayer._Name then
		item.onClick:Add(squad_OnDetail);
		item.data = memberData.sender.data;
		count = count + 1;
	end
	item.enabled = memberData.sender.data.RoleName ~= GamePlayer._Name;

	item = popMenu:GetChildAt(0);
	if GuildSystem.MyJob() == 3 and memberData.sender.data.RoleName ~= GamePlayer._Name then
		item.onClick:Add(squad_OnRenming);
		item.data = memberData.sender.data;
		count = count + 1;
	end
	item.enabled = GuildSystem.MyJob() == 3 and memberData.sender.data.RoleName ~= GamePlayer._Name;

	item = popMenu:GetChildAt(1);
	if (GuildSystem.MyJob() == 2 or GuildSystem.MyJob() == 3) and memberData.sender.data.RoleName ~= GamePlayer._Name and GuildSystem.MyJob() > memberData.sender.data.Job then
		item.onClick:Add(squad_OnKick);
		item.data = memberData.sender.data;
		count = count + 1;
	end
	item.enabled = (GuildSystem.MyJob() == 2 or GuildSystem.MyJob() == 3) and memberData.sender.data.RoleName ~= GamePlayer._Name and GuildSystem.MyJob() > memberData.sender.data.Job;

	allPop.visible = count ~= 0;
end

function squad_OnDetail(memberData)
	Proxy4Lua.QueryPlayerInfo(memberData.sender.data.RoleId);
end

function squad_OnRenming(memberData)
	if memberData.sender.data == nil then
		return;
	end

	UIParamHolder.Set("squadRenmingPlayer", memberData.sender.data.RoleId);
	UIManager.Show("squadRenming");
end

function squad_OnKick(memberData)
	if memberData.sender.data == nil then
		return;
	end

	willKickPlayer = memberData.sender.data.RoleId;
	local MessageBox = UIManager.ShowMessageBox();
	MessageBox:SetData("提示", "是否将该玩家踢出家族？", false, squad_OnConfirmKick);
end

function squad_OnSetting()
	UIManager.Show("squadSetting");
end

function squad_OnQuit()
	local MessageBox = UIManager.ShowMessageBox();
	MessageBox:SetData("提示", "是否退出家族？", false, squad_OnConfirmQuit);
end

function squad_OnConfirmQuit()
	Proxy4Lua.LeaveGuild();
	UIManager.HideMessageBox();
end

function squad_OnConfirmKick()
	Proxy4Lua.KickOut(willKickPlayer);
	UIManager.HideMessageBox();
end

function squad_OnHelp(context)
	UIManager.Show("squadHelp");
end

function squad:OnUpdate()
	if UIManager.IsDirty("squad") then
		squad_FlushData();
		UIManager.ClearDirty("squad");
	end
end

function squad:OnTick()
	local left = TimerManager.GetCountDownSecond("AssistantCoolDown");
	local str = TimerManager.GetCountDown("AssistantCoolDown");
	if left < 0 then
		left = 0;
		str = "";
		helpBtn.enabled = true;
	else
		helpBtn.enabled = false;
	end
	helpTime.text = str;
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
	member.numItems = GuildSystem.guildMemberList.Count;

	local isScrollBottom = contentList.scrollPane.isBottomMost;
	local type = squadliaotian_GetChatType(crtType);
	if type == 5 then
		crtList = ChatSystem._Assistant;
	else
		crtList = ChatSystem.MsgByType(type);
	end
	contentList.numItems = crtList.Count;

	if isScrollBottom then
		contentList.scrollPane:ScrollBottom();
	end

	if type == 0 or type == 5 then
		sendBtn.visible = false;
		emojiBtn.visible = false;
		yyBtn.visible = false;
		content.visible = false;
		contentBg.visible = false;
	else
		sendBtn.visible = true;
		emojiBtn.visible = true;
		yyBtn.visible = true;
		content.visible = true;
		contentBg.visible = true;
	end

	settingBtn.visible = GuildSystem.MyJob() == 2 or GuildSystem.MyJob() == 3;
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
	if squadliaotian_GetChatType(crtType) == 5 then
		return assisCom;
	else
		if GamePlayer.IsMe(crtList[index].PlayerInstId) then
			return selfCom;
		elseif crtList[index].Type == 0 then
			return sysCom;
		else
			return otherCom;
		end
	end
end

function squadliaotian_OnEmojiItem(context)
	content:ReplaceSelection("[:" .. context.data.gameObjectName .. "]");
end

function squadliaotian_OnRenderListItem(index, obj)
	if squadliaotian_GetChatType(crtType) == 5 then
		local name = obj:GetChild("n2");
		local status = obj:GetChild("n4");
		local statusBar = obj:GetChild("n5");
		local itemIconCom = obj:GetChild("n1");
		local itemIcon = itemIconCom:GetChild("n5");
		local itemIconBack = itemIconCom:GetChild("n11");
		local assBtn = obj:GetChild("n6");
		assBtn.data = crtList[index].Id;
		assBtn.onClick:Add(squad_OnAssistant);

		name.text = crtList[index].PlayerName;
		status.text = crtList[index].CrtCount .. "/" .. crtList[index].MaxCount;
		statusBar.value = crtList[index].CrtCount / crtList[index].MaxCount * 100;
		local iData = ItemData.GetData(crtList[index].ItemId);
		if iData ~= nil then
			itemIcon.url = "ui://" .. iData._Icon;
			itemIconBack.url = "ui://" .. iData._IconBack;
		else
			itemIcon.url = "";
			itemIconBack.url = "";
		end
		local meAssistanted = false;
		if crtList[index].IsAssistanted ~= nil then
			for i=0, crtList[index].IsAssistanted.Length - 1 do
				if GamePlayer.IsMe(crtList[index].IsAssistanted[i]) then
					meAssistanted = true;
					break;
				end
			end
		end

		assBtn.enabled = not meAssistanted and crtList[index].PlayerName ~= GamePlayer._Name and crtList[index].CrtCount < crtList[index].MaxCount;
	else
		if crtList[index].Type == 0 then --系统
			local content = obj:GetChild("n4").asTextField;
			content.text = Proxy4Lua.ChangeColor("[系统]", "yellow") .. EmojiParser.inst:Parse(crtList[index].Content);
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
			if crtList[index].Type == 1 then
				name.text = Proxy4Lua.ChangeColor("[世界]", "red") .. Proxy4Lua.ChangeColor(crtList[index].PlayerName, "blue");
			elseif crtList[index].Type == 4 then
				name.text = Proxy4Lua.ChangeColor("[家族]", "green") .. Proxy4Lua.ChangeColor(crtList[index].PlayerName, "blue");
			end
			lv.text = crtList[index].Level;
		end
	end
end

function squad_OnAssistant(context)
	Proxy4Lua.Assistant(context.sender.data);
--	local chat = COM_Chat.New();
--	chat.Type = 4;
--	chat.PlayerInstId = GamePlayer._InstID;
--	chat.PlayerName = GamePlayer._Name;
--	chat.HeadIcon = GamePlayer.GetMyDisplayData()._HeadIcon;
--	chat.Level = GamePlayer._Data.IProperties[9];
	local assData = ChatSystem.GetAss(context.sender.data);
	if assData ~= nil then
		local iData = ItemData.GetData(assData.ItemId);
		if iData ~= nil then
--			chat.Content = "我捐助了" .. assData.PlayerName .. "1个" .. iData._Name;
--			Proxy4Lua.SendChat(chat);
			Proxy4Lua.PopMsg("我捐助了" .. assData.PlayerName .. "1个" .. iData._Name);
		end
	end
end

function squadliaotian_OnYYBegin()
	yyAnim.visible = true;
	YYSystem.StartRecord();
end

function squadliaotian_OnYYEnd()
	yyAnim.visible = false;
	YYSystem.StopRecord(false, squadliaotian_GetChatType(crtType));
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

	if chatType == 4 and GamePlayer._Data.IProperties[9] < 3 then
		Proxy4Lua.PopMsg("3级开启帮派权限");
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
		return 4;--帮派
	elseif uitype == 4 then
		return 5;--帮派求助
	end
end