require "FairyGUI"

squadRenming = fgui.window_class(WindowBase)
local Window;

local typeList;
local crtType;

function squadRenming:OnEntry()
	Window = squadRenming.New();
	Window:Show();
end

function squadRenming:GetWindow()
	return Window;
end

function squadRenming:OnInit()
	self.contentPane = UIPackage.CreateObject("bangpai", "bangpairenming_").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n16");

	typeList = self.contentPane:GetChild("n24").asList;
	typeList.onClickItem:Add(squadRenming_OnTypeSelected);

	local confirmBtn = self.contentPane:GetChild("n14").asButton;
	confirmBtn.onClick:Add(squadRenming_OnConfirm);

	local cancelBtn = self.contentPane:GetChild("n15").asButton;
	cancelBtn.onClick:Add(squadRenming_OnCancel);

	crtType = 0;
	typeList.selectedIndex = -1;
	squadRenming_FlushData();
end

function squadRenming_OnTypeSelected()
	crtType = typeList.selectedIndex;
end

function squadRenming_OnConfirm()
	local playerInstid = UIParamHolder.Get("squadRenmingPlayer");
	Proxy4Lua.ChangeMemberPosition(playerInstid, crtType + 1);
	squadRenming_OnCancel();
end

function squadRenming_OnCancel()
	UIManager.Hide("squadRenming");
end

function squadRenming:OnUpdate()
	if UIManager.IsDirty("squadRenming") then
		squadRenming_FlushData();
		UIManager.ClearDirty("squadRenming");
	end
end

function squadRenming:OnTick()
	
end

function squadRenming:isShow()
	return Window.isShowing;
end

function squadRenming:OnDispose()
	Window:Dispose();
end

function squadRenming:OnHide()
	crtType = 0;
	typeList.selectedIndex = -1;
	Window:Hide();
end

function squadRenming_FlushData()
	
end