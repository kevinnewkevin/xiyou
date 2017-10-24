-- buff测试用脚本 掉血
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff162")

function buff_162_add(battleid, unitid, buffinstid) 
	sys.log("buff_162_add 掉血buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_162_update(battleid, buffinstid, unitid)	
	buff_id = 162 --配置表中的buffid
	
	Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_162_update掉血buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_162_delete(battleid, unitid, buffinstid, data)

	-- sys.log("buff_1_delete"..battleid..unitid..data)

	-- Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK") 	-- 修改属性
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减去护盾
	
	sys.log("buff_162_delete 掉血buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end