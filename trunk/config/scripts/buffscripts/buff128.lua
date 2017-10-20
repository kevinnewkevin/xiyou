-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加 暴击几率
sys.log("buff128")

function buff_128_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_KILL") 
	
	--sys.log("buff_128_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_128_add  添加 增加暴击几率属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_128_update(battleid, buffinstid, unitid)	
	buff_id = 128 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_128_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_128_update  更新增加暴击几率属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_128_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_KILL")

	--sys.log("buff_128_delete "..","..battleid..","..buffinstid..","..data)
	sys.log("buff_128_delete  删除 增加暴击几率属性 buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end