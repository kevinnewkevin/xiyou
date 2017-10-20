-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减 物理强度
sys.log("buff141")

function buff_141_add(battleid, unitid, buffinstid, data) 
	
	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK") 
	
	--sys.log("buff_141_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_141_add  添加 减 物理强度 被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_141_update(battleid, buffinstid, unitid)	
	buff_id = 141 --配置表中的buffid
	--sys.log("buff_140_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_141_update  更新减 物理强度被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_141_delete(battleid, unitid, buffinstid, data)

	 --sys.log("buff_141_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_ATK")  

	--sys.log("buff_141_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_141_delete  删除 减 物理强度被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
end