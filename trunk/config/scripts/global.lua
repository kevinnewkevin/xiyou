
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
		MessageBox:SetData("提示", "匹配中 请稍后...", true);
	end
end

function RegGlobalValue()
	Define.Set("UIModelScale", 200);
	Define.Set("MaxFee", 5);
	Define.Set("MoveSpeed", 8);
	Define.Set("PointLight", "Effect/chushengdian");
	Define.Set("DestLight", "Effect/dianjiguangquan");
	Define.Set("BornPos", "20,-14.3,4");
	Define.Set("BattleCamera_plus", "5,5,5"); 		--偏移坐标
end