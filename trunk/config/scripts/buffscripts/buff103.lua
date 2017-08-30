-- buff测试用脚本 加护盾
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff1")

function buff_103_add(battleid, unitid, buffinstid) 
	Player.AddSheld(battleid, unit, buffinstid)
end

function buff_103_update(battleid, buffinstid, unitid)	
	buff_id = 103 --配置表中的buffid
	
	Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_1_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_103_delete(battleid, unitid, data)

	-- sys.log("buff_1_delete"..battleid..unitid..data)

	-- Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK") 	-- 修改属性
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减去护盾

end