-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加荆棘反射
sys.log("buff133")

function buff_133_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_REFLEX") 
	 
	sys.log("buff_133_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_133_update(battleid, buffinstid, unitid)	
	buff_id = 133 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_133_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_133_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_REFLEX") 
	
	
	sys.log("buff_133_delete "..","..battleid..","..buffinstid..","..data)

end