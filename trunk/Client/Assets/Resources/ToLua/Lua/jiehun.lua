require "FairyGUI"

jiehun = fgui.window_class(WindowBase)
local Window;

local cardGroupUrl = "ui://jiehun/guanka_Button";
local cardGroupList;
local guankaId = 0;
local stamaPoint;
local starLab;
local crtTab;
local rewardShow;
local getRewardBtn;
local rewardOkBtn;
function jiehun:OnEntry()
	Define.LaunchUIBundle("guankatupian");
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
	rewardShow = self.contentPane:GetChild("n14");
	rewardShow.visible= false;
	local feeList = self.contentPane:GetChild("n13").asList;
	getRewardBtn = rewardShow:GetChild("n13");
	getRewardBtn = rewardShow:GetChild("n16");
	local feeMax = feeList.numItems;
	local feeItem;
	for i=1, feeMax do
		feeItem = feeList:GetChildAt(i-1);
		feeItem.data = i;
		feeItem.onClick:Add(jiehun_OnFeeItemClick);
	end
	crtTab = 1;
	feeList.selectedIndex = crtTab-1;

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
	local comData;
	if crtTab == 1 then
		comData  = JieHunSystem.instance.ChapterEasyDataList[index];
	else
		comData  = JieHunSystem.instance.ChapterHardDataList[index];
	end

	obj.onClick:Add(jiehun_OnSelectGroup);
	local data = HeroStroyData.GetData(comData.ChapterId);
	if crtTab ~= data.Type_ then
		return;
	end
	obj.data =comData.ChapterId;
	local name = obj:GetChild("n9");
	name.text = data.Name_;
	local desc = obj:GetChild("n14");
	desc.text = data.Desc_;
	local img = obj:GetChild("n5");
	img.asLoader.url = "ui://" .. data.Icon_;
	local starLab = obj:GetChild("n12");
	local starLab0 = obj:GetChild("n20");
	local starLab1 = obj:GetChild("n24");
	local starLab2 = obj:GetChild("n28");
	local starBar = obj:GetChild("n15");
	local box0 = obj:GetChild("n18");
	local box1 = obj:GetChild("n22");
	local box2 = obj:GetChild("n26");
	local len  = comData.SmallChapters.Length;
	local starNum = 0;
	for i=1, len do
		if comData.SmallChapters[i-1].Star1 == true then
			starNum = starNum + 1;
		end
		if comData.SmallChapters[i-1].Star2 == true then
			starNum = starNum + 1;
		end
		if comData.SmallChapters[i-1].Star3 == true then 
			starNum = starNum +1;
		end
	end
	starBar.value = starNum/data.Star_[2]*100;
	starLab.text = starNum .. "/" .. data.Star_[2];
	starLab0.text = "" .. data.Star_[0];
	starLab1.text = "" .. data.Star_[1];
	starLab2.text = "" .. data.Star_[2];
	 local suo = obj:GetChild("n32");
	 suo.visible  = false;
	
end

function jiehun_OnSelectGroup(context)
	guankaId = context.sender.data;
	Proxy4Lua.RequestChapterData(guankaId );
	UIManager.Show("guanka");
	jiehun_FlushData();
end

function jiehun_OnFeeItemClick(context)
	crtTab = context.sender.data;
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
		if crtTab == 1 then
		cardGroupList.numItems =  JieHunSystem.instance.ChapterEasyDataList.Count;
	else
		cardGroupList.numItems =  JieHunSystem.instance.ChapterHardDataList.Count;
	end

	stamaPoint.text = GamePlayer._Data.IProperties[2];
end