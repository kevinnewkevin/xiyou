-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff1")

function buff_113_add(battleid, unitid, buffinstid, data) 
	sys.log("buff_113_add "..","..battleid..","..buffinstid..","..unitid)
	Player.ChangeUnitProperty(battleid, unitid, data, "BF_WFHFSM")  --加 无法恢复生命
end

function buff_113_update(battleid, buffinstid, unitid)	
	buff_id = 113 --配置表中的buffid
	sys.log("buff_113_update "..","..battleid..","..buffinstid..","..unitid)
end

function buff_113_delete(battleid, unitid, buffinstid, data)

	 sys.log("buff_113_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, -data, "BF_WFHFSM")
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减 无法恢复生命

	sys.log("buff_113_delete "..","..battleid..","..buffinstid..","..unitid)
	
end