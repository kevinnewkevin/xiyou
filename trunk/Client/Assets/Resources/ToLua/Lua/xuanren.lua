require "FairyGUI"

xuanren = fgui.window_class(WindowBase)
local Window;

local name;
local createBtn;
local selectList;
local randNameBtn;
local descCom;
local descAnim;

function xuanren:OnEntry()
	Window = xuanren.New();
	Window:Show();
end

function xuanren:GetWindow()
	return Window;
end

function xuanren:OnInit()
	self.contentPane = UIPackage.CreateObject("xuanren", "xuanren_com").asCom;
	self:Center();

	name = self.contentPane:GetChild("n9");
	createBtn = self.contentPane:GetChild("n7");
	selectList = self.contentPane:GetChild("n10");
	randNameBtn = self.contentPane:GetChild("n4");
	randNameBtn.onClick:Add(xuanren_RandName);
	descCom = self.contentPane:GetChild("n20");
	descCom.visible = false;
	descAnim = descCom:GetTransition("t0");

	selectList.onClickItem:Add(xuanren_OnSelectRole);

	name.text = RandNameData.Rand();

	createBtn.onClick:Add(xuanren_OnCreate);
end

function xuanren_RandName()
	name.text = RandNameData.Rand();
end

function xuanren_OnSelectRole()
	if selectList.selectedIndex == Proxy4Lua.CrtSelect() then
		return;
	end

	Proxy4Lua.SelectRole(selectList.selectedIndex);
	descCom.visible = selectList.selectedIndex ~= -1;
	if descAnim.playing then
		descAnim:Stop();
	end
	descAnim:Play();
end

function xuanren:OnUpdate()
	if UIManager.IsDirty("xuanren") then
		xuanren_FlushData();
		UIManager.ClearDirty("xuanren");
	end
end

function xuanren_FlushData()
	selectList.selectedIndex = Proxy4Lua.CrtSelect();
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
	local select = Proxy4Lua.CrtSelect();
	if select == -1 then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "请选择角色", true);
		return;
	end

	if name.text == "" then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "请输入昵称", true);
		return;
	end
	Proxy4Lua.CreatePlayer(select + 1, name.text);
end