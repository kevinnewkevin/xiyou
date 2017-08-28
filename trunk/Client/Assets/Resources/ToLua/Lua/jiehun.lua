require "FairyGUI"

jiehun = fgui.window_class(WindowBase)
local Window;

local cardGroupUrl = "ui://jiehun/guanka_Button";
local cardGroupList;
local guankaId = 0;
local stamaPoint;
function jiehun:OnEntry()      
	Window = jiehun.New();
	Window:Show();
end

function jiehun:OnInit()
	self.contentPane = UIPackage.CreateObject("jiehun", "jiehun_com").asCom;
	self:Center();
	self.modal = true;
	
	self.closeButton = self.contentPane:GetChild("n5").asButton;
	
		cardGroupList = self.contentPane:GetChild("n1").asList;
	cardGroupList:SetVirtual();
	cardGroupList.itemRenderer = jiehun_RenderListItem;

	stamaPoint = self.contentPane:GetChild("n12");
	
	jiehun_FlushData();
end

function jiehun:GetWindow()
	return Window;
end

function jiehun:OnUpdate()
	if UIManager.IsDirty("jiehun") then
		jiehun_FlushData();
		UIManager.ClearDirty("jiehun");
	end
end



function jiehun_RenderListItem(index, obj)

	obj.onClick:Add(jiehun_OnSelectGroup);
	local data = HeroStroyData.GetData(index+1);
	obj.data =index+1;
	 local name = obj:GetChild("n9");
	 name.text = data.Name_;
	 local desc = obj:GetChild("n14");
	 desc.text = data.Desc_;
	 local img = obj:GetChild("n5");
	img.asLoader.url = "ui://" .. data.Icon_;
	 local suo = obj:GetChild("n32");
	 suo.visible  = false;
end

function jiehun_OnSelectGroup(context)
	guankaId = context.sender.data;
	Proxy4Lua.RequestChapterData(guankaId );
	UIManager.Show("guanka");
	jiehun_FlushData();
end

function jiehun:GetGuankaId()
	return guankaId;
end

function jiehun:OnTick()
	
end

function jiehun:isShow()
	return Window.isShowing;
end

function jiehun:OnDispose()
	Window:Dispose();
end

function jiehun:OnHide()
	Window:Hide();
end

function jiehun_FlushData()
	cardGroupList.numItems = HeroStroyData.metaData.Count;

	stamaPoint.text = GamePlayer._Data.IProperties[2];
end