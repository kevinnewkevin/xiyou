require "FairyGUI"

zhanbao = fgui.window_class(WindowBase)
local Window;

local reportCom;
local side;
local quali;
local icon;
local skillIcon;
local skillName;
local targetList;
local targetListItemUrl = "ui://zhanbao/touxiangxiao_com";
local targetBuffItemUrl = "";

local outCom;
local outSide;
local outQuali;
local outIcon;

local outSkillCom;
local outSkillHeadSide;
local outSkillHeadQuali;
local outSkillHeadIcon;
local outSkillIcon;
local outSkillName;

function zhanbao:OnEntry()
	Define.LaunchUIBundle("icon");
	Window = zhanbao.New();
	Window:Show();
end

function zhanbao:GetWindow()
	return Window;
end

function zhanbao:OnInit()
	self.contentPane = UIPackage.CreateObject("zhanbao", "zhanbao_com").asCom;
	self:Center();

	self.closeButton = self.contentPane:GetChild("n4").asButton;

	reportCom = self.contentPane:GetChild("n6").asCom;
	local casterCom = reportCom:GetChild("n6").asCom;
	side = casterCom:GetChild("n2").asLoader;
	quali = casterCom:GetChild("n0").asLoader;
	icon = casterCom:GetChild("n1").asLoader;

	local skillCom = reportCom:GetChild("n7").asCom;
	skillIcon = skillCom:GetChild("n8").asLoader;
	skillName = skillCom:GetChild("n12").asTextField;

	targetList = reportCom:GetChild("n19").asList;

	outCom = self.contentPane:GetChild("n5").asCom;
	local outGroupCom = outCom:GetChild("n0").asCom;
	outSide = outGroupCom:GetChild("n2").asLoader;
	outQuali = outGroupCom:GetChild("n0").asLoader;
	outIcon = outGroupCom:GetChild("n1").asLoader;

	outSkillCom = self.contentPane:GetChild("n22").asCom;
	local outSkillHead = outSkillCom:GetChild("n0").asCom;
	outSkillHeadSide = outSkillHead:GetChild("n2").asLoader;
	outSkillHeadQuali = outSkillHead:GetChild("n0").asLoader;
	outSkillHeadIcon = outSkillHead:GetChild("n1").asLoader;

	local outSkillInfoCom = outSkillCom:GetChild("n4").asCom;
	outSkillIcon = outSkillInfoCom:GetChild("n8").asLoader;
	outSkillName = outSkillInfoCom:GetChild("n12").asTextField;

	zhanbao_FlushData();
end

function zhanbao:OnUpdate()
	if UIManager.IsDirty("zhanbao") then
		zhanbao_FlushData();
		UIManager.ClearDirty("zhanbao");
	end
end

function zhanbao:OnTick()
	
end

function zhanbao:isShow()
	return Window.isShowing;
end

function zhanbao:OnDispose()
	Window:Dispose();
end

function zhanbao:OnHide()
	Window:Hide();
end

