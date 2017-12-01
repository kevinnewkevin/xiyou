require "FairyGUI"

guankaluxiang = fgui.window_class(WindowBase)
local Window;

local guankaluxiangList;
local guankaluxiangitem = "ui://guankaluxiang/luxiangguanka_com";

function guankaluxiang:OnEntry()
	Define.LaunchUIBundle("guankatupian");
	Window = guankaluxiang.New();
	Window:Show();
end

function guankaluxiang:GetWindow()
	return Window;
end

function guankaluxiang:OnInit()
	self.contentPane = UIPackage.CreateObject("guankaluxiang", "guankaluxiang_com").asCom;
	self:Center();
	self.modal = true;
	self.closeButton = self.contentPane:GetChild("n45");

	guankaluxiangList = self.contentPane:GetChild("n44").asList;

	guankaluxiang_FlushData();
end

function guankaluxiang:OnUpdate()
	if UIManager.IsDirty("guankaluxiang") then
		guankaluxiang_FlushData();
		UIManager.ClearDirty("guankaluxiang");
	end
end

function guankaluxiang:OnTick()
	
end

function guankaluxiang:isShow()
	return Window.isShowing;
end

function guankaluxiang:OnDispose()
	Window:Dispose();
end

function guankaluxiang:OnHide()
	BattleRecordSystem._BrDetail = nil;
	Window:Hide();
end

function guankaluxiang_OnPlay(context)
	Proxy4Lua.RequestRecord(context.sender.data);
end

--function guankaluxiang_OnShare(context)
--	if context.sender.data == nil then
--		return;
--	end
--
--	local chat = COM_Chat.New();
--	chat.Type = 1;--世界频道
--	chat.PlayerInstId = GamePlayer._InstID;
--	chat.PlayerName = GamePlayer._Name;
--	chat.HeadIcon = GamePlayer.GetMyDisplayData()._HeadIcon;
--	chat.Level = GamePlayer._Data.IProperties[9];
--	chat.Content = Proxy4Lua.ChangeColor(GamePlayer._Name, "blue") .. "分享一段战斗录像,[url]点击观看[/url]";
--	chat.AudioId = context.sender.data;
--	Proxy4Lua.SendChat(chat);
--end

function guankaluxiang_FlushData()
	Window:ShowModalWait();
	guankaluxiangList:RemoveChildrenToPool();
--	local shareBtnVisible = Proxy4Lua.LongIsEqual(BattleRecordSystem.MirrorPlayerId, GamePlayer._InstID);
	if BattleRecordSystem._BrDetail ~= nil then
		for i=0, BattleRecordSystem._BrDetail.Length -1 do
			------------------------getobjbegin-----------------------------
			local obj = guankaluxiangList:AddItemFromPool(guankaluxiangitem);
			local playBtn = obj:GetChild("n46").asButton;
			local shareBtn = obj:GetChild("n47").asButton;
			playBtn.data = BattleRecordSystem._BrDetail[i].ReportId;
			playBtn.onClick:Add(guankaluxiang_OnPlay);
--			shareBtn.data = BattleRecordSystem._BrDetail[i].ReportId;
--			shareBtn.onClick:Add(guankaluxiang_OnShare);
			shareBtn.visible = false;

			--leftpart
			local lhead = obj:GetChild("n37").asLoader;
			local llv = obj:GetChild("n39").asTextField;
			local lname = obj:GetChild("n8").asTextField;
			local lscore = obj:GetChild("n7").asTextField;
			local lscoreimg = obj:GetChild("n32").asLoader;
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
			local rfbimg = obj:GetChild("n45").asLoader;
			rfbimg.url = "";
			------------------------getobjend-------------------------------

			for m=0, BattleRecordSystem._BrDetail[i].Players.Length - 1 do
				if Proxy4Lua.LongIsEqual(BattleRecordSystem._BrDetail[i].Players[m].InstId, BattleRecordSystem.MirrorPlayerId) then
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
					local cpData = CheckpointData.GetDataByBattleID(BattleRecordSystem._BrDetail[i].Battleid);
					if cpData ~= nil then
						rfbimg.url = "ui://" .. cpData._pic;
					end
				end
			end
		end
		Window:CloseModalWait();
	end
end