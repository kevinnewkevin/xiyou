-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data

--免死
sys.log("buff106")

function buff_106_add(battleid, unitid, buffinstid,data) 
	 Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_UNDEAD")  --加眩晕免除一次造成死亡的伤害

	
	sys.log("buff_106_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_106_update(battleid, buffinstid, unitid)	
	buff_id = 106 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_106_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_106_delete(battleid, unitid, buffinstid,data)

	Player.PopSpec(battleid, unitid, buffinstid,"BF_UNDEAD")   --减眩晕免除一次造成死亡的伤害
	
	sys.log("buff_106_delete "..","..battleid..","..buffinstid..","..unitid)

end

