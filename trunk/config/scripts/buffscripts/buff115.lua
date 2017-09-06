-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减物理属性
sys.log("buff115")

function buff_115_add(battleid, unitid, buffinstid, data) 
	sys.log("buff_115_add "..","..battleid..","..buffinstid..","..unitid)
	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK")  --减属性值  物理
end

function buff_115_update(battleid, buffinstid, unitid)	
	buff_id = 115 --配置表中的buffid
	sys.log("buff_115_update "..","..battleid..","..buffinstid..","..unitid)
end

function buff_115_delete(battleid, unitid, buffinstid, data)

	 sys.log("buff_115_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_ATK")
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减属性值 物理

	sys.log("buff_115_delete "..","..battleid..","..buffinstid..","..unitid)
	
end