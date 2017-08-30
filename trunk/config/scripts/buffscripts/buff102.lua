-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff1")

function buff_101_add(battleid, unitid, buffinstid) 
	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_ATK")
end

function buff_101_update(battleid, buffinstid, unitid)	
	buff_id = 101 --配置表中的buffid
	
end

function buff_101_delete(battleid, unitid, data)

	-- sys.log("buff_1_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK")
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减去护盾

end