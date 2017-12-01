-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff182")

function buff_182_add(battleid, unitid, buffinstid, data) 
	--sys.log("buff_146_add "..","..battleid..","..buffinstid..","..unitid)
	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_ATK")  --加 物理属性值 
	sys.log("buff_182_add  添加 加物理属性 buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_182_update(battleid, buffinstid, unitid)	
	buff_id = 182 --配置表中的buffid
	--sys.log("buff_146_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_182_update  更新 加物理属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_182_delete(battleid, unitid, buffinstid, data)

	 sys.log("buff_182_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK")
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减属性值 物理

	--sys.log("buff_146_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_182_delete  删除 加物理属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
end