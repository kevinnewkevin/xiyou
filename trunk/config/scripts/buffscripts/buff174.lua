-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--加荆棘
sys.log("buff174")

function buff_174_add(battleid, unitid, buffinstid,data) 
	 --Player.ChangeSpecial(battleid, unitid, buffinstid,data,"CPT_AGILE") 
	 
	 Player.ChangeUnitProperty(battleid, unitid, data, "CPT_REFLEX") --加 荆棘
	
	--sys.log("buff_117_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_174_add  添加 加荆棘buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_173_update(battleid, buffinstid, unitid)	
	buff_id = 117 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_117_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_174_update  更新加荆棘buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_174_delete(battleid, unitid, buffinstid,data)

	--Player.PopSpec(battleid, unitid, buffinstid,-data,"CPT_AGILE")   
	
	 Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_REFLEX")--减 荆棘
	
	--sys.log("buff_117_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_174_delete  删除加荆棘buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end