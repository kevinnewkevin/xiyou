require "FairyGUI"

BattlePanel = fgui.window_class(WindowBase)
local Window;

local autoBtn;
local stateIcon;
local startCom;
local battStartIcon;
--local battleTurn;
local countDown;

local cards = {};
local cardsInGroup = {};
local cardsInGroupNum;
local fees = {};

local skillActor;
local skillActor_modRes;
local skillObj = nil;

local skillTransCom;
local skillTrans;
local skillTransName;

local turnTransCom;
local turnTrans;
local turnTransName;

local selectMainRolePos;

local skillList;
local skills;

local reportList;
local reportListBg;
local reportBtnUrl = "ui://BattlePanel/zhanbao_Button";
local reportTransition;

local throwCardCom;
local throwCardTrans;
local throwCardTimer;

function BattlePanel:OnEntry()
	Define.LaunchUIBundle("Card");
	Define.LaunchUIBundle("zhandoushuzi");
	Window = BattlePanel.New();
	Window:Show();
end

function BattlePanel:OnInit()
	self.contentPane = UIPackage.CreateObject("BattlePanel", "BattlePanel_com").asCom;
	self:Center();

	countDown = {};
	countDown.ui = self.contentPane:GetChild("n45").asTextField;
	countDown.count = 30;
	countDown.ui.visible = false;

	startCom = self.contentPane:GetChild("n7").asCom;
	battStartIcon = startCom:GetTransition("t0");
	startCom.visible = false;

--	battleTurn = self.contentPane:GetChild("n36");

	local returnBtn = self.contentPane:GetChild("n8").asButton;
	returnBtn.onClick:Add(BattlePanel_OnReturnBtn);

	autoBtn = self.contentPane:GetChild("n12").asButton;
	--autoBtn.onClick:Add(BattlePanel_OnAutoBtn);
	autoBtn.visible = false;

	stateIcon = self.contentPane:GetChild("n16").asLoader;
	stateIcon.onClick:Add(BattlePanel_OnTurnOver);

	skillList = self.contentPane:GetChild("n50").asList;
	skillList.onClickItem:Add(BattlePanel_OnSelectSkill);

	reportList = self.contentPane:GetChild("n77").asList;
	reportListBg = self.contentPane:GetChild("n51");
	reportTransition = self.contentPane:GetTransition("t0");
	reportList.visible = false;
	reportListBg.visible = false;

	for i=1, 5 do
		cards[i] = {};
		cards[i]["card"] = self.contentPane:GetChild("n" .. (16 + i)).asCom;
		cards[i]["card"].draggable = true;
		cards[i]["card"].onDragStart:Add(BattlePanel_OnDragStart);
		cards[i]["selected"] = cards[i]["card"]:GetChild("n5");
		cards[i]["selected"].visible = false;
		cards[i]["power"] = cards[i]["card"]:GetChild("power");
		cards[i]["cost"] = cards[i]["card"]:GetChild("cost");
		cards[i]["icon"] = cards[i]["card"]:GetChild("card").asLoader;
		cards[i]["card"].data = i;
		cards[i]["card"].onClick:Add(BattlePanel_OnCardClick);
	end

	for i=1, 10 do
		cardsInGroup[i] = self.contentPane:GetChild("n" .. (23 + i));
	end

	for i=1, 5 do
		fees[i] = {};
		fees[i]["com"] = self.contentPane:GetChild("n" .. (36 + i));
		fees[i]["controller"] = fees[i]["com"]:GetController("c1");
	end

	cardsInGroupNum = self.contentPane:GetChild("n43");

	skillActor = self.contentPane:GetChild("n44").asGraph;

	skillTransCom = self.contentPane:GetChild("n49").asCom;
	skillTransName = skillTransCom:GetChild("n45").asTextField;
	skillTrans = skillTransCom:GetTransition("t0");
	skillTransCom.visible = false;

	turnTransCom = self.contentPane:GetChild("n46").asCom;
	turnTransName = turnTransCom:GetChild("n47").asTextField;
	turnTrans = turnTransCom:GetTransition("t2");
	turnTransCom.visible = false;

	selectMainRolePos = self.contentPane:GetChild("n47").asCom;
	selectMainRolePos.visible = false;

	throwCardCom = self.contentPane:GetChild("n78").asCom;
	throwCardCom.visible = false;
	throwCardTrans = self.contentPane:GetTransition("t1");
	throwCardTimer = {};

	BattlePanel_FlushData();
