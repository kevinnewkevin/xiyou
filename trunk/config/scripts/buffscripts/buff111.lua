-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff111")

function buff_111_add(battleid, unitid, buffinstid) 
	 Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_AOE")  --加必定溅射
	
	--sys.log("buff_111_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_111_add  添加必定溅射buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_111_update(battleid, buffinstid, unitid)	
	buff_id = 111 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_111_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_111_update  更新必定溅射buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_106_delete(battleid, unitid, buffinstid)

	Player.PopSpec(battleid, unitid, buffinstid,"BF_AOE")   --减必定溅射
	
	--sys.log("buff_111_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_111_delete  删除必定溅射buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end