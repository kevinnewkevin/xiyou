-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--加 物理强度
sys.log("buff139")

function buff_139_add(battleid, unitid, buffinstid, data) 
	
	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_ATK")  --  减物理
	
	--sys.log("buff_139_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_139_add  添加 加 物理强度 被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_139_update(battleid, buffinstid, unitid)	
	buff_id = 139 --配置表中的buffid
	--sys.log("buff_139_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_139_update  更新加 物理强度 被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_139_delete(battleid, unitid, buffinstid, data)

	 --sys.log("buff_139_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK")--  减物理

	--sys.log("buff_139_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_139_delete  删除 加 物理强度 被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
end