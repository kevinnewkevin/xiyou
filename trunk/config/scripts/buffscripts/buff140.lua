-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--加 法术强度
sys.log("buff140")

function buff_140_add(battleid, unitid, buffinstid, data) 
	
	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_MAGIC_ATK")  --  增法术
	
	sys.log("buff_140_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_140_update(battleid, buffinstid, unitid)	
	buff_id = 140 --配置表中的buffid
	sys.log("buff_140_update "..","..battleid..","..buffinstid..","..unitid)
end

function buff_140_delete(battleid, unitid, buffinstid, data)

	 sys.log("buff_140_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_MAGIC_ATK")--  减法术

	sys.log("buff_140_delete "..","..battleid..","..buffinstid..","..unitid)
	
end