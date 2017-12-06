require "FairyGUI"

luxiang = fgui.window_class(WindowBase)
local Window;

local modalTimer;

local luxiangList;
local luxiangitem = "ui://luxiang/duizhan_com";

function luxiang:OnEntry()
	Window = luxiang.New();
	Window:Show();
end

function luxiang:GetWindow()
	return Window;
end

function luxiang:OnInit()
	self.contentPane = UIPackage.CreateObject("luxiang", "luxiang_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n45");

	luxiangList = self.contentPane:GetChild("n44").asList;

	luxiang_FlushData();
end

function luxiang:OnUpdate()
	if UIManager.IsDirty("luxiang") then
		luxiang_FlushData();
		UIManager.ClearDirty("luxiang");
	end
end

function luxiang:OnTick()
	if modalTimer ~= nil then
		modalTimer.crt = modalTimer.crt + 1;
		if modalTimer.crt > modalTimer.max then
			Window:CloseModalWait();
			modalTimer = nil;
		end
	end
end

function luxiang:isShow()
	return Window.isShowing;
end

function luxiang:OnDispose()
	Window:Dispose();
end

function luxiang:OnHide()
	BattleRecordSystem._BrDetail = nil;
	Window:Hide();
end

function luxiang_OnPlay(context)
	Proxy4Lua.RequestRecord(context.sender.data);
end

function luxiang_OnShare(context)
	if context.sender.data == nil then
		return;
	end

	local chat = COM_Chat.New();
	chat.Type = 1;--世界频道
	chat.PlayerInstId = GamePlayer._InstID;
	chat.PlayerName = GamePlayer._Name;
	chat.HeadIcon = GamePlayer.GetMyDisplayData()._HeadIcon;
	chat.Level = GamePlayer._Data.IProperties[9];
	chat.Content = Proxy4Lua.ChangeColor(GamePlayer._Name, "blue") .. "分享一段战斗录像,[url]点击观看[/url]";
	chat.RecordId = context.sender.data;
	Proxy4Lua.SendChat(chat);

	Proxy4Lua.PopMsg("分享成功");
end

