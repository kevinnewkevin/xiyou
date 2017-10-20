-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff113")

function buff_113_add(battleid, unitid, buffinstid,data) 
	--sys.log("buff_113_add "..","..battleid..","..buffinstid..","..unitid)
	 Player.ChangeSpecial(battleid, unitid,buffinstid, "BF_UNCURE")  --加 无法恢复生命
	 sys.log("buff_113_add  添加无法恢复生命buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_113_update(battleid, buffinstid, unitid)	
	buff_id = 113 --配置表中的buffid
	--sys.log("buff_113_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_113_update  更新无法恢复生命buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_113_delete(battleid, unitid, buffinstid,data)

	 --sys.log("buff_113_delete"..battleid..unitid)

	 Player.ChangeSpecial(battleid, unitid,buffinstid, "BF_UNCURE")-- 减 无法恢复生命
	-- Player.ChangeSheld(battleid, unit, -data)						 	

	--sys.log("buff_113_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_113_delete  删除无法恢复生命buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
end