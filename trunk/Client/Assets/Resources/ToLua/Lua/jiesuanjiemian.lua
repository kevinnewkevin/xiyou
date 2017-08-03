require "FairyGUI"

jiesuanjiemian = fgui.window_class(WindowBase)
local Window;

local okBtn;
local resultImg;

function jiesuanjiemian:OnEntry()
	Window = jiesuanjiemian.New();
	Window:Show();
end

function jiesuanjiemian:GetWindow()
	return Window;
end

function jiesuanjiemian:OnInit()
	self.contentPane = UIPackage.CreateObject("jiesuanjiemian", "jiesuanjiemian_com").asCom;
	self:Center();

	resultImg = self.contentPane:GetChild("n36").asLoader;

	okBtn = self.contentPane:GetChild("n33").asButton;
	okBtn.onClick:Add(jiesuanjiemian_OnOkBtn);

	jiesuanjiemian_FlushData();
end

function jiesuanjiemian_FlushData()
	if Battle.IsWin then
		resultImg.url = UIPackage.GetItemURL("jiesuanjiemian", "shengli");
	else
		resultImg.url = UIPackage.GetItemURL("jiesuanjiemian", "shibai");
	end
end

function jiesuanjiemian:OnUpdate()
	
end

function jiesuanjiemian:OnTick()
	
end

function jiesuanjiemian:isShow()
	return Window.isShowing;
end

function jiesuanjiemian:OnDispose()
	Window:Dispose();
end

function jiesuanjiemian:OnHide()
	Window:Hide();
end

function jiesuanjiemian_OnOkBtn()
	SceneLoader.LoadScene("main");
end