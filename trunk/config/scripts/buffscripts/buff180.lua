-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加双防
sys.log("buff180")

function buff_180_add(battleid, unitid, buffinstid,data) 
	-- Player.ChangeUnitProperty(battleid, unitid,data,"CPT_DEF")  --加物理防御
	 Player.ChangeUnitProperty(battleid, unitid,data,"CPT_MAGIC_DEF")  --加法术防御

	
	--sys.log("buff_114_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_180_add  添加 加法术防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_180_update(battleid, buffinstid, unitid)	
	buff_id = 180 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_114_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_180_update  更新加法术防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_180_delete(battleid, unitid, buffinstid,data)

	--Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_DEF")   --减物理防御
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_MAGIC_DEF")   --减法术防御
	
	--sys.log("buff_114_delete "..","..battleid..","..buffinstid..","..data)
	sys.log("buff_180_delete  删除加法术防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end