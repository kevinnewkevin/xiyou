-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减 物理防御
sys.log("buff112")

function buff_112_add(battleid, unitid, buffinstid, data) 
	sys.log("buff_112_add "..","..battleid..","..buffinstid..","..unitid)
	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_DEF")  --  减物理防御
end

function buff_112_update(battleid, buffinstid, unitid)	
	buff_id = 112 --配置表中的buffid
	sys.log("buff_112_update "..","..battleid..","..buffinstid..","..unitid)
end

function buff_112_delete(battleid, unitid, buffinstid, data)

	 sys.log("buff_112_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_DEF")
	-- Player.ChangeSheld(battleid, unit, -data)						 	--  减物理防御

	sys.log("buff_112_delete "..","..battleid..","..buffinstid..","..unitid)
	
end