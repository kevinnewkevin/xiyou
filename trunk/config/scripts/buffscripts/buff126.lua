-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减法术防御
sys.log("buff126")

function buff_126_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_DEF")  --物理防御
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_MAGIC_DEF")  --法术防御

	
	sys.log("buff_126_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_126_update(battleid, buffinstid, unitid)	
	buff_id = 126 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_126_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_126_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_DEF")   --物理防御
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_MAGIC_DEF")   --法术防御
	
	sys.log("buff_126_delete "..battleid)
	sys.log("buff_126_delete "..buffinstid)
	sys.log("buff_126_delete "..data)
	
	sys.log("buff_126_delete "..","..battleid..","..buffinstid..","..data)

end