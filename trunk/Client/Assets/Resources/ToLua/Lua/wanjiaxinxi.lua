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
local tatle;
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
	tatle = self.contentPane:GetChild("n11");
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
	tatle.text = info.Name.."的信息";
	guildLab.text = "帮会：" .. info.ClanName;
	if info.TiatiRank <= 0 then
		rankLab.text = "未入榜";
	else
		rankLab.text = info.TiatiRank.."";
	end

	levelLab.text = info.Level.."";

	local eData = EntityData.GetData(info.UnitID);
	if eData ~= nil then
		local dData = DisplayData.GetData(eData._DisplayId);
		if dData ~= nil then
			icon.url = "ui://" .. dData._HeadIcon;
		else
			icon.url = "";
		end
	end
	duanLab.asLoader.url = "ui://wanjiaxinxi/xiao_duanwei" .. GamePlayer.RankLevel(info.TiatiVal);
	if info.UnitLIst ~= nil then
		cardGroupList.numItems = info.UnitLIst.Length;
	else
		cardGroupList.numItems = 0;
	end
end

function wanjiaxinxi_RenderListItem(indx, obj)
	local palyer = FriendSystem.friendInfo.UnitLIst[indx];
	local level = obj:GetChild("n6");
	level.text = palyer.Level .."";
	local edata = EntityData.GetData(palyer.UnitId);
	local ddata = DisplayData.GetData(edata._DisplayId);
	obj:GetChild("n10").visible = false;
	local fee = obj:GetChild("n7");
	fee.text = edata._Cost.."";
	obj:GetChild("n5").asLoader.url = "ui://" .. ddata._HeadIcon;
	obj.onClick:Add(wanjiaxinxi_OnCard);
	obj.data = palyer.UnitId;
end 

function wanjiaxinxi_OnCard(context)
	UIParamHolder.Set("showBoosId", context.sender.data);
	UIParamHolder.Set("showBoos", true);
	UIManager.Show("xiangxiziliao");
	
end