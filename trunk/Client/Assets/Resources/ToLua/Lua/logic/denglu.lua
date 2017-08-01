require "FairyGUI"

denglu = fgui.window_class(WindowBase)
local Window;

local enterBtn;

function denglu:OnEntry()
	Window = denglu.New();
	Window:Show();
end

function denglu:OnInit()
	UIPackage.AddPackage("UI/UI_Fairy/export/denglu");
	self.contentPane = UIPackage.CreateObject("denglu", "denglu").asCom;
	self:Center();

	enterBtn = self.contentPane:GetChild("n3").asButton;
	enterBtn.onClick:Add(OnEnterGame);

	FlushData();
end

function denglu:OnUpdate()
	if UIManager.IsDirty("denglu") then
		FlushData();
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

function FlushData()
	enterBtn.enabled = not DataLoader._IsLoading;
end

function OnEnterGame()
	Proxy4Lua.CreatePlayer(1, "guowengui");
end