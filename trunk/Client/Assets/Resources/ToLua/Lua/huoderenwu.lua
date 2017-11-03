require "FairyGUI"

huoderenwu = fgui.window_class(WindowBase)
local Window;


local nameLab;
local race;
local needPower;
local modelRes;
local back;
local holder;
local hitNextBtn;
local infoBtn;
function huoderenwu:OnEntry()
	Window = huoderenwu.New();
	Window:Show();
end

function huoderenwu:GetWindow()
	return Window;
end

function huoderenwu:OnInit()
	self.contentPane = UIPackage.CreateObject("huoderenwu", "huoderenwu_com").asCom;
	self:Center();
	self.modal = true;
    self.closeButton = self.contentPane:GetChild("n10");
    nameLab = self.contentPane:GetChild("n4");
    race = self.contentPane:GetChild("n7");
    needPower = self.contentPane:GetChild("n8");
    holder = self.contentPane:GetChild("n1").asGraph;
    infoBtn = self.contentPane:GetChild("n12");
    infoBtn.onClick:Add(huderenwu_OnInfo);
	huoderenwu_FlushData();
end

function huoderenwu:OnUpdate()  
	if UIManager.IsDirty("huoderenwu") then
		huoderenwu_FlushData();
		UIManager.ClearDirty("huoderenwu");
	end
end

function huderenwu_OnExit(context)
	holder:SetNativeObject(Proxy4Lua.GetAssetGameObject("", false));
	Proxy4Lua.UnloadAsset(modelRes);
	modelRes = "";
	SceneLoader.LoadScene("main");
end

function huderenwu_OnInfo(context)
  local card = GamePlayer.newCard;
    local displayData = GamePlayer.GetDisplayDataByInstID(card.InstId);
	UIParamHolder.Set("qiecuo1", card.InstId);
	UIParamHolder.Set("qiecuo2", true);
	UIParamHolder.Set("showBoos",false);
	UIManager.Show("xiangxiziliao");
end


function huoderenwu:OnTick()
	
end

function huoderenwu:isShow()
	return Window.isShowing;
end

function huoderenwu:OnDispose()
	Window:Dispose();
end

function huoderenwu:OnHide()
	holder:SetNativeObject(Proxy4Lua.GetAssetGameObject("", false));
	Proxy4Lua.UnloadAsset(modelRes);
	modelRes = "";
	SceneLoader.LoadScene("main");
	Window:Hide();
end

function huoderenwu_FlushData()

    local card = GamePlayer.newCard;
    local displayData = GamePlayer.GetDisplayDataByInstID(card.InstId);
    modelRes = displayData._AssetPath;
    race.asLoader.url = "ui://" .. displayData._Race;
   -- local wrapper = Proxy4Lua.GetAssetGameObject(modelRes, true);
    --local anim = wrapper.wrapTarget:GetComponent("Animation");
    --anim.Play("attack");
    --holder:SetNativeObject(wrapper);
    holder:SetNativeObject(Proxy4Lua.GetAssetGameObject(modelRes, true));
    local entityData = GamePlayer.GetEntityDataByInstID(card.InstId);
    needPower.text = entityData._Cost;
    nameLab.text = entityData._Name;

end
