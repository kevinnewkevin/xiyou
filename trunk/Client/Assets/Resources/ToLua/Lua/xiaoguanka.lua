require "FairyGUI"

xiaoguanka = fgui.window_class(WindowBase)
local Window;


local guankaID;
local nameLab;

local teamBtn;
local smallChapters;
local fubenArr;
local showNum;
local playerNum;
local rewardBarPanel;
local rewardBar;
local rewardStarNum;
local rewardStar0;
local rewardStar1;
local rewardStar2;
local rewardBox0;
local rewardBox1;
local rewardBox2;
local showRewardStar;
local showRewardId;
local rewardShow;
local rewardList;
local getRewardBtn;
local closeRewardBtn
local rewardNeedNum;
local rewardOkBtn;
local nowCanBattle;
local oneShow;
local chaptersList;

function xiaoguanka:OnEntry()
	Define.LaunchUIBundle("guankatupian");
	Define.LaunchUIBundle("guankatupian");
	Window = xiaoguanka.New();
	Window:Show();
end

function xiaoguanka:OnInit()
	self.contentPane = UIPackage.CreateObject("xiaoguanka", "xiaoguanka_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n19").asButton;
    nameLab = self.contentPane:GetChild("n6");
    teamBtn = self.contentPane:GetChild("n30");
    playerPos = self.contentPane:GetChild("n27");
    rewardBarPanel = self.contentPane:GetChild("n4");
	rewardBar = rewardBarPanel:GetChild("n4");
	rewardShow = self.contentPane:GetChild("n38");
	rewardList = rewardShow:GetChild("n12");
	rewardStarNum = rewardBarPanel:GetChild("n14");
	rewardStar0 = rewardBarPanel:GetChild("n16");
	rewardStar1 = rewardBarPanel:GetChild("n17");
	rewardStar2 = rewardBarPanel:GetChild("n18");
	rewardBox0 = rewardBarPanel:GetChild("n10");
	rewardBox1 = rewardBarPanel:GetChild("n11");
	rewardBox2 = rewardBarPanel:GetChild("n12");
	getRewardBtn = rewardShow:GetChild("n18");
	closeRewardBtn = rewardShow:GetChild("n19");
	rewardNeedNum = rewardShow:GetChild("n15");
	rewardOkBtn = rewardShow:GetChild("n17");

	chaptersList = self.contentPane:GetChild("n39");
	chaptersList.itemRenderer = xiaoguanka_RenderListItem;
	rewardBox0.onClick:Add(xiaoguanka_OnBox);
	rewardBox1.onClick:Add(xiaoguanka_OnBox1);
	rewardBox2.onClick:Add(xiaoguanka_OnBox2);
	teamBtn.onClick:Add(xiaoguanka_OnTeamBtn);
	closeRewardBtn.onClick:Add(xiaoguanka_OnRewardOkBtn);
	getRewardBtn.onClick:Add(xiaoguanka_OnGetRewardBtn);
	rewardOkBtn.onClick:Add(xiaoguanka_OnRewardOkBtn);
    rewardShow.visible= false;

    guankaID = UIManager.GetWindow("daguanka").GetGuankaId();
    showNum = 0;
    playerNum = 0;
    oneShow = false;
	xiaoguanka_FlushData();
end


function xiaoguanka:GetWindow()
	return Window;
end

function xiaoguanka:OnUpdate()
	if UIManager.IsDirty("xiaoguanka") then
		xiaoguanka_FlushData();
		UIManager.ClearDirty("xiaoguanka");
	end
end

function xiaoguanka:OnTick()
	
end

function xiaoguanka:isShow()
	return Window.isShowing;
end

function xiaoguanka:OnDispose()
	Window:Dispose();
end

function xiaoguanka:OnHide()
	oneShow = false; 
	 playerNum = 0;
	Proxy4Lua.ClearToDeleteAsset("xiaoguanka");
	Window:Hide();
	GuideSystem.CloseUI("xiaoguanka");
end

function xiaoguanka_OnTeamBtn(context)
	UIManager.Show("paiku");
end



function xiaoguanka_OnBox(context)
	showRewardId = context.sender.data;
	showRewardStar =0;
	rewardShow.visible= true;
	updateReward();
end

function xiaoguanka_OnBox1(context)
	showRewardId = context.sender.data;
	showRewardStar =1;
	rewardShow.visible= true;
	updateReward();
end

function xiaoguanka_OnBox2(context)
	showRewardId = context.sender.data;
	showRewardStar =2;
	rewardShow.visible= true;
	updateReward();
end


function xiaoguanka_OnGetRewardBtn(context)
	local data = HeroStroyData.GetData(showRewardId);
	Proxy4Lua.RequestChapterStarReward(showRewardId,data.Star_[showRewardStar]);
	JieHunSystem.instance.chapterID = showRewardId;
	JieHunSystem.instance.chapterBox = showRewardStar;
	rewardShow.visible= false;
end

function xiaoguanka_OnRewardOkBtn(context)
	rewardShow.visible= false;
end


