require "FairyGUI"

denglu = fgui.window_class(WindowBase)
local Window;

local selectInfo;
local loginBtn;
local enterBtn;
local selectServ;
local selectServCloseBtn;

local crtServer;
local crtServerText;
local serverList;
local serverItem = "ui://denglu/xuanfu_Button";

local accountGroup;
local account;
local password;

local finalized;

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
	serverList = selectServ:GetChild("n10").asList;
	crtServer = self.contentPane:GetChild("n13").asButton;
	crtServer.onClick:Add(denglu_OnShowServ);
	crtServerText = crtServer:GetChild("n3");
	crtServer.visible = false;

	accountGroup = self.contentPane:GetChild("n10").asCom;
	account = accountGroup:GetChild("n11").asTextField;
	password = accountGroup:GetChild("n12").asTextField;
	loginBtn = accountGroup:GetChild("n13").asButton;
	loginBtn.onClick:Add(denglu_OnLogin);

	enterBtn.enabled = false;
	loginBtn.enabled = false;

	finalized = false;

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
	accountGroup.visible = true;
	acc = "";
	crtServer.visible = false;
	finalized = false;
	Window:Hide();
end

function denglu_ServerSelect(context)
	local servInfo = Proxy4Lua._ServList[context.sender.data];
	Proxy4Lua._ServerIP = servInfo.serverIP;
	crtServerText.text = servInfo.serverName;
	selectServ.visible = false;
end

function denglu_FlushData()
	if Proxy4Lua._ServerIP == nil or Proxy4Lua._ServerIP == "" then
		crtServerText.text = "选择服务器";
	end
	if Proxy4Lua._ServList ~= nil then
		serverList:RemoveChildrenToPool();
		for i=0, Proxy4Lua._ServList.Count - 1 do
			local si = serverList:AddItemFromPool(serverItem);
			si:GetChild("n3").text = Proxy4Lua._ServList[i].serverName;
			si:GetChild("n4").text = Proxy4Lua._ServList[i].serverIP;
			si.data = i;
			si.onClick:Add(denglu_ServerSelect);
		end
	end

	if finalized == true then
		return;
	end
	if DataLoader._LoadFinish then
		DataLoader._LoadFinish = false;
		Proxy4Lua.ResUpdate._UpdateFinish = false;
		print("_LoadFinish");
		enterBtn.enabled = true;
		loginBtn.enabled = true;
		finalized = true;
		return;
	end

	if Proxy4Lua.ResUpdate._UpdateFinish and not DataLoader._IsLoading then
		--读表
		DataLoader.BeginLoad();
		print("读表");
	end

	if Proxy4Lua.ResUpdate._IsDoing == false and not Proxy4Lua.ResUpdate._UpdateFinish then
		--监测更新
		Proxy4Lua.ResUpdate:Init();
		print("监测更新");
	end
end

function denglu_OnShowServ()
	selectServ.visible = true;
end

function denglu_OnCloseServ()
	selectServ.visible = false;
end

local acc;
function denglu_OnLogin()
	if account.text == nil or account.text == "" then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "输入账号", true);
		return;
	end

	acc = account.text;
	accountGroup.visible = false;
	selectServ.visible = true;
	Proxy4Lua.CheckServer();
	crtServer.visible = true;
end

function denglu_OnEnterGame()
	if Proxy4Lua.ReconnectServer() == true then
		Proxy4Lua.Login(acc, password.text);
		enterBtn.enabled = false;
		loginBtn.enabled = false;
	end
end