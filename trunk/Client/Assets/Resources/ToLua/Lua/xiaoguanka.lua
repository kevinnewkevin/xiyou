require "FairyGUI"

xiaoguanka = fgui.window_class(WindowBase)
local Window;

local fuben0;
local fuben1;
local fuben2;
local fuben3;
local guankaID;
local nameLab;
local leftBtn;
local rightBtn;
local teamBtn;
local challengePanel;
local challengeBtn;
local needPowerLab;
local playerPos;
local smallChapters;
local fubenArr;
local showNum;
local smallId;
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
function xiaoguanka:OnEntry()
	Define.LaunchUIBundle("guankatupian");
	Window = xiaoguanka.New();
	Window:Show();
end

function xiaoguanka:OnInit()
	self.contentPane = UIPackage.CreateObject("xiaoguanka", "xiaoguanka_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n19").asButton;
	local back= self.contentPane:GetChild("n29");
	back.onClick:Add(xiaoguanka_OnBack);
    leftBtn = self.contentPane:GetChild("n21");
    rightBtn = self.contentPane:GetChild("n20");
    leftBtn.onClick:Add(xiaoguanka_OnLeftBtn);
    rightBtn.onClick:Add(xiaoguanka_OnRightBtn);
    fuben0 = self.contentPane:GetChild("n22");
    fuben1 = self.contentPane:GetChild("n23");
    fuben2 = self.contentPane:GetChild("n24");
    fuben3 = self.contentPane:GetChild("n25");
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
	rewardBox0.onClick:Add(xiaoguanka_OnBox);
	rewardBox1.onClick:Add(xiaoguanka_OnBox1);
	rewardBox2.onClick:Add(xiaoguanka_OnBox2);
	teamBtn.onClick:Add(xiaoguanka_OnTeamBtn);
	closeRewardBtn.onClick:Add(xiaoguanka_OnRewardOkBtn);
	getRewardBtn.onClick:Add(xiaoguanka_OnGetRewardBtn);
	rewardOkBtn.onClick:Add(xiaoguanka_OnRewardOkBtn);
    challengePanel = self.contentPane:GetChild("n26");
    challengePanel.visible = false;
    challengeBtn = challengePanel:GetChild("n2");
    needPowerLab = challengePanel:GetChild("n1");
    challengeBtn.onClick:Add(xiaoguanka_OnChallengeBtn);
    fuben0.onClick:Add(xiaoguanka_OnfunbenOne);
    fuben1.onClick:Add(xiaoguanka_OnfunbenTwo);
    fuben2.onClick:Add(xiaoguanka_OnfunbenThree);
    fuben3.onClick:Add(xiaoguanka_OnfunbenFour);
    rewardShow.visible= false;
    fubenArr = {};
    fubenArr[0] = fuben0;
    fubenArr[1] = fuben1;
    fubenArr[2] = fuben2;
    fubenArr[3] = fuben3;

    guankaID = UIManager.GetWindow("daguanka").GetGuankaId();
    showNum = 0;
    playerNum = 0;
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
	challengePanel.visible = false;
	Window:Hide();
end


function xiaoguanka_OnLeftBtn(context)
	challengePanel.visible = false;
	rightBtn.visible = true;
	showNum = showNum - 4;
    if showNum  <= 0 then
        showNum = 0;
        leftBtn.visible = false;
    end
     xiaoguanka_UpdataInfo();
end


function xiaoguanka_OnBack(context)
	challengePanel.visible = false;
end


function xiaoguanka_OnRightBtn(context) 
	challengePanel.visible = false;
	leftBtn.visible = true
	showNum = showNum + 4;
    local chapterData =  JieHunSystem.instance:GetChapterData(guankaID);
    smallChapters = chapterData.SmallChapters;
    local len = smallChapters.Length / 4;
    len = (len -1 ) * 4;
    if showNum >= len then
        showNum = len;
        rightBtn.visible = false;
     end
     xiaoguanka_UpdataInfo();
end

function xiaoguanka_OnTeamBtn(context)
	--Window:Hide();
	UIManager.Show("paiku");
end

function xiaoguanka_OnfunbenOne(context)

	smallId = context.sender.data;
  --  local len = smallChapters.Length;
    --for i = 1, len do
  	  --if smallChapters[i -1]._ID == smallId then    
  	  	--	local smallData  = smallChapters[i -1];
		--end
	--end

	local data = CheckpointData.GetSmallData(guankaID,smallId);
	needPowerLab.text =  data._Main .. "";
    challengePanel.visible = true;
    challengePanel:SetXY(fuben0.x,fuben0.y +365);
end

function xiaoguanka_OnfunbenTwo(context)
	smallId = context.sender.data;
	local data = CheckpointData.GetSmallData(guankaID,smallId);
	needPowerLab.text =  data._Main .. "";
    challengePanel.visible = true;
    challengePanel:SetXY(fuben1.x,fuben1.y - 350);
end

function xiaoguanka_OnfunbenThree(context)
	smallId = context.sender.data;
	local data = CheckpointData.GetSmallData(guankaID,smallId);
	needPowerLab.text =  data._Main .. "";
    challengePanel.visible = true;
    challengePanel:SetXY(fuben2.x,fuben2.y + 365);
