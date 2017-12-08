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
local hunBi;
local boxInfo;
local infoMoneyLab;
local infoBoxNameLab;
local infoBoxCloseBtn;
local greenLab;
local buleLab;
local purpleLab;
local orangeLab;
local greenLabBack;
local buleLabBack;
local purpleLabBack;
local orangeLabBack;
local greenImg;
local buleImg;
local purpleImg;
local orangeImg;
local buyShopId;
local infoBuyBtnLab;
local infoBuyBoxImg;


local blackList;
local blackRefreshBtn;
local hunMoneyLab;
local refreshNumLab;
local refreshTimeLab;

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
	shopList.scrollItemToViewOnClick = false;

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

	greenLabBack =  boxInfo:GetChild("n25");
	buleLabBack  =  boxInfo:GetChild("n26");
	purpleLabBack  =  boxInfo:GetChild("n27");
	orangeLabBack  =  boxInfo:GetChild("n28");

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

	local moneyGroup = self.contentPane:GetChild("n10").asCom;
	gold = moneyGroup:GetChild("n5");
	chopper = moneyGroup:GetChild("n7");
	stamaPoint = moneyGroup:GetChild("n12");
	hunBi = moneyGroup:GetChild("n16");


	local black = shopList:GetChildAt(1);
	blackList =  black:GetChild("n42").asList;
	--blackList:SetVirtual();
	blackList.itemRenderer = cangbaoge_RenderListItem;
	blackRefreshBtn = black:GetChild("n34");
	hunMoneyLab = black:GetChild("n38");
	refreshNumLab = black:GetChild("n39");
	--refreshTimeLab = black:GetChild("n41");
	--refreshTimeLab.text = "每天00:00 6:00 12:00 18:00刷新" 
	blackRefreshBtn.onClick:Add(cangbaoge_OnRefreshBtn);

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
	hunBi.text = GamePlayer._Data.IProperties[11];
	local blackMarket = ShopSystem.BlackMarket;
	blackList.numItems = 6;
	-- blackMarket:ShopItems.length;
	hunMoneyLab.text = GamePlayer._Data.IProperties[11] ;
	local refreshNum = ShopSystem.GetBlackRefreshNum();
	refreshNumLab.text =  "还可刷新 " .. refreshNum .. " 次(消耗1魂币)";
end


function cangbaoge_RenderListItem(index, obj)
	local name =  obj:GetChild("n4");
	local price =  obj:GetChild("n3");
	local icon = obj:GetChild("n2");
	local isBuy = obj:GetChild("n10");
	local itemNum = obj:GetChild("n11");

	local back = obj:GetChild("n5");
	local sBack = obj:GetChild("n6");
	sBack.visible = false; 
	back.visible = false; 
	local shopId = ShopSystem.GetBlackMarketId(index);
	local bIsBuy = ShopSystem.GetBlackMarketIsBuy(index);
	if bIsBuy == false then
		isBuy.visible = true; 
		--obj.enabled = false;
		obj.onClick:Remove(cangbaoge_OnBlackBuyClick);
	else
		isBuy.visible = false; 
		obj.enabled = true;
		obj.onClick:Add(cangbaoge_OnBlackBuyClick);
	end
	local shopData = ShopData.GetData(shopId);
	name.text = "" .. shopData._Name;
	price.text = "" .. shopData._Price;
	itemNum.text = "X" .. shopData._Num;
	obj.data = shopData._ShopId;
	itemdata = ItemData.GetData(shopData._itemId);
	icon.asLoader.url = "ui://" .. itemdata._Icon;
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
	Window:ShowModalWait();
	boxInfo.visible = false;
end 


function cangbaoge_OnCloseClick(context)
	boxInfo.visible = false;
end

function cangbaoge_UpdateInfo()
	local shopData = ShopData.GetData(buyShopId);
	infoBoxNameLab.text = shopData._Name;
	infoBuyBtnLab.text = shopData._Price;
	infoMoneyLab.text = shopData._copper .. "铜币";
	greenLab.visible = false;
	greenLabBack.visible = false;
	greenImg.visible = false;	
	buleLab.visible = false;
	buleLabBack.visible = false;
	buleImg.visible = false;
	purpleLab.visible = false;
	purpleLabBack.visible = false;
	purpleImg.visible = false;
	orangeLab.visible = false;
	orangeLabBack.visible = false;
	orangeImg.visible = false;

	local cardData = CardcloseData.GetData(shopData._CardId);
	if cardData._Greennum > 0 then
		greenImg.visible = true;
		greenLab.visible = true;	
		greenLabBack.visible = true;	
		greenLab.text = "绿色碎片*" ..cardData._Greennum;	
	end
	if cardData._Bluenum > 0 then
		buleImg.visible = true;
		buleLab.visible = true;
		buleLabBack.visible = true;	
		buleLab.text = "蓝色碎片*" ..cardData._Bluenum;		
	end
	if cardData._Purplenum > 0 then
		purpleImg.visible = true;
		purpleLab.visible = true;
		purpleLabBack.visible = true;
		purpleLab.text = "紫色碎片*" ..cardData._Purplenum;  	
	end
	if cardData._Orangenum > 0 then
		orangeImg.visible = true;
		orangeLab.visible = true;
		orangeLabBack.visible = true;
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

function cangbaoge_OnRefreshBtn(context)
	local refreshNum = ShopSystem.GetBlackRefreshNum();
	if refreshNum <= 0 then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "刷新次数不够", true);
		return;
	end
	if GamePlayer._Data.IProperties[11] < 1 then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "魂币不够", true);
		return;
	end
	Proxy4Lua.RefreshBlackMarkte();
	Proxy4Lua.PopMsg("刷新成功");

end 

function cangbaoge_OnBlackBuyClick(context)
	buyShopId = context.sender.data;
	local shopData = ShopData.GetData(buyShopId);
	if shopData._Price > GamePlayer._Data.IProperties[11] then
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "魂币不够", true);
		return;
	end

	local MessageBox = UIManager.ShowMessageBox();
	MessageBox:SetData("提示", "是否花费 " .. shopData._Price .. " 魂币?", false, cangbaoge_OnBuyBlack);
end 

function cangbaoge_OnBuyBlack(context)
	UIManager.HideMessageBox();
	ShopSystem.buyType = buyShopId;
	Proxy4Lua.BuyShopItem(buyShopId);
	Window:ShowModalWait();
end
