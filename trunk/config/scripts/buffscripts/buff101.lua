-- buff测试用脚本 加血
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff101")

function buff_101_add(battleid, unitid, buffinstid) 
	sys.log("buff_101_add "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_101_update(battleid, buffinstid, unitid)	
	buff_id = 101 --配置表中的buffid
	
	Battle.BuffCureHp(battleid, unitid, buffinstid)
	
	sys.log("buff_101_update "..","..battleid..","..buffinstid..","..unitid)
	
	
end

function buff_101_delete(battleid, unitid, buffinstid, data)

	-- Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK") 	-- 修改属性
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减去护盾
	
	sys.log("buff_101_delete "..","..battleid..","..buffinstid..","..unitid)

end