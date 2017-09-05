require "FairyGUI"

duibaikuang = fgui.window_class(WindowBase)
local Window;

local name;
local content;
local displayId;

local talks = {};
local crtIdx = 0;
local talkNums;
local typeWriter;

function duibaikuang:OnEntry()
	Window = duibaikuang.New();
	Window:Show();
end

function duibaikuang:GetWindow()
	return Window;
end

function duibaikuang:OnInit()
	self.contentPane = UIPackage.CreateObject("duibaikuang", "duibaikuang_com").asCom;
	self:Center();

	local skipBtn = self.contentPane:GetChild("n9").asButton;
	skipBtn.onClick:Add(duibaikuang_OnNext);

	self.contentPane.onClick:Add(duibaikuang_OnNext);

	name = self.contentPane:GetChild("n8").asTextField;
	content = self.contentPane:GetChild("n6").asTextField;
	typeWriter = TypingEffect.New(content);

	duibaikuang_FlushData();
end

function duibaikuang_OnSkip()
	Proxy4Lua.ContinueOpra();
	UIManager.Hide("duibaikuang");
	crtIdx = 0;
end

function duibaikuang_OnNext()
	crtIdx = crtIdx + 1;
	UIManager.SetDirty("duibaikuang");
end

function duibaikuang:OnUpdate()
	if UIManager.IsDirty("duibaikuang") then
		duibaikuang_FlushData();
		UIManager.ClearDirty("duibaikuang");
	end
end

function duibaikuang:OnTick()
	
end

function duibaikuang:isShow()
	return Window.isShowing;
end

function duibaikuang:OnDispose()
	Window:Dispose();
end

function duibaikuang:OnHide()
	Window:Hide();
end

function duibaikuang_FlushData()
	talks = Proxy4Lua.GetTalk();
	talkNum = Proxy4Lua.GetTalkNum();

	if talkNum <= crtIdx then
		Proxy4Lua.ContinueOpra();
		UIManager.Hide("duibaikuang");
		crtIdx = 0;
	else
		local tData = TalkData.GetData(talks[crtIdx]);
		name.text = tData._Name;
		content.text = tData._Content;
		typeWriter:Start();
		typeWriter:PrintAll(0.05);
	end
end