require "FairyGUI"

denglu = fgui.window_class(WindowBase)
local Window;

local enterBtn;

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

function denglu_OnEnterGame()
	Proxy4Lua.CreatePlayer(1, "小西游名字七字");
end