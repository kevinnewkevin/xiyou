require "FairyGUI"

fenjie = fgui.window_class(WindowBase)
local Window;
local nameLab;

local iteminst;
local itemName;
local itemIcon;
local itemIconBack;
local itemNum;

local moneyName;
local moneyIcon;
local moneyIconBack;
local moneyNum;
local inputLab;

local okBtn;
local cancelBtn;
local minusBtn;
local addBtn;
local numLab;
local delNum;
local changeNum;
function fenjie:OnEntry()
	Window = fenjie.New();
	Window:Show();
end

function fenjie:GetWindow()
	return Window;
end

function fenjie:OnInit()
	self.contentPane = UIPackage.CreateObject("fenjie", "fenjie_com").asCom;
	self:Center();
	self.modal = true;
	--self.onClick:Add(fenjie_OnHide);

    local icon = self.contentPane:GetChild("n4");
    local mIcon = self.contentPane:GetChild("n6");
    itemName = self.contentPane:GetChild("n14");
    itemIcon = icon:GetChild("n1");
    itemIconBack = icon:GetChild("n0");
    itemNum = icon:GetChild("n2");
    moneyName = self.contentPane:GetChild("n15");
    moneyIcon = mIcon:GetChild("n1");
    moneyIconBack = mIcon:GetChild("n0");
    moneyNum = mIcon:GetChild("n2");


    numLab = self.contentPane:GetChild("n11");
    minusBtn = self.contentPane:GetChild("n8");
    addBtn = self.contentPane:GetChild("n9");
    minusBtn.onClick:Add(fenjie_OnMinusBtn);
    addBtn.onClick:Add(fenjie_OnAddBtn);

    okBtn = self.contentPane:GetChild("n18");
	okBtn.onClick:Add(fenjie_OnOk);
	cancelBtn = self.contentPane:GetChild("n19");
	cancelBtn.onClick:Add(fenjie_OnCancelBtn);
    delNum = 1;
	fenjie_FlushData();
end

function fenjie:OnUpdate()
	if UIManager.IsDirty("fenjie") then
		fenjie_FlushData();
		UIManager.ClearDirty("fenjie");
	end

	if numLab.text == nil or numLab.text == "" then
		return;
	end


	local num = numLab.text;
	local dNum = tonumber(num);	
	if dNum >iteminst.Stack then
		dNum = iteminst.Stack;
	end

	if dNum <1 then
		dNum = 1;
	end

	delNum = dNum; 
	numLab.text = "" .. delNum;
	moneyNum.text = "" .. delNum * changeNum;

	
end

function fenjie:OnTick()
	
end

function fenjie:isShow()
	return Window.isShowing;
end

function fenjie:OnDispose()
	Window:Dispose();
end

function fenjie:OnHide()
	Window:Hide();
end

function fenjie_FlushData()
    local bagwin = UIManager.GetWindow("bagui");
    iteminst = BagSystem.GetItemInstByIndex(bagwin.GetClickItemIdx(), bagwin.GetCrtTab());
    local itemdata;
    if iteminst ~= null then
        itemdata = ItemData.GetData(iteminst.ItemId);
    end
    local iconpath = "";
    if itemdata == nil then
        return;
    end
    itemIcon.asLoader.url = "ui://" .. itemdata._Icon;
    itemNum.text ="" .. iteminst.Stack;
    changeNum = itemdata._Price;
    itemIconBack.asLoader.url = "ui://" .. itemdata._IconBack;
    itemName.text = itemdata._Name;

    moneyIcon.asLoader.url = "ui://icon/hunbi_icon";

    moneyName.text = "魂币";
    moneyNum.text = "" .. delNum * changeNum;
    numLab.text = "" .. delNum;
end

function fenjie_OnOk(context)
	Proxy4Lua.ResolveItem(iteminst.InstId,delNum);
	UIManager.Hide("fenjie");
    delNum = 1;
end

function fenjie_OnCancelBtn(context)
    UIManager.Hide("fenjie");
    delNum = 1;
end



function fenjie_OnMinusBtn(context)
    if delNum <=1 then
        delNum  = 1;
    else
        delNum = delNum-1;
    end
    numLab.text = "" .. delNum;
   -- itemNum.text ="" .. iteminst.Stack - delNum;
    moneyNum.text = "" .. delNum * changeNum;
end

function fenjie_OnAddBtn(context)
    if delNum >=iteminst.Stack then
        delNum = iteminst.Stack; 
    else
        delNum = delNum+1;
    end
    numLab.text = "" .. delNum;
   -- itemNum.text ="" .. iteminst.Stack - delNum;
    moneyNum.text = "" .. delNum * changeNum;
end


