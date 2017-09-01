require "FairyGUI"

qiecuo = fgui.window_class(WindowBase)
local Window;

local allCardGroupList;
local cardGroupUrl = "ui://qiecuo/paizuanniu_Button";


function qiecuo:OnEntry()
	Window = qiecuo.New();
	Window:Show();
end

function qiecuo:GetWindow()
	return Window;
end

function qiecuo:OnInit()
 	self.contentPane = UIPackage.CreateObject("qiecuo", "qiecuo_com").asCom;
	self:Center();
	self.closeButton = self.contentPane:GetChild("n7").asButton;

	local rightPart = self.contentPane:GetChild("n5").asCom;
	local bg = rightPart:GetChild("n3");
	allCardGroupList = bg:GetChild("n5").asList;

	for i=1, 5 do
		local groupItem = allCardGroupList:AddItemFromPool(cardGroupUrl);
		groupItem.onClick:Add(qiecuo_OnSelectGroup);
		groupItem.data = i - 1;
	end

	qiecuo_FlushData();
end



function qiecuo_RenderListItem(index, obj)

end

function qiecuo_OnDeleteGroup(context)
	local MessageBox = UIManager.ShowMessageBox();
end
 
function qiecuo_OnSetBattle(context)
	local MessageBox = UIManager.ShowMessageBox();
end





function qiecuo:OnUpdate()
	if UIManager.IsDirty("qiecuo") then
		qiecuo_FlushData();
		UIManager.ClearDirty("qiecuo");
	end
end

function qiecuo:OnTick()
	
end

function qiecuo:isShow()
	return Window.isShowing;
end

function qiecuo:OnDispose()
	Window:Dispose();
end

function qiecuo:OnHide()
	Window:Hide();
end

function qiecuo_FlushData()

	local groupItem;
	local groupName;
	for i=1, 5 do
			groupItem = allCardGroupList:GetChildAt(i-1);
			groupName = GamePlayer.GetGroupName(i - 1);
			if groupName == "" then
				groupName = "卡组" .. i;
			end
			groupItem:GetChild("n3").text = groupName;
			if crtGroupIdx == i - 1 then
				crtCardName.text = groupName;
			end
			if i - 1 == GamePlayer._CrtBattleGroupIdx then
				groupItem:GetChild("n4").visible = true;
			else
				groupItem:GetChild("n4").visible = false;
			end
		end
end



function qiecuo_OnSelectGroup(context)

	qiecuo_FlushData();
end

