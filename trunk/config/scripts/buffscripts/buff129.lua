-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加 暴击伤害
sys.log("buff129")

function buff_129_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_CRIT")  
	
	sys.log("buff_129_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_129_update(battleid, buffinstid, unitid)	
	buff_id = 129 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_129_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_129_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_CRIT")  
	
	
	sys.log("buff_129_delete "..","..battleid..","..buffinstid..","..data)

end