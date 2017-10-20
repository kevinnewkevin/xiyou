-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data

--免伤
sys.log("buff151")

function buff_151_add(battleid, unitid, buffinstid,data) 
	 Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_UNDAMAGE")  --免除一切伤害

	
	--sys.log("buff_106_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_151_add  添加免伤buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_151_update(battleid, buffinstid, unitid)	
	buff_id = 151 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_151_update  更新免伤buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_151_delete(battleid, unitid, buffinstid,data)

	Player.PopSpec(battleid, unitid, buffinstid,"BF_UNDEAD")   --减眩晕免除一次造成死亡的伤害
	
	sys.log("buff_151_delete  删除免伤buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end

