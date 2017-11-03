require "FairyGUI"

baowu = fgui.window_class(WindowBase)
local Window;

local hitNextLbl;
local hitNextBtn;

local dropList;
local openTrans;
local outTrans;

local openCom;
local outCom;

local dropItemUrl = "url://baowu/daoju_com";

local readyOutDrop;
local canExit;

local timeOut;
local boxEff;
local boxEffHolder;

function baowu:OnEntry()
	Define.LaunchUIBundle("icon");
	Window = baowu.New();
	Window:Show();
end

function baowu:GetWindow()
	return Window;
end

function baowu:OnInit()
	self.contentPane = UIPackage.CreateObject("baowu", "baowu_com").asCom;
	self:Center();
	self.modal = true;
	hitNextLbl = self.contentPane:GetChild("n0").asTextField;
	hitNextBtn = self.contentPane:GetChild("n9");
	hitNextBtn.onClick:Add(baowu_OnExit);
	hitNextBtn.enabled = false;
	hitNextLbl.visible = false;

	dropList = self.contentPane:GetChild("n2").asList;
	dropList.visible = false;
	openCom = self.contentPane:GetChild("n4").asCom;
	outCom = self.contentPane:GetChild("n7").asCom;
	outCom.visible = false;
	boxEffHolder = outCom:GetChild("n2").asGraph;
	openTrans = self.contentPane:GetTransition("t1");
	outTrans = outCom:GetTransition("t0");

	openCom.onClick:Add(baowu_OnBoxBtn);

	timeOut = {};
	timeOut.max = 2;
	timeOut.count = 0;

	baowu_FlushData();
end

function baowu_OnBoxBtn(context)
	timeOut = nil;
	outCom.visible = true;
	openCom.visible = false;
	outTrans:Play();

	readyOutDrop = {};
	readyOutDrop.max = 12;
	readyOutDrop.count = 0;

	boxEff = "effect/baoxiangguangmang";
	boxEffHolder:SetNativeObject(Proxy4Lua.GetAssetGameObject(boxEff, false));
end

function baowu_OnExit(context)
	local showChapters = UIParamHolder.Get("showChaptersDrop");
	if showChapters ==true then
		openCom.visible = true;
		hitNextBtn.enabled = false;
		hitNextLbl.visible = false;
		dropList.visible = false;
		outCom.visible = false;
		--UIManager.Show("jiehun");
		Window:Hide();			
	else
		if GamePlayer.showNewCard == true then
			UIManager.Show("huoderenwu");
			GamePlayer.showNewCard = false;
		else
			SceneLoader.LoadScene("main");
		end
	end 
	Proxy4Lua.UnloadAsset(boxEff);
	boxEff = "";
	boxEffHolder:SetNativeObject(Proxy4Lua.GetAssetGameObject("", false));
end

function baowu_FlushData()
	local dropItem; 
	local dropItemCount;
	local itemInst;
	local stack;
	local name;
	local quality;
	local icon;
	local iData;

	local showChapters = UIParamHolder.Get("showChaptersDrop");
	if showChapters == true then
		local dropId= UIParamHolder.Get("chaptersDrop");
		local drop = DropData.GetData(dropId);
		dropItemCount = 1;
		dropList:RemoveChildrenToPool();
		local cDropItem = dropList:AddItemFromPool(dropItemUrl);
		local cIData = ItemData.GetData(drop.item1_);
		local cQuality = cDropItem:GetChild("n0").asLoader;
		local cIcon = cDropItem:GetChild("n1").asLoader;
		local cStack = cDropItem:GetChild("n2").asTextField;
		local cName = cDropItem:GetChild("n3").asTextField;

		cStack.text = drop.itemNum1_ .. "";
		if cIData ~= nil then
			cName.text = cIData._Name;
			cIcon.url = "ui://" .. cIData._Icon;
			cQuality.url = "ui://" .. cIData._IconBack;
		end
	else
		dropItemCount = Battle.DropItemCount;
		dropList:RemoveChildrenToPool();
		for i=1, dropItemCount do
			itemInst = Battle.DropItem(i-1);
			dropItem = dropList:AddItemFromPool(dropItemUrl);
			if itemInst ~= nil then
				iData = ItemData.GetData(itemInst.ItemId);
				quality = dropItem:GetChild("n0").asLoader;
				icon = dropItem:GetChild("n1").asLoader;
				stack = dropItem:GetChild("n2").asTextField;
				name = dropItem:GetChild("n3").asTextField;

				stack.text = itemInst.Stack;
				if iData ~= nil then
					name.text = iData._Name;
					icon.url = "ui://" .. iData._Icon;
					quality.url = "ui://" .. iData._IconBack;
				end
			end
		end
	end
end

function baowu:OnUpdate()
	if UIManager.IsDirty("baowu") then
		baowu_FlushData();
		UIManager.ClearDirty("baowu");
	end

	if readyOutDrop ~= nil then
		readyOutDrop.count = readyOutDrop.count + 1;
		if readyOutDrop.count >= readyOutDrop.max then
			dropList.visible = true;
			openTrans:Play();
			readyOutDrop = nil;
			canExit = {};
			canExit.max = 20;
			canExit.count = 0;
		end
	end

	if canExit ~= nil then
		canExit.count = canExit.count + 1;
		if canExit.count >= canExit.max then
			hitNextLbl.visible = true;
			hitNextBtn.enabled = true;

			canExit = nil;
		end
	end
end

function baowu:OnTick()
	if timeOut ~= nil then
		timeOut.count = timeOut.count + 1;
		if timeOut.count >= timeOut.max then
			baowu_OnBoxBtn();
			timeOut = nil;
		end
	end
end

function baowu:isShow()
	return Window.isShowing;
end

function baowu:OnDispose()
	Window:Dispose();
end

function baowu:OnHide()
	openCom.visible = true;
	hitNextBtn.enabled = false;
	hitNextLbl.visible = false;
	dropList.visible = false;
	outCom.visible = false;
	timeOut = {};
	timeOut.max = 2;
	timeOut.count = 0;
	Window:Hide();
end