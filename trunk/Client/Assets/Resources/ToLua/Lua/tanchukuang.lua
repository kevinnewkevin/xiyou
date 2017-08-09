require "FairyGUI"

tanchukuang = fgui.window_class(WindowBase)

local Window;

local _Title;
local _Content;
local _Confirm1;
local _Confirm2;
local _Cancel;

local ttitle;
local tcontent;
local tconfirm1;
local tconfirm2;
local tcancel;
local tsingle;

function tanchukuang:OnEntry()
	Window = tanchukuang.New();
	Window:Show();
end

function tanchukuang:GetWindow()
	return Window;
end

function tanchukuang:OnInit()
	self.contentPane = UIPackage.CreateObject("tanchukuang", "tanchukuang_com").asCom;
	self:Center();
	self.modal = true;

	_Title = self.contentPane:GetChild("n4");
    _Content = self.contentPane:GetChild("n5");
    _Confirm1 = self.contentPane:GetChild("n6");
    _Confirm2 = self.contentPane:GetChild("n8");
    _Cancel = self.contentPane:GetChild("n7");

    self.closeButton = _Cancel;
end

function tanchukuang:SetData(title, content, isSingle, confirm, cancel)
	_Confirm1.onClick:Clear();
	_Confirm2.onClick:Clear();

	ttitle = title;
	tcontent = content;
	tconfirm1 = confirm;
	tconfirm2 = confirm;
	tcancel = cancel;
	tsingle = isSingle;

--	UIManager.SetDirty("tanchukuang");
	tanchukuang_FlushData();
end

function tanchukuang:OnUpdate()
	if UIManager.IsDirty("tanchukuang") then
		tanchukuang_FlushData();
		UIManager.ClearDirty("tanchukuang");
	end
end

function tanchukuang:OnTick()
	
end

function tanchukuang:isShow()
	return Window.isShowing;
end

function tanchukuang:OnDispose()
	Window:Dispose();
end

function tanchukuang:OnHide()
	Window:Hide();
end

function tanchukuang_FlushData()
	_Title.text = ttitle;
	_Content.text = tcontent;
	if tsingle then
		_Confirm1.visible = false;
		_Cancel.visible = false;
		_Confirm2.visible = true;
	else
		_Confirm1.visible = true;
		_Cancel.visible = true;
		_Confirm2.visible = false;
	end

	if tconfirm1 ~= nil then
		_Confirm1.onClick:Add(tconfirm1);
	end
	if tconfirm2 ~= nil then
		_Confirm2.onClick:Add(tconfirm2);
	end
	if tcancel ~= nil then
		_Cancel.onClick:Add(tcancel);
	end
end