end

function xiaoguanka_OnfunbenFour(context)
	smallId = context.sender.data;
	local data = CheckpointData.GetSmallData(guankaID,smallId);
	needPowerLab.text =  data._Main .. "";
    challengePanel.visible = true;
    challengePanel:SetXY(fuben3.x,fuben3.y - 350);
end

function xiaoguanka_OnChallengeBtn(context)
	local data = CheckpointData.GetSmallData(guankaID,smallId);
	if  GamePlayer._Data.IProperties[10] < data._Main then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "体力不够", true);
	else
		Proxy4Lua.ChallengeSmallChapter(smallId);
	end
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

	local starNum = 0;
	 for i = 1, len do
  	  if smallChapters[i -1].Star1 == true then
  	  	starNum = starNum +1;
  	  end

  	  if smallChapters[i -1].Star2 == true then
  	  	starNum = starNum +1;
  	  end

  	  if  smallChapters[i -1].Star3 == true  then    
		starNum = starNum +1;
		end
	end
	rewardStarNum.text = starNum .. "";
	rewardStar0.text = "" .. heroData .Star_[0];
	rewardStar1.text = "" .. heroData .Star_[1];
	rewardStar2.text = "" .. heroData .Star_[2];
	rewardBar.value = starNum /heroData .Star_[2]*100;

	rewardBox0.data = guankaID;
	rewardBox1.data = guankaID;
	rewardBox2.data = guankaID;

    xiaoguanka_UpdataInfo();
    updateReward()

end

function xiaoguanka_UpdataInfo()

    local chapterData =  JieHunSystem.instance:GetChapterData(guankaID);
    smallChapters = chapterData.SmallChapters;
    local hData = HeroStroyData.GetData(guankaID);
    nameLab.text = hData .Name_;
    local len = smallChapters.Length;
	playerPos.visible = false;


	fuben0.onClick:Set(xiaoguanka_OnBack);
    fuben1.onClick:Set(xiaoguanka_OnBack);
    fuben2.onClick:Set(xiaoguanka_OnBack);
    fuben3.onClick:Set(xiaoguanka_OnBack);

    for i =1, 4 do 
    	local bBattle = false;
        if showNum +i -1 >= len then
            fubenArr[i-1].visible  = false;
        else
            fubenArr[i-1].visible  = true;
            local name = fubenArr[i-1]:GetChild("n7");
            local player = fubenArr[i-1]:GetChild("n8");
            local star0 = fubenArr[i-1]:GetChild("n3");
	 		local star1 = fubenArr[i-1]:GetChild("n5");
	 		local star2 = fubenArr[i-1]:GetChild("n4");
            local smallData = smallChapters[showNum +i - 1];
            local data = CheckpointData.GetSmallData(guankaID,smallData.SmallChapterId);
            name.text = data._Name;
            local entityData = EntityData.GetData(data._EntityID);
            local displayData = DisplayData.GetData(entityData._DisplayId);
            local modelRes = displayData._AssetPath;
            player:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes, false));
            --Proxy4Lua.ColorGameObject(player,0.3,0.3,0.3);
            local lock = fubenArr[i-1]:GetChild("n9");
            Proxy4Lua.AddToDelete(modelRes);
            fubenArr[i-1].data = data._ID;

            star0.enabled = false;
            star1.enabled = false;
            star2.enabled = false;
             lock.visible = true;
            if smallData.Star1 == true then 
		 		star0.enabled = true;
		 		-- Proxy4Lua.WhiteGameObject(player);
		 		 bBattle = true;
		 		 lock.visible  = false;
		  	end
			if smallData.Star2 == true then 
		 		star1.enabled = true;
		 		-- Proxy4Lua.WhiteGameObject(player);
		 		 bBattle = true;
		 		 lock.visible  = false;
			  end
		    if smallData.Star3 == true then 
		   		star2.enabled = true;
		   		-- Proxy4Lua.WhiteGameObject(player);
		   		 bBattle = true;
		   		 lock.visible  = false;
			end
            local Trans = fubenArr[i-1]:GetTransition("t0");
            Trans:Play();
            if playerNum == showNum +i - 1 then
            	playerPos.visible = true;
            	playerPos:SetXY(fubenArr[i-1].x + 160 ,fubenArr[i-1].y-100 );
            	--Proxy4Lua.WhiteGameObject(player);
            	bBattle = true;
            	lock.visible  = false;
            --else
            	--playerPos.visible = false;
            end

            if bBattle == true then
            	if i == 1 then
            		fuben0.onClick:Set(xiaoguanka_OnfunbenOne);
            	end
            	if i ==2 then
            	  fuben1.onClick:Set(xiaoguanka_OnfunbenTwo);
            	end
            	if i ==3 then
            	 fuben2.onClick:Set(xiaoguanka_OnfunbenThree);
            	end

            	if i ==4 then
            	fuben3.onClick:Set(xiaoguanka_OnfunbenFour);
            	end
            end
        end
    end

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