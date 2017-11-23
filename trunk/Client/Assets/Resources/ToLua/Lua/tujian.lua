require "FairyGUI"

tujian = fgui.window_class(WindowBase)
local Window;

local cardItemUrl = "ui://tujian/touxiangkuang_Label";
local crtCardsFee = 0;
local crtCardIdx = 0; --当前所有卡的起始索引
local countPerPage = 9; --所有卡组列表 每页的最大数量
local crtCardsType = 0;
local selectCardId;

local infoPanel;
local modelPanel;
local propPanel;

local nameLab;
local holder;
local modelRes;
local fee;
local race;
local skillList;
local hp;
local agility;
local atk;
local def;
local matk;
local mdef;
local crit;
local critdmg;
local combo;
local split;
local rec;
local reflect;
local suck;
local desc;
local icon;
local leftBtn;
local rightBtn;
local raceBack;
local infoList;
local crtSelectIdx;
local leftBtnImg;
local rightBtnImg;

local shihunBtn;

function tujian:OnEntry()
	Define.LaunchUIBundle("Card");
	Window = tujian.New();
	Window:Show();
end

function tujian:GetWindow()
	return Window;
end

function tujian:OnInit()
	self.contentPane = UIPackage.CreateObject("tujian", "tujian_com").asCom;
	self:Center(); 
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n4");
	local leftPart = self.contentPane:GetChild("n97").asCom;
	selectCardId = 3;
	crtSelectIdx = 0;
    allCardList = leftPart:GetChild("n27").asList;
	--allCardList.onDrop:Add(tujian_OnDropCard);
	allCardList.data = 0;
	allCardList.scrollItemToViewOnClick = false;
	--total = leftPart:GetChild("n33").asTextField;
	prevBtn = leftPart:GetChild("n36");
	nextBtn = leftPart:GetChild("n37");
	pageText = leftPart:GetChild("n39");
	prevBtn.onClick:Add(tujian_OnPrevPage);
	nextBtn.onClick:Add(tujian_OnNextPage);
	infoList = self.contentPane:GetChild("n109").asList;
	leftBtn = self.contentPane:GetChild("n107");
	rightBtn = self.contentPane:GetChild("n108");
	leftBtnImg = leftBtn:GetChild("n2");
	rightBtnImg = rightBtn:GetChild("n2");
	leftBtn.onClick:Add(tujian_OnLeftBtn);
	rightBtn.onClick:Add(tujian_OnRightBtn);
	raceBack = self.contentPane:GetChild("n111");
	infoPanel = infoList:GetChildAt(1);
	modelPanel = infoList:GetChildAt(0);
	propPanel = infoList:GetChildAt(2);
	holder = modelPanel:GetChild("n63");
	fee = modelPanel:GetChild("n58");
	race = modelPanel:GetChild("n89");
	nameLab = modelPanel:GetChild("n60");

	skillList = propPanel:GetChild("n237").asList;
	hp = propPanel:GetChild("n242");
	agility = propPanel:GetChild("n246");
	atk = propPanel:GetChild("n213");
	def = propPanel:GetChild("n214");
	matk = propPanel:GetChild("n215");
	mdef = propPanel:GetChild("n216");
	crit = propPanel:GetChild("n217");
	critdmg = propPanel:GetChild("n218");
	combo = propPanel:GetChild("n219");
	split = propPanel:GetChild("n220");
	rec = propPanel:GetChild("n221");
	reflect = propPanel:GetChild("n222");
	suck = propPanel:GetChild("n223");
	icon = infoPanel:GetChild("n98");
	local feeList = leftPart:GetChild("n34").asList;
	local feeMax = feeList.numItems;
	local feeItem;
	for i=1, feeMax do
		feeItem = feeList:GetChildAt(i-1);
		feeItem.data = i-1;
		feeItem.onClick:Add(tujian_OnFeeItemClick);
	end
	feeList.selectedIndex = crtCardsFee;

	local typeList = leftPart:GetChild("n35").asList;
	local typeMax = typeList.numItems;
	local typeItem;
	for i=1, typeMax do
		typeItem = typeList:GetChildAt(i-1);
		typeItem.data = i-1;
		typeItem.onClick:Add(tujian_OnTypeItemClick);
	end
	for i=2, 4 do
		local skill = skillList:GetChildAt(i - 2);
		skill.onClick:Add(tujian_OnSkillBtn);
	end
	typeList.selectedIndex = crtCardsType;
	shihunBtn = self.contentPane:GetChild("n106");
	shihunBtn.onClick:Add(tujian_OnShiHunClick);
	infoList.scrollPane.onScroll:Add(DoSpecialEffect);
	infoList.scrollItemToViewOnClick = false;
	tujian_FlushData();
