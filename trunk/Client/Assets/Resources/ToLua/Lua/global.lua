
--主入口函数。从这里开始lua逻辑
function Main()					
	 		
end

--场景切换通知
function OnLevelWasLoaded(level)
	collectgarbage("collect")
	Time.timeSinceLevelLoad = 0
end

function ExcuteNpc(id)
	if id == 1 then
		Proxy4Lua.BattleJoin();
		local MessageBox = UIManager.ShowMessageBox();
		MessageBox:SetData("提示", "匹配中...", true);
	end
end

function RegGlobalValue()
	Define.Set("UIModelScale", 200);
end