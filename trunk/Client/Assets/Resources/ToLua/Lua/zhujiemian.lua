require "FairyGUI"

zhujiemian = fgui.window_class(WindowBase)
local Window;

local playerName;
local playerLevel;
local playerExp;
local gold;
local chopper;
local stamaPoint;
local expBar;
local headIcon;

local stateIcon;
local stateBar;
local listGroup;
local rankList;

local minChatList;
local minChatItem = "ui://zhujiemian/wenzi_com";

local maxChat;
local maxChatList;
local maxChatInput;
local maxChatSendBtn;

local friendRad;
local mailRad;

local chatBtn;
local guildBtn;

function zhujiemian:OnEntry()
	UIManager.RegIDirty("zhujiemian");
	Define.LaunchUIBundle("icon");
	Window = zhujiemian.New();
	Window:Show();
end

function zhujiemian:OnInit()
	self.contentPane = UIPackage.CreateObject("zhujiemian", "zhujiemian_com").asCom;
	self:Center();

	--stateIcon = self.contentPane:GetChild("n17").asButton;
	--stateIcon.onClick:Add(zhujiemian_OnFolder);

	local bottomGroup = self.contentPane:GetChild("n18");
	stateBar = bottomGroup:GetController("c1");
	local bottomBtnList = bottomGroup:GetChild("n15").asList;
	local skillBtn = bottomBtnList:GetChildAt(1);
	skillBtn.onClick:Add(zhujiemian_OnSkillBtn);
	local cardCargo = bottomBtnList:GetChildAt(2);
	cardCargo.onClick:Add(zhujiemian_OnCardCargo);

	listGroup =  self.contentPane:GetChild("n43");

	local bagBtn = bottomBtnList:GetChildAt(0);
	bagBtn.onClick:Add(zhujiemian_OnBagBtn);

	guildBtn = bottomBtnList:GetChildAt(3);
	guildBtn.onClick:Add(zhujiemian_OnGuild);

	local shopBtn = self.contentPane:GetChild("n30");
	shopBtn.onClick:Add(zhujiemian_OnShop);

	local friendBtn = self.contentPane:GetChild("n44");
	friendBtn.onClick:Add(zhujiemian_OnFriend);

	local tujianBtn = self.contentPane:GetChild("n45");
	tujianBtn.onClick:Add(zhujiemian_OnTuJian);

	local emailBtn = self.contentPane:GetChild("n7");
	emailBtn.onClick:Add(zhujiemian_OnEmailBtn);
	mailRad = emailBtn:GetChild("n3");
	mailRad.visible = false;
	friendRad = self.contentPane:GetChild("n56");
	friendRad.visible = false;

	local infoGroup = self.contentPane:GetChild("n15").asCom;
	playerName = infoGroup:GetChild("n9");
	playerExp = infoGroup:GetChild("n10");
	playerLevel = infoGroup:GetChild("n11");
	expBar = infoGroup:GetChild("n8").asProgress;
	headIcon = infoGroup:GetChild("n13").asLoader;
	local headBtn = infoGroup:GetChild("n2");
	headIcon.onClick:Add(zhujiemian_OnHead);

	local moneyGroup = self.contentPane:GetChild("n16").asCom;
	gold = moneyGroup:GetChild("n5");
	chopper = moneyGroup:GetChild("n7");
	stamaPoint = moneyGroup:GetChild("n12");
--	listGroup:SetVirtualAndLoop();
	listGroup.itemRenderer = zhujiemain_RenderListItem;
	--listGroup.scrollPane.onScroll:Add(DoSpecialEffect);

	rankList =  self.contentPane:GetChild("n42").asList;
	rankList:SetVirtual();
	rankList.itemRenderer = zhujiemian_RankRenderListItem;
	rankList.onClick:Add(zhujiemian_OnRank);

	local rankBtn = self.contentPane:GetChild("n37").asButton;
	rankBtn.onClick:Add(zhujiemian_OnRank);

	minChatList = self.contentPane:GetChild("n55").asList;
	minChatList.onClick:Add(zhujiemian_OnMinChat);

	chatBtn = self.contentPane:GetChild("n8").asButton;
	chatBtn.onClick:Add(zhujiemian_OnChatBtn);

	zhujiemian_FlushData();
	zhujiemian_FlushChatData();
	zhujiemian_FlushRankData();
	UIManager.AddExDirty("zhujiemian_liaotian");
	UIManager.AddExDirty("zhujiemian_paihang");
	--DoSpecialEffect();
end

