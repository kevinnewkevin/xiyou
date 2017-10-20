-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减物理属性
sys.log("buff143")

function buff_143_add(battleid, unitid, buffinstid, data) 
	--sys.log("buff_143_add "..","..battleid..","..buffinstid..","..unitid)
	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK")  --减属性值  物理
	sys.log("buff_143_add  添加 减物理属性 buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_143_update(battleid, buffinstid, unitid)	
	buff_id = 143 --配置表中的buffid
	--sys.log("buff_143_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_143_update  更新 减物理属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_143_delete(battleid, unitid, buffinstid, data)

	 sys.log("buff_143_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_ATK")
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减属性值 物理

	--sys.log("buff_143_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_143_delete  删除 减物理属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
end