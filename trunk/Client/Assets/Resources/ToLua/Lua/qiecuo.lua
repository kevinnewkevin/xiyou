require "FairyGUI"

qiecuo = fgui.window_class(WindowBase)
local Window;

local allCardGroupList;
local cardGroupList;
local cardGroupUrl = "ui://qiecuo/paizuanniu_Button";
local cardItemGapUrl = "ui://qiecuo/kong_com";
local cardItemUrl = "ui://qiecuo/touxiangkuang_Label";
local crtCardName;
local crtGroupIdx = 0;
local starting;
local cancelBtn;
local cardGroupList;
local tiantiFen0;
local tiantiFen1;
local levelImg;
local battleTeamId = 0;
local okTrans; 
local anim;
local pipeiOk;
local nameArr;
local overTime;

local ttRewardList;
local rewardList;
local ttRewardClose;


function qiecuo:OnEntry()
	Window = qiecuo.New();
	Window:Show();
end

function qiecuo:GetWindow()
	return Window;
end

function qiecuo:OnInit()
 	self.contentPane = UIPackage.CreateObject("qiecuo", "qiecuo_com").asCom;
	self:Center();
	self.closeButton = self.contentPane:GetChild("n7").asButton;

	local rightPart = self.contentPane:GetChild("n5").asCom;
	local bg = rightPart:GetChild("n3");
	allCardGroupList = bg:GetChild("n5").asList;

	local cardGroup = rightPart:GetChild("n6");
	--crtCardName = cardGroup:GetChild("n23");
	cardGroupList = cardGroup:GetChild("n27").asList;
	local setBattleBtn = cardGroup:GetChild("n29").asButton;
	setBattleBtn.onClick:Add(qiecuo_OnSetBattle);
	local delBtn = cardGroup:GetChild("n24").asButton;
	delBtn.visible = false;
	starting = self.contentPane:GetChild("n9").asCom;
	cancelBtn = starting:GetChild("n8").asButton;
	cancelBtn.onClick:Add(qiecuo_OnCancel);
	starting.visible = false;
	pipeiOk = starting:GetChild("n13");
	anim = starting:GetChild("n14");
	okTrans = pipeiOk:GetTransition("t1");

	local name0 = pipeiOk:GetChild("n24");
	local name1 = pipeiOk:GetChild("n23");
	local name2 = pipeiOk:GetChild("n22");
	local name3 = pipeiOk:GetChild("n21");
	local name4 = pipeiOk:GetChild("n19");
	local name5 = pipeiOk:GetChild("n18");
	local name6 = pipeiOk:GetChild("n17");
	local name7 = pipeiOk:GetChild("n16");
	local name8 = pipeiOk:GetChild("n15");
	local name9 = pipeiOk:GetChild("n14");
	local name10 = pipeiOk:GetChild("n13");

	ttRewardList = self.contentPane:GetChild("n12").asCom;
	ttRewardClose = ttRewardList:GetChild("n16");
	ttRewardClose.onClick:Add(qiecuo_ttRewardClose);
	ttRewardList.visible = false;
	rewardList = ttRewardList:GetChild("n13");
	rewardList.itemRenderer = qiecuo_RenderListItem;
	nameArr ={};
	nameArr[0] = name0;
	nameArr[1] = name1;
	nameArr[2] = name2;
	nameArr[3] = name3;
	nameArr[4] = name4;
	nameArr[5] = name5;
	nameArr[6] = name6;
	nameArr[7] = name7;
	nameArr[8] = name8;
	nameArr[9] = name9;
	nameArr[10] = name10;



	for i=1, 5 do
		local groupItem = allCardGroupList:AddItemFromPool(cardGroupUrl);
		groupItem.onClick:Add(qiecuo_OnSelectGroup);
		groupItem.data = i - 1;
	end

	allCardGroupList.selectedIndex = crtGroupIdx;

	local liftPanel = self.contentPane:GetChild("n6").asCom;
	local startBtn = liftPanel:GetChild("n35");
	startBtn.onClick:Add(qiecuo_OnStart);
	levelImg = liftPanel:GetChild("n34");
	local descLab = liftPanel:GetChild("n33");
	descLab.text = "天梯会匹配实力相近的对手，战斗胜利将获得各种英雄碎片，段位越高奖励越好！";
	tiantiFen0 = liftPanel:GetChild("n41");
	--tiantiFen1 = liftPanel:GetChild("n42");
	overTime = liftPanel:GetChild("n43");
	local listBtn = liftPanel:GetChild("n45");
	listBtn.onClick:Add(qiecuo_OnList);
	local battleVideoBtn = liftPanel:GetChild("n44");
	battleVideoBtn.onClick:Add(qiecuo_OnVideo);
	qiecuo_FlushData();
