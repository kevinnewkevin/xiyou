-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--加 法术强度
sys.log("buff159")

function buff_159_add(battleid, unitid, buffinstid, data) 
	
	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_MAGIC_ATK")  --  增法术
	
	--sys.log("buff_140_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_159_add  添加 加 法术强度 被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_159_update(battleid, buffinstid, unitid)	
	buff_id = 159 --配置表中的buffid
	--sys.log("buff_140_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_159_update  更新加 法术强度被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_159_delete(battleid, unitid, buffinstid, data)

	 sys.log("buff_159_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_MAGIC_ATK")--  减法术

	--sys.log("buff_140_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_159_delete  删除 加 法术强度被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
end