-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加回复
sys.log("buff132")

function buff_132_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_RECOVERY") 
	
	--sys.log("buff_132_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_132_add  添加 增加回复生命属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_132_update(battleid, buffinstid, unitid)	
	buff_id = 132 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	---sys.log("buff_132_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_132_update  更新增加回复生命属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_132_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_RECOVERY")  
	
	
	--sys.log("buff_132_delete "..","..battleid..","..buffinstid..","..data)
	sys.log("buff_132_delete  删除 增加回复生命属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end