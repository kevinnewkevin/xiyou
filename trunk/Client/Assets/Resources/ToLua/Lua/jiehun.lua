require "FairyGUI"

jiehun = fgui.window_class(WindowBase)
local Window;

local cardGroupUrl = "ui://jiehun/daoju_Button";
local cardGroupList;
local guankaId = 0;
local stamaPoint;
local starLab;
local crtTab;
local rewardShow;
local getRewardBtn;
local closeRewardBtn;
local rewardOkBtn;
local showRewardStar;
local showRewardId;
local rewardList;
local rewardNeedNum;
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
	closeRewardBtn = rewardShow:GetChild("n19");
	getRewardBtn = rewardShow:GetChild("n18");
	rewardOkBtn = rewardShow:GetChild("n17");
	rewardList = rewardShow:GetChild("n12");
	rewardNeedNum = rewardShow:GetChild("n15");
	getRewardBtn.onClick:Add(jiehun_OnGetRewardBtn);
	rewardOkBtn.onClick:Add(jiehun_OnRewardOkBtn);
	closeRewardBtn.onClick:Add(jiehun_OnRewardOkBtn);
	rewardShow.visible= false;
	local feeList = self.contentPane:GetChild("n13").asList;
	local feeMax = feeList.numItems;
	local feeItem;
	for i=1, feeMax do
		feeItem = feeList:GetChildAt(i-1);
		feeItem.data = i;
		feeItem.onClick:Add(jiehun_OnFeeItemClick);
	end
	crtTab = 1;
	feeList.selectedIndex = crtTab-1;
	local teamBtn = self.contentPane:GetChild("n7");
	teamBtn.onClick:Add(jiehun_OnTeamBtn);


	cardGroupList.scrollPane.onScroll:Add(DoSpecialEffect);

	DoSpecialEffect();

	jiehun_FlushData();
end



function DoSpecialEffect()
		local midX = cardGroupList.scrollPane.posX + cardGroupList.viewWidth / 2;
		local cnt = cardGroupList.numChildren;
		for  i = 1, cnt do
			local obj = cardGroupList:GetChildAt(i-1);
			local dist = Mathf.Abs(midX - obj.x - obj.width / 2);
			if dist > obj.width then
				obj:SetScale(1, 1);
			else
				local ss = 1 + (1 - dist / obj.width) * 0.24;
				obj:SetScale(ss, ss);
			end
		end
		--_mainView.GetChild("n3").text = "" + ((cardGroupList.GetFirstChildInView() + 1) % cardGroupList.numItems);
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
		comData  = JieHunSystem.instance.ChapterEasyDataList[0];
	else
		comData  = JieHunSystem.instance.ChapterHardDataList[0];
	end

	local objBack= obj:GetChild("n34");
	objBack.onClick:Add(jiehun_OnSelectGroup);
	local objIcon= obj:GetChild("n5");
	objIcon.onClick:Add(jiehun_OnSelectGroup);
	objIcon.data =comData.ChapterId;
	local data = HeroStroyData.GetData(comData.ChapterId);
	objBack.data =comData .ChapterId;
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
	
	local boxOpen0 = box0:GetChild("n2");
	local boxNoOpen0 = box0:GetChild("n1");
	local boxOpen1 = box1:GetChild("n2");
	local boxNoOpen1 = box1:GetChild("n1");
	local boxOpen2 = box2:GetChild("n2");
	local boxNoOpen2 = box2:GetChild("n1");

	box0.onClick:Add(jiehun_OnBox);
	box0.data =comData.ChapterId;
	box1.onClick:Add(jiehun_OnBox1);
	box1.data =comData.ChapterId;
	box2.onClick:Add(jiehun_OnBox2);
	box2.data =comData.ChapterId;

	boxOpen0.visible =false;
	boxOpen1.visible =false;
	boxOpen2.visible =false;
	boxNoOpen0.visible = true;
	boxNoOpen1.visible = true;
	boxNoOpen2.visible = true;
	if comData.StarReward ~= nil then
		local len = comData.StarReward.Length;
		for i=1, len do

			if comData.StarReward[i-1] == 10 then
				boxOpen0.visible =true;
				boxNoOpen0.visible =false;
			end
			if comData.StarReward[i-1] == 20 then
				boxOpen1.visible =true;
				boxNoOpen1.visible =false;
			end
			if comData.StarReward[i-1] == 30 then
				boxOpen2.visible =true;
				boxNoOpen2.visible =false;
			end
		end
	else
		boxOpen0.visible =false;
		boxOpen1.visible =false;
		boxOpen2.visible =false;
		boxNoOpen0.visible = true;
		boxNoOpen1.visible = true;
		boxNoOpen2.visible = true;
	end
	

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

