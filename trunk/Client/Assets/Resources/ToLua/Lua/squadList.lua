require "FairyGUI"

squadList = fgui.window_class(WindowBase)
local Window;

local squads;

local createBtn;
local createCom;
local createName;
local createCancel;
local createConfirm;

local searchBtn;
local searchCom;
local searchCancel;
local searchConfirm;

function squadList:OnEntry()
	Window = squadList.New();
	Window:Show();
end

function squadList:GetWindow()
	return Window;
end

function squadList:OnInit()
	self.contentPane = UIPackage.CreateObject("squadList", "squadList_com").asCom;
	self:Center();

	squads = self.contentPane:GetChild("").asList;
	squads:SetVirtual();
	squads.itemRenderer = squadList_RenderListItem;

	createBtn = self.contentPane:GetChild("").asButton;
	createBtn.onClick:Add(squadList_OnCreate);
	createCom = self.contentPane:GetChild("").asCom;
	createCom.visible = false;
	createName = createCom:GetChild("");
	createCancel = createCom:GetChild("");
	createConfirm = createCom:GetChild("");
	createCancel.onClick:Add(squadList_OnCreateCancel);
	createConfirm.onClick:Add(squadList_OnCreateConfirm);

	searchBtn = self.contentPane:GetChild("").asButton;
	searchBtn.onClick:Add(squadList_OnSearch);
	searchCom = self.contentPane:GetChild("").asCom;
	searchCom.visible = false;
	searchName = createCom:GetChild("");
	searchCancel = createCom:GetChild("");
	searchConfirm = createCom:GetChild("");
	searchCancel.onClick:Add(squadList_OnSearchCancel);
	searchConfirm.onClick:Add(squadList_OnSearchConfirm);

	squadList_FlushData();
end

function squadList_RenderListItem(index, obj)
	local checkBtn = obj:GetChild("").asButton;
	checkBtn.data = index;
	checkBtn.onClick:Add(squadList_OnCheckSquad);
end

function squadList_OnCheckSquad(context)
	
end

function squadList_OnCreate()
	createCom.visible = true;
end

function squadList_OnCreateCancel()
	createCom.visible = false;
end

function squadList_OnCreateConfirm()
	createCom.visible = false;
end

function squadList_OnSearch()
	searchCom.visible = true;
end

function squadList_OnCreateCancel()
	searchCom.visible = false;
end

function squadList_OnCreateConfirm()
	searchCom.visible = false;
end

function squadList:OnUpdate()
	if UIManager.IsDirty("squadList") then
		squadList_FlushData();
		UIManager.ClearDirty("squadList");
	end
end

function squadList:OnTick()
	
end

function squadList:isShow()
	return Window.isShowing;
end

function squadList:OnDispose()
	Window:Dispose();
end

function squadList:OnHide()
	searchName.text = "";
	searchCom.visible = false;
	createName.text = "";
	createCom.visible = false;
	Window:Hide();
end

function squadList_FlushData()
	squads.numItems = 100;
end