require "FairyGUI"

zhujiemian = fgui.window_class(WindowBase)
local Window;

local playerName;
local playerLevel;
local playerExp;
local gold;
local chopper;
local stamaPoint;
local expBar;
local headIcon;

local stateIcon;
local stateBar;
local listGroup;
function zhujiemian:OnEntry()
	UIManager.RegIDirty("zhujiemian");
	Define.LaunchUIBundle("icon");
	Window = zhujiemian.New();
	Window:Show();
end

function zhujiemian:OnInit()
	self.contentPane = UIPackage.CreateObject("zhujiemian", "zhujiemian_com").asCom;
	self:Center();

	--stateIcon = self.contentPane:GetChild("n17").asButton;
	--stateIcon.onClick:Add(zhujiemian_OnFolder);

	local bottomGroup = self.contentPane:GetChild("n18");
	stateBar = bottomGroup:GetController("c1");
	local bottomBtnList = bottomGroup:GetChild("n15").asList;
	local skillBtn = bottomBtnList:GetChildAt(1);
	skillBtn.onClick:Add(zhujiemian_OnSkillBtn);
	local cardCargo = bottomBtnList:GetChildAt(2);
	cardCargo.onClick:Add(zhujiemian_OnCardCargo);

	listGroup =  self.contentPane:GetChild("n43");

	local bagBtn = bottomBtnList:GetChildAt(0);
	bagBtn.onClick:Add(zhujiemian_OnBagBtn);

	local shopBtn = self.contentPane:GetChild("n30");
	shopBtn.onClick:Add(zhujiemian_OnShop);

	local infoGroup = self.contentPane:GetChild("n15").asCom;
	playerName = infoGroup:GetChild("n9");
	playerExp = infoGroup:GetChild("n10");
	playerLevel = infoGroup:GetChild("n11");
	expBar = infoGroup:GetChild("n8").asProgress;
	headIcon = infoGroup:GetChild("n13").asLoader;
	local headBtn = infoGroup:GetChild("n2");
	headIcon.onClick:Add(zhujiemian_OnHead);

	local moneyGroup = self.contentPane:GetChild("n16").asCom;
	gold = moneyGroup:GetChild("n5");
	chopper = moneyGroup:GetChild("n7");
	stamaPoint = moneyGroup:GetChild("n12");
	listGroup:SetVirtualAndLoop();
	listGroup.itemRenderer = zhujiemain_RenderListItem;
	listGroup.scrollPane.onScroll:Add(DoSpecialEffect);
	zhujiemian_FlushData();
	--DoSpecialEffect();
end


function DoSpecialEffect()
	local midX = listGroup.scrollPane.posX + listGroup.viewWidth / 2;
	local cnt = listGroup.numChildren;
	for  i = 1, cnt do
		local obj = listGroup:GetChildAt(i-1);
		local dist = Mathf.Abs(midX - obj.x - obj.width / 2) ;
		local effect = obj:GetChild("n6");
		if dist > obj.width then
			obj:SetScale(1, 1);
			local colorObj = obj:GetChild("n5");
			colorObj.color = Color.gray;
			effect.visible =false;
			obj.onClick:Set(zhujiemian_OnClose);
		else
			local ss = 1 + (1 - dist / obj.width) * 1;
			obj:SetScale(ss, ss);
			local colorObj = obj:GetChild("n5");
			colorObj.color = Color.white;
			effect.visible = true;
			if obj.data == 0 then
				obj.onClick:Set(zhujiemian_OnQieCuoBtn);
			elseif obj.data == 1 then
				obj.onClick:Set(zhujiemian_OnTaskBtn);
			end

		end
	end
end


function zhujiemian:GetWindow()
	return Window;
end

function zhujiemian:OnUpdate()
	if UIManager.IsDirty("zhujiemian") then
		zhujiemian_FlushData();
		UIManager.ClearDirty("zhujiemian");
	end
end

function zhujiemian:OnTick()
	
end

function zhujiemian:isShow()
	return Window.isShowing;
end

function zhujiemian:OnDispose()
	Window:Dispose();
end

function zhujiemian:OnHide()
	Proxy4Lua.ClearToDeleteAsset("zhujiemian");
	Window:Hide();
end

function zhujiemian_FlushData()
	local displayData = GamePlayer.GetMyDisplayData();
	headIcon.url = "ui://" .. displayData._HeadIcon .. "_yuan";
	local needExp = ExpData.NeedExp(GamePlayer._Data.IProperties[9]);
	playerName.text = GamePlayer._Name;
	playerExp.text = GamePlayer._Data.IProperties[4] .. "/" .. needExp;
	playerLevel.text = GamePlayer._Data.IProperties[9];

	gold.text = GamePlayer._Data.IProperties[8];
	chopper.text = GamePlayer._Data.IProperties[6];
	stamaPoint.text = GamePlayer._Data.IProperties[10];

	expBar.value = GamePlayer._Data.IProperties[4] / needExp * 100;
	listGroup.numItems = 3;
	DoSpecialEffect();
end

function zhujiemain_RenderListItem(index, obj)
	obj:SetPivot(0.5, 0.5);
	local img = obj:GetChild("n5");
	local effect = obj:GetChild("n6");
	local effRes = "" ;
	img:SetScale(0.5,0.5);
	if index == 0 then
		img.url = "ui://zhujiemian/zjm_31";
		obj.onClick:Set(zhujiemian_OnQieCuoBtn);
		obj.data = 0;
		effRes = "Effect/sunwukong_ui";
		effect :SetNativeObject(Proxy4Lua.GetEffectAssetGameObject(effRes)); 
		Proxy4Lua.AddToDelete("zhujiemian",effRes);
	elseif index == 1 then
		img.url = "ui://zhujiemian/zjm_30";
		obj.onClick:Set(zhujiemian_OnTaskBtn);
		effRes = "Effect/xihailongwang_ui";
		effect :SetNativeObject(Proxy4Lua.GetEffectAssetGameObject(effRes));
		Proxy4Lua.AddToDelete("zhujiemian",effRes);
		obj.data = 1;
	elseif index == 2 then
		img.url = "ui://zhujiemian/zjm_28";
		obj.onClick:Set(zhujiemian_OnClose);
		obj.data = 2;
		effect :SetNativeObject(Proxy4Lua.GetEffectAssetGameObject(""));
	elseif index == 3 then
		img.url = "ui://zhujiemian/zjm_28";
		obj.onClick:Set(zhujiemian_OnClose);
		obj.data = 3;
		effect :SetNativeObject(Proxy4Lua.GetEffectAssetGameObject(""));
	end
	
end

function zhujiemian_OnClose()
end

function zhujiemian_OnFolder()
	stateIcon:GetController("icon").selectedIndex = (stateIcon:GetController("icon").selectedIndex + 1) % 2;
	stateBar.selectedIndex = stateIcon:GetController("icon").selectedIndex;
end

function zhujiemian_OnCardCargo()
	UIManager.Show("paiku");
end

function zhujiemian_OnTaskBtn()
	UIManager.Show("daguanka");
end

function zhujiemian_OnSkillBtn()
	UIManager.Show("jineng");
end

function zhujiemian_OnQieCuoBtn()
	UIManager.Show("qiecuo");
end

function zhujiemian_OnHead()
	UIManager.Show("renwuziliao");
end

function zhujiemian_OnBagBtn()
	UIManager.Show("bagui");
end

function zhujiemian_OnShop()
	UIManager.Show("cangbaoge");
end