function xiaoguanka_FlushData()

    guankaID = UIManager.GetWindow("daguanka").GetGuankaId();
    local chapterData =  JieHunSystem.instance:GetChapterData(guankaID);
    local heroData = HeroStroyData.GetData(guankaID);
    smallChapters = chapterData.SmallChapters;
    local len = smallChapters.Length;
    for i = 1, len do
 	 if smallChapters[i -1].Star1 == true or smallChapters[i -1].Star2 == true or smallChapters[i -1].Star3 == true  then    
			playerNum = i;
		end
	end
	if playerNum == len  then
	 	playerNum  = len  -1;
	end
	nowCanBattle = 0;
	local starNum = 0;
	 for i = 1, len do
  	  if smallChapters[i -1].Star1 == true then
  	  	starNum = starNum +1;
  	  	nowCanBattle = i -1;
  	  end

  	  if smallChapters[i -1].Star2 == true then
  	  	starNum = starNum +1;
  	  	nowCanBattle = i -1;
  	  end

  	  if  smallChapters[i -1].Star3 == true  then    
		starNum = starNum +1;
		nowCanBattle = i -1;
		end
	end
	rewardStarNum.text = starNum .. "";
	rewardStar0.text = "" .. heroData .Star_[0];
	rewardStar1.text = "" .. heroData .Star_[1];
	rewardStar2.text = "" .. heroData .Star_[2];
	rewardBar.value = starNum /heroData .Star_[2]*100;

	local boxOpen0 = rewardBox0:GetChild("n5");
	local boxNoOpen0 = rewardBox0:GetChild("n4");
	local boxOpen1 = rewardBox1:GetChild("n5");
	local boxNoOpen1 = rewardBox1:GetChild("n4");
	local boxOpen2 = rewardBox2:GetChild("n5");
	local boxNoOpen2 = rewardBox2:GetChild("n4");
	boxOpen0.visible =false;
	boxOpen1.visible =false;
	boxOpen2.visible =false;
	boxNoOpen0.visible = true;
	boxNoOpen1.visible = true;
	boxNoOpen2.visible = true;
	local openArr= {};
	openArr[0] = boxOpen0;
	openArr[1] = boxOpen1;
	openArr[2] = boxOpen2;
	local noOpenArr= {};
	noOpenArr[0] = boxNoOpen0;
	noOpenArr[1] = boxNoOpen1;
	noOpenArr[2] = boxNoOpen2;
	if chapterData.StarReward ~= nil then
		local len = chapterData.StarReward.Length;
		print("len" .. len);
		for i=1, len do
			if chapterData.StarReward[i-1] ~= 0 then
				openArr[i-1].visible =true;
				noOpenArr[i-1].visible =false;
			end
		end
	end

	rewardBox0.data = guankaID;
	rewardBox1.data = guankaID;
	rewardBox2.data = guankaID;

	if oneShow == false then
		oneShow = true;
		showNum = nowCanBattle/4*4;
	end	 

 	local chapterData =  JieHunSystem.instance:GetChapterData(guankaID);
    smallChapters = chapterData.SmallChapters;
    local hData = HeroStroyData.GetData(guankaID);
    nameLab.text = hData .Name_;
    chaptersList.numItems = chapterData.SmallChapters.Length;

    updateReward()

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

function xiaoguanka_RenderListItem(index, obj)
	local chapterData =  JieHunSystem.instance:GetChapterData(guankaID);
    smallChapters = chapterData.SmallChapters;
    local smallData = smallChapters[index];
    local data = CheckpointData.GetSmallData(guankaID,smallData.SmallChapterId);
    obj:GetChild("n7").text = data._Name;
    local gImg = obj:GetChild("n11");
    local Img = obj:GetChild("n10");
    Img .asLoader.url = "ui://" .. data._pic;
   	gImg .asLoader.url = "ui://" .. data._pic;
    local lock = obj:GetChild("n9");
    local star0 = obj:GetChild("n3");
	local star1 = obj:GetChild("n5");
	local star2 = obj:GetChild("n4");
	star0.enabled = false;
	star1.enabled = false;
	star2.enabled = false;
	lock.visible = true;
	gImg.visible = true;
	local RecordBtn = obj:GetChild("n15");
	if smallData.Star1 == true then 
	 		star0.enabled = true;
	 		 bBattle = true;
	 		 lock.visible  = false;
	 		 gImg.visible  = false;
	 		 RecordBtn.visible = true;
	  	end
		if smallData.Star2 == true then 
	 		star1.enabled = true;
	 		 bBattle = true;
	 		 lock.visible  = false;
	 		 RecordBtn.visible = true;
	 		 gImg.visible  = false;
		  end
	    if smallData.Star3 == true then 
	   		star2.enabled = true;
	   		 bBattle = true;
	   		 lock.visible  = false;
	   		 gImg.visible  = false;
	   		 RecordBtn.visible = true;
		end
	RecordBtn.onClick:Add(xiaoguanka_OnRecord);
	RecordBtn.data = smallData.SmallChapterId;
	--local img = obj:GetChild("n10");
	Img.onClick:Add(xiaoguanka_OnChallengeBtn);
	Img.data = smallData.SmallChapterId;
	RecordBtn.visible = false;
	if index == playerNum then
		lock.visible  = false;
		gImg.visible  = false;
		RecordBtn.visible = true;
	end

end

function xiaoguanka_OnChallengeBtn(context)
	local sId = context.sender.data;
	local data = CheckpointData.GetSmallData(guankaID, sId);
	if  GamePlayer._Data.IProperties[10] < data._Main then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "体力不够", true);
	else
		Proxy4Lua.ChallengeSmallChapter(sId);
		Proxy4Lua.RegHoldUI("main", "daguanka");
		Proxy4Lua.RegHoldUI("main", "xiaoguanka");
	end
end


function xiaoguanka_OnRecord(context)
	local sId = context.sender.data;
	Proxy4Lua.RequestFBRecordData(sId);
end