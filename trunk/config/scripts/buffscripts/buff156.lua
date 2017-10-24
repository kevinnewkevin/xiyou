-- buff测试用脚本 加护盾
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff156")

function buff_156_add(battleid, unitid, buffinstid,data) 
	Player.AddSheld(battleid, unitid, buffinstid)  --加护盾
	--sys.log("buff_156_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_156_add  加护盾buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_156_update(battleid, buffinstid, unitid)	
	buff_id = 156 --配置表中的buffid
	
	--Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_156_update 2回合 加护盾"..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_156_update  更新护盾buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)

	
end

function buff_156_delete(battleid, unitid, buffinstid,data)

	-- Player.PopSheld(battleid, unitid, buffinstid)						 	-- 减去护盾
	--sys.log("buff_156_delete 2回合 加护盾"..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_156_delete 删除加护盾buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
	
end