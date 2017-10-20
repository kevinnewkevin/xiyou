-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff142")

function buff_142_add(battleid, unitid, buffinstid,data) 
	Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_LOCK")  --所目标
	
	--sys.log("buff_142_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_142_add  添加 锁定目标 buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_142_update(battleid, buffinstid, unitid)	
	buff_id = 142 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_142_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_142_update  更新 锁定目标buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_142_delete(battleid, unitid, buffinstid,data)

	Player.PopSpec(battleid, unitid, buffinstid,"BF_LOCK")   --所目标
	
	--sys.log("buff_142_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_142_delete  删除 锁定目标buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end