require "FairyGUI"

wanjiaxinxi = fgui.window_class(WindowBase)
local Window;
local nameLab;
local levelLab;
local numLab;
local guildLab;
local duanLab;
local rankLab;
local icon;
local cardGroupList;


function wanjiaxinxi:OnEntry()
	Window = wanjiaxinxi.New();
	Window:Show();
end

function wanjiaxinxi:GetWindow()
	return Window;
end

function wanjiaxinxi:OnInit()
	self.contentPane = UIPackage.CreateObject("wanjiaxinxi", "wanjiaxinxi_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n2");
    nameLab = self.contentPane:GetChild("n5");
    levelLab =  self.contentPane:GetChild("n10");
    icon = self.contentPane:GetChild("n8");
    numLab = self.contentPane:GetChild("n3");
	guildLab = self.contentPane:GetChild("n6");
	duanLab = self.contentPane:GetChild("n4");
	rankLab = self.contentPane:GetChild("n14");
	cardGroupList = self.contentPane:GetChild("n13");

	cardGroupList.itemRenderer = wanjiaxinxi_RenderListItem;
	wanjiaxinxi_FlushData();
end



function wanjiaxinxi:OnUpdate()
	if UIManager.IsDirty("wanjiaxinxi") then
		wanjiaxinxi_FlushData();
		UIManager.ClearDirty("wanjiaxinxi");
	end
end

function wanjiaxinxi:OnTick()
	
end

function wanjiaxinxi:isShow()
	return Window.isShowing;
end

function wanjiaxinxi:OnDispose()
	Window:Dispose();
end

function wanjiaxinxi:OnHide()
	Window:Hide();
end

function wanjiaxinxi_FlushData()

	local info = FriendSystem.friendInfo;
	if info == nil then
		return;
	end
	nameLab.text = info.Name;
	numLab.text = info.TiatiVal .."";    
	guildLab.text = info.ClanName;
	duanLab.text = info.TiatiRank.."";
	rankLab.text = info.TiatiRank.."";
	local displayData = DisplayData.GetData(info.DisplayID);
	icon.asLoader.url = "ui://" .. displayData._HeadIcon;
	cardGroupList.numItems = info.UnitLIst.length;
end

function wanjiaxinxi_RenderListItem(indx, obj)
	local palyer = FriendSystem.friendInfo.UnitLIst[indx];
	local level = obj:GetChild("n6");
	level.text = palyer.Level .."";
	local edata = EntityData.GetData(unit.UnitId);
	local ddata = DisplayData.GetData(edata._DisplayId);
	obj:GetChild("n5").asLoader.url = "ui://" .. ddata._HeadIcon;
end