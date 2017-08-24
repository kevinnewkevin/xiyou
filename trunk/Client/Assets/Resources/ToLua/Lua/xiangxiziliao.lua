require "FairyGUI"

xiangxiziliao = fgui.window_class(WindowBase)
local Window;

local holder;

local controller;
local isInGroup;

local fee;
local name;

function xiangxiziliao:OnEntry()
	Window = xiangxiziliao.New();
	Window:Show();
end

function xiangxiziliao:GetWindow()
	return Window;
end

function xiangxiziliao:OnInit()
	self.contentPane = UIPackage.CreateObject("xiangxiziliao", "xiangxiziliao_com").asCom;
	self:Center();
	self.modal = true;

	self.closeButton = self.contentPane:GetChild("n4");
	holder = self.contentPane:GetChild("n63").asGraph;

	local addGroupBtn = self.contentPane:GetChild("n61").asButton;
	addGroupBtn.onClick:Add(xiangxiziliao_OnAddGroup);
	controller = addGroupBtn:GetController("huang");

	fee = self.contentPane:GetChild("n58");
	name = self.contentPane:GetChild("n60");

	xiangxiziliao_FlushData();
end

function xiangxiziliao_OnAddGroup()
	xiangxiziliao_OnMessageConfirm();
	--[[local MessageBox = UIManager.ShowMessageBox();
	isInGroup = UIManager.GetWindow("paiku").IsInGroup();
	if isInGroup then
		MessageBox:SetData("提示", "是否取出卡组？", false, xiangxiziliao_OnMessageConfirm, xiangxiziliao_OnMessageCancel);
	else
		MessageBox:SetData("提示", "是否加入卡组？", false, xiangxiziliao_OnMessageConfirm, xiangxiziliao_OnMessageCancel);
	end--]]
end

function xiangxiziliao_OnMessageConfirm()
	local crtCardInstID = UIManager.GetWindow("paiku").GetCrtCard();
	local crtGroupIdx = UIManager.GetWindow("paiku").GetCrtGroup();
	isInGroup = UIManager.GetWindow("paiku").IsInGroup();
	if isInGroup then
		GamePlayer.TakeOffCard(crtCardInstID, crtGroupIdx);
	else
		GamePlayer.PutInCard(crtCardInstID, crtGroupIdx);
	end
	--UIManager.HideMessageBox();
	UIManager.Hide("xiangxiziliao");
	UIManager.SetDirty("paiku");
end

function xiangxiziliao_OnMessageCancel()
	UIManager.HideMessageBox();
	UIManager.SetDirty("paiku");
end

function xiangxiziliao:OnUpdate()
	if UIManager.IsDirty("xiangxiziliao") then
		xiangxiziliao_FlushData();
		UIManager.ClearDirty("xiangxiziliao");
	end
end

function xiangxiziliao:OnTick()
	
end

function xiangxiziliao:isShow()
	return Window.isShowing;
end

function xiangxiziliao:OnDispose()
	Window:Dispose();
end

function xiangxiziliao:OnHide()
	Window:Hide();
end

function xiangxiziliao_FlushData()
	local instId = UIManager.GetWindow("paiku").GetCrtCard();
	local displayData = GamePlayer.GetDisplayDataByInstID(instId);
	local modelRes = displayData._AssetPath;
	holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes));

	isInGroup = UIManager.GetWindow("paiku").IsInGroup();
	if isInGroup then
		controller.selectedIndex = 1;
	else
		controller.selectedIndex = 0;
	end

	local entityData = GamePlayer.GetEntityDataByInstID(instId);
	fee.text = entityData._Cost;
	name.text = entityData._Name;
end