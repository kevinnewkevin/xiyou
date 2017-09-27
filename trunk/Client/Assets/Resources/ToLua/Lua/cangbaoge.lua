require "FairyGUI"

cangbaoge = fgui.window_class(WindowBase)
local Window;


local cardGroupList;
local crtGroupIdx = 0;
local crtCardInstID = 0;
local crtCardsFee = 0;
local crtCardName;
local isInGroup;
local gold;
local chopper;
local stamaPoint;
function cangbaoge:OnEntry()
	Define.LaunchUIBundle("icon");
	Window = cangbaoge.New();
	Window:Show();
end

function cangbaoge:GetWindow()
	return Window;
end

function cangbaoge:OnInit()
	self.contentPane = UIPackage.CreateObject("cangbaoge", "cangbaoge_com").asCom;
	self:Center();

	self.closeButton = self.contentPane:GetChild("n2").asButton;
	local shopList = self.contentPane:GetChild("n5").asList;
	local feeMax = shopList.numItems;
	local feeItem;
	feeItem = shopList:GetChildAt(0);



	local shopBtn0 = feeItem:GetChild("n9");
	local itemLab0 = feeItem:GetChild("n6");
	local moenyLab0 = shopBtn0 :GetChild("n3");
	shopBtn0.data = 1000;
	shopBtn0.onClick:Add(cangbaoge_OnBuyClick);
	local shopData = ShopData.GetData(1000);
	itemLab0.text = shopData._Name;
	moenyLab0.text = shopData._Price;

	local shopBtn1 = feeItem:GetChild("n11");
	local itemLab1 = feeItem:GetChild("n7");
	local moenyLab1 = shopBtn1 :GetChild("n3");
	shopBtn1.data = 1001;
	shopBtn1.onClick:Add(cangbaoge_OnBuyClick);
	local shopData = ShopData.GetData(1001);
	itemLab1.text = shopData._Name;
	moenyLab1.text = shopData._Price;

	local shopBtn2 = feeItem:GetChild("n12");
	local itemLab2 = feeItem:GetChild("n8");
	local moenyLab2 = shopBtn2 :GetChild("n3");
	shopBtn2.data = 1002;
	shopBtn2.onClick:Add(cangbaoge_OnBuyClick);
	local shopData = ShopData.GetData(1002);
	itemLab2.text = shopData._Name;
	moenyLab2.text = shopData._Price;

	local moneyGroup = self.contentPane:GetChild("n0").asCom;
	gold = moneyGroup:GetChild("n5");
	chopper = moneyGroup:GetChild("n7");
	stamaPoint = moneyGroup:GetChild("n12");
	cangbaoge_FlushData();
end




function cangbaoge:OnUpdate() 
	if UIManager.IsDirty("cangbaoge") then
		cangbaoge_FlushData();
		UIManager.ClearDirty("cangbaoge");
	end
end

function cangbaoge:OnTick()
	
end

function cangbaoge:isShow()
	return Window.isShowing;
end

function cangbaoge:OnDispose()
	Window:Dispose();
end

function cangbaoge:OnHide()
	Window:Hide();
end

function cangbaoge_FlushData()
	gold.text = GamePlayer._Data.IProperties[8];
	chopper.text = GamePlayer._Data.IProperties[6];
	stamaPoint.text = GamePlayer._Data.IProperties[10];

end

function cangbaoge_OnBuyClick(context)

	local shopData = ShopData.GetData(context.sender.data);
	if shopData._Price > GamePlayer._Data.IProperties[8] then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "金不够", true);
		return;
	end
	ShopSystem.buyType = context.sender.data;
	Proxy4Lua.BuyShopItem(context.sender.data);
end 