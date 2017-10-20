-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data

--锁定己方随机目标
sys.log("buff152")

function buff_152_add(battleid, unitid, buffinstid,data) 
	 Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_FRIENDLOCK")  --

	
	--sys.log("buff_106_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_152_add  添加免伤buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_152_update(battleid, buffinstid, unitid)	
	buff_id = 152 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_152_update  更新免伤buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_152_delete(battleid, unitid, buffinstid,data)

	Player.PopSpec(battleid, unitid, buffinstid,"BF_FRIENDLOCK")   --
	
	sys.log("buff_152_delete  删除免伤buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end

