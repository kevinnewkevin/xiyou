-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减速
sys.log("buff117")

function buff_117_add(battleid, unitid, buffinstid,data) 
	 --Player.ChangeSpecial(battleid, unitid, buffinstid,data,"CPT_AGILE") 
	 
	 Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_AGILE") --加 减速
	
	sys.log("buff_117_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_117_update(battleid, buffinstid, unitid)	
	buff_id = 117 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_117_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_117_delete(battleid, unitid, buffinstid,data)

	--Player.PopSpec(battleid, unitid, buffinstid,-data,"CPT_AGILE")   
	
	 Player.ChangeUnitProperty(battleid, unitid, data, "CPT_AGILE")--减 减速
	
	sys.log("buff_117_delete "..","..battleid..","..buffinstid..","..unitid)

end