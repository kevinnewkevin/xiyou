require "FairyGUI"

shengli_Component = fgui.window_class(WindowBase)
local Window;

local OkBtn;

function shengli_Component:OnEntry()
	Window = shengli_Component.New();
	Window:Show();
end

function shengli_Component:OnInit()
	UIPackage.AddPackage("UI/UI_Fairy/export/jiesuanjiemian");
	self.contentPane = UIPackage.CreateObject("jiesuanjiemian", "shengli_Component").asCom;
	self:Center();

	OkBtn = self.contentPane:GetChild("n33").asButton;
	OkBtn.onClick:Add(OnOkBtn);
end

function shengli_Component:OnUpdate()
	
end

function shengli_Component:OnTick()
	
end

function shengli_Component:isShow()
	return Window.isShowing;
end

function shengli_Component:OnDispose()
	Window:Dispose();
end

function OnOkBtn()
	SceneLoader.LoadScene("main");
end