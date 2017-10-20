-- buff测试用脚本 掉血
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff150")

function buff_150_add(battleid, unitid, buffinstid) 
	sys.log("buff_150_add 中毒掉血buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_150_update(battleid, buffinstid, unitid)	
	buff_id = 150 --配置表中的buffid
	
	Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_150_update 中毒掉血buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_150_delete(battleid, unitid, buffinstid, data)

	-- sys.log("buff_1_delete"..battleid..unitid..data)

	-- Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK") 	-- 修改属性
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减去护盾
	
	sys.log("buff_150_delete 中毒掉血buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end