end

function BattlePanel_OnDragStart(context)
	BattlePanel_OnCardClick(context);
	context:PreventDefault();
	local eData = Battle.GetHandCard(context.sender.data - 1);
	local dData = DisplayData.GetData(eData._DisplayId);
	DragDrop.inst:StartDrag(context.sender, dData._AssetPath, context.sender, context.data);
end

function BattlePanel:GetWindow()
	return Window;
end

function BattlePanel:OnUpdate()
	if UIManager.IsDirty("BattlePanel") then
		BattlePanel_FlushData();
		UIManager.ClearDirty("BattlePanel");
	end
end

function BattlePanel:OnTick()
	if skillObj ~= nil then
		skillObj.count = skillObj.count + 1;
		if skillObj.count > skillObj.max then
			skillActor:SetNativeObject(Proxy4Lua.GetAssetGameObject("", false));
			Proxy4Lua.UnloadAsset(skillActor_modRes);
			skillActor_modRes = "";
			skillObj = nil;
		end
	end

	if Battle._CurrentState == Battle.BattleState.BS_Oper then
		if countDown.ui.visible == false then
			countDown.ui.visible = true;
		end
		countDown.count = countDown.count - 1;
		if countDown.count <= 0 then
			BattlePanel_OnTurnOver();
		end
	else
		if countDown.ui.visible == true then
			countDown.ui.visible = false;
		end
		countDown.count = 30;
	end
	countDown.ui.text = countDown.count;
end

function BattlePanel:isShow()
	return Window.isShowing;
end

function BattlePanel:OnDispose()
	Window:Dispose();
end

function BattlePanel:OnHide()
	selectMainRolePos.visible = false;
	Window:Hide();
end