function zhanbao_FlushData()
	reportCom.visible = false;
	side.url = "";
	quali.url = "";
	icon.url = "";
	skillIcon.url = "";
	skillName.text = "";
	targetList:RemoveChildrenToPool();

	outCom.visible = false;
	outSide.url = "";
	outQuali.url = "";
	outIcon.url = "";

	outSkillCom.visible = false;
	outSkillHeadSide.url = "";
	outSkillHeadQuali.url = "";
	outSkillHeadIcon.url = "";
	outSkillIcon.url = "";
	outSkillName.text = "";

	print("Battle._SelectReportIdx: " .. Battle._SelectReportIdx);
	if Battle._SelectReportIdx >= Battle._ReportTips.Count or Battle._SelectReportIdx < 0 then
		return;
	end

	if Battle._ReportTips[Battle._SelectReportIdx]._RBType == ReportBase.RBType.RBT_SelfSkill then
		outSkillCom.visible = true;
		local dData = GamePlayer.GetMyDisplayData();
		outSkillHeadSide.url = "ui://BattlePanel/zb_wo";
		outSkillHeadQuali.url = "ui://icon/touxiangkuang_bai";
		outSkillHeadIcon.url = "ui://" .. dData._HeadIcon;
		local sData = SkillData.GetData(Battle._ReportTips[Battle._SelectReportIdx]._SkillID);
		if sData ~= nil then
			outSkillIcon.url = "ui://" .. sData._Icon;
			outSkillName.text = sData._Name;
		end
	elseif Battle._ReportTips[Battle._SelectReportIdx]._RBType == ReportBase.RBType.RBT_SelfAppear then
		outCom.visible = true;
		outSide.url = "ui://BattlePanel/zb_wo";
		local eData = EntityData.GetData(Battle._ReportTips[Battle._SelectReportIdx]._CasterEntityID);
		if eData ~= nil then
			local dData = DisplayData.GetData(eData._DisplayId);
			if dData ~= nil then
				outIcon.url = "ui://" .. dData._HeadIcon;
				outQuali.url = "ui://icon/touxiangkuang_bai";
			else
				outQuali.url = "ui://icon/touxiangkuang_bai";
			end
		end
	elseif Battle._ReportTips[Battle._SelectReportIdx]._RBType == ReportBase.RBType.RBT_AllAppear then
		outCom.visible = true;
		outSide.url = "ui://BattlePanel/zb_di";
		local eData = EntityData.GetData(Battle._ReportTips[Battle._SelectReportIdx]._CasterEntityID);
		if eData ~= nil then
			local dData = DisplayData.GetData(eData._DisplayId);
			if dData ~= nil then
				outIcon.url = "ui://" .. dData._HeadIcon;
				outQuali.url = "ui://icon/touxiangkuang_bai";
			else
				outQuali.url = "ui://icon/touxiangkuang_bai";
			end
		end
	elseif Battle._ReportTips[Battle._SelectReportIdx]._RBType == ReportBase.RBType.RBT_AllSkill then
		reportCom.visible = true;
		if Battle._ReportTips[Battle._SelectReportIdx]._Self then
			sideStr = "ui://BattlePanel/zb_wo";
		else
			sideStr = "ui://BattlePanel/zb_di";
		end
		side.url = sideStr;
		quali.url = "ui://icon/touxiangkuang_bai";
		local eData = EntityData.GetData(Battle._ReportTips[Battle._SelectReportIdx]._CasterEntityID);
		if eData ~= nil then
			local dData = DisplayData.GetData(eData._DisplayId);
			if dData ~= nil then
				icon.url = "ui://" .. dData._HeadIcon;
			end
		end
		local sData = SkillData.GetData(Battle._ReportTips[Battle._SelectReportIdx]._SkillID);
		if sData ~= nil then
			skillIcon.url = "ui://" .. sData._Icon;
			skillName.text = sData._Name;
		end

		if Battle._ReportTips[Battle._SelectReportIdx]._Targets ~= nil then
			local target;
			local tEid;
			local tSelf;
			for i=0, Battle._ReportTips[Battle._SelectReportIdx]._Targets.Length - 1 do
				target = Battle._ReportTips[Battle._SelectReportIdx]._Targets[i];
				tEid = Battle._ReportTips[Battle._SelectReportIdx]._TargetEntityID[i];
				tSelf = Battle._ReportTips[Battle._SelectReportIdx]._TargetSelf[i];
				if target ~= nil then
					local targetItem = targetList:AddItemFromPool(targetListItemUrl);
					local tivalbar = targetItem:GetChild("n19");
					local tiside = targetItem:GetChild("n18").asLoader;
					local tiqua = targetItem:GetChild("n16").asLoader;
					local tiicon = targetItem:GetChild("n17").asLoader;
					local tichada = targetItem:GetChild("n23");--手牌删除
					local tidead = targetItem:GetChild("n24");--是否死亡
					local buffList = targetItem:GetChild("n25").asList;
					local dmg = targetItem:GetChild("n20").asTextField;
					local addhp = targetItem:GetChild("n22").asTextField;
					tiside.url = "";
					tiqua.url = "";
					tiicon.url = "";
					tivalbar.visible = false;
					if target.ActionParam > 0 then
						addhp.text = "+" .. target.ActionParam;
						dmg.text = "";
						tivalbar.visible = true;
					elseif target.ActionParam < 0 then
						dmg.text = target.ActionParam;
						addhp.text = "";
						tivalbar.visible = true;
					else
						dmg.text = "";
						addhp.text = "";
					end

					tidead.visible = target.Dead;
					print(target.ThrowCard.EntityId);
					if target.ThrowCard.EntityId ~= 0 then
						tEid = target.ThrowCard.EntityId;
					end
					tichada.visible = target.ThrowCard.EntityId ~= 0;

					local sideStr = "";
					if tSelf then
						sideStr = "ui://BattlePanel/zb_wo";
					else
						sideStr = "ui://BattlePanel/zb_di";
					end
					tiside.url = sideStr;
					local eData = EntityData.GetData(tEid);
					if eData ~= nil then
						local dData = DisplayData.GetData(eData._DisplayId);
						if dData ~= nil then
							tiicon.url = "ui://" .. dData._HeadIcon;
							tiqua.url = "ui://icon/touxiangkuang_bai";
						else
							tiqua.url = "ui://icon/touxiangkuang_bai";
						end
					end

					buffList:RemoveChildrenToPool();
					if target.BuffAdd ~= nil then
						for j=0, target.BuffAdd.Length - 1 do
							if target.BuffAdd[j].Change == 1 then
								local bufficon = buffList:AddItemFromPool(targetBuffItemUrl);
								local bloader = bufficon:GetChild("n0").asLoader;
								local bData = BuffData.GetData(target.BuffAdd[j].BuffId);
								bloader.url = "ui://" .. bData._Icon;
							end
						end
					end
				end
			end
		end
	end
end