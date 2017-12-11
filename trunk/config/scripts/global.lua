
--主入口函数。从这里开始lua逻辑
function Main()					
	 		
end

--场景切换通知
function OnLevelWasLoaded(level)
	collectgarbage("collect")
	Time.timeSinceLevelLoad = 0
end

--现在当选择人物场景的角色选择回调用
function ExcuteNpc(id)
	Proxy4Lua.FocusSelectRoleObject(id);
	--Proxy4Lua.SelectRole(id);
	--[[if id == 2 then
		Proxy4Lua.FocusNpcObject(2);
		return;
	end

	if id == 11 then
		UIManager.Show("jiehun");
		return;
	end

	if id == 12 then
		UIManager.Show("qiecuo");
		return;
	end

	local MessageBox = UIManager.ShowMessageBox();
	MessageBox:SetData("提示", "暂未开放,敬请期待", true);--]]
end

--网络异常处理
function NetWorkException(errCode)
	local MessageBox = UIManager.ShowMessageBox();
	if errCode == 10061 then
		MessageBox:SetData("提示", "服务器连接不上，请检查网络或服务器状态", true, NetWorkReconnect);
	else
		MessageBox:SetData("提示", "网络连接已断开", true, NetWorkReconnect);
	end
end

--网络重连
function NetWorkReconnect()
	UIManager.HideMessageBox();
	if Proxy4Lua.ReconnectServer() == true then
		if Proxy4Lua._LoginInfo.Username ~= nil or Proxy4Lua._LoginInfo.Username ~= "" then
        	Proxy4Lua.Login(Proxy4Lua._LoginInfo.Username, Proxy4Lua._LoginInfo.Password);
        end
	end
end

--错误处理
function ErrorMessage(errCode)
	local MessageBox = UIManager.ShowMessageBox();
	MessageBox:SetData("提示", errCode, true);
end

--I属性更新
function UpdateUnitIProperty(instid, type, value)
	if type == 11 then
		if value > GamePlayer._Data.IProperties[type] then
			Proxy4Lua.PopMsg("获得魂币" .. (value - GamePlayer._Data.IProperties[type]));
			local chatTip = COM_Chat.New();
			chatTip.Type = 0;
			chatTip.Content = "获得魂币" .. (value - GamePlayer._Data.IProperties[type]);
			ChatSystem.AddMsg(chatTip);
		end
	end
end

--C属性更新
function UpdateUnitCProperty(instid, type, value)
	
end

--获得物品
function GainItem(inst)
	local iData = ItemData.GetData(inst.ItemId);
	if iData ~= nil then
		local finalStr = "";
		local content = iData._Name;
		if iData._Quality == 2 then
			finalStr = "[color=#72ff5b]" .. content .. "[/color]";
		elseif iData._Quality == 3 then
			finalStr = "[color=#08e5ff]" .. content .. "[/color]";
		elseif iData._Quality == 4 then
			finalStr = "[color=#fa68ff]" .. content .. "[/color]";
		elseif iData._Quality == 5 then
			finalStr = "[color=#ff7a38]" .. content .. "[/color]";
		elseif iData._Quality == 6 then
			finalStr = "[color=#ff6868]" .. content .. "[/color]";
		else
			finalStr = "[color=#ffffff]" .. content .. "[/color]";
		end

		Proxy4Lua.PopMsg("获得" .. finalStr .. "X" .. inst.Stack);
		local chatTip = COM_Chat.New();
		chatTip.Type = 0;
		chatTip.Content = "获得" .. finalStr .. "X" .. inst.Stack;
		ChatSystem.AddMsg(chatTip);
	end
end

--谁支援了我
function WhoAssistantMe(who, itemid)
	local iData = ItemData.GetData(itemid);
	if iData ~= nil then
		local finalStr = "";
		local content = iData._Name;
		if iData._Quality == 2 then
			finalStr = "[color=#72ff5b]" .. content .. "[/color]";
		elseif iData._Quality == 3 then
			finalStr = "[color=#08e5ff]" .. content .. "[/color]";
		elseif iData._Quality == 4 then
			finalStr = "[color=#fa68ff]" .. content .. "[/color]";
		elseif iData._Quality == 5 then
			finalStr = "[color=#ff7a38]" .. content .. "[/color]";
		elseif iData._Quality == 6 then
			finalStr = "[color=#ff6868]" .. content .. "[/color]";
		else
			finalStr = "[color=#ffffff]" .. content .. "[/color]";
		end

		Proxy4Lua.PopMsg(who .. "支援了我1个" .. finalStr);
	end
