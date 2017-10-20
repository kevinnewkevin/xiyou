-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减速
sys.log("buff145")

function buff_145_add(battleid, unitid, buffinstid,data) 
	 --Player.ChangeSpecial(battleid, unitid, buffinstid,data,"CPT_AGILE") 
	 
	 Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_AGILE") --加 减速
	
	--sys.log("buff_145_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_145_add  添加 减速 buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_145_update(battleid, buffinstid, unitid)	
	buff_id = 145 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_145_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_145_update  更新 减速buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_145_delete(battleid, unitid, buffinstid,data)

	--Player.PopSpec(battleid, unitid, buffinstid,-data,"CPT_AGILE")   
	
	 Player.ChangeUnitProperty(battleid, unitid, data, "CPT_AGILE")--减 减速
	
	--sys.log("buff_145_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_145_delete  删除 减速buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end