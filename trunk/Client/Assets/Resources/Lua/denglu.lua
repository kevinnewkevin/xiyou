require "FairyGUI"

denglu = fgui.window_class(WindowBase)
local Window;

function denglu:OnEntry()
	Window = denglu.New();
	Window:Show();
end

function denglu:OnInit()
	UIPackage.AddPackage("UI/UI_Fairy/export/Package1");
	self.contentPane = UIPackage.CreateObject("Package1", "denglu").asCom;
	self:Center();

	local enterBtn = self.contentPane:GetChild("n3").asButton;
	enterBtn.onClick:Add(OnEnterGame);

	FlushData();
end

function denglu:OnUpdate()
	
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
	
end

function OnEnterGame()
	Proxy4Lua.CreatePlayer(1, "guowengui");
end