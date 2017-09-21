-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加法术防御力
sys.log("buff137")

function buff_137_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_MAGIC_DEF")  
	sys.log("buff_137_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_137_update(battleid, buffinstid, unitid)	
	buff_id = 137 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_137_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_137_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_MAGIC_DEF")  
	 
	
	sys.log("buff_137_delete "..","..battleid..","..buffinstid..","..data)

end