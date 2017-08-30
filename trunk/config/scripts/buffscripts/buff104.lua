-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff1")

function buff_104_add(battleid, unitid, buffinstid) 
	Player.ChangeSpecial(battleid, unit, buffinstid)
end

function buff_104_update(battleid, buffinstid, unitid)	
	buff_id = 104 --配置表中的buffid
	
	Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_1_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_104_delete(battleid, unitid, buffinstid, data)


end