require "FairyGUI"

kaikabao = fgui.window_class(WindowBase)
local Window;

local hitNextLbl;
local hitNextBtn;

local dropList;
local openTrans;
local outTrans;

local openCom;
local outCom;
local backCom;

local dropItemUrl = "url://kaikabao/daoju_com";

local readyOutDrop;
local canExit;

local timeOut;
local overEffStr;
local boxEff;
local boxEffHolder;

local itemList;
local isInitItemList;

local cardOutDrop;

local showBuyItem;
local showPayBoxEffect;
local showOuTCard;
local topPanel;
local overPanel;
local overTrans;
local overEffect;

local itemInfoName;
local itemInfolevel;
local itemInfoIcon;
local itemInfoIcon_Icon;
local itemInfoIconDebris;
local itemInfoIconNum;
local itemInfoIconTeam;
local itemInfoIconPower;
local itemInfoIconlevel;
local itemInfoIconRed;
local barLab;
local itemBar;
local overShowBox;
function kaikabao:OnEntry()
	Define.LaunchUIBundle("icon");
	Window = kaikabao.New();
	Window:Show();
end

function kaikabao:GetWindow()
	return Window;
end

function kaikabao:OnInit()
	self.contentPane = UIPackage.CreateObject("kaikabao", "kaikabao_com").asCom;
	self:Center();

	hitNextLbl = self.contentPane:GetChild("l3").asTextField;
	topPanel = self.contentPane:GetChild("n8");
	backCom = self.contentPane:GetChild("n23");
	backCom.onClick:Add(kaikabao_OnBackBtn);
	backCom.enabled = false;
	hitNextLbl.visible = false;
	itemInfoName = topPanel:GetChild("n15");
	itemInfolevel = topPanel:GetChild("n14");
	itemInfoIcon = topPanel:GetChild("n8");
	barLab = topPanel:GetChild("n12");
	itemBar = topPanel:GetChild("n11");
	itemInfoIcon_Icon = itemInfoIcon:GetChild("n5");
	itemInfoIconDebris = itemInfoIcon:GetChild("n11");
	itemInfoIconNum = itemInfoIcon:GetChild("n12");
	itemInfoIconTeam = itemInfoIcon:GetChild("n9");
	itemInfoIconPower = itemInfoIcon:GetChild("n7");
	itemInfoIconlevel = itemInfoIcon:GetChild("n6");
	itemInfoIconRed = itemInfoIcon:GetChild("n10");


	itemInfoIconTeam.visible= false;
	itemInfoIconRed.visible= false;
	itemInfoIconDebris.visible= false; 
	topPanel.visible= false; 

	overPanel = self.contentPane:GetChild("n22");
	overShowBox = overPanel:GetChild("n19");
	overTrans = overPanel:GetTransition("t2");
	dropList = overPanel:GetChild("n20").asList;
	hitNextBtn = overPanel:GetChild("n22");
	overEffect = overPanel:GetChild("n21").asGraph;
	hitNextBtn.onClick:Add(kaikabao_OnExit);
	overPanel.visible= false; 


	openCom = self.contentPane:GetChild("n0").asCom;
	outCom = self.contentPane:GetChild("n3");
	outCom.visible = false;
	boxEffHolder =  self.contentPane:GetChild("n19").asGraph;
	openTrans = self.contentPane:GetTransition("t1");
--	outTrans = outCom:GetTransition("t0");

	---openCom.onClick:Add(kaikabao_OnBoxBtn);


	timeOut = {};
	timeOut.max =2;
	timeOut.count = 0;
	showBuyItem = false;
	showPayBoxEffect = false;
	showOuTCard = false;
	isInitItemList = false;

	cardOutDrop = {};
	cardOutDrop.max = 1;
	cardOutDrop.count = 0; 
	showOuTCard = true;

	kaikabao_FlushData();
end


