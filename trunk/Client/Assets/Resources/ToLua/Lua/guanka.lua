require "FairyGUI"

guanka = fgui.window_class(WindowBase)
local Window;
local tatle;
local desc;
local smallList;
local smallId;
local stamaPoint;
function guanka:OnEntry()      
	Window = guanka.New();
	Window:Show();
end

function guanka:OnInit()
	self.contentPane = UIPackage.CreateObject("guanka", "guanka_com").asCom;
	self:Center();
	self.modal = true;

	self.closeButton = self.contentPane:GetChild("n1").asButton;
	tatle = self.contentPane:GetChild("n15");
	desc = self.contentPane:GetChild("n8");
	local btn = self.contentPane:GetChild("n11");
	btn.onClick:Add(guanka_OnBattle);


	smallList = self.contentPane:GetChild("n13").asList;
	smallList :SetVirtual();
	smallList .itemRenderer = guakan_RenderListItem;
	stamaPoint = self.contentPane:GetChild("n19");
	guanka_FlushData();
end

function guanka:GetWindow()
	return Window;
end

function guanka:OnUpdate() 
	if UIManager.IsDirty("guanka") then
		guanka_FlushData();
		UIManager.ClearDirty("guanka");
	end
end



function guanka_OnBattle(context)
	smallId= 1;
	Proxy4Lua.ChallengeSmallChapter(smallId);
end

function guanka:OnTick()
	
end

function guanka:isShow()
	return Window.isShowing;
end

function guanka:OnDispose()
	Window:Dispose();
end

function guanka:OnHide()
	Window:Hide();
end

function guanka_FlushData()

	local guankaID = UIManager.GetWindow("jiehun").GetGuankaId();
	local data = HeroStroyData.GetData(guankaID);
	tatle.text = data.Name_;
	desc.text = data.Desc_;
	local data = CheckpointData.GetData(guankaID);
	smallList.numItems = data.Count;
	stamaPoint.text = GamePlayer._Data.IProperties[2];
end

function guakan_RenderListItem(index, obj)
	

end
