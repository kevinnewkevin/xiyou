-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加溅射
sys.log("buff131")

function buff_131_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_SPUTTERING")  
	  
	sys.log("buff_131_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_131_update(battleid, buffinstid, unitid)	
	buff_id = 131 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_131_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_131_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_SPUTTERING")  
	
	sys.log("buff_131_delete "..","..battleid..","..buffinstid..","..data)

end