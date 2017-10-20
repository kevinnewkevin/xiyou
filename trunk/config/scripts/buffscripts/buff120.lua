-- buff测试用脚本 减护盾
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减护盾值
sys.log("buff120")

function buff_120_add(battleid, unitid, buffinstid,data) 
	Player.AddSheld(battleid, unit, buffinstid)  --减护盾
	--sys.log("buff_120_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_120_add  添加 减护盾buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_120_update(battleid, buffinstid, unitid)	
	buff_id = 120 --配置表中的buffid
	
	--Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_120_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_120_update  更新减护盾buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)

	
end

function buff_120_delete(battleid, unitid, buffinstid,data)

	Player.ChangeSheld(battleid, unit, buffinstid)						 	-- 减去护盾
	--sys.log("buff_120_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_120_delete  删除 减护盾buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
	
end