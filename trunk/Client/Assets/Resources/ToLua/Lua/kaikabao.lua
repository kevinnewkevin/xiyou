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

local dropItemUrl = "url://kaikabao/daoju_com";

local readyOutDrop;
local canExit;

local timeOut;
local boxEff;
local boxEffHolder;

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

	hitNextLbl = self.contentPane:GetChild("n3").asTextField;
	hitNextBtn = self.contentPane:GetChild("n4");
	hitNextBtn.onClick:Add(kaikabao_OnExit);
	hitNextBtn.enabled = false;
	hitNextLbl.visible = false;

	dropList = self.contentPane:GetChild("n2").asList;
	dropList.visible = false;
	openCom = self.contentPane:GetChild("n0").asCom;
	outCom = self.contentPane:GetChild("n5").asCom;
	outCom.visible = false;
	boxEffHolder = outCom:GetChild("n6").asGraph;
	openTrans = self.contentPane:GetTransition("t1");
	outTrans = outCom:GetTransition("t0");

	openCom.onClick:Add(kaikabao_OnBoxBtn);


	timeOut = {};
	timeOut.max = 2;
	timeOut.count = 0;

	kaikabao_FlushData();
end

function kaikabao_OnBoxBtn(context)
	timeOut = nil;
	outCom.visible = true;
	openCom.visible = false;
	outTrans:Play();

	readyOutDrop = {};
	readyOutDrop.max = 12;
	readyOutDrop.count = 0;

	boxEff = "effect/baoxiangguangmang";
	boxEffHolder:SetNativeObject(Proxy4Lua.GetAssetGameObject(boxEff));
end

function kaikabao_OnExit(context)

	Window:Hide();			
	
	Proxy4Lua.UnloadAsset(boxEff);
	boxEff = "";
	boxEffHolder:SetNativeObject(Proxy4Lua.GetAssetGameObject(""));
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
	local btnImg1 = outCom:GetChild("n5");
	if shopType == 1000 then
		btnImg.asLoader.url  = "ui://kaikabao/kabao1";
		btnImg1.asLoader.url  = "ui://kaikabao/kabao1";
	end
	if shopType == 1001 then
		btnImg.asLoader.url  = "ui://kaikabao/kabao2";
		btnImg1.asLoader.url  = "ui://kaikabao/kabao2";
	end
	if shopType == 1002 then
		btnImg.asLoader.url  = "ui://kaikabao/kabao3";
		btnImg1.asLoader.url  = "ui://kaikabao/kabao3";
	end


	local showChapters = UIParamHolder.Get("showChaptersDrop");

		dropItemCount = ShopSystem.BuyItems.Length;
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

function kaikabao:OnUpdate()
	if UIManager.IsDirty("kaikabao") then
		kaikabao_FlushData();
		UIManager.ClearDirty("kaikabao");
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

function kaikabao:OnTick()
	if timeOut ~= nil then
		timeOut.count = timeOut.count + 1;
		if timeOut.count >= timeOut.max then
			kaikabao_OnBoxBtn();
			timeOut = nil;
		end
	end
end

function kaikabao:isShow()
	return Window.isShowing;
end

function kaikabao:OnDispose()
	Window:Dispose();
end

function kaikabao:OnHide()
	openCom.visible = true;
	hitNextBtn.enabled = false;
	hitNextLbl.visible = false;
	dropList.visible = false;
	outCom.visible = false;
	Window:Hide();
end