function BattlePanel_FlushData()
	reportList:RemoveChildrenToPool();
	for i=0, Battle._ReportTips.Count - 1 do
		local reportBtn = reportList:AddItemFromPool(reportBtnUrl);
		reportBtn.data = i;
		reportBtn.onClick:Add(BattlePanel_OnReport);
		local side = reportBtn:GetChild("n8").asLoader;
		local oper = reportBtn:GetChild("n7").asLoader;
		local icon = reportBtn:GetChild("n6").asLoader;
		local qubg = reportBtn:GetChild("n5").asLoader;
		side.url = "";
		oper.url = "";
		icon.url = "";
		qubg.url = "";

		if Battle._ReportTips[i]._RBType == ReportBase.RBType.RBT_SelfSkill then
			local sData = SkillData.GetData(Battle._ReportTips[i]._SkillID);
			if sData ~= nil then
				icon.url = "ui://" .. sData._Icon;
			end
		elseif Battle._ReportTips[i]._RBType == ReportBase.RBType.RBT_SelfAppear then
			side.url = "ui://BattlePanel/zb_wo";
			oper.url = "ui://BattlePanel/zb_zhan";
			local eData = EntityData.GetData(Battle._ReportTips[i]._CasterEntityID);
			if eData ~= nil then
				local dData = DisplayData.GetData(eData._DisplayId);
				if dData ~= nil then
					icon.url = "ui://" .. dData._HeadIcon;
					qubg.url = "ui://icon/touxiangkuang_bai";
				else
					qubg.url = "ui://icon/touxiangkuang_bai";
				end
			end
		elseif Battle._ReportTips[i]._RBType == ReportBase.RBType.RBT_AllAppear then
			side.url = "ui://BattlePanel/zb_di";
			oper.url = "ui://BattlePanel/zb_zhan";
			local eData = EntityData.GetData(Battle._ReportTips[i]._CasterEntityID);
			if eData ~= nil then
				local dData = DisplayData.GetData(eData._DisplayId);
				if dData ~= nil then
					icon.url = "ui://" .. dData._HeadIcon;
					qubg.url = "ui://icon/touxiangkuang_bai";
				else
					qubg.url = "ui://icon/touxiangkuang_bai";
				end
			end
		elseif Battle._ReportTips[i]._RBType == ReportBase.RBType.RBT_AllSkill then
			local sideStr = "";
			local eData = EntityData.GetData(Battle._ReportTips[i]._CasterEntityID);
			if eData ~= nil then
				if Battle._ReportTips[i]._Self then
					sideStr = "ui://BattlePanel/zb_wo";
				else
					sideStr = "ui://BattlePanel/zb_di";
				end
			end
			side.url = sideStr;
			oper.url = "ui://BattlePanel/zb_ji";
			local eData = EntityData.GetData(Battle._ReportTips[i]._CasterEntityID);
			if eData ~= nil then
				local dData = DisplayData.GetData(eData._DisplayId);
				if dData ~= nil then
					icon.url = "ui://" .. dData._HeadIcon;
					qubg.url = "ui://icon/touxiangkuang_bai";
				else
					qubg.url = "ui://icon/touxiangkuang_bai";
				end
			end
		end
	end

	local operating = Battle._CurrentState == Battle.BattleState.BS_Oper;
	if operating then
		stateIcon.url = UIPackage.GetItemURL("BattlePanel", "battle_jieshuhuihe");
		stateIcon.touchable = true;
	else
		stateIcon.url = UIPackage.GetItemURL("BattlePanel", "battle_dengdaizhong");
		stateIcon.touchable = false;
	end

	skills = GamePlayer.GetMyActiveSkill();
	local skill;
	local sData;
	skillList.selectedIndex = -1;
	for i=0, skills.Length - 1 do
		skill = skillList:GetChildAt(i);
		local icon = skill:GetChild("n8").asLoader;
		local fee = skill:GetChild("n16").asTextField;
		local lv = skill:GetChild("n7").asTextField;
		local lock = skill:GetChild("n12");
		local hide1 = skill:GetChild("n10");
		local hide2 = skill:GetChild("n13");
		hide1.visible = false;
		hide2.visible = false;
		sData = SkillData.GetData(skills[i]);
		if sData ~= nil then
			icon.url = "ui://" .. sData._Icon;
			lv.text = sData._Level;
			fee.text = sData._Fee;
			lock.visible = false;
			skill.enabled = Battle._Fee >= sData._Fee and Battle.SelectSkillID == 0;
		else
			icon.url = "";
			lv.text = "";
			fee.text = "";
			lock.visible = true;
			skill.enabled = false;
		end
		if Battle.SelectSkillID == skills[i] then
			skillList.selectedIndex = i;
		end
	end

	local cardNum = Battle._LeftCardNum;
	local eData;
	local dData;
	local noPos = Battle.FindEmptyPos() == -1;
	for i=1, 5 do
		if i <= cardNum then
			eData = Battle.GetHandCard(i-1);
			dData = Battle.GetHandCardDisplay(i-1);
			cards[i]["icon"].url = "ui://" .. dData._CardIcon;
			cards[i]["power"].text = Battle.GetHandCardStrLv(i-1);
			cards[i]["cost"].text = eData._Cost;
			cards[i]["card"].visible = true;

			if noPos then
				cards[i]["card"].enabled = false;
			else
