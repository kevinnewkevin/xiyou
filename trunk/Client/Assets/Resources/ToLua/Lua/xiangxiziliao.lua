require "FairyGUI"

xiangxiziliao = fgui.window_class(WindowBase)
local Window;

local holder;

function xiangxiziliao:OnEntry()
	Window = xiangxiziliao.New();
	Window:Show();
end

function xiangxiziliao:GetWindow()
	return Window;
end

function xiangxiziliao:OnInit()
	self.contentPane = UIPackage.CreateObject("xiangxiziliao", "xiangxiziliao_com").asCom;
	self:Center();
	self.modal = true;

	self.closeButton = self.contentPane:GetChild("n4");
	holder = self.contentPane:GetChild("n63").asGraph;

	xiangxiziliao_FlushData();
end

function xiangxiziliao:OnUpdate()
	if UIManager.IsDirty("xiangxiziliao") then
		xiangxiziliao_FlushData();
		UIManager.ClearDirty("xiangxiziliao");
	end
end

function xiangxiziliao:OnTick()
	
end

function xiangxiziliao:isShow()
	return Window.isShowing;
end

function xiangxiziliao:OnDispose()
	Window:Dispose();
end

function xiangxiziliao:OnHide()
	Window:Hide();
end

function xiangxiziliao_FlushData()
	holder:SetNativeObject(Proxy4Lua.GetAssetGameObject("Player/longtaizi"));
end