function kaikabao_OnExit(context)

	topPanel.visible= false; 
	overPanel.visible= false; 
	openCom.visible = true;
	outCom.visible = false;
	isInitItemList = false;
	cardOutDrop = {};
	cardOutDrop.max = 1;
	cardOutDrop.count = 0; 
	showOuTCard = true;

	Proxy4Lua.UnloadAsset(boxEff);
	boxEff = "";
	Proxy4Lua.UnloadAsset(overEffStr);
	overEffStr = "";
	boxEffHolder:SetNativeObject(Proxy4Lua.GetAssetGameObject("", false));
	overEffect:SetNativeObject(Proxy4Lua.GetAssetGameObject("", false));
	Window:Hide();			


end

function kaikabao_FlushData()
	local dropItem;
	local dropItemCount;
	local itemInst;
	local stack;
	local name;
	local quality;
	local icon;
	local iData;
	local shopType = ShopSystem.buyType;
	local btnImg = openCom:GetChild("n0");
	if shopType == 1000 then
		btnImg.asLoader.url  = "ui://kaikabao/kabao1";
		outCom.asLoader.url  = "ui://kaikabao/kabao1";
		overShowBox.asLoader.url  = "ui://kaikabao/kabao1";
	end
	if shopType == 1001 then
		btnImg.asLoader.url  = "ui://kaikabao/kabao2";
		outCom.asLoader.url  = "ui://kaikabao/kabao2";
		overShowBox.asLoader.url  = "ui://kaikabao/kabao2";
	end
	if shopType == 1002 then
		btnImg.asLoader.url  = "ui://kaikabao/kabao3";
		outCom.asLoader.url  = "ui://kaikabao/kabao3";
		overShowBox.asLoader.url  = "ui://kaikabao/kabao3";
	end


	local showChapters = UIParamHolder.Get("showChaptersDrop");

	dropItemCount = ShopSystem.BuyItems.Length;



	initItemList();

	dropList:RemoveChildrenToPool();
	for i=1, dropItemCount do
		itemInst = ShopSystem.BuyItems[i-1];
		dropItem = dropList:AddItemFromPool(dropItemUrl);
		if itemInst ~= nil then
			iData = ItemData.GetData(itemInst.ItemId);
			quality = dropItem:GetChild("n0").asLoader;
			icon = dropItem:GetChild("n1").asLoader;
			stack = dropItem:GetChild("n2").asTextField;
			--name = dropItem:GetChild("n3").asTextField;

			stack.text = itemInst.Stack;
			if iData ~= nil then
				--name.text = iData._Name;
				icon.url = "ui://" .. iData._Icon;
				quality.url = "ui://" .. iData._IconBack;
			end
		end
	end
end


function kaikabao_OnBoxBtn(context)
	print("kaikabao_OnBoxBtn");
	outCom.visible = true;
	openCom.visible = false;

	readyOutDrop = {};
	readyOutDrop.max = 1;
	readyOutDrop.count = 0;
	showPayBoxEffect = true;


end



function kaikabao:OnUpdate()

	if UIManager.IsDirty("kaikabao") then
		kaikabao_FlushData();
		UIManager.ClearDirty("kaikabao");
	end


	if readyOutDrop ~= nil and showPayBoxEffect == true then
		readyOutDrop.count = readyOutDrop.count + 1;
		if readyOutDrop.count >= readyOutDrop.max then

			boxEff = "effect/kaibao_liziguangdian";
			boxEffHolder:SetNativeObject(Proxy4Lua.GetAssetGameObject(boxEff, false));
			outCom.visible = true;
			openCom.visible = false;

			readyOutDrop = {};
			readyOutDrop.max = 1;
			readyOutDrop.count = 0;
			showPayBoxEffect = false;

			topPanel.visible = true;
			openTrans:Play();
			setItemInfo();
			backCom.enabled = true;
			hitNextLbl.visible = true;

		end
	end



