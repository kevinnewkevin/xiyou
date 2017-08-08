require "FairyGUI"

xiangxiziliao = fgui.window_class(WindowBase)
local Window;

local holder;

local controller;
local isInGroup;

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

	xiangxiziliao_FlushData();
end

function xiangxiziliao_OnAddGroup()
	local MessageBox = UIManager.ShowMessageBox();
	if isInGroup then
		MessageBox:SetData("提示", "是否取出卡组？", false, xiangxiziliao_OnMessageConfirm);
	else
		MessageBox:SetData("提示", "是否加入卡组？", false, xiangxiziliao_OnMessageConfirm);
	end
end

function xiangxiziliao_OnMessageConfirm()
	local index = UIParamHolder.Get("paiku_OnCardItem");
	if index ~= nil then
		local instid = GamePlayer.GetInstIDInMyCards(index);
	end

	local gindex = UIParamHolder.Get("paiku_OnCardInGroupGroupIdx");
	local cindex = UIParamHolder.Get("paiku_OnCardInGroupCardIdx");
	if gindex ~= nil and cindex ~= nil then
		local instid = GamePlayer.GetInstIDInMyGroup(gindex, cindex);
	end

	if isInGroup then
		print(" isInGroup ");
	else
		print(" not isInGroup ");
	end
	UIManager.HideMessageBox();
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
	local index = UIParamHolder.Get("paiku_OnCardItem");
	local instId;
	if index ~= nil then
		local modelRes = GamePlayer.GetResPathInMyCards(index);
		holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes));
		instId = GamePlayer.GetInstIDInMyCards(index);
	end

	local gindex = UIParamHolder.Get("paiku_OnCardInGroupGroupIdx");
	local cindex = UIParamHolder.Get("paiku_OnCardInGroupCardIdx");
	if gindex ~= nil and cindex ~= nil then
		local modelRes = GamePlayer.GetResPathInMyGroup(gindex, cindex);
		holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes));
		instId = GamePlayer.GetInstIDInMyGroup(gindex, cindex);
	end

	isInGroup = GamePlayer.IsInGroup(instId);
	if isInGroup then
	print("controller.selectedIndex = 1;");
		controller.selectedIndex = 1;
	else
	print("controller.selectedIndex = 0;");
		controller.selectedIndex = 0;
	end
end