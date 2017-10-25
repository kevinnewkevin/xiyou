require "FairyGUI"

daguanka = fgui.window_class(WindowBase)
local Window;

local cardGroupList;
local infoPanel;
local infoName;
local infoDesc;
local leftBtn;
local rightBtn;
local teamBtn;
local boosInfoBtn;
local starBtn;
local rewardIcon;
local crtTab;
local guankaId = 0;
local assetArr;

function daguanka:OnEntry()
	Define.LaunchUIBundle("guankatupian");
	Window = daguanka.New();
	Window:Show();
end

function daguanka:OnInit()
	self.contentPane = UIPackage.CreateObject("daguanka", "daguanka_com").asCom;
	self:Center();
	self.modal = true;
	assetArr = {};
	self.closeButton = self.contentPane:GetChild("n3").asButton;
	infoPanel = self.contentPane:GetChild("n4");
	infoName = infoPanel:GetChild("n13");
	infoDesc = infoPanel:GetChild("n12");
	boosInfoBtn = infoPanel:GetChild("n4");
	leftBtn = self.contentPane:GetChild("n5");
	rightBtn = self.contentPane:GetChild("n6");
	leftBtn.onClick:Add(daguanka_OnLeftBtn);
	rightBtn.onClick:Add(daguanka_OnRightBtn);
	teamBtn = self.contentPane:GetChild("n7");
	teamBtn.onClick:Add(daguanka_OnTeamBtn);
	boosInfoBtn.onClick:Add(daguanka_OnBoosInfo);
	rewardIcon = infoPanel:GetChild("n16");
	starBtn = infoPanel:GetChild("n3");
	starBtn.onClick:Add(daguanka_OnStart);
	infoPanel.visible  = false;
	local dList = self.contentPane:GetChild("n12");
	cardGroupList = dList:GetChild("n11").asList;
	cardGroupList:SetVirtualAndLoop();
	cardGroupList.itemRenderer = daguanka_RenderListItem;
	cardGroupList.scrollPane.onScroll:Add(DoSpecialEffect);
	cardGroupList.onTouchBegin:Add(daguanka_ListTouchBegin);
	daguanka_FlushData();
	DoSpecialEffect();
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
				local ss = 1 + (1 - dist / obj.width) * 0.5;
				obj:SetScale(ss, ss);

				--obj:Center();
			end
		end
		--_mainView.GetChild("n3").text = "" + ((cardGroupList.GetFirstChildInView() + 1) % cardGroupList.numItems);
end



function daguanka:GetWindow()
	return Window;
end

function daguanka:OnUpdate()
	if UIManager.IsDirty("daguanka") then
		daguanka_FlushData();
		UIManager.ClearDirty("daguanka");
	end
end



function daguanka_RenderListItem(index, obj)
	local comData;
	if crtTab == 1 then
		comData  = JieHunSystem.instance.ChapterEasyDataList[0];
	else
		comData  = JieHunSystem.instance.ChapterHardDataList[0];
	end

	local data = HeroStroyData.GetData(comData.ChapterId);
	local entityData = EntityData.GetData(data.EntityID_);
	local displayData = DisplayData.GetData(entityData._DisplayId);
	obj:SetPivot(0.5, 0.5);
	local mode = obj:GetChild("n8");
	local modelRes = displayData._AssetPath;
	mode:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes, false));
	obj.data = comData .ChapterId;
	obj.onClick:Set(daguanka_OnSelectGroup);
	if assetArr ~= nil then
		assetArr[index] = modelRes;
	end
end

function daguanka_OnSelectGroup(context)
	guankaId = context.sender.data;
	--Proxy4Lua.RequestChapterData(guankaId );
	--UIManager.Show("guanka");
	local data = HeroStroyData.GetData(guankaId);
	infoPanel.visible  = true;
	infoName.text = data.Name_;
	infoDesc.text = data.Desc_;
	local entityData = EntityData.GetData(data.EntityID_);
	local displayData = DisplayData.GetData(entityData._DisplayId);
	rewardIcon.asLoader.url = "ui://" .. displayData._HeadIcon;


	daguanka_FlushData();
end

function  daguanka_OnLeftBtn(context)
	infoPanel.visible  = false;

	---local midX = cardGroupList.scrollPane.posX + cardGroupList.viewWidth / 2;
	--local num = (cardGroupList:GetFirstChildInView() + 1) % cardGroupList.numItems;
	--local obj = cardGroupList:GetChildAt(num);
	--local dist = Mathf.Abs(midX - obj.x - obj.width / 2);

	cardGroupList.scrollPane:SetPosX(cardGroupList.scrollPane.posX + 535,true);
	print(cardGroupList.scrollPane.posX );
end

function daguanka_OnRightBtn(context) 
	infoPanel.visible  = false;
	cardGroupList.scrollPane:SetPosX(cardGroupList.scrollPane.posX - 535,true);
	print(cardGroupList.scrollPane.posX );
end


function daguanka_OnBoosInfo(context)
	local data = HeroStroyData.GetData(guankaId);
	UIParamHolder.Set("showBoosId", data.EntityID_);
	UIParamHolder.Set("showBoos", true);
	UIManager.Show("xiangxiziliao");
end

function daguanka_OnTeamBtn(context)
	Window:Hide();
	UIManager.Show("paiku");
end


function daguanka_OnStart(context)
	Proxy4Lua.RequestChapterData(guankaId );
	UIManager.Show("guanka");
end

function daguanka_ListTouchBegin(context)
	infoPanel.visible  = false;
end

function daguanka:GetGuankaId()
	return guankaId;
end

function daguanka:OnTick()
	
end

function daguanka:isShow()
	return Window.isShowing;
end

function daguanka:OnDispose()
	Window:Dispose();
end

function daguanka:OnHide()
	infoPanel.visible  = false;
	local k;
	local v;
	for k,v in ipairs(assetArr) do
		Proxy4Lua.ForceUnloadAsset(v);
	end
	assetArr = nil;
	print('del  assetArr');
	Window:Hide();
end

function daguanka_FlushData()
	if assetArr == nil then
		assetArr = {};	
	end
	local a = cardGroupList.scrollPane.scrollStep ;
	if crtTab == 1 then
		cardGroupList.numItems = 10;-- JieHunSystem.instance.ChapterEasyDataList.Count;
	else
		cardGroupList.numItems = 10;-- JieHunSystem.instance.ChapterHardDataList.Count;
	end
end


