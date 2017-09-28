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
local boxInfo;
local infoMoneyLab;
local infoBoxNameLab;
local infoBoxCloseBtn;
local greenLab;
local buleLab;
local purpleLab;
local orangeLab;
local greenImg;
local buleImg;
local purpleImg;
local orangeImg;
local buyShopId;
local infoBuyBtnLab;
local infoBuyBoxImg;
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

	boxInfo = self.contentPane:GetChild("n8");
	boxInfo.visible = false;
	local infoBuyBtn = boxInfo:GetChild("n18");
	infoBuyBtn.onClick:Add(cangbaoge_OnInfoBuyClick);
	infoMoneyLab = boxInfo:GetChild("n13");
	infoBoxNameLab = boxInfo:GetChild("n12");
	infoBoxCloseBtn = boxInfo:GetChild("n20");
	infoBoxCloseBtn.onClick:Add(cangbaoge_OnCloseClick);
	greenLab =  boxInfo:GetChild("n14");
	buleLab =  boxInfo:GetChild("n15");
	purpleLab =  boxInfo:GetChild("n16");
	orangeLab =  boxInfo:GetChild("n17");

	greenImg =  boxInfo:GetChild("n7");
	buleImg =  boxInfo:GetChild("n8");
	purpleImg =  boxInfo:GetChild("n10");
	orangeImg =  boxInfo:GetChild("n9");

	infoBuyBtnLab = infoBuyBtn:GetChild("n3");
	infoBuyBoxImg = boxInfo:GetChild("n22");
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
	buyShopId = context.sender.data;
	boxInfo.visible = true;
	cangbaoge_UpdateInfo();
end 

function cangbaoge_OnInfoBuyClick(context)
	local shopData = ShopData.GetData(buyShopId);
	if shopData._Price > GamePlayer._Data.IProperties[8] then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "金不够", true);
		return;
	end
	ShopSystem.buyType = buyShopId;
	Proxy4Lua.BuyShopItem(buyShopId);
	boxInfo.visible = false;
end 


function cangbaoge_OnCloseClick(context)
	boxInfo.visible = false;
end

function cangbaoge_UpdateInfo()
	local shopData = ShopData.GetData(buyShopId);
	infoBoxNameLab.text = shopData._Name;
	infoBuyBtnLab.text = shopData._Price;
	infoMoneyLab.text = shopData._copper;
	greenLab.visible = false;
	greenImg.visible = false;	
	buleLab.visible = false;
	buleImg.visible = false;
	purpleLab.visible = false;
	purpleImg.visible = false;
	orangeLab.visible = false;
	orangeImg.visible = false;

	local cardData = CardcloseData.GetData(shopData._CardId);
	if cardData._Greennum > 0 then
		greenImg.visible = true;
		greenLab.visible = true;	
		greenLab.text = "绿色碎片*" ..cardData._Greennum;	
	end
	if cardData._Bluenum > 0 then
		buleImg.visible = true;
		buleLab.visible = true;
		buleLab.text = "蓝色碎片*" ..cardData._Bluenum;		
	end
	if cardData._Purplenum > 0 then
		purpleImg.visible = true;
		purpleLab.visible = true;
		purpleLab.text = "紫色碎片*" ..cardData._Purplenum;  	
	end
	if cardData._Orangenum > 0 then
		orangeImg.visible = true;
		orangeLab.visible = true;
		orangeLab.text = "橙色碎片*" ..cardData._Orangenum;	
	end

	if buyShopId == 1000 then
		infoBuyBoxImg.asLoader.url  = "ui://cangbaoge/renxian";
	end
	if buyShopId == 1001 then
		infoBuyBoxImg.asLoader.url  = "ui://cangbaoge/dixian";
	end
	if buyShopId == 1002 then
		infoBuyBoxImg.asLoader.url  = "ui://cangbaoge/tianxian";
	end

end