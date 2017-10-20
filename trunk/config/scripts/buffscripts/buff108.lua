-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff108")

function buff_108_add(battleid, unitid, buffinstid,data) 
	 Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_AOE")  --加必定溅射
	 Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_COMBO")  --加必定连击

	
	--sys.log("buff_108_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_108_add  添加必定溅射|必定连击buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_108_update(battleid, buffinstid, unitid)	
	buff_id = 108 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_108_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_108_update  更新必定溅射|必定连击buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_108_delete(battleid, unitid, buffinstid,data)

	Player.PopSpec(battleid, unitid, buffinstid,"BF_AOE")   --减必定溅射
	Player.PopSpec(battleid, unitid, buffinstid,"BF_COMBO")   --减必定连击
	
	--sys.log("buff_108_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_108_delete  删除必定溅射|必定连击buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end