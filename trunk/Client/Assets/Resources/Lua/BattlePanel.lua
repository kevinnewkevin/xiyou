--[[local GameObject = UnityEngine.GameObject;

function BattlePanel_start()
	RegistEvt();
	FlushData();
end

function BattlePanel_update()
	if UIManager.IsDirty("BattlePanel") then
		FlushData();
		UIManager.ClearDirty("BattlePanel");
	end
end

function BattlePanel_tick()
	
end

local Card0 = {}
local Card1 = {}
local Card2 = {}
local Card3 = {}
local Card4 = {}
local Skip = {}

function RegistEvt()
	--SkipBtn
	Skip.obj = GameObject.Find("BattlePanel(Clone)/SkipBtn");
	local btn = GameObject.GetComponent(Skip.obj, "Button");
	local evt = btn.onClick;
	evt.AddListener(evt, skip);

	--HandCards
	Card0.obj = GameObject.Find("BattlePanel(Clone)/Card0");
	Card0.plus = GameObject.Find("BattlePanel(Clone)/Card0/Plus");
	Card0.cost = GameObject.Find("BattlePanel(Clone)/Card0/Cost");
	Card0.plus = GameObject.GetComponent(Card0.plus, "Text");
	Card0.cost = GameObject.GetComponent(Card0.cost, "Text");
	btn = GameObject.GetComponent(Card0.obj, "Button");
	evt = btn.onClick;
	evt.AddListener(evt, HandCards_0);

	Card1.obj = GameObject.Find("BattlePanel(Clone)/Card1");
	Card1.plus = GameObject.Find("BattlePanel(Clone)/Card1/Plus");
	Card1.cost = GameObject.Find("BattlePanel(Clone)/Card1/Cost");
	Card1.plus = GameObject.GetComponent(Card1.plus, "Text");
	Card1.cost = GameObject.GetComponent(Card1.cost, "Text");
	btn = GameObject.GetComponent(Card1.obj, "Button");
	evt = btn.onClick;
	evt.AddListener(evt, HandCards_1);

	Card2.obj = GameObject.Find("BattlePanel(Clone)/Card2");
	Card2.plus = GameObject.Find("BattlePanel(Clone)/Card2/Plus");
	Card2.cost = GameObject.Find("BattlePanel(Clone)/Card2/Cost");
	Card2.plus = GameObject.GetComponent(Card2.plus, "Text");
	Card2.cost = GameObject.GetComponent(Card2.cost, "Text");
	btn = GameObject.GetComponent(Card2.obj, "Button");
	evt = btn.onClick;
	evt.AddListener(evt, HandCards_2);

	Card3.obj = GameObject.Find("BattlePanel(Clone)/Card3");
	Card3.plus = GameObject.Find("BattlePanel(Clone)/Card3/Plus");
	Card3.cost = GameObject.Find("BattlePanel(Clone)/Card3/Cost");
	Card3.plus = GameObject.GetComponent(Card3.plus, "Text");
	Card3.cost = GameObject.GetComponent(Card3.cost, "Text");
	btn = GameObject.GetComponent(Card3.obj, "Button");
	evt = btn.onClick;
	evt.AddListener(evt, HandCards_3);

	Card4.obj = GameObject.Find("BattlePanel(Clone)/Card4");
	Card4.plus = GameObject.Find("BattlePanel(Clone)/Card4/Plus");
	Card4.cost = GameObject.Find("BattlePanel(Clone)/Card4/Cost");
	Card4.plus = GameObject.GetComponent(Card4.plus, "Text");
	Card4.cost = GameObject.GetComponent(Card4.cost, "Text");
	btn = GameObject.GetComponent(Card4.obj, "Button");
	evt = btn.onClick;
	evt.AddListener(evt, HandCards_4);
end

function FlushData()
	local lblObj = GameObject.Find("BattlePanel(Clone)/Top/Turn");
	local lblText = GameObject.GetComponent(lblObj, "Text");
	lblText.text = "第1回合";

	lblObj = GameObject.Find("BattlePanel(Clone)/CardPoolNum");
	lblText = GameObject.GetComponent(lblObj, "Text");
	lblText.text = "8";

	Card0.plus.text = 1;
	Card0.cost.text = 2;

	Card1.plus.text = 3;
	Card1.cost.text = 4;

	Card2.plus.text = 5;
	Card2.cost.text = 6;

	Card3.plus.text = 7;
	Card3.cost.text = 8;

	Card4.plus.text = 9;
	Card4.cost.text = 0;
end

function skip()
	Proxy4Lua.BattleSetup();
end

function HandCards_0()
	Proxy4Lua.SelectCard4Ready(0);
end
function HandCards_1()
	Proxy4Lua.SelectCard4Ready(1);
end
function HandCards_2()
	Proxy4Lua.SelectCard4Ready(2);
end
function HandCards_3()
	Proxy4Lua.SelectCard4Ready(3);
end
function HandCards_4()
	Proxy4Lua.SelectCard4Ready(4);
end--]]

