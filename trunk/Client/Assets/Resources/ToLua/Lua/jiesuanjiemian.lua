require "FairyGUI"

jiesuanjiemian = fgui.window_class(WindowBase)
local Window;

local okBtn;
local resultImg;

local star1;
local star2;
local star3;

local star1_ac;
local star2_ac;
local star3_ac;

local star1_des;
local star2_des;
local star3_des;

local gold;
local coin;

local headicon;
local level;
local exp;
local expBar;

function jiesuanjiemian:OnEntry()
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

	resultImg = self.contentPane:GetChild("n36").asLoader;

	okBtn = self.contentPane:GetChild("n33").asButton;
	okBtn.onClick:Add(jiesuanjiemian_OnOkBtn);

	star1 = self.contentPane:GetChild("n29");
	star2 = self.contentPane:GetChild("n30");
	star3 = self.contentPane:GetChild("n32");

	star1_ac = self.contentPane:GetChild("n23");
	star2_ac = self.contentPane:GetChild("n27");
	star3_ac = self.contentPane:GetChild("n28");

	star1_des = self.contentPane:GetChild("n20");
	star2_des = self.contentPane:GetChild("n21");
	star3_des = self.contentPane:GetChild("n22");

	gold = self.contentPane:GetChild("n10");
	coin = self.contentPane:GetChild("n11");

	headicon = self.contentPane:GetChild("n12").asLoader;
	level = self.contentPane:GetChild("n16");
	exp = self.contentPane:GetChild("n18");
	expBar = self.contentPane:GetChild("n17").asProgress;

	jiesuanjiemian_FlushData();
end

function jiesuanjiemian_FlushData()
	star1_des.text = "击杀";
	star2_des.text = "回合内胜利";
	star3_des.text = "阵亡卡牌少于";

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
		headicon.url = "ui://" .. display._HeadIcon;
	else
		headicon.url = "";
	end
	level.text = GamePlayer._Data.IProperties[4];
	exp.text = GamePlayer._Data.IProperties[5];

	if Battle._Result ~= nil then
		gold.text = Battle._Result.Money;
		coin.text = Battle._Result.Money;
	else
		gold.text = 0;
		coin.text = 0;
	end

	if Battle.IsWin then
		resultImg.url = UIPackage.GetItemURL("jiesuanjiemian", "shengli");
		star1.visible = Proxy4Lua.IsAchieve1;
		star2.visible = star1.visible and Proxy4Lua.IsAchieve2;
		star3.visible = star1.visible and star2.visible and Proxy4Lua.IsAchieve3;
	else
		resultImg.url = UIPackage.GetItemURL("jiesuanjiemian", "shibai");
		star1.visible = false;
		star2.visible = false;
		star3.visible = false;
	end

	if star1.visible then
		star1_ac.text = "达成";
	else
		star1_ac.text = "未达成";
	end
	if star2.visible then
		star2_ac.text = "达成";
	else
		star2_ac.text = "未达成";
	end
	if star3.visible then
		star3_ac.text = "达成";
	else
		star3_ac.text = "未达成";
	end
end

function jiesuanjiemian:OnUpdate()
	if UIManager.IsDirty("jiesuanjiemian") then
		jiesuanjiemian_FlushData();
		UIManager.ClearDirty("jiesuanjiemian");
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
	SceneLoader.LoadScene("main");
end