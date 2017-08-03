require "FairyGUI"

paiku_Component = fgui.window_class(WindowBase)
local Window;

function paiku_Component:OnEntry()
	Window = paiku_Component.New();
	Window:Show();
end

function paiku_Component:GetWindow()
	return Window;
end

function paiku_Component:OnInit()
	UIPackage.AddPackage("UI/UI_Fairy/export/paiku");
	self.contentPane = UIPackage.CreateObject("paiku", "paiku_Component").asCom;
	self:Center();

	local closeBtn = self.contentPane:GetChild("n7").asButton;
	closeBtn.onClick:Add(paiku_Component_OnClose);

	paiku_Component_FlushData();
end

function paiku_Component:OnUpdate()
	if UIManager.IsDirty("paiku_Component") then
		paiku_Component_FlushData();
		UIManager.ClearDirty("paiku_Component");
	end
end

function paiku_Component:OnTick()
	
end

function paiku_Component:isShow()
	return Window.isShowing;
end

function paiku_Component:OnDispose()
	Window:Dispose();
end

function paiku_Component:OnHide()
	Window:Hide();
end

function paiku_Component_FlushData()
	
end

function paiku_Component_OnClose()
print("paiku_Component_OnClose");
	UIManager.Hide("paiku_Component");
end