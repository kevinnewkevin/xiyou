require "FairyGUI"

xuanzepaizu = fgui.window_class(WindowBase)
local Window;

local allCardGroupList;
local cardGroupList;
local cardItemUrl = "ui://xuanzepaizu/touxiangkuang_Label";
local cardGroupUrl = "ui://xuanzepaizu/paizuanniu_Button";

local crtGroupIdx = 0;

function xuanzepaizu:OnEntry()
	Window = xuanzepaizu.New();
	Window:Show();
end

function xuanzepaizu:GetWindow()
	return Window;
end

function xuanzepaizu:OnInit()
	self.contentPane = UIPackage.CreateObject("xuanzepaizu", "xuanzepaizu_com").asCom;
	self:Center();

	local bg = self.contentPane:GetChild("n3");
	allCardGroupList = bg:GetChild("n5").asList;

	local cardGroup = self.contentPane:GetChild("n6");
	cardGroupList = cardGroup:GetChild("n27").asList;

	for i=1, 5 do
		local groupItem = allCardGroupList:AddItemFromPool(cardGroupUrl);
		groupItem.onClick:Add(xuanzepaizu_OnSelectGroup);
		groupItem.data = i - 1;
	end
	allCardGroupList.selectedIndex = crtGroupIdx;

	xuanzepaizu_FlushData();
end

function xuanzepaizu_OnSelectGroup(context)
	crtGroupIdx = context.sender.data;
	xuanzepaizu_FlushData();
end

function xuanzepaizu:OnUpdate()
	if UIManager.IsDirty("paiku") then
		xuanzepaizu_FlushData();
		UIManager.ClearDirty("paiku");
	end
end

function xuanzepaizu_FlushData()
	cardGroupList:RemoveChildrenToPool();
	local groupCards = GamePlayer.GetGroupCards(crtGroupIdx);
	if groupCards == nil then
		return;
	end
	for i=1, groupCards.Count do
		local itemBtn = cardGroupList:AddItemFromPool(cardItemUrl);
	end
end

function xuanzepaizu:OnTick()
	
end

function xuanzepaizu:isShow()
	return Window.isShowing;
end

function xuanzepaizu:OnDispose()
	Window:Dispose();
end

function xuanzepaizu:OnHide()
	Window:Hide();
end