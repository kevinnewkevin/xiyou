-- buff测试用脚本 减护盾
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减护盾值
sys.log("buff1")

function buff_120_add(battleid, unitid, buffinstid,data) 
	Player.AddSheld(battleid, unit, buffinstid,-data)  --减护盾
	sys.log("buff_120_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_120_update(battleid, buffinstid, unitid)	
	buff_id = 120 --配置表中的buffid
	
	--Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_120_update "..","..battleid..","..buffinstid..","..unitid)

	
end

function buff_120_delete(battleid, unitid, buffinstid,data)

	Player.ChangeSheld(battleid, unit, buffinstid,data)						 	-- 减去护盾
	sys.log("buff_120_delete "..","..battleid..","..buffinstid..","..unitid)
	
	
end