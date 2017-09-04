require "FairyGUI"

bagtips = fgui.window_class(WindowBase)
local Window;

function bagtips:OnEntry()
	Window = bagtips.New();
	Window:Show();
end

function bagtips:GetWindow()
	return Window;
end

function bagtips:OnInit()
	self.contentPane = UIPackage.CreateObject("bagtips", "bagtips_com").asCom;
	self:Center();
	self.modal = true;

	self.closeButton = self.contentPane:GetChild("n4");

	bagtips_FlushData();
end

function bagtips:OnUpdate()
	if UIManager.IsDirty("bagtips") then
		bagtips_FlushData();
		UIManager.ClearDirty("bagtips");
	end
end

function bagtips:OnTick()
	
end

function bagtips:isShow()
	return Window.isShowing;
end

function bagtips:OnDispose()
	Window:Dispose();
end

function bagtips:OnHide()
	Window:Hide();
end

function bagtips_FlushData()
	local bagwin = UImanager.GetWindow("bagui");
	local iteminst = BagSystem.GetItemInstByIndex(bagwin.crtClickItemIdx, bagwin.crtTab);
end