--	if canExit ~= nil then
--		canExit.count = canExit.count + 1;
--		if canExit.count >= canExit.max then
--			hitNextLbl.visible = true;
--			hitNextBtn.enabled = true;
--			canExit = nil;
--		end
--	end
end


function kaikabao_OnBackBtn(context)
		if ShopSystem._ShowBuyItems.Count > 0 then
			topPanel.visible = false;
			outCom.visible = false;
			openCom.visible = true;

			cardOutDrop = {};
			cardOutDrop.max = 1;
			cardOutDrop.count = 0; 
			showOuTCard = true;

		else
			showOuTCard = false;
			overPanel.visible = true;
			overEffStr = "effect/kaibao_guangquan";
			overEffect:SetNativeObject(Proxy4Lua.GetAssetGameObject(overEffStr, false));
			overTrans:Play();
		end
		backCom.enabled = false;
		hitNextLbl.visible = false;
end


function kaikabao:OnTick()

	if cardOutDrop ~= nil and showOuTCard == true then
		cardOutDrop.count = cardOutDrop.count + 1;
		if cardOutDrop.count >= cardOutDrop.max then
			showOuTCard = false;
			outCom.visible = true;
			openCom.visible = false;
			readyOutDrop = {};
			readyOutDrop.max = 1;
			readyOutDrop.count = 0;
			showPayBoxEffect = true;
		end
	end

	--if timeOut ~= nil then
	--	timeOut.count = timeOut.count + 1;
	--	if timeOut.count >= timeOut.max then
	--		kaikabao_OnBoxBtn();
	--		timeOut = nil;
	--	end
	--end
end

function kaikabao:isShow()
	return Window.isShowing;
end

function kaikabao:OnDispose()
	Window:Dispose();
end

function kaikabao:OnHide()

	topPanel.visible= false; 
	overPanel.visible= false; 
	openCom.visible = true;
	outCom.visible = false;
	isInitItemList = false;
	cardOutDrop = {};
	cardOutDrop.max = 1;
	cardOutDrop.count = 0; 
	showOuTCard = true;
	Proxy4Lua.UnloadAsset(boxEff);
	boxEff = "";
	Proxy4Lua.UnloadAsset(overEffStr);
	overEffStr = "";
	Window:Hide();
end

function initItemList()
	if isInitItemList == true then
		return;
	end

  	itemList = ShopSystem.BuyItems;
	 isInitItemList = true; 
end

function setItemInfo()
	local itemInst = ShopSystem.GetShowItem();
	if itemInst ~= nil then
		local iData = ItemData.GetData(itemInst.ItemId);
		local eData = EntityData.GetData(iData._EntityID);
		itemInfoName.text = eData._Name; 
		local entityInst = GamePlayer.GetCardByEntityID(iData._EntityID);
		if entityInst ~= nil then
			itemInfolevel.text = entityInst.IProperties[9] .. "";
			local displayData = DisplayData.GetData(eData._DisplayId);
			itemInfoIcon_Icon.asLoader.url = "ui://" .. displayData._HeadIcon;
			itemInfoIconNum.text =  "" .. itemInst.Stack;
			itemInfoIconPower.text = eData._Cost;
			itemInfoIconlevel.text = entityInst.IProperties[9] .. "";
			local  levelData =  StrengthenData.GetData( entityInst.UnitId, entityInst.IProperties[9]+1);
			local itemNum = BagSystem.GetItemMaxNum(itemInst.ItemId);
			barLab.text = itemNum .. "/" .. levelData._ItemNum ;
			itemBar.value = 0;
			itemBar:TweenValue( (itemNum/levelData._ItemNum*100),20 );
		else
			itemInfoIconDebris.visible= true; 
			itemInfoIconNum.text =  "" .. itemInst.Stack;
			itemInfoIconlevel.text = "0";
		end

		ShopSystem.DelShowItem();
	else

	end
end

