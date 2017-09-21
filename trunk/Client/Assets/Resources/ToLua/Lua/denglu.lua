require "FairyGUI"

denglu = fgui.window_class(WindowBase)
local Window;

local selectInfo;
local enterBtn;
local selectServ;
local selectServCloseBtn;

local accountGroup;
local account;
local password;

function denglu:OnEntry()
	Window = denglu.New();
	Window:Show();
end

function denglu:GetWindow()
	return Window;
end

function denglu:OnInit()
	self.contentPane = UIPackage.CreateObject("denglu", "denglu_com").asCom;
	self:Center();

	enterBtn = self.contentPane:GetChild("n3").asButton;
	enterBtn.onClick:Add(denglu_OnEnterGame);

	selectInfo = self.contentPane:GetChild("n9");
	selectInfo.onClick:Add(denglu_OnShowServ);

	selectServ = self.contentPane:GetChild("n8").asCom;
	selectServ.visible = false;
	selectServCloseBtn = selectServ:GetChild("n5").asButton;
	selectServCloseBtn.onClick:Add(denglu_OnCloseServ);

	accountGroup = self.contentPane:GetChild("n10").asCom;
	account = accountGroup:GetChild("n11").asTextField;
	password = accountGroup:GetChild("n12").asTextField;

	denglu_FlushData();
end

function denglu:OnUpdate()
	if UIManager.IsDirty("denglu") then
		denglu_FlushData();
		UIManager.ClearDirty("denglu");
	end
end

function denglu:OnTick()
	
end

function denglu:isShow()
	return Window.isShowing;
end

function denglu:OnDispose()
	Window:Dispose();
end

function denglu:OnHide()
	Window:Hide();
end

function denglu_FlushData()
	enterBtn.enabled = not DataLoader._IsLoading;
end

function denglu_OnShowServ()
	selectServ.visible = true;
end

function denglu_OnCloseServ()
	selectServ.visible = false;
end

function denglu_OnEnterGame()
	if account.text == nil or account.text == "" then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "输入账号", true);
		return;
	end
	if Proxy4Lua.ReconnectServer() == true then
		Proxy4Lua.Login(account.text, password.text);
	end
end