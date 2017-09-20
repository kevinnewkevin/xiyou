-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加 暴击几率
sys.log("buff128")

function buff_128_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_KILL") 
	
	sys.log("buff_128_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_128_update(battleid, buffinstid, unitid)	
	buff_id = 128 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_128_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_128_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_KILL")

	sys.log("buff_128_delete "..","..battleid..","..buffinstid..","..data)

end