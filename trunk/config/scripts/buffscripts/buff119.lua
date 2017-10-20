-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减法术防御
sys.log("buff119")

function buff_119_add(battleid, unitid, buffinstid,data)
	-- Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_DEF")  --物理防御
	 Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_MAGIC_DEF")  --法术防御

	
	--sys.log("buff_119_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_119_add  添加 减法术防御 buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_119_update(battleid, buffinstid, unitid)	
	buff_id = 119 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_119_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_119_update  更新减法术防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_119_delete(battleid, unitid, buffinstid,data)

	--Player.ChangeUnitProperty(battleid, unitid,data,"CPT_DEF")   --物理防御
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_MAGIC_DEF")   --法术防御
	
	--sys.log("buff_119_delete "..battleid)
	--sys.log("buff_119_delete "..buffinstid)
	--sys.log("buff_119_delete "..data)
	
	--sys.log("buff_119_delete "..","..battleid..","..buffinstid..","..data)
	sys.log("buff_119_delete  删除 减法术防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end