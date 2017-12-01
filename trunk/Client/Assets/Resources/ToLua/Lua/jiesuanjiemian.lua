require "FairyGUI"

jiesuanjiemian = fgui.window_class(WindowBase)
local Window;

local okBtn;
local okLbl;
local resultImg;

local star1;
local star2;
local star3;

local star1_eff;
local star2_eff;
local star3_eff;

local star1_ac;
local star2_ac;
local star3_ac;

local star1_des;
local star2_des;
local star3_des;

local headicon;
local level;
local exp;
local expBar;

local starEff = nil;
local starIdx = 0;

local pvpHideItem = {};

function jiesuanjiemian:OnEntry()
	UIManager.RegIDirty("zhujiemian");
	Define.LaunchUIBundle("icon");
	Window = jiesuanjiemian.New();
	Window:Show();
end

function jiesuanjiemian:GetWindow()
	return Window;
end

function jiesuanjiemian:OnInit()
	self.contentPane = UIPackage.CreateObject("jiesuanjiemian", "jiesuanjiemian_com").asCom;
	self:Center();
	self.modal = true;

	resultImg = self.contentPane:GetChild("n36").asLoader;

	okBtn = self.contentPane:GetChild("n33").asButton;
	okBtn.onClick:Add(jiesuanjiemian_OnOkBtn);
	okBtn.enabled = false;

	okLbl = self.contentPane:GetChild("n31");
	okLbl.visible = false;

	star1 = self.contentPane:GetChild("n29");
	star2 = self.contentPane:GetChild("n30");
	star3 = self.contentPane:GetChild("n32");

	star1_eff = self.contentPane:GetChild("n40");
	star2_eff = self.contentPane:GetChild("n41");
	star3_eff = self.contentPane:GetChild("n42");

	star1_ac = self.contentPane:GetChild("n23");
	star2_ac = self.contentPane:GetChild("n27");
	star3_ac = self.contentPane:GetChild("n28");

	star1_des = self.contentPane:GetChild("n20");
	star2_des = self.contentPane:GetChild("n21");
	star3_des = self.contentPane:GetChild("n22");

	headicon = self.contentPane:GetChild("n12").asLoader;
	level = self.contentPane:GetChild("n16");
	exp = self.contentPane:GetChild("n18");
	expBar = self.contentPane:GetChild("n17").asProgress;

	pvpHideItem.a = self.contentPane:GetChild("n7");
	pvpHideItem.b = self.contentPane:GetChild("n2");
	pvpHideItem.c = self.contentPane:GetChild("n4");
	pvpHideItem.d = self.contentPane:GetChild("n5");
	pvpHideItem.e = self.contentPane:GetChild("n6");
	pvpHideItem.f = self.contentPane:GetChild("n12");
	pvpHideItem.g = self.contentPane:GetChild("n14");
	pvpHideItem.h = self.contentPane:GetChild("n15");
	pvpHideItem.i = self.contentPane:GetChild("n16");
	pvpHideItem.j = self.contentPane:GetChild("n17");
	pvpHideItem.k = self.contentPane:GetChild("n18");
	pvpHideItem.l = self.contentPane:GetChild("n20");
	pvpHideItem.m = self.contentPane:GetChild("n21");
	pvpHideItem.n = self.contentPane:GetChild("n22");
	pvpHideItem.o = self.contentPane:GetChild("n23");
	pvpHideItem.p = self.contentPane:GetChild("n27");
	pvpHideItem.q = self.contentPane:GetChild("n28");
	pvpHideItem.r = self.contentPane:GetChild("n29");
	pvpHideItem.s = self.contentPane:GetChild("n30");
	pvpHideItem.t = self.contentPane:GetChild("n32");

	jiesuanjiemian_FlushData();
end