function jiehun_OnTeamBtn(context)
	Window:Hide();
	UIManager.Show("paiku");
end


function jiehun_OnGetRewardBtn(context)
	local data = HeroStroyData.GetData(showRewardId);
	Proxy4Lua.RequestChapterStarReward(showRewardId,data.Star_[showRewardStar]);
	rewardShow.visible= false;
end

function jiehun_OnRewardOkBtn(context)
	rewardShow.visible= false;
end

function jiehun_OnBox(context)
	showRewardId = context.sender.data;
	showRewardStar =0;
	rewardShow.visible= true;
	updateReward();
end

function jiehun_OnBox1(context)
	showRewardId = context.sender.data;
	showRewardStar =1;
	rewardShow.visible= true;
	updateReward();
end

function jiehun_OnBox2(context)
	showRewardId = context.sender.data;
	showRewardStar =2;
	rewardShow.visible= true;
	updateReward();
end

function updateReward()

		if rewardShow.visible == false then
			return;
		end
		JieHunSystem.instance.chapterID = showRewardId;
		JieHunSystem.instance.chapterBox = showRewardStar;
		rewardList:RemoveChildrenToPool(); 
		local data = HeroStroyData.GetData(showRewardId);
		local itemid = data.Rewards_[showRewardStar];
		local drop = DropData.GetData(itemid);
		local itemdata =  ItemData.GetData(drop.item1_);

		local exp = rewardList:AddItemFromPool(cardGroupUrl); 
		local expIcon0 = exp:GetChild("n5");
		local expName = exp:GetChild("n4");
		local expicon = expIcon0:GetChild("n1");
		local expiconBack = expIcon0:GetChild("n0");
		local expiconLab = expIcon0:GetChild("n2");
		expicon.asLoader.url ="ui://icon/jingyan_icon" ;
		expiconBack.asLoader.url = "ui://" .. itemdata._IconBack;
		expiconLab.text =  drop.exp_ .. "";
		expName.text ="";

		local money = rewardList:AddItemFromPool(cardGroupUrl); 
		local moneyIcon0 = money:GetChild("n5");
		local moneyName = money:GetChild("n4");
		local moneyicon = moneyIcon0 :GetChild("n1");
		local moneyiconBack = moneyIcon0 :GetChild("n0");
		local moneyiconLab = moneyIcon0 :GetChild("n2");
		moneyicon.asLoader.url ="ui://icon/jinbi_icon" ;
		moneyiconBack.asLoader.url = "ui://" .. itemdata._IconBack;
		moneyiconLab.text =  drop.money_ .. "";
		moneyName.text ="";

		local Item = rewardList:AddItemFromPool(cardGroupUrl); 
		local ItemIcon = Item:GetChild("n5");
		local ItemName = Item:GetChild("n4");
		local icon = ItemIcon:GetChild("n1");
		local iconBack = ItemIcon:GetChild("n0");
		local iconLab = ItemIcon:GetChild("n2");
		icon.asLoader.url = "ui://" .. itemdata._Icon;
		iconBack.asLoader.url = "ui://" .. itemdata._IconBack;
		iconLab.text = drop.itemNum1_ .. "";
		ItemName.text = itemdata._Name;
		rewardNeedNum.text = "达到" .. data.Star_[showRewardStar] .. "魂可领取";
		
		local comData  = JieHunSystem.instance:GetChapterData(showRewardId);
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
		if  starNum >= data.Star_[showRewardStar]  then
			if comData.StarReward == nil then
				getRewardBtn.visible = true;
				rewardOkBtn.visible = false;
				UIParamHolder.Set("chaptersDrop", drop.Id_);
				UIParamHolder.Set("showChaptersDrop", true);
			else
				getRewardBtn.visible = true;
				rewardOkBtn.visible = false;
				local max = comData.StarReward.Length;
				for i=1, max do
					if comData.StarReward[i-1] == data.Star_[showRewardStar] then
						getRewardBtn.visible = false;
						rewardOkBtn.visible = true;
						UIParamHolder.Set("chaptersDrop", drop.Id_);
						UIParamHolder.Set("showChaptersDrop", true);
					end
				end
			end
		else
			getRewardBtn.visible = false;
			rewardOkBtn.visible = true;
		end

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
		cardGroupList.numItems = 10;-- JieHunSystem.instance.ChapterEasyDataList.Count;
	else
		cardGroupList.numItems = 10;-- JieHunSystem.instance.ChapterHardDataList.Count;
	end
	stamaPoint.text = GamePlayer._Data.IProperties[10];
	updateReward();
end