require "FairyGUI"

yemiantishi = fgui.window_class(WindowBase)

local Window;

local contentList;
local itemUrl = "ui://yemiantishi/tishi_com";

local checkTime;

local lifeTime;

function yemiantishi:OnEntry()
	Window = yemiantishi.New();
	Window:Show();
end

function yemiantishi:GetWindow()
	return Window;
end

function yemiantishi:OnInit()
	self.contentPane = UIPackage.CreateObject("yemiantishi", "yemiantishi_com").asCom;
	self:Center();

	contentList = self.contentPane:GetChild("n2").asList;
	contentList.touchable = false;

	self.sortingOrder = 2147482647;

	checkTime = {};
	checkTime.crt = 0;
	checkTime.max = 0.5;
end

function yemiantishi:OnUpdate()
	if checkTime ~= nil then
		checkTime.crt = checkTime.crt + Proxy4Lua.DeltaTime;
		if checkTime.crt >= checkTime.max then
			yemiantishi_FlushData();
			checkTime.crt = 0;
			checkTime.max = 0.5;
		end
	end

	if lifeTime ~= nil then
		lifeTime.crt = lifeTime.crt + Proxy4Lua.DeltaTime;
		if lifeTime.crt >= lifeTime.max then
			if contentList.numItems > 0 then
				contentList:RemoveChildToPoolAt(0);
			end
			if contentList.numItems > 0 then
				lifeTime.crt = 0;
				lifeTime.max = 0.5;
			end
		end
	end
end

function yemiantishi:OnTick()
	
end

function yemiantishi:isShow()
	return Window.isShowing;
end

function yemiantishi:OnDispose()
	Window:Dispose();
end

function yemiantishi:OnHide()
	--Window:Hide();
end

function yemiantishi_FlushData()
	Window:BringToFront();
	if Proxy4Lua.Message.Count > 0 then
		local msg = Proxy4Lua.Message:Dequeue();
		local msgCom = contentList:AddItemFromPool(itemUrl);
		msgCom.touchable = false;
		local trans = msgCom:GetTransition("t0");
		trans:Play();
		local msgt = msgCom:GetChild("n1").asTextField;
		msgt.text = msg;

		if contentList.numItems == 1 then
			lifeTime = {};
			lifeTime.crt = 0;
			lifeTime.max = 2;
		end
	end

	if contentList.numItems > 5 then
		contentList:RemoveChildToPoolAt(0);
	end
end