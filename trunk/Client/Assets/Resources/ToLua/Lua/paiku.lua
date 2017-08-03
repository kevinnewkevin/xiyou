require "FairyGUI"

paiku = fgui.window_class(WindowBase)
local Window;

function paiku:OnEntry()
	Window = paiku.New();
	Window:Show();
end

function paiku:GetWindow()
	return Window;
end

function paiku:OnInit()
	self.contentPane = UIPackage.CreateObject("paiku", "paiku_com").asCom;
	self:Center();

	local closeBtn = self.contentPane:GetChild("n7").asButton;
	closeBtn.onClick:Add(paiku_OnClose);

	paiku_FlushData();
end

function paiku:OnUpdate()
	if UIManager.IsDirty("paiku") then
		paiku_FlushData();
		UIManager.ClearDirty("paiku");
	end
end

function paiku:OnTick()
	
end

function paiku:isShow()
	return Window.isShowing;
end

function paiku:OnDispose()
	Window:Dispose();
end

function paiku:OnHide()
	Window:Hide();
end

function paiku_FlushData()
	
end

function paiku_OnClose()
	UIManager.Hide("paiku");
end