end


function DoSpecialEffect()
		local midX = infoList.scrollPane.posX + infoList.viewWidth / 2;
		local cnt = infoList.numChildren;

			for  i = 1, cnt do
				local obj = infoList:GetChildAt(i-1);
				local dist = Mathf.Abs(midX - obj.x - obj.width / 2);
				if dist > obj.width then

				else
					crtSelectIdx = i-1;
				end
			end

		tujian_UpdateLRBtn();
		print(crtSelectIdx );
end

function tujian:OnUpdate()
	if UIManager.IsDirty("tujian") then
		tujian_FlushData();
		UIManager.ClearDirty("tujian");
	end
	
end

function tujian:OnTick()
	
end

function tujian:isShow()
	return Window.isShowing;
end

function tujian:OnDispose()
	Window:Dispose();
end

function tujian:OnHide()
	crtCardsFee = 0;
	crtCardIdx = 0;
	crtCardsType = 0;
	selectCardId = 3;
	crtSelectIdx = 0;
	allCardList.data = 0;
	Proxy4Lua.ClearToDeleteAsset("tujian");
	Window:Hide();
end

function tujian_FlushData()
	shihunBtn.enabled = JieHunSystem.instance._NextDrawData ~= nil;

	--total.text = "(数量:" .. GamePlayer.CardsByFeeAndType(0, 0).Count .. ")";
	local page = crtCardIdx / countPerPage + 1;
	pageText.text = "第" .. Proxy4Lua.ConvertToChineseNumber(page) .. "页";

	local cardList = EntityData.CardsByFeeAndType(crtCardsFee, crtCardsType);
	allCardList:RemoveChildrenToPool();

	if cardList ~= nil then
		local max = crtCardIdx + countPerPage;
		if max > cardList.Count then
			max = cardList.Count;
		end
		for i=crtCardIdx, max - 1 do
			tujian_RenderListItem(i, allCardList:AddItemFromPool(cardItemUrl));
		end
		nextBtn.enabled = crtCardIdx + countPerPage < cardList.Count;
		prevBtn.enabled = crtCardIdx ~= 0;
	else
		nextBtn.enabled = false;
		prevBtn.enabled = false;
	end

	tujian_updateModelInfo();
	tujian_updatePropInfo();
	tujian_updateInfoInfo();
	tujian_UpdateLRBtn();
end 


function tujian_OnNextPage()
	local cardList = EntityData.CardsByFeeAndType(crtCardsFee, crtCardsType);
	if cardList ~= nil then
		if crtCardIdx + countPerPage <= cardList.Count then
			crtCardIdx = crtCardIdx + countPerPage;
		end
	end
	UIManager.SetDirty("tujian");
end

function tujian_OnPrevPage()
	crtCardIdx = crtCardIdx - countPerPage;
	if crtCardIdx < 0 then
		crtCardIdx = 0;
	end
	UIManager.SetDirty("tujian");
end

function tujian_OnFeeItemClick(context)
	crtCardsFee = context.sender.data;
	crtCardIdx = 0;
	UIManager.SetDirty("tujian");
end

function tujian_OnTypeItemClick(context)
	crtCardsType = context.sender.data;
	crtCardIdx = 0;
	UIManager.SetDirty("tujian");
end

function tujian_RenderListItem(index, obj)
	local edata = EntityData.GetDisplayDataByIndex(crtCardsFee, crtCardsType, index);
	local displayData = DisplayData.GetData(edata._DisplayId);

	local cdata = JieHunSystem.instance:GetChapterData(edata._ChapterID);

	if cdata ~= nil then
		obj:GetChild("n13").visible = false;
		local smallChapters = cdata.SmallChapters;
		local len = smallChapters.Length;
		local finish = true;
   		for i = 1, len do
			if smallChapters[i -1].Star1 ~= true or smallChapters[i -1].Star2 ~= true or smallChapters[i -1].Star3 ~= true  then    
				finish = false;
			end
	 	end
	 	obj:GetChild("n12").visible = finish;
	else
		obj:GetChild("n13").visible = true;
		obj:GetChild("n12").visible = false;
	end

	obj:GetChild("n5").asLoader.url = "ui://" .. displayData._HeadIcon;
	obj:GetChild("n7").text = edata._Cost.."";
	obj.onClick:Add(tujian_OnCardItemClick);
	obj.data = edata._UnitId;

end

function tujian_OnCardItemClick(context)
	selectCardId = context.sender.data;
	tujian_updateModelInfo();
	tujian_updatePropInfo();
	tujian_updateInfoInfo();
end

function tujian_updateModelInfo()
	local entityData = EntityData.GetData(selectCardId);
	local displayData =  DisplayData.GetData(entityData._DisplayId);
	modelRes = displayData._AssetPath; 
	holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes, false, 1200, 1.5));
	race.asLoader.url = "ui://" .. displayData._Race;
	Proxy4Lua.AddToDelete("tujian",modelRes);
	fee.text = entityData._Cost;
	nameLab.text = entityData._Name;
	raceBack.asLoader.url = "ui://" .. displayData._Race.. "_bj";
end 

function tujian_updatePropInfo()
	local entityData = EntityData.GetData(selectCardId);
	hp.text = entityData.CPT_HP .. "";
	agility.text = entityData.CPT_AGILE .. "";
	atk.text = entityData.CPT_ATK .. "";
	def.text = entityData.CPT_DEF .. "";
	matk.text = entityData.CPT_MAGIC_ATK .. "";
	mdef.text = entityData.CPT_MAGIC_DEF .. "";
	crit.text = entityData.CPT_CRIT .. "";
	critdmg.text =  "0";--entityData.;
	combo.text = "0";--entityData.;
	split.text = "0";--entityData.;
	rec.text = "0";--entityData.;
	reflect.text = "0";--entityData.;
	suck.text = "0";--entityData.;

	for i=2, 4 do
		local skill = skillList:GetChildAt(i - 2);
		local sData = SkillData.GetData(entityData._Skills[i-1]);
		local loader = skill:GetChild("n8").asLoader;
		local lv = skill:GetChild("n7").asTextField;
		if sData ~= nil then
			loader.url = "ui://" .. sData._Icon;
			lv.text = sData._Level;
			skill.data = sData._Id;
		else
			loader.url = "";
			lv.text = "";
		end
	end
end

function tujian_updateInfoInfo()
	local entityData = EntityData.GetData(selectCardId);
	local displayData =  DisplayData.GetData(entityData._DisplayId);
	icon.asLoader.url = "ui://" .. displayData._CardIcon;
end

function tujian_OnSkillBtn(context)
	if context.sender.data == nil or context.sender.data == 0 then
		return;
	end

	UIParamHolder.Set("jinengxiangqing", context.sender.data);
	UIManager.Show("jinengxiangqing");
end
  
function tujian_OnShiHunClick(context)
	UIManager.Show("shihun");
end

function tujian_OnRightBtn(context) 
	if crtSelectIdx == 2 then
		return;
	end
	crtSelectIdx = crtSelectIdx+1;
	infoList:ScrollToView(crtSelectIdx, false);
	tujian_UpdateLRBtn();
end

function  tujian_OnLeftBtn(context)
	if crtSelectIdx == 0 then
		return;
	end
	crtSelectIdx = crtSelectIdx-1;
	infoList:ScrollToView(crtSelectIdx, false);
	tujian_UpdateLRBtn();
end


function  tujian_UpdateLRBtn()
	if crtSelectIdx == 0 then
		leftBtn.visible = false;
		rightBtn.visible = true;
		rightBtnImg.asLoader.url = "ui://tujian/56";
	elseif crtSelectIdx == 1 then
		leftBtn.visible = true;
		rightBtn.visible = true;
		leftBtnImg.asLoader.url = "ui://tujian/58";
		rightBtnImg.asLoader.url = "ui://tujian/57";
	elseif crtSelectIdx == 2 then
		leftBtn.visible = true;
		rightBtn.visible = false;
		leftBtnImg.asLoader.url = "ui://tujian/56";
	end
end