require "FairyGUI"

local BattleWindow = fgui.window_class(WindowBase)

local bagWindow;
local List;

function Init()
	bagWindow = BattleWindow.New();
	bagWindow:Show();
end

function BattleWindow:OnInit()
	self.contentPane = UIPackage.CreateObject("Battle", "BattlePanel").asCom;
	self:Center();
	self.modal = true;

	List = self.contentPane:GetChild("list").asList;
	List.onClickItem:Add(__clickItem);
	List.itemRenderer = RenderListItem;
	List.numItems = 45;
end

function RenderListItem(index, button)
	button.icon = "i" .. UnityEngine.Random.Range(0, 10);
	button.title = "" .. UnityEngine.Random.Range(0, 100);
end

function __clickItem(context)
	local item = context.data;
	bagWindow.contentPane:GetChild("n11").asLoader.url = item.icon;
	bagWindow.contentPane:GetChild("n13").text = item.icon;
end
--]]

require "FairyGUI"

BattlePanel = fgui.window_class(WindowBase)
local Window;

local autoBtn;
local stateIcon;

local cards = {};

function BattlePanel:OnEntry()
	Window = BattlePanel.New();
	Window:Show();
end

function BattlePanel:OnInit()
	UIPackage.AddPackage("UI/UI_Fairy/export/Battle");
	self.contentPane = UIPackage.CreateObject("Battle", "BattlePanel").asCom;
	self:Center();

	local returnBtn = self.contentPane:GetChild("n8").asButton;
	returnBtn.onClick:Add(OnReturnBtn);

	autoBtn = self.contentPane:GetChild("n12").asButton;
	autoBtn.onClick:Add(OnAutoBtn);

	stateIcon = self.contentPane:GetChild("n16").asLoader;
	stateIcon.onClick:Add(OnTurnOver);

	for i=1, 5 do
		cards[i] = {};
		cards[i]["card"] = self.contentPane:GetChild("n" .. (16 + i)).asCom;
		cards[i]["power"] = cards[i]["card"]:GetChild("power");
		cards[i]["cost"] = cards[i]["card"]:GetChild("cost");
	end

	FlushData();
end

function BattlePanel:OnUpdate()
	if UIManager.IsDirty("BattlePanel") then
		FlushData();
		UIManager.ClearDirty("BattlePanel");
	end
end

function BattlePanel:OnTick()
	
end

function BattlePanel:isShow()
	return Window.isShowing;
end

function BattlePanel:OnDispose()
	Window:Dispose();
end

function FlushData()
	if Battle._CurrentState == Battle.BattleState.BS_Oper then
		stateIcon.url = UIPackage.GetItemURL("Battle", "battle_jieshuhuihe");
		stateIcon.touchable = true;
	else
		stateIcon.url = UIPackage.GetItemURL("Battle", "battle_dengdaizhong");
		stateIcon.touchable = false;
	end

	local cardNum = Battle._LeftCardNum;
	for i=1, 5 do
		if i <= cardNum then
			cards[i]["card"].data = i;
			cards[i]["card"].onClick:Add(OnCardClick);
			cards[i]["power"].text = i;
			cards[i]["cost"].text = i;
			cards[i]["card"].visible = true;
			if Battle._Turn == 1 and not Battle.IsSelfCard(i-1) then
				cards[i]["card"].enabled = false;
			else
				cards[i]["card"].enabled = true;
			end
		else
			cards[i]["card"].visible = false;
		end
	end
end

function OnCardClick(context)
	print(context.sender.data);
	Proxy4Lua.SelectCard4Ready(context.sender.data - 1);
end

function OnReturnBtn()
	print("OnReturnBtn");
	SceneLoader.LoadScene("main");
end

function OnTurnOver()
	print("OnTurnOver");
	Proxy4Lua.BattleSetup();
end

function OnAutoBtn()
	print("OnAutoBtn");
	GamePlayer._IsAuto = not GamePlayer._IsAuto;
	if GamePlayer._IsAuto then
		autoBtn:GetController("icon").selectedIndex = 1;
	else
		autoBtn:GetController("icon").selectedIndex = 0;
	end
end