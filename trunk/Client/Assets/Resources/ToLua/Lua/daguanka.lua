require "FairyGUI"

daguanka = fgui.window_class(WindowBase)
local Window;

local cardGroupList;
local infoPanel;
local infoName;
local infoDesc;
local leftBtn;
local rightBtn;
local playerName;
local boosInfoBtn;
local starBtn;
local rewardIcon;
local crtTab;
local guankaId = 0;
local ready;
local Trans0;
local Trans1;
local Trans2;
local crtSelectIdx;

function daguanka:OnEntry()
	Define.LaunchUIBundle("guankatupian");
	Window = daguanka.New();
	Window:Show();
end

function daguanka:OnInit()
	self.contentPane = UIPackage.CreateObject("daguanka", "daguanka_com").asCom;
	self:Center();
	self.modal = true;
	crtTab = 1;
	self.closeButton = self.contentPane:GetChild("n3").asButton;
	infoPanel = self.contentPane:GetChild("n4");
	infoName = infoPanel:GetChild("n13");
	infoDesc = infoPanel:GetChild("n12");
	boosInfoBtn = infoPanel:GetChild("n4");
	leftBtn = self.contentPane:GetChild("n5");
	rightBtn = self.contentPane:GetChild("n6");
	leftBtn.onClick:Add(daguanka_OnLeftBtn);
	rightBtn.onClick:Add(daguanka_OnRightBtn);
	playerName = self.contentPane:GetChild("n17");
	boosInfoBtn.onClick:Add(daguanka_OnBoosInfo);
	rewardIcon = infoPanel:GetChild("n16");
	starBtn = infoPanel:GetChild("n3");
	starBtn.onClick:Add(daguanka_OnStart);
	infoPanel.visible  = false;

	Trans0 = self.contentPane:GetTransition("t0");
	Trans1 = self.contentPane:GetTransition("t1");
	Trans2 = self.contentPane:GetTransition("t2");

	local dList = self.contentPane:GetChild("n12");
	cardGroupList = dList:GetChild("n11").asList;
	cardGroupList:SetVirtualAndLoop();
	cardGroupList.itemRenderer = daguanka_RenderListItem;
	cardGroupList.scrollPane.onScroll:Add(DoSpecialEffect);
	--cardGroupList.onTouchBegin:Add(daguanka_ListTouchBegin);
	daguanka_FlushData();
	DoSpecialEffect();
	crtSelectIdx = 0;
end



function DoSpecialEffect()
		local midX = cardGroupList.scrollPane.posX + cardGroupList.viewWidth / 2;
		local cnt = cardGroupList.numChildren;
		for  i = 1, cnt do
			local obj = cardGroupList:GetChildAt(i-1);
			local dist = Mathf.Abs(midX - obj.x - obj.width / 2);
			if dist > obj.width then
				obj:SetScale(1, 1);
				obj.onClick:Remove(daguanka_OnSelectGroup);
				obj.visible = false;
			else
				local ss = 1 + (1 - dist / obj.width) * 0.5;
				obj:SetScale(ss, ss);
				local data = HeroStroyData.GetData(obj.data);
				local entityData = EntityData.GetData(data.EntityID_);
				playerName.text = entityData._Name;
				obj.onClick:Set(daguanka_OnSelectGroup);

				obj.visible = true; 
				local data = HeroStroyData.GetData(obj.data);
				local entityData = EntityData.GetData(data.EntityID_);
				local displayData = DisplayData.GetData(entityData._DisplayId);
				local mode = obj:GetChild("n8");
				local modelRes = displayData._AssetPath;
				mode:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes, false));
				Proxy4Lua.AddToDelete(modelRes);
				infoPanel.visible  = true;
				infoName.text = data.Name_;
				infoDesc.text = data.Desc_;
				rewardIcon.asLoader.url = "ui://" .. displayData._HeadIcon;
				guankaId = obj.data;
				local cData = JieHunSystem.instance:GetChapterData(guankaId);
				if  cData == nil then
					starBtn.visible = false;
				else
					starBtn.visible = true;
				end 
			end
		end
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

	if index == 0 then
		index = HeroStroyData.GetEasyListNum() -1;
	else
		index = index -1
	end 

	local data = HeroStroyData.easyList[index]; 
	--local entityData = EntityData.GetData(data.EntityID_);
	--local displayData = DisplayData.GetData(entityData._DisplayId);
	obj:SetPivot(0.5, 0.5);
--	local mode = obj:GetChild("n8");
--	local modelRes = displayData._AssetPath;
	--mode:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes, false));
	obj.data = data.Id_;-- comData.ChapterId;
	--Proxy4Lua.AddToDelete(modelRes);
end

function daguanka_OnSelectGroup(context)
	guankaId = context.sender.data;
	local data = HeroStroyData.GetData(guankaId);
	--infoPanel.visible  = true;
	--Trans0:Play();
	infoName.text = data.Name_;
	infoDesc.text = data.Desc_;
	local entityData = EntityData.GetData(data.EntityID_);
	local displayData = DisplayData.GetData(entityData._DisplayId);
	rewardIcon.asLoader.url = "ui://" .. displayData._HeadIcon;
	local cData = JieHunSystem.instance:GetChapterData(guankaId);
	if  cData == nil then
		starBtn.visible = false;
	else
		starBtn.visible = true;
	end 

	daguanka_FlushData();
end

function  daguanka_OnLeftBtn(context)
	--infoPanel.visible  = false;
	crtSelectIdx = (crtSelectIdx - 1) % cardGroupList.numItems;
	cardGroupList:ScrollToView(crtSelectIdx, false);
	print(crtSelectIdx);
end

function daguanka_OnRightBtn(context) 
--	infoPanel.visible  = false;
	crtSelectIdx = (crtSelectIdx + 1) % cardGroupList.numItems;
	cardGroupList:ScrollToView(crtSelectIdx, false);
	print(crtSelectIdx);
end


function daguanka_OnBoosInfo(context)
	local data = HeroStroyData.GetData(guankaId);
	UIParamHolder.Set("showBoosId", data.EntityID_);
	UIParamHolder.Set("showBoos", true);
	UIManager.Show("xiangxiziliao");
end


function daguanka_OnStart(context)
	Proxy4Lua.RequestChapterData(guankaId );

	ready = {};
	ready.max = 1;
	ready.count = 0;
	Trans2:Play();
	--UIManager.Show("xiaoguanka");
end

function daguanka_ListTouchBegin(context)
	--infoPanel.visible  = false;
end

function daguanka:GetGuankaId()
	return guankaId;
end

function daguanka:OnTick()
	if ready ~= nil and ready.count < ready.max  then
		ready.count = ready.count + 1;
		if ready.count >= ready.max then
			UIManager.Show("xiaoguanka");
		end
	end
end

function daguanka:isShow()
	return Window.isShowing;
end

function daguanka:OnDispose()
	Window:Dispose();
end

function daguanka:OnHide()
	--infoPanel.visible  = false;
	Proxy4Lua.ClearToDeleteAsset();
	Window:Hide();
end

function daguanka_FlushData()
		cardGroupList.numItems = HeroStroyData.GetEasyListNum();
end