--				if Battle._Turn == 1 and Battle.IsSelfCard(i-1) then
--					cards[i]["card"].enabled = true;
--				else
					if Battle._Fee < eData._Cost or operating == false then
						cards[i]["card"].enabled = false;
					else
						cards[i]["card"].enabled = true;
					end
--				end
			end
		else
			cards[i]["card"].visible = false;
		end
	end

	BattlePanel_DisableSkills(not operating or Battle.SelectSkillID ~= 0);

	if Battle._ThrowCardInst ~= nil then
		if throwCardCom.visible == false then
			throwCardCom.visible = true;
		end
		local power = throwCardCom:GetChild("power");
		local cost = throwCardCom:GetChild("cost");
		local icon = throwCardCom:GetChild("card");
		local eData = EntityData.GetData(Battle._ThrowCardInst.UnitId);
		local dData = DisplayData.GetData(eData._DisplayId);
		if dData ~= nil then
			icon.url = "ui://" .. dData._CardIcon;
		else
			icon.url = "";
		end
		if eData ~= nil then
			cost.text = eData._Cost;
		else
			cost.text = "";
		end
		power.text = Battle._ThrowCardInst.IProperties[9];

		if throwCardTrans.playing then
			throwCardTrans:Stop();
		end
		throwCardTrans:Play();

		Battle._ThrowCardInst = nil;
	end

--	local mainActor = Battle.GetActor(GamePlayer._InstID);
--	if Battle._Turn == 1 and mainActor == nil then
--		if Battle._CurrentState == Battle.BattleState.BS_Oper then
--			stateIcon.enabled = false;
----			selectMainRolePos.visible = true;
--			Battle.PutMainInBattle();
--			BattlePanel_FlushData();
--		end
--	else 
--		stateIcon.enabled = true;
--		selectMainRolePos.visible = false;
--	end

	for i=1, 10 do
		if i <= Battle.CardsInGroupCount then
			cardsInGroup[i].visible = true;
		else
			cardsInGroup[i].visible = false;
		end
	end

	cardsInGroupNum.text = Battle.CardsInGroupCount;
	if Battle.CardsInGroupCount <= 0 then
		cardsInGroupNum.visible = false;
	else
		cardsInGroupNum.visible = true;
	end

	if GamePlayer._IsAuto then
		autoBtn:GetController("icon").selectedIndex = 1;
	else
		autoBtn:GetController("icon").selectedIndex = 0;
	end

--	battleTurn.text = "第" .. Battle._Turn .. "回合";
	BattlePanel_SetFeeCount(Battle._Fee);
end

function BattlePanel_OnReport(context)
	Battle._SelectReportIdx = context.sender.data;
	UIManager.Show("zhanbao");
end

function BattlePanel_OnSelectSkill()
--	if Battle.SelectSkillID == skills[skillList.selectedIndex] then
--		skillList.selectedIndex = -1;
--		Battle.SelectSkillID = 0;
--	else
		Battle.SelectSkillID = skills[skillList.selectedIndex];
		local sData = SkillData.GetData(Battle.SelectSkillID);
		if sData ~= nil then
			Battle.CostFee(sData._Fee);
		end
		BattlePanel_DisableSkills(true);
--	end
end

function BattlePanel_DisableSkills(yes)
	skills = GamePlayer.GetMyActiveSkill();
	local skill;
	local sData;
	for i=0, skills.Length - 1 do
		skill = skillList:GetChildAt(i);
		skill.enabled = not yes;
		sData = SkillData.GetData(skills[0]);
		if sData ~= nil  and skill.enabled then
			skill.enabled = sData._Fee <= Battle._Fee;
		end
	end
end

function BattlePanel_SetFeeCount(count)
	for i=1, 5 do
		if i <= count then
			fees[i]["com"].visible = true;
		else
			fees[i]["com"].visible = false;
		end
		fees[i]["controller"].selectedIndex = 1;
	end
end