function DoSpecialEffect()
	local midX = listGroup.scrollPane.posX + listGroup.viewWidth / 2;
	local cnt = listGroup.numChildren;
	for  i = 1, cnt do
		local obj = listGroup:GetChildAt(i-1);
		local dist = Mathf.Abs(midX - obj.x - obj.width / 2) ;
		local effect = obj:GetChild("n6");
		if dist > obj.width then
			obj:SetScale(1, 1);
			local colorObj = obj:GetChild("n5");
			colorObj.color = Color.gray;
			effect.visible =false;
			obj.onClick:Set(zhujiemian_OnClose);
		else
			local ss = 1 + (1 - dist / obj.width) * 1;
			obj:SetScale(ss, ss);
			local colorObj = obj:GetChild("n5");
			colorObj.color = Color.white;
			effect.visible = true;
			if obj.data == 0 then
				obj.onClick:Set(zhujiemian_OnQieCuoBtn);
			elseif obj.data == 1 then
				obj.onClick:Set(zhujiemian_OnTaskBtn);
			end

		end
	end
end

function zhujiemian:GetWindow()
	return Window;
end

function zhujiemian:OnUpdate()
	if UIManager.IsDirty("zhujiemian") then
		zhujiemian_FlushData();
		UIManager.ClearDirty("zhujiemian");
	end

	if UIManager.IsDirty("zhujiemian_liaotian") then
		zhujiemian_FlushChatData();
		UIManager.ClearDirty("zhujiemian_liaotian");
	end

	if UIManager.IsDirty("zhujiemian_paihang") then
		zhujiemian_FlushRankData();
		UIManager.ClearDirty("zhujiemian_paihang");
	end
end

function zhujiemian:OnTick()
	
end

function zhujiemian:isShow()
	return Window.isShowing;
end

function zhujiemian:OnDispose()
	Window:Dispose();
end

function zhujiemian:OnHide()
	Proxy4Lua.ClearToDeleteAsset("zhujiemian");
	Window:Hide();
end

function zhujiemian_FlushData()
	local displayData = GamePlayer.GetMyDisplayData();
	headIcon.url = "ui://" .. displayData._HeadIcon .. "_yuan";
	local needExp = ExpData.NeedExp(GamePlayer._Data.IProperties[9]);
	playerName.text = GamePlayer._Name;
	playerExp.text = GamePlayer._Data.IProperties[4] .. "/" .. needExp;
	playerLevel.text = GamePlayer._Data.IProperties[9];

	gold.text = GamePlayer._Data.IProperties[8];
	chopper.text = GamePlayer._Data.IProperties[6];
	stamaPoint.text = GamePlayer._Data.IProperties[10];

	expBar.value = GamePlayer._Data.IProperties[4] / needExp * 100;
	listGroup.numItems = 2;
	friendRad.visible = FriendSystem.isApplyFriend;

	local num = FriendSystem.GetNewCahtListNum();
  	if num > 0 then 
  		friendRad.visible = true;
  	else
  			if FriendSystem.isApplyFriend == false then   
  				friendRad.visible = false;
  			end
  	end

  	local mails = MailSystem.mailList;
  	mailRad.visible = false;
  	if mails.Count > 0 then 
		for i = 1 , mails.Count do
			if mails[i-1].IsRead == false or mails[i-1].Items ~= nil then
				mailRad.visible = true;
				break;
			end
		end
	end
	--DoSpecialEffect();
	guildBtn.enabled = GamePlayer._Data.IProperties[9] >= 1;
end

function zhujiemian_FlushChatData()
	local list = ChatSystem.LastestMsgByType(-1, 3);
	minChatList:RemoveChildrenToPool();
	for i=0, list.Count - 1 do
		local content = minChatList:AddItemFromPool(minChatItem);
		local lbl = content:GetChild("n0");
		local yyCom = content:GetChild("n3");
		local frontPlus = "";
		if list[i].Type == 0 then
			frontPlus = Proxy4Lua.ChangeColor("系统:", "yellow");
		else
			if list[i].PlayerName ~= nil then
				frontPlus = Proxy4Lua.ChangeColor(list[i].PlayerName, "blue") .. ":";
			end
		end
		lbl.width = lbl.initWidth;
		if list[i].AudioUrl ~= nil then
			yyCom.visible = true;
			yyCom.onClick:Add(zhujiemian_liaotian_OnPlayRecord);
			yyCom.data = list[i].AudioId;
			yyCom:GetChild("n3").visible = not list[i].AudioOld;
			yyCom:GetChild("n2").text = list[i].AudioLen .. "\"";
			lbl.text = frontPlus;
		else
			yyCom.visible = false;
			lbl.visible = true;
			if list[i].Content ~= nil then
				lbl.text = frontPlus .. EmojiParser.inst:Parse(list[i].Content);
			end
		end
		lbl.width = lbl.textWidth;
	end
	if minChatList.numItems > 0 then
		if minChatList.numItems >= 3 then
			minChatList:ScrollToView(2, false);
		else
			minChatList:ScrollToView(minChatList.numItems - 1, false);
		end
	end
