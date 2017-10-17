-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff146")

function buff_146_add(battleid, unitid, buffinstid, data) 
	sys.log("buff_146_add "..","..battleid..","..buffinstid..","..unitid)
	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_ATK")  --加属性值  物理
end

function buff_146_update(battleid, buffinstid, unitid)	
	buff_id = 146 --配置表中的buffid
	sys.log("buff_146_update "..","..battleid..","..buffinstid..","..unitid)
end

function buff_146_delete(battleid, unitid, buffinstid, data)

	 sys.log("buff_146_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK")
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减属性值 物理

	sys.log("buff_146_delete "..","..battleid..","..buffinstid..","..unitid)
	
end