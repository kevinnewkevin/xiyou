require "FairyGUI"

guanka = fgui.window_class(WindowBase)
local Window;


function guanka:OnEntry()      
	Window = guanka.New();
	Window:Show();
end

function guanka:OnInit()
	self.contentPane = UIPackage.CreateObject("guanka", "guanka_com").asCom;
	self:Center();
	self.modal = true;

	self.closeButton = self.contentPane:GetChild("n1").asButton;

	local guankaID = UIManager.GetWindow("jiehun").GetGuankaId();
	local data = HeroStroyData.GetData(guankaID);
	local tatle = self.contentPane:GetChild("n15");
	 tatle.text = data.Name_;
	local desc = self.contentPane:GetChild("n8");
	desc.text = data.Desc_;
	local btn = self.contentPane:GetChild("n11");
	btn.onClick:Add(guanka_OnBattle);
	guanka_FlushData();
end

function guanka:GetWindow()
	return Window;
end

function guanka:OnUpdate()
	if UIManager.IsDirty("guanka") then
		guanka_FlushData();
		UIManager.ClearDirty("guanka");
	end
end



function guanka_OnBattle(context)
	Proxy4Lua.BattleJoin();
end

function guanka:OnTick()
	
end

function guanka:isShow()
	return Window.isShowing;
end

function guanka:OnDispose()
	Window:Dispose();
end

function guanka:OnHide()
	Window:Hide();
end

function guanka_FlushData()
end