end



function qiecuo_RenderListItem(index, obj)

end

function qiecuo_OnDeleteGroup(context)
	local MessageBox = UIManager.ShowMessageBox();
end
 
function qiecuo_OnSetBattle(context)
	local MessageBox = UIManager.ShowMessageBox();
end





function qiecuo:OnUpdate()
	if UIManager.IsDirty("qiecuo") then
		qiecuo_FlushData();
		UIManager.ClearDirty("qiecuo");
	end

	if Proxy4Lua.ReadyToJoinBattle == true then
		anim.visible = false;
		cancelBtn.enabled = false;
		okTrans:Play();
	end


end

function qiecuo:OnTick()
	 
end

function qiecuo:isShow()
	return Window.isShowing;
end

function qiecuo:OnDispose()
	Window:Dispose();
end

function qiecuo:OnHide()
	starting.visible = false;
	anim.visible = true;
	cancelBtn.enabled = true;
	Window:Hide();
end

function qiecuo_FlushData()

	cardGroupList:RemoveChildrenToPool(); 
	local groupCards = GamePlayer.GetGroupCards(crtGroupIdx);
	if groupCards == nil then
		return;
	end
	local displayData;
	local entityData;
	for i=1, groupCards.Count do
		if i == 1 or i == 4 or i == 8 then
			cardGroupList:AddItemFromPool(cardItemGapUrl);
		end
		displayData = GamePlayer.GetDisplayDataByIndexFromGroup(crtGroupIdx, i - 1);
		local itemBtn = cardGroupList:AddItemFromPool(cardItemUrl);
		if displayData ~= nil then
			itemBtn:GetChild("n5").asLoader.url = "ui://" .. displayData._HeadIcon;
		else
			itemBtn:GetChild("n5").asLoader.url = "";
		end
		entityData = GamePlayer.GetEntityDataByIndexFromGroup(crtGroupIdx, i - 1);
		if entityData ~= nil then
			local fee = itemBtn:GetChild("n7");
			fee.text = entityData._Cost
		else
			fee.text = "";
		end
		local radImg = itemBtn:GetChild("n10");
		local level = itemBtn:GetChild("n6");
		local instId = GamePlayer.GetInstIDFromGroup(crtGroupIdx,  i - 1);
		local entityInst = GamePlayer.GetCardByInstID(instId);
		local  levelData =  StrengthenData.GetData( entityInst.UnitId,  entityInst.IProperties[9]+1);
		level.text = entityInst.IProperties[9] .. "";
		local itemNum = BagSystem.GetItemMaxNum(levelData._ItemId);
		if itemNum >= levelData._ItemNum then
			radImg.visible = true;
		else
			radImg.visible = false;
		end

		itemBtn.onClick:Add(qiecuo_OnCardInGroup);
		itemBtn.data = GamePlayer.GetInstIDFromGroup(crtGroupIdx, i - 1);
	end


	local groupItem;
	local groupName;
	for i=1, 5 do
			groupItem = allCardGroupList:GetChildAt(i-1);
--			groupName = GamePlayer.GetGroupName(i - 1);
--			if groupName == "" then
				groupName = "" .. i;
--			end
			groupItem:GetChild("n3").text = groupName;
			if crtGroupIdx == i - 1 then
			--	crtCardName.text = groupName;
			end
			if i - 1 == GamePlayer._CrtBattleGroupIdx then
				groupItem:GetChild("n4").visible = true;
			else
				groupItem:GetChild("n4").visible = false;
			end
		end
		local month = GamePlayer.nowTimeMonth();
		local year = GamePlayer.nowTimeYear();
		if month  == 12 then
			month = 1;
			year  = year +1;
		else
			month = month +1; 
		end

		overTime.text =  year.."年"..month.."月1日"
		tiantiFen0.text =  "积分 " .. GamePlayer._TianTiVal;
		--tiantiFen1.text =  "积分 " .. GamePlayer._TianTiVal;
		local level = GamePlayer.GetTianTiLevel();
		levelImg.asLoader.url = "ui://qiecuo/duanwei" .. level;
		--tiantiFen1.text =  Proxy4Lua.ConvertToChineseNumber(level).."段";
