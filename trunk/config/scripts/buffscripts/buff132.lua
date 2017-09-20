-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加回复
sys.log("buff132")

function buff_132_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_RECOVERY") 
	
	sys.log("buff_132_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_132_update(battleid, buffinstid, unitid)	
	buff_id = 132 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_132_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_132_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_RECOVERY")  
	
	
	sys.log("buff_132_delete "..","..battleid..","..buffinstid..","..data)

end