-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加速度
sys.log("buff109")

function buff_109_add(battleid, unitid, buffinstid,data) 
	 --Player.ChangeSpecial(battleid, unitid, buffinstid,data,"CPT_AGILE") 
	 
	 Player.ChangeUnitProperty(battleid, unitid, data, "CPT_AGILE") --加 速
	
	--sys.log("buff_109_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_109_add  添加加 速buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_109_update(battleid, buffinstid, unitid)	
	buff_id = 109 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_109_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_109_update  更新加 速buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_109_delete(battleid, unitid, buffinstid,data)

	--Player.PopSpec(battleid, unitid, buffinstid,-data,"CPT_AGILE")   
	
	 Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_AGILE")--减 速
	
	--sys.log("buff_109_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_109_delete  删除加 速buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end