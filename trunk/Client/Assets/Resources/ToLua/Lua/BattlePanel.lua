require "FairyGUI"

BattlePanel = fgui.window_class(WindowBase)
local Window;

local autoBtn;
local stateIcon;
local battStartIcon;
local battleTurn;

local battleStart = {};
battleStart.max = 1;
battleStart.count = 0;

local cards = {};
local cardsInGroup = {};
local cardsInGroupNum;
local fees = {};

local skillActor;
local skillActor_modRes;
local skillObj = nil;

function BattlePanel:OnEntry()
	Define.LaunchUIBundle("Card");
	Define.LaunchUIBundle("zhandoushuzi");
	Window = BattlePanel.New();
	Window:Show();
end

function BattlePanel:OnInit()
	self.contentPane = UIPackage.CreateObject("BattlePanel", "BattlePanel_com").asCom;
	self:Center();

	battStartIcon = self.contentPane:GetChild("n7").asImage;

	battleTurn = self.contentPane:GetChild("n36");

	local returnBtn = self.contentPane:GetChild("n8").asButton;
	returnBtn.onClick:Add(BattlePanel_OnReturnBtn);

	autoBtn = self.contentPane:GetChild("n12").asButton;
	autoBtn.onClick:Add(BattlePanel_OnAutoBtn);

	stateIcon = self.contentPane:GetChild("n16").asLoader;
	stateIcon.onClick:Add(BattlePanel_OnTurnOver);

	for i=1, 5 do
		cards[i] = {};
		cards[i]["card"] = self.contentPane:GetChild("n" .. (16 + i)).asCom;
		cards[i]["card"].draggable = true;
		cards[i]["card"].onDragStart:Add(BattlePanel_OnDragStart);
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

	if Battle._CasterDisplayID ~= 0 then
		local dData = DisplayData.GetData(Battle._CasterDisplayID);
		if dData ~= nil then
			skillObj = {};
			skillObj.max = 2;
			skillObj.count = 0;
			skillActor_modRes = dData._AssetPath;
			skillActor:SetNativeObject(Proxy4Lua.GetAssetGameObject(skillActor_modRes, dData._BattleSkillScale, dData._BattleSkillHeight));
		else
			skillActor:SetNativeObject(Proxy4Lua.GetAssetGameObject(""));
			Proxy4Lua.UnloadAsset(skillActor_modRes);
			skillActor_modRes = "";
		end
		Battle._CasterDisplayID = 0;
	end
end

function BattlePanel:OnTick()
	--2秒倒计时
	if battleStart ~= nil then
		battleStart.count = battleStart.count + 1;
		if battleStart.count > battleStart.max then
			BattlePanel_OnHiddenBattleStart();
			battleStart = nil;
		end
	end

	if skillObj ~= nil then
		skillObj.count = skillObj.count + 1;
		if skillObj.count > skillObj.max then
			skillActor:SetNativeObject(Proxy4Lua.GetAssetGameObject(""));
			Proxy4Lua.UnloadAsset(skillActor_modRes);
			skillActor_modRes = "";
			skillObj = nil;
		end
	end
end

function BattlePanel:isShow()
	return Window.isShowing;
end

function BattlePanel:OnDispose()
	Window:Dispose();
end

function BattlePanel:OnHide()
	Window:Hide();
end

function BattlePanel_FlushData()
	local operating = Battle._CurrentState == Battle.BattleState.BS_Oper;
	if operating then
		stateIcon.url = UIPackage.GetItemURL("BattlePanel", "battle_jieshuhuihe");
		stateIcon.touchable = true;
	else
		stateIcon.url = UIPackage.GetItemURL("BattlePanel", "battle_dengdaizhong");
		stateIcon.touchable = false;
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
			cards[i]["power"].text = i;
			cards[i]["cost"].text = eData._Cost;
			cards[i]["card"].visible = true;

			if noPos then
				cards[i]["card"].enabled = false;
			else
				if Battle._Turn == 1 and Battle.IsSelfCard(i-1) then
					cards[i]["card"].enabled = true;
				else
					if Battle._Fee < eData._Cost or operating == false then
						cards[i]["card"].enabled = false;
					else
						cards[i]["card"].enabled = true;
					end
				end
			end
		else
			cards[i]["card"].visible = false;
		end
	end

	local mainActor = Battle.GetActor(GamePlayer._InstID);
	if Battle._Turn == 1 and mainActor == nil then
		if Battle._CurrentState == Battle.BattleState.BS_Oper then
			stateIcon.enabled = false;
		end
	else 
		stateIcon.enabled = true;
	end

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

	battleTurn.text = "第" .. Battle._Turn .. "回合";
	BattlePanel_SetFeeCount(Battle._Fee);
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
	Proxy4Lua.SelectCard4Ready(context.sender.data - 1);
	local eData = Battle.GetHandCard(context.sender.data - 1);
	if eData ~= nil then
		BattlePanel_SetFeeCostCount(eData._Cost);
	end

	for i=1, 5 do
		cards[i]["card"]:SetScale(0.5, 0.5);
	end
end

function BattlePanel_OnReturnBtn()
	print("OnReturnBtn");
end

function BattlePanel_OnTurnOver()
	Proxy4Lua.BattleSetup();
	for i=1, 5 do
		cards[i]["card"]:SetScale(1, 1);
	end
end

function BattlePanel_OnAutoBtn()
	GamePlayer._IsAuto = not GamePlayer._IsAuto;
	UIManager.SetDirty("BattlePanel")
end

function BattlePanel_OnHiddenBattleStart()
	battStartIcon.visible = false;
end