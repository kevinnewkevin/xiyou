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

	local shopBtn = self.contentPane:GetChild("n30");
	shopBtn.onClick:Add(zhujiemian_OnShop);

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
	listGroup:SetVirtualAndLoop();
	listGroup.itemRenderer = zhujiemain_RenderListItem;
	listGroup.scrollPane.onScroll:Add(DoSpecialEffect);

	rankList =  self.contentPane:GetChild("n42").asList;
	rankList:SetVirtual();
	rankList.itemRenderer = zhujiemian_RankRenderListItem;
	rankList.onClick:Add(zhujiemian_OnRank);

	local rankBtn = self.contentPane:GetChild("n37").asButton;
	rankBtn.onClick:Add(zhujiemian_OnRank);

	minChatList = self.contentPane:GetChild("n55").asList;
	minChatList.onClick:Add(zhujiemian_OnMinChat);

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
	listGroup.numItems = 3;
	DoSpecialEffect();
end

function zhujiemian_FlushChatData()
	local list = ChatSystem.LastestMsgByType(-1, 3);
	minChatList:RemoveChildrenToPool();
	for i=0, list.Count - 1 do
		if list[i].Content ~= nil and list[i].PlayerName ~= nil then
			local content = minChatList:AddItemFromPool(minChatItem);
			local lbl = content:GetChild("n0");
			local yyCom = content:GetChild("n3");
			local frontPlus = "";
			if list[i].Type == 0 then
				frontPlus = "系统:";
			else
				frontPlus = Proxy4Lua.ChangeColor(list[i].PlayerName, "blue") .. ":";
			end
--			if list[i].AudioId ~= 0 then
--				lbl.visible = false;
--				yyCom.visible = true;
--				yyCom.onClick:Add(zhujiemian_liaotian_OnPlayRecord);
--				yyCom.data = list[i].AudioId;
--				yybtn:GetChild("n3").visible = list[i].Audio == nil;
--			else
				yyCom.visible = false;
				lbl.visible = true;
				lbl.text = frontPlus .. EmojiParser.inst:Parse(list[i].Content);
--			end
		end
	end
end

function zhujiemian_liaotian_OnPlayRecord(context)
	local record = ChatSystem.GetRecord(context.sender.data);
	if record == nil then
		Proxy4Lua.PlayAudio(context.sender.data);
	else
		YYSystem.PlayRecord(record);
	end
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
	obj:SetPivot(0.5, 0.5);
	local img = obj:GetChild("n5");
	local effect = obj:GetChild("n6");
	local effRes = "" ;
	img:SetScale(0.5,0.5);
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
	local dData = DisplayData.GetData(RankSystem._FirendRank[index].DisplayID);
	if dData ~= nil then
		icon.url = "ui://" .. dData._HeadIcon;
	else
		icon.url = "";
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