require "FairyGUI"

shengli_Component = fgui.window_class(WindowBase)
local Window;

local okBtn;
local resultImg;

function shengli_Component:OnEntry()
	Window = shengli_Component.New();
	Window:Show();
end

function shengli_Component:GetWindow()
	return Window;
end

function shengli_Component:OnInit()
	UIPackage.AddPackage("UI/UI_Fairy/export/jiesuanjiemian");
	self.contentPane = UIPackage.CreateObject("jiesuanjiemian", "shengli_Component").asCom;
	self:Center();

	resultImg = self.contentPane:GetChild("n36").asLoader;

	okBtn = self.contentPane:GetChild("n33").asButton;
	okBtn.onClick:Add(shengli_Component_OnOkBtn);

	shengli_Component_FlushData();
end

function shengli_Component_FlushData()
	if Battle.IsWin then
		resultImg.url = UIPackage.GetItemURL("jiesuanjiemian", "shengli");
	else
		resultImg.url = UIPackage.GetItemURL("jiesuanjiemian", "shibai");
	end
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

function shengli_Component:OnHide()
	Window:Hide();
end

function shengli_Component_OnOkBtn()
	SceneLoader.LoadScene("main");
end