end

function zhujiemian_liaotian_OnPlayRecord(context)
	local record = ChatSystem.GetRecord(context.sender.data);
	if record == nil then
		return;
	end
	YYSystem.PlayRecord(record.AudioPath, record.AudioUrl);
	ChatSystem.SetRecord(record.AudioId);
end

function zhujiemian_FlushRankData()
	rankList.numItems = RankSystem._FirendRank.Count;
end

function zhujiemian_OnMinChat()
	UIManager.Show("liaotian");
end

function zhujiemian_OnRank()
	UIManager.Show("paihangbang");
end

function zhujiemain_RenderListItem(index, obj)
	--obj:SetPivot(0.9, 0.9);
	local img = obj:GetChild("n5");
	local effect = obj:GetChild("n6");
	local effRes = "" ;
	img:SetScale(0.9,0.9);
	if index == 0 then
		img.url = "ui://zhujiemian/zjm_31";
		obj.onClick:Set(zhujiemian_OnQieCuoBtn);
		obj.data = 0;
		effRes = "Effect/sunwukong_ui";
		effect :SetNativeObject(Proxy4Lua.GetEffectAssetGameObject(effRes)); 
		Proxy4Lua.AddToDelete("zhujiemian",effRes);
	elseif index == 1 then
		img.url = "ui://zhujiemian/zjm_30";
		obj.onClick:Set(zhujiemian_OnTaskBtn);
		effRes = "Effect/xihailongwang_ui";
		effect :SetNativeObject(Proxy4Lua.GetEffectAssetGameObject(effRes));
		Proxy4Lua.AddToDelete("zhujiemian",effRes);
		obj.data = 1;
	elseif index == 2 then
		img.url = "ui://zhujiemian/zjm_28";
		obj.onClick:Set(zhujiemian_OnClose);
		obj.data = 2;
		effect :SetNativeObject(Proxy4Lua.GetEffectAssetGameObject(""));
	elseif index == 3 then
		img.url = "ui://zhujiemian/zjm_28";
		obj.onClick:Set(zhujiemian_OnClose);
		obj.data = 3;
		effect :SetNativeObject(Proxy4Lua.GetEffectAssetGameObject(""));
	end
	
end

function zhujiemian_RankRenderListItem(index, obj)
	local rank = obj:GetChild("n59");
	local lv = obj:GetChild("n58");
	local icon = obj:GetChild("n60").asLoader;

	lv.text = RankSystem._FirendRank[index].Level;
	rank.text = index + 1 .. "";
	local eData = EntityData.GetData(RankSystem._FirendRank[index].UnitID);
	if eData ~= nil then
		local dData = DisplayData.GetData(eData._DisplayId);
		if dData ~= nil then
			icon.url = "ui://" .. dData._HeadIcon;
		else
			icon.url = "";
		end
	end
end

function zhujiemian_OnClose()
end

function zhujiemian_OnFolder()
	stateIcon:GetController("icon").selectedIndex = (stateIcon:GetController("icon").selectedIndex + 1) % 2;
	stateBar.selectedIndex = stateIcon:GetController("icon").selectedIndex;
end

function zhujiemian_OnCardCargo()
	UIManager.Show("paiku");
end

function zhujiemian_OnTaskBtn()
	UIManager.Show("daguanka");
end

function zhujiemian_OnSkillBtn()
	UIManager.Show("jineng");
end

function zhujiemian_OnQieCuoBtn()
	UIManager.Show("qiecuo");
end

function zhujiemian_OnHead()
	UIManager.Show("renwuziliao");
end

function zhujiemian_OnBagBtn()
	UIManager.Show("bagui");
end

function zhujiemian_OnShop()
	UIManager.Show("cangbaoge");
end

function zhujiemian_OnFriend()
	UIManager.Show("haoyou");
end

function zhujiemian_OnChatBtn()
	UIManager.Show("liaotian");
end


function zhujiemian_OnTuJian()
	UIManager.Show("tujian");
end

function zhujiemian_OnGuild()
	if GamePlayer._iGuildId == 0 then
		Proxy4Lua.QueryGuildList();
		UIManager.Show("squadList");
	else
		Proxy4Lua.QueryGuildData();
		UIManager.Show("squad");
	end
end

function zhujiemian_OnEmailBtn()
	UIManager.Show("youxiang");
end

