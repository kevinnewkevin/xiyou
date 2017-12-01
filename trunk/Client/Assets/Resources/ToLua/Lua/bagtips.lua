require "FairyGUI"

bagtips = fgui.window_class(WindowBase)
local Window;
local nameLab;
local descLab;
local levelLab;
local typeLab;
local icon;
local iconBack;
local stackLab;
local iteminst;
local useBtn;
local useAllBtn;
local delBtn;
local unlockBtn;
local btns; 
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
	self.onClick:Add(bagtips_OnHide);
	--self.closeButton = self.contentPane:GetChild("n11");
	nameLab = self.contentPane:GetChild("n9");
	descLab= self.contentPane:GetChild("n24");
	levelLab= self.contentPane:GetChild("n21");
	typeLab= self.contentPane:GetChild("n22");

	local itemIcon = self.contentPane:GetChild("n10");
	icon = itemIcon:GetChild("n1");
	iconBack = itemIcon:GetChild("n0");
	stackLab = itemIcon:GetChild("n2");

	btns = self.contentPane:GetChild("n28").asList;

	useBtn = btns:GetChildAt(2);
	useBtn.visible = false;
	useAllBtn = btns:GetChildAt(3);
	useAllBtn.visible = false;
	delBtn = btns:GetChildAt(1);
	delBtn.onClick:Add(bagtips_OnDelItem);
	unlockBtn = btns:GetChildAt(0);
	unlockBtn.onClick:Add(bagtips_OnUnlockBtn);
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
	local tipsItemId = UIParamHolder.Get("tipsItem");
	if  tipsItemId ~= 0 then
		btns.visible = false;
		local iData = ItemData.GetData(tipsItemId);
		icon.asLoader.url = "ui://" .. iData._Icon;
		stackLab.text ="";
		iconBack.asLoader.url = "ui://" .. iData._IconBack;
		descLab.text = iData._Desc;
		nameLab.text = iData._Name;
		if  iData._Type == "IMT_Debris" then
			typeLab.text = "碎片";
		elseif iData._Type == "IMT_Consumables" then
			typeLab.text = "消耗品";
		end
	else
		local bagwin = UIManager.GetWindow("bagui");
		iteminst = BagSystem.GetItemInstByIndex(bagwin.GetClickItemIdx(), bagwin.GetCrtTab());
		local itemdata;
		if iteminst ~= null then
			itemdata = ItemData.GetData(iteminst.ItemId);
		end
		btns.visible = true;
		local iconpath = "";
		if itemdata == nil then
			return;
		end
		icon.asLoader.url = "ui://" .. itemdata._Icon;
		stackLab.text ="" .. iteminst.Stack;
		iconBack.asLoader.url = "ui://" .. itemdata._IconBack;
		descLab.text = itemdata._Desc;
		nameLab.text = itemdata._Name;
		if  itemdata._Type == "IMT_Debris" then
			typeLab.text = "碎片";
			unlockBtn.visible = true;
		elseif itemdata._Type == "IMT_Consumables" then
			typeLab.text = "消耗品";
			unlockBtn.visible = false;
		end
	end
end

function bagtips_OnDelItem(context)
	Proxy4Lua.DeleteItem(iteminst.InstId,iteminst.Stack);
	UIManager.Hide("bagtips");
end

function bagtips_OnUnlockBtn(context)
	UIManager.Show("fenjie");
end

function bagtips_OnHide(context)
	UIManager.Hide("bagtips");
end
