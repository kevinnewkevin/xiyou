require "FairyGUI"

guanka = fgui.window_class(WindowBase)
local Window;
local tatle;
local desc;
local smallList;
local smallId;
local stamaPoint;
local img;
local guankaID;
local needPower;
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
	img = self.contentPane:GetChild("n5");
	needPower= self.contentPane:GetChild("n10");
	local btn = self.contentPane:GetChild("n11");
	btn.onClick:Add(guanka_OnBattle);
	smallList = self.contentPane:GetChild("n13").asList;
	smallList :SetVirtual();
	smallList .itemRenderer = guakan_RenderListItem;
	stamaPoint = self.contentPane:GetChild("n19");
	guankaID = UIManager.GetWindow("jiehun").GetGuankaId();
	local smalldata = CheckpointData.GetData(guankaID);
	smallId = smalldata[0]._ID;
	smallList.selectedIndex = 0;
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
	guankaID = UIManager.GetWindow("jiehun").GetGuankaId();
	local data = HeroStroyData.GetData(guankaID);
	tatle.text = data.Name_;
	desc.text = data.Desc_;

	img.asLoader.url = "ui://" .. data.Icon_;
	local data = CheckpointData.GetData(guankaID);
	smallList.numItems = data.Count;

	stamaPoint.text = GamePlayer._Data.IProperties[2];
end

function guakan_RenderListItem(index, obj)
	 local data = CheckpointData.GetData(guankaID);
	 local name = obj:GetChild("n6");
	 name.text = data[index]._Name;
	 local open = obj:GetChild("n12");
	 open.visible  = false;
	 obj.data = data[index]._ID;
	 obj.onClick:Add(guakan_OnSelectGroup);
end

function guakan_OnSelectGroup(context)
	smallId = context.sender.data;
end