end

--收到聊天
function ReceivedAssistant(ass)
	Proxy4Lua.PopMsg(ass.PlayerName .. "在支援频道发布了支援请求,快去帮助他吧!");
end

--创建帮派
function CreateGuild()
	Proxy4Lua.PopMsg("创建成功");
end

--离开帮派
function LeaveGuild()
	ChatSystem.ClearMsgByType(4);
	ChatSystem._Assistant:Clear();
end

--改变颜色ubb语法
function ChangeColor(content, color)
	if color == "blue" then
		return "[color=#08e5ff]" .. content .. "[/color]";
	elseif color == "red" then
		return "[color=#ff6868]" .. content .. "[/color]";
	elseif color == "yellow" then
		return "[color=#fff44a]" .. content .. "[/color]";
	elseif color == "green" then
		return "[color=#72ff5b]" .. content .. "[/color]";
	elseif color == "purple" then
		return "[color=#fa68ff]" .. content .. "[/color]";
	elseif color == "orange" then
		return "[color=#ff7a38]" .. content .. "[/color]";
	elseif color == "white" then
		return "[color=#ffffff]" .. content .. "[/color]";
	end
end

--常量定义
function RegGlobalValue()
	Define.Set("LogUrl", "http://106.75.78.151:8080/log");
	Define.Set("DebugServerAddress", "本机#127.0.0.1|外网#106.75.78.151");
	Define.Set("DebugServerPort", 10999);
	Define.Set("CenterServer", "");
	Define.Set("UIModelScale", 200);
	Define.Set("MaxFee", 5);
	Define.Set("MaxReportTips", 13);
	Define.Set("MoveSpeed_InBattle", 8);
	Define.Set("MoveSpeed_InWorld", 2.5);
	Define.Set("PointLight", "Effect/chushengdian");
	Define.Set("DestLight", "Effect/dianjiguangquan");
	Define.Set("BornPos", "8.55,-14.34,4");
	Define.Set("BattleCamera_plus", "5,5,5"); 		--偏移坐标
	Define.Set("WorldCamera_focusPlus", "0,0.8,0");	--主场景公告牌偏移坐标
	Define.Set("WeatherCheckTime", 60); --天气监测间隔 秒
	Define.Set("CreateMalePos", "-0.1065,0.2,-0.8456"); --男主角初始位置
	Define.Set("CreateMaleRotY", -29.17); --男主角初始旋转
	Define.Set("CreateFemalePos", "0.9614,0.33,-1.3568"); --女主角初始位置
	Define.Set("CreateFemaleRot", -6.651); --女主角初始旋转
	Define.Set("CreateSelectPos", "-1.52,-0.09,0.9"); --选中位置
	Define.Set("CreateDefaultClip", "stand"); --默认动画
	Define.Set("CreateSelectClip", "attack"); --选中动画
	Define.Set("BattleScenePool", "huangmozhandou,haidizhandou"); --战斗场景随机库
	Define.Set("AssistantCoolDown", 86400);
	Define.Set("EmojiTags","001,002,003,004,005,006,007,008,009,010,011,012,013,014,015,016,017,018,019,020,021,022,023,024,025,026,027,028");
	Define.Set("CreateGuild", 1);
	Define.Set("ChatMaxCount", 1000); --前端聊天最大保存数量
	Define.Set("BattleMaxTime", 30); --战斗倒计时
	
	--注册不弹popmsg提示的界面
	Define.RegNoPopUI("cangbaoge");
	Define.RegNoPopUI("BattlePanel");
end

--ui资源和ui逻辑关联map
function RegUIResMap()
	UIManager.RegUIResMap("jinengshengji", "qianghuachenggong");
	UIManager.RegUIResMap("squad", "bangpai");
	UIManager.RegUIResMap("squadHelp", "bangpai");
	UIManager.RegUIResMap("squadSetting", "bangpai");
	UIManager.RegUIResMap("squadRenming", "bangpai");
end