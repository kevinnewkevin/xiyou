require "FairyGUI"

youxiang = fgui.window_class(WindowBase)
local Window;


local mailList;
local contentLab;
local getBtn;
local timeLab;
local itemList;
local titleLab;
local _mailId;
local delBtn;

local itemUrl = "ui://youxiang/daoju_com";

function youxiang:OnEntry()
	Window = youxiang.New();
	Window:Show();
end

function youxiang:GetWindow()
	return Window;
end

function youxiang:OnInit()
	self.contentPane = UIPackage.CreateObject("youxiang", "youxiang_com").asCom;
	self:Center(); 
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n3");
	mailList =  self.contentPane:GetChild("n5");
	contentLab = self.contentPane:GetChild("n13");
	getBtn = self.contentPane:GetChild("n4");
	timeLab = self.contentPane:GetChild("n14");
	itemList = self.contentPane:GetChild("n18");
	titleLab = self.contentPane:GetChild("n12");
	delBtn = self.contentPane:GetChild("n21");
	timeLab.text ="";
	contentLab.text ="";
	titleLab.text ="";
	getBtn.visible = false;
	delBtn.visible = false;
	_mailId = 0;
	mailList.itemRenderer = youxiang_RenderListItem;
	--itemList.itemRenderer = youxiang_ItemListItem;
	getBtn.onClick:Add(youxiang_OnGetReward);
	delBtn.onClick:Add(youxiang_OnDel);
	youxiang_FlushData();
end




function youxiang:OnUpdate()
	if UIManager.IsDirty("youxiang") then
		youxiang_FlushData();
		UIManager.ClearDirty("youxiang");
	end
end

function youxiang:OnTick()
end

function youxiang:isShow()
	return Window.isShowing;
end

function youxiang:OnDispose()
	Window:Dispose();
end

function youxiang:OnHide()
	_mailId = 0;
	getBtn.visible = false;
	delBtn.visible = false;
	itemList:RemoveChildrenToPool(); 
	Window:Hide();
end

function youxiang_FlushData()
	 mailList.numItems = MailSystem.mailList.Count;
	 updateMainInfo();
end 

function youxiang_RenderListItem(index, obj)
	local mail = MailSystem.mailList[index];
	obj:GetChild("n4").text = mail.Title;
	obj.data = mail.MailId;
	obj.onClick:Add(youxiang_OnMail);
	if mail.IsRead then
		obj:GetChild("n3").asLoader.url = "ui://youxiang/02";
	else
		obj:GetChild("n3").asLoader.url = "ui://youxiang/01";
	end
end


function youxiang_ItemListItem(index, obj)
	local data = MailSystem.GetMail(_mailId);
	local item = ItemData.GetData(data.Items[index].ItemId);
	local itemIcon = obj:GetChild("n17");
	itemIcon.asLoader.url = "ui://" .. item._Icon;
	obj:GetChild("n16").asLoader.url = "ui://" .. item._IconBack;  
	itemIcon.onClick:Add(youxiang_OnTtItem);
	itemIcon.data = data.Items[index].ItemId;
end

function youxiang_OnTtItem(context)
	local iId = context.sender.data;
	UIParamHolder.Set("tipsItem", iId);
	UIManager.Show("bagtips");
end

function youxiang_OnMail(content)
	_mailId = content.sender.data;
	updateMainInfo();
end

function youxiang_OnGetReward(content)
	Proxy4Lua.GetMailItem(_mailId);
end

function youxiang_OnDel(content)
	Proxy4Lua.DelMail(_mailId);
end

function updateMainInfo()

	local data = MailSystem.GetMail(_mailId);
	if data  == nil then
		timeLab.text ="";
		contentLab.text ="";
		titleLab.text ="";
		--itemList.numItems = 0;
		getBtn.visible = false;
		delBtn.visible = false;
	else
		if data.IsRead == false then
			Proxy4Lua.ReadMail(_mailId);
		end

		--timeLab.text = data.MailTimestamp.."";
		contentLab.text =data.Content;
		titleLab.text =data.SendPlayerName;
		if data.Items ~= nil  then
			--itemList.numItems = data.Items.Length;
			getBtn.visible = true;
		else
			--itemList.numItems = 0;
			getBtn.visible = false;
			delBtn.visible = true;
		end 


		itemList:RemoveChildrenToPool(); 

		if data.Gold ~= 0 then
			local item = itemList:AddItemFromPool(itemUrl);
		 		item:GetChild("n16").asLoader.url =  "ui://icon/daojukuang_bai";
		 	item:GetChild("n17").asLoader.url = "ui://icon/jinbi_icon" ;
		 	item:GetChild("n17").onClick:Remove(youxiang_OnTtItem); 
		 	getBtn.visible = true;
		end

		if data.Copper ~= 0 then
			local item1 = itemList:AddItemFromPool(itemUrl);
		 	item1:GetChild("n16").asLoader.url =  "ui://icon/daojukuang_bai";
		 	item1:GetChild("n17").asLoader.url = "ui://icon/jinbi_icon" ;
		 	item1:GetChild("n17").onClick:Remove(youxiang_OnTtItem); 
		 	getBtn.visible = true;
		end

		if data.Hero ~= 0 then
			local item2 = itemList:AddItemFromPool(itemUrl);
		 	local entityData = EntityData.GetData(data.Hero);
			local displayData = DisplayData.GetData(entityData._DisplayId);
			item2:GetChild("n16").asLoader.url = "ui://" .. displayData._Quality;
			item2:GetChild("n17").asLoader.url = "ui://" .. displayData._HeadIcon;
			item2:GetChild("n17").onClick:Remove(youxiang_OnTtItem); 
			getBtn.visible = true;
		end

		if data.Items ~= nil then
			for i=1, data.Items.Length do
				local item3 = itemList:AddItemFromPool(itemUrl);
				local data = MailSystem.GetMail(_mailId);
				local item = ItemData.GetData(data.Items[i-1].ItemId);
				local itemIcon1 = item3:GetChild("n17");
				itemIcon1.asLoader.url = "ui://" .. item._Icon;
				item3:GetChild("n16").asLoader.url = "ui://" .. item._IconBack;  
				itemIcon1.onClick:Add(youxiang_OnTtItem);
				itemIcon1.data = data.Items[i-1].ItemId;
			end
		end

	end

end