end



function qiecuo_OnSelectGroup(context)
	crtGroupIdx = context.sender.data;
	qiecuo_FlushData();
end


function qiecuo_OnStart(context)

	local groupCards = GamePlayer.GetGroupCards(crtGroupIdx);
	if groupCards == nil then
		return;
	end
	Proxy4Lua.NextBattleDelay = 3;
	Proxy4Lua.StartMatching(GamePlayer._CrtBattleGroupIdx+1);
	starting.visible = true;
	qiecuo_RandName();
	qiecuo_FlushData();
end

function qiecuo_OnCardInGroup(context)
	crtCardInstID = context.sender.data;
	UIParamHolder.Set("qiecuo1", crtCardInstID);
	UIParamHolder.Set("qiecuo2", true);
	UIParamHolder.Set("showBoos",false);
	UIManager.Show("xiangxiziliao");
end

function qiecuo_OnCancel(context)
	Proxy4Lua.StopMatching();
	starting.visible = false;
	Proxy4Lua.NextBattleDelay = 0;
	Proxy4Lua._CancelMatch = true;
end

function qiecuo_OnSetBattle(context)
	local MessageBox = UIManager.ShowMessageBox();
	MessageBox:SetData("提示", "是否设置当前卡组为出战卡组？", false, qiecuo_OnSet);
end

function qiecuo_OnSet(context)
	GamePlayer._CrtBattleGroupIdx = crtGroupIdx;
	battleTeamId = crtGroupIdx;
	UIManager.HideMessageBox();
	UIManager.SetDirty("qiecuo");
end

function qiecuo_RandName()

	for i=1, 11 do
	nameArr[i -1].text = RandNameData.Rand(); 
	end
end

function qiecuo_ttRewardClose(context)
	ttRewardList.visible = false;
end

function qiecuo_OnList(context)
	rewardList.numItems = LadderData.metaData.Count;
	ttRewardList.visible = true;
end

function qiecuo_OnVideo(context)
	ttRewardList.visible = false;
	Proxy4Lua.RequestPlayerRecordData(GamePlayer._InstID);
	--Proxy4Lua.RequestSelfRecordData();
end

function qiecuo_OnTtItem(context)
	local iId = context.sender.data;
	UIParamHolder.Set("tipsItem", iId);
	UIManager.Show("bagtips");
end

function qiecuo_RenderListItem(index, obj)
	local data = LadderData.metaData[index+1];
	obj:GetChild("n6").text = "积分：" .. data._ScoreL .. "-" .. data._ScoreH;
	obj:GetChild("n5").asLoader.url = "ui://qiecuo/xiao_duanwei" .. (index +1);

	if GamePlayer._TianTiVal >= data._ScoreL and GamePlayer._TianTiVal <= data._ScoreH then
		obj:GetChild("n13").visible = true;
	else
		obj:GetChild("n13").visible = false;
	end
	local drop = DropData.GetData(data._LadderDrop);

	if drop.money_ ~= 0 then
		obj:GetChild("n8").asLoader.url = "ui://icon/jinbi_icon" ;
		--obj:GetChild("n7").asLoader.url = "ui://" .. item._IconBack;  
	else
		obj:GetChild("n8").visible = false;
		obj:GetChild("n7").visible = false; 
	end

	local itemIcon = obj:GetChild("n10");
	if drop.item1_ ~= 0 then
		local item1 = ItemData.GetData(drop.item1_);
		itemIcon.asLoader.url = "ui://" .. item1._Icon;
		obj:GetChild("n9").asLoader.url = "ui://" .. item1._IconBack;  
		itemIcon.onClick:Add(qiecuo_OnTtItem);
		itemIcon.data = drop.item1_;
	else
		itemIcon.visible = false;
		itemIcon.onClick:Remove(qiecuo_OnTtItem);
		obj:GetChild("n9").visible = false; 
	end

	if drop.hero_ ~= 0 then
		local entityData = EntityData.GetData(drop.hero_);
		local displayData = DisplayData.GetData(entityData._DisplayId);
		obj:GetChild("n11").asLoader.url = "ui://" .. displayData._Quality;
		obj:GetChild("n12").asLoader.url = "ui://" .. displayData._HeadIcon; 
	else
		obj:GetChild("n11").visible = false;
		obj:GetChild("n12").visible = false; 
	end


end