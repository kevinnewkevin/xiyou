
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
	end
end

function RegGlobalValue()
	Define.Set("123", 5);
	Define.Set("1234", 3.4);
	Define.Set("1235", "水电费");
end