function luxiang_FlushData()
	Window:ShowModalWait();
	modalTimer = {};
	modalTimer.crt = 0;
	modalTimer.max = 5;
	luxiangList:RemoveChildrenToPool();
	local shareBtnVisible = Proxy4Lua.LongIsEqual(BattleRecordSystem.MirrorPlayerId, GamePlayer._InstID);
	if BattleRecordSystem._BrDetail ~= nil then
		for i=0, BattleRecordSystem._BrDetail.Length -1 do
			------------------------getobjbegin-----------------------------
			local obj = luxiangList:AddItemFromPool(luxiangitem);
			local playBtn = obj:GetChild("n44").asButton;
			local shareBtn = obj:GetChild("n45").asButton;
			playBtn.data = BattleRecordSystem._BrDetail[i].ReportId;
			playBtn.onClick:Add(luxiang_OnPlay);
			shareBtn.data = BattleRecordSystem._BrDetail[i].ReportId;
			shareBtn.onClick:Add(luxiang_OnShare);
			shareBtn.visible = shareBtnVisible;

			--leftpart
			local lwinlose = obj:GetChild("n34").asLoader;
			local lhead = obj:GetChild("n37").asLoader;
			local llv = obj:GetChild("n39").asTextField;
			local lname = obj:GetChild("n8").asTextField;
			local lscore = obj:GetChild("n7").asTextField;
			local lscoreimg = obj:GetChild("n32").asLoader;
			lwinlose.url = "";
			lhead.url = "";
			llv.text = "";
			lname.text = "";
			lscore.text = "";
			lscoreimg.url = "";
			local ltouxiangkuang = {};
			for j=0, 9 do
				ltouxiangkuang[j] = obj:GetChild("n" .. (12 + j));
				ltouxiangkuang[j].visible = false;
			end

			--rightpart
			local rwinlose = obj:GetChild("n35").asLoader;
			local rhead = obj:GetChild("n41").asLoader;
			local rlv = obj:GetChild("n43").asTextField;
			local rname = obj:GetChild("n10").asTextField;
			local rscore = obj:GetChild("n9").asTextField;
			local rscoreimg = obj:GetChild("n33").asLoader;
			rwinlose.url = "";
			rhead.url = "";
			rlv.text = "";
			rname.text = "";
			rscore.text = "";
			rscoreimg.url = "";
			local rtouxiangkuang = {};
			for j=0, 9 do
				rtouxiangkuang[j] = obj:GetChild("n" .. (22 + j));
				rtouxiangkuang[j].visible = false;
			end
			------------------------getobjend-------------------------------

			for m=0, BattleRecordSystem._BrDetail[i].Players.Length - 1 do
				if Proxy4Lua.LongIsEqual(BattleRecordSystem._BrDetail[i].Players[m].InstId, BattleRecordSystem.MirrorPlayerId) then
					if Proxy4Lua.LongIsEqual(BattleRecordSystem._BrDetail[i].Winner, BattleRecordSystem._BrDetail[i].Players[m].InstId) then
						lwinlose.url = "ui://luxiang/013";
					else
						lwinlose.url = "ui://luxiang/014";
					end
					llv.text = BattleRecordSystem._BrDetail[i].Players[m].MainUnit.Level;
					lname.text = BattleRecordSystem._BrDetail[i].Players[m].MainUnit.Name;
					lscore.text = BattleRecordSystem._BrDetail[i].Players[m].TianTi;
					local lvurl = GamePlayer.RankLevel(BattleRecordSystem._BrDetail[i].Players[m].MainUnit.Level);
					lscoreimg.url = "ui://luxiang/xiao_duanwei" .. lvurl;

					local eData = EntityData.GetData(BattleRecordSystem._BrDetail[i].Players[m].MainUnit.UnitId);
					if eData ~= nil then
						local dData = DisplayData.GetData(eData._DisplayId);
						if dData ~= nil then
							lhead.url = "ui://" .. dData._HeadIcon;
						end
					end

					if BattleRecordSystem._BrDetail[i].Players[m].Units ~= nil then
						for z=0, BattleRecordSystem._BrDetail[i].Players[m].Units.Length - 1 do
							local lv = ltouxiangkuang[z]:GetChild("n6").asTextField;
							local fee = ltouxiangkuang[z]:GetChild("n7").asTextField;
							local head = ltouxiangkuang[z]:GetChild("n5").asLoader;
							local quality = ltouxiangkuang[z]:GetChild("n11").asLoader;
							fee.text = "";
							head.url = "";
							quality.url = "";
							ltouxiangkuang[z].visible = true;

							lv.text = BattleRecordSystem._BrDetail[i].Players[m].Units[z].Level;
							local eData = EntityData.GetData(BattleRecordSystem._BrDetail[i].Players[m].Units[z].UnitId);
							if eData ~= nil then
								fee.text = eData._Cost;
								local dData = DisplayData.GetData(eData._DisplayId);
								if dData ~= nil then
									head.url = "ui://" .. dData._HeadIcon;
									quality.url = "ui://" .. dData._Quality;
								end
							end
						end
					end
				else
					if Proxy4Lua.LongIsEqual(BattleRecordSystem._BrDetail[i].Winner, BattleRecordSystem._BrDetail[i].Players[m].InstId) then
						rwinlose.url = "ui://luxiang/013";
					else
						rwinlose.url = "ui://luxiang/014";
					end
					rlv.text = BattleRecordSystem._BrDetail[i].Players[m].MainUnit.Level;
					rname.text = BattleRecordSystem._BrDetail[i].Players[m].MainUnit.Name;
					rscore.text = BattleRecordSystem._BrDetail[i].Players[m].TianTi;
					local lvurl = GamePlayer.RankLevel(BattleRecordSystem._BrDetail[i].Players[m].MainUnit.Level);
					rscoreimg.url = "ui://luxiang/xiao_duanwei" .. lvurl;

					local eData = EntityData.GetData(BattleRecordSystem._BrDetail[i].Players[m].MainUnit.UnitId);
					if eData ~= nil then
						local dData = DisplayData.GetData(eData._DisplayId);
						if dData ~= nil then
							rhead.url = "ui://" .. dData._HeadIcon;
						end
					end

					if BattleRecordSystem._BrDetail[i].Players[m].Units ~= nil then
						for z=0, BattleRecordSystem._BrDetail[i].Players[m].Units.Length - 1 do
							local lv = rtouxiangkuang[z]:GetChild("n6").asTextField;
							local fee = rtouxiangkuang[z]:GetChild("n7").asTextField;
							local head = rtouxiangkuang[z]:GetChild("n5").asLoader;
							local quality = rtouxiangkuang[z]:GetChild("n11").asLoader;
							fee.text = "";
							head.url = "";
							quality.url = "";
							rtouxiangkuang[z].visible = true;

							lv.text = BattleRecordSystem._BrDetail[i].Players[m].Units[z].Level;
							local eData = EntityData.GetData(BattleRecordSystem._BrDetail[i].Players[m].Units[z].UnitId);
							if eData ~= nil then
								fee.text = eData._Cost;
								local dData = DisplayData.GetData(eData._DisplayId);
								if dData ~= nil then
									head.url = "ui://" .. dData._HeadIcon;
									quality.url = "ui://" .. dData._Quality;
								end
							end
						end
					end
				end
			end
		end
		if modalTimer ~= nil then
			modalTimer = nil;
		end
		Window:CloseModalWait();
	end
end