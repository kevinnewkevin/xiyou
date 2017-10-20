-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加 连击
sys.log("buff130")

function buff_130_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_DOUBLE_HIT")  
	 
	--sys.log("buff_130_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_130_add  添加 增加连击属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_130_update(battleid, buffinstid, unitid)	
	buff_id = 130 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_130_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_130_update  更新增加连击属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_130_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_DOUBLE_HIT")  
	
	
	--sys.log("buff_130_delete "..","..battleid..","..buffinstid..","..data)
	sys.log("buff_130_delete  删除 增加连击属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end