function jiesuanjiemian_FlushData()

	if Battle._BattleId == 0 or Battle._IsRecord then
		pvpHideItem.a.visible = false;
		pvpHideItem.b.visible = false;
		pvpHideItem.c.visible = false;
		pvpHideItem.d.visible = false;
		pvpHideItem.e.visible = false;
		pvpHideItem.f.visible = false;
		pvpHideItem.g.visible = false;
		pvpHideItem.h.visible = false;
		pvpHideItem.i.visible = false;
		pvpHideItem.j.visible = false;
		pvpHideItem.k.visible = false;
		pvpHideItem.l.visible = false;
		pvpHideItem.m.visible = false;
		pvpHideItem.n.visible = false;
		pvpHideItem.o.visible = false;
		pvpHideItem.p.visible = false;
		pvpHideItem.q.visible = false;
		pvpHideItem.r.visible = false;
		pvpHideItem.s.visible = false;
		pvpHideItem.t.visible = false;
	else
		pvpHideItem.a.visible = true;
		pvpHideItem.b.visible = true;
		pvpHideItem.c.visible = true;
		pvpHideItem.d.visible = true;
		pvpHideItem.e.visible = true;
		pvpHideItem.f.visible = true;
		pvpHideItem.g.visible = true;
		pvpHideItem.h.visible = true;
		pvpHideItem.i.visible = true;
		pvpHideItem.j.visible = true;
		pvpHideItem.k.visible = true;
		pvpHideItem.l.visible = true;
		pvpHideItem.m.visible = true;
		pvpHideItem.n.visible = true;
		pvpHideItem.o.visible = true;
		pvpHideItem.p.visible = true;
		pvpHideItem.q.visible = true;
		pvpHideItem.r.visible = true;
		pvpHideItem.s.visible = true;
		pvpHideItem.t.visible = true;
	end

	star1_des.text = "击杀";
	star2_des.text = "回合内胜利";
	star3_des.text = "阵亡卡牌不超过";

	star1.visible = false;
	star2.visible = false;
	star3.visible = false;

	star1_ac.text = "";
	star2_ac.text = "";
	star3_ac.text = "";

	local cpData = CheckpointData.GetDataByBattleID(Battle._BattleId);
	if cpData ~= nil then
		local eData = EntityData.GetData(cpData._Star1Need);
		if eData ~= nil then
			star1_des.text = star1_des.text .. eData._Name;
		end
		star2_des.text = cpData._Star2Need .. star2_des.text;
		star3_des.text = star3_des.text .. cpData._Star3Need .. "张";
	end
	
	local display = GamePlayer.GetMyDisplayData();
	if display ~= nil then
		headicon.url = "ui://" .. display._HeadIcon .. "_yuan";
	else
		headicon.url = "";
	end
	level.text = GamePlayer._Data.IProperties[9] .. "";
	if Battle._Result ~= nil then
		exp.text = "+" .. Battle._Result.Exp;
	else
		exp.text = "";
	end
	local needExp = ExpData.NeedExp(GamePlayer._Data.IProperties[9]);
	expBar.value = GamePlayer._Data.IProperties[4] / needExp * 100;

	if Battle.IsWin then
		resultImg.url = UIPackage.GetItemURL("jiesuanjiemian", "shengli");
		starEff = {};
		starEff.max = 15;
		starEff.count = 0;
		if Proxy4Lua.IsAchieve1 == true then
			star1_eff:SetNativeObject(Proxy4Lua.GetAssetGameObject("effect/lanhuoqiu", false));
		end
	else
		resultImg.url = UIPackage.GetItemURL("jiesuanjiemian", "shibai");
		okBtn.enabled = true;
		okLbl.visible = true;
		star1_ac.text = "未达成";
		star2_ac.text = "未达成";
		star3_ac.text = "未达成";
	end
end

function jiesuanjiemian:OnUpdate()
	if UIManager.IsDirty("jiesuanjiemian") then
		jiesuanjiemian_FlushData();
		UIManager.ClearDirty("jiesuanjiemian");
	end

	if starEff ~= nil then
		starEff.count = starEff.count + 1;
		if starEff.count > starEff.max then
			if starIdx == 0 then
				star1.visible = Proxy4Lua.IsAchieve1;
				starIdx = starIdx + 1;
				starEff.count = 0;
				if star1.visible then
					star1_ac.text = "达成";
					Proxy4Lua.ShakeCamera(Vector3.New(0.1, 0.1, 0), 0.3);
				end
				if Proxy4Lua.IsAchieve2 == true then
					star2_eff:SetNativeObject(Proxy4Lua.GetAssetGameObject("effect/lanhuoqiu", false));
				end
				return;
			end
			if starIdx == 1 then
				star2.visible = Proxy4Lua.IsAchieve2;
				starIdx = starIdx + 1;
				if star2.visible == false then
					starEff.count = starEff.max;
					return;
				end
				starEff.count = 0;
				if star2.visible then
					star2_ac.text = "达成";
					Proxy4Lua.ShakeCamera(Vector3.New(0.1, 0.1, 0), 0.3);
				end
				if Proxy4Lua.IsAchieve3 == true then
					star3_eff:SetNativeObject(Proxy4Lua.GetAssetGameObject("effect/lanhuoqiu", false));
				end
				return;
			end
			if starIdx == 2 then
				star3.visible = Proxy4Lua.IsAchieve3;
				if star2.visible == false then
					starIdx = 0;
					starEff = nil;
					okBtn.enabled = true;
					okLbl.visible = true;
					return;
				end
				starIdx = 0;
				starEff = nil;
				okBtn.enabled = true;
				okLbl.visible = true;
				if star3.visible then
					star3_ac.text = "达成";
					Proxy4Lua.ShakeCamera(Vector3.New(0.1, 0.1, 0), 0.3);
				end
			end
		end
	end
end

function jiesuanjiemian:OnTick()
	
end

function jiesuanjiemian:isShow()
	return Window.isShowing;
end

function jiesuanjiemian:OnDispose()
	Window:Dispose();
end

function jiesuanjiemian:OnHide()
	Window:Hide();
end

function jiesuanjiemian_OnOkBtn()
	starIdx = 0;
	okBtn.enabled = false;
	okLbl.visible = false;

	if Proxy4Lua.IsAchieve1 == true then
		star1_eff:SetNativeObject(Proxy4Lua.GetAssetGameObject("", false));
		Proxy4Lua.UnloadAsset("effect/lanhuoqiu");
	end
	if Proxy4Lua.IsAchieve2 == true then
		star2_eff:SetNativeObject(Proxy4Lua.GetAssetGameObject("", false));
		Proxy4Lua.UnloadAsset("effect/lanhuoqiu");
	end
	if Proxy4Lua.IsAchieve3 == true then
		star3_eff:SetNativeObject(Proxy4Lua.GetAssetGameObject("", false));
		Proxy4Lua.UnloadAsset("effect/lanhuoqiu");
	end

	if Battle.DropItemCount > 0 then
		UIManager.Hide("jiesuanjiemian");
		UIParamHolder.Set("showChaptersDrop", false);
		UIManager.Show("baowu");
	else
		if GamePlayer.showNewCard == true then
			UIManager.Show("huoderenwu");
			GamePlayer.showNewCard = false;
			Window:Hide();
		else
			SceneLoader.LoadScene("main");
		end
	end
end