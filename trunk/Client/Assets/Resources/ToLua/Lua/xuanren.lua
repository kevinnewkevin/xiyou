require "FairyGUI"

xuanren = fgui.window_class(WindowBase)
local Window;

local name;
local createBtn;
local selectList;

function xuanren:OnEntry()
	Window = denglu.New();
	Window:Show();
end

function xuanren:GetWindow()
	return Window;
end

function denglu:OnInit()
	self.contentPane = UIPackage.CreateObject("xuanren", "xuanren_com").asCom;
	self:Center();

	name = self.contentPane:GetChild("n9");
	createBtn = self.contentPane:GetChild("n7");
	selectList = self.contentPane:GetChild("n10");

	createBtn.onClick:Add(xuanren_OnCreate);
end

function xuanren:OnUpdate()
	
end

function xuanren:OnTick()
	
end

function xuanren:isShow()
	return Window.isShowing;
end

function xuanren:OnDispose()
	Window:Dispose();
end

function xuanren:OnHide()
	Window:Hide();
end

function xuanren_OnCreate()
	if selectList.selectedIndex == -1 then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "请选择角色", true);
		return;
	end

	if name.text == "" then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "请输入昵称", true);
		return;
	end
	Proxy4Lua.CreatePlayer(selectList.selectedIndex + 1, name.text);
end