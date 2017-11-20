require "FairyGUI"

paihangbang = fgui.window_class(WindowBase)
local Window;

local rankList;
local typeList;

local oneUrl = "ui://paihangbang/07";
local twoUrl = "ui://paihangbang/08";
local threeUrl = "ui://paihangbang/09";

local myRank;

local crtType;

function paihangbang:OnEntry()
	Window = paihangbang.New();
	Window:Show();
end

function paihangbang:GetWindow()
	return Window;
end

function paihangbang:OnInit()
	self.contentPane = UIPackage.CreateObject("paihangbang", "paihangbang_com").asCom;
	self:Center();
	self.closeButton = self.contentPane:GetChild("n1");

	myRank = self.contentPane:GetChild("n4").asTextField;

	rankList = self.contentPane:GetChild("n3").asList;
	rankList:SetVirtual();
	rankList.itemRenderer = paihangbang_RenderListItem;

	typeList = self.contentPane:GetChild("n2").asList;
	typeList.onClickItem:Add(paihangbang_OnSelectType);
	crtType = 0;
	typeList.selectedIndex = crtType;

	paihangbang_FlushData();
end

function paihangbang_OnSelectType()
	crtType = typeList.selectedIndex;
	UIManager.SetDirty("paihangbang");
end

function paihangbang_RenderListItem(index, obj)
	local rank = obj:GetChild("n3");
	local name = obj:GetChild("n1");
	local score = obj:GetChild("n2");
	local lv = obj:GetChild("n10");
	local flower = obj:GetChild("n8").asLoader;
	local icon = obj:GetChild("n7").asLoader;
	local duan = obj:GetChild("n11").asLoader;
	obj.onClick:Add(paihangbang_OnInfoClick);
	if crtType == 0 then
		obj.data =RankSystem._FirendRank[index].InstId;
		name.text = RankSystem._FirendRank[index].Name;
		lv.text = RankSystem._FirendRank[index].Level;
		score.text = RankSystem._FirendRank[index].TianTi;

		duan.url = "ui://paihangbang/xiao_duanwei" .. GamePlayer.RankLevel(RankSystem._FirendRank[index].TianTi);
		rank.text = index + 1 .. "";
		if index == 0 then
			flower.url = oneUrl;
		elseif index == 1 then
			flower.url = twoUrl;
		elseif index == 2 then
			flower.url = threeUrl;
		else
			flower.url = "";
		end
		local dData = DisplayData.GetData(RankSystem._FirendRank[index].DisplayID);
		if dData ~= nil then
			icon.url = "ui://" .. dData._HeadIcon;
		else
			icon.url = "";
		end
	elseif crtType == 1 then
		obj.data =RankSystem._AllRank[index].InstId;
		name.text = RankSystem._AllRank[index].Name;
		lv.text = RankSystem._AllRank[index].Level;
		score.text = RankSystem._AllRank[index].TianTi;
		duan.url = "ui://paihangbang/xiao_duanwei" .. GamePlayer.RankLevel(RankSystem._AllRank[index].TianTi);
		rank.text = index + 1 .. "";
		if index == 0 then
			flower.url = oneUrl;
		elseif index == 1 then
			flower.url = twoUrl;
		elseif index == 2 then
			flower.url = threeUrl;
		else
			flower.url = "";
		end
		local dData = DisplayData.GetData(RankSystem._AllRank[index].DisplayID);
		if dData ~= nil then
			icon.url = "ui://" .. dData._HeadIcon;
		else
			icon.url = "";
		end
	end
end

function paihangbang:OnUpdate()
	if UIManager.IsDirty("paihangbang") then
		paihangbang_FlushData();
		UIManager.ClearDirty("paihangbang");
	end
end

function paihangbang:OnTick()
	
end

function paihangbang:isShow()
	return Window.isShowing;
end

function paihangbang:OnDispose()
	Window:Dispose();
end

function paihangbang:OnHide()
	crtType = 0;
	typeList.selectedIndex = crtType;
	Window:Hide();
end

function paihangbang_FlushData()
	local count = 0;
	local myRankStr = "未上榜";
	if crtType == 0 then
		count = RankSystem._FirendRank.Count;
		if RankSystem._MyFirendRank ~= -1 then
			myRankStr = (RankSystem._MyFirendRank + 1) .. "";
		end
	elseif crtType == 1 then
		count = RankSystem._AllRank.Count;
		if RankSystem._MyAllRank ~= -1 then
			myRankStr = (RankSystem._MyAllRank + 1) .. "";
		end
	end

	myRank.text = myRankStr;
	rankList.numItems = count;
end

function paihangbang_OnInfoClick(context)
	Proxy4Lua.QueryPlayerInfo(context.sender.data);
end