function BattlePanel_SetFeeCostCount(count)
	for i=1, 5 do
		if i <= count then
			fees[i]["controller"].selectedIndex = 0;
		else
			fees[i]["controller"].selectedIndex = 1;
		end
	end
end

function BattlePanel_OnCardClick(context)
	if Proxy4Lua.SameCardSelected(context.sender.data - 1) == true then
		UIManager.Show("shuxing");
	end

	Proxy4Lua.SelectCard4Ready(context.sender.data - 1);
	local eData = Battle.GetHandCard(context.sender.data - 1);
	if eData ~= nil then
		BattlePanel_SetFeeCostCount(eData._Cost);
	end

	for i=1, 5 do
		cards[i]["selected"].visible = context.sender.data == i;
		cards[i]["card"]:SetScale(0.5, 0.5);
	end
end

function BattlePanel_OnReturnBtn()
	local MessageBox = UIManager.ShowMessageBox();
	MessageBox:SetData("提示", "是否退出战斗？", false, BattlePanel_OnReturn);
end

function BattlePanel_OnReturn()
	--SceneLoader.LoadScene("main");
end

function BattlePanel_OnTurnOver()
	if Battle._CurrentState ~= Battle.BattleState.BS_Oper then
		return;
	end
--	if Battle._Turn == 1 and mainActor == nil then
--		Battle.PutCardInBattle();
--	else
		Proxy4Lua.BattleSetup();
--	end
	for i=1, 5 do
		cards[i]["selected"].visible = false;
		cards[i]["card"]:SetScale(1, 1);
	end
	Battle.SwitchPoint(false);
	Battle._SelectedHandCardInstID = 0;
end

function BattlePanel:NormalCard()
	for i=1, 5 do
		cards[i]["selected"].visible = false;
		cards[i]["card"]:SetScale(1, 1);
	end
--	Battle._SelectedHandCardInstID = 0;
end

function BattlePanel_OnAutoBtn()
	GamePlayer._IsAuto = not GamePlayer._IsAuto;
	UIManager.SetDirty("BattlePanel")
	for i=1, 5 do
		cards[i]["selected"].visible = false;
		cards[i]["card"]:SetScale(1, 1);
	end
end

function BattlePanel:ShowBattleStart()
	if startCom.visible == false then
		startCom.visible = true;
	end
	if battStartIcon.playing then
		battStartIcon:Stop();
	end
	battStartIcon:Play();

	if reportList.visible == false then
		reportList.visible = true;
	end
	if reportListBg.visible == false then
		reportListBg.visible = true;
	end
	if reportTransition.playing then
		reportTransition:Stop();
	end
	reportTransition:Play();
end

function BattlePanel:ShowTurn()
	Battle.SelectSkillID = 0;
	skillList.selectedIndex = -1;
	BattlePanel_DisableSkills(false);
	if turnTransCom.visible == false then
		turnTransCom.visible =true;
	end
	turnTransName.text = "第" .. Battle._Turn .. "回合";
	if turnTrans.playing then
		turnTrans:Stop();
	end
	turnTrans:Play();
end

function BattlePanel:ShowSkill()
	skillTransName.text = Battle._CasterSkillName;
	if skillTransCom.visible == false then
		skillTransCom.visible = true;
	end
	if skillTrans.playing then
		skillTrans:Stop();
	end
	skillTrans:Play();

	local dData = DisplayData.GetData(Battle._CasterDisplayID);
	if dData ~= nil then
		skillObj = {};
		skillObj.max = 2;
		skillObj.count = 0;
		skillActor_modRes = dData._AssetPath;
		skillActor:SetNativeObject(Proxy4Lua.GetAssetGameObject(skillActor_modRes, dData._BattleSkillScale, dData._BattleSkillHeight, false));
	else
		skillActor:SetNativeObject(Proxy4Lua.GetAssetGameObject("", false));
		Proxy4Lua.UnloadAsset(skillActor_modRes);
		skillActor_modRes = "";
	end
end