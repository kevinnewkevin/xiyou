-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减伤盾
sys.log("buff124")

function buff_124_add(battleid, unitid, buffinstid,data) 
	sys.log("buff_124_add")
	 --Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_WEAK")  --加增伤
	Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_SHELD")  --减伤盾

	
	--sys.log("buff_124_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_124_add  添加 减伤盾buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_124_update(battleid, buffinstid, unitid)	
	buff_id = 124 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_124_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_124_update  更新减伤盾buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_124_delete(battleid, unitid, buffinstid,data)

	--Player.PopSpec(battleid, unitid, buffinstid,"BF_WEAK")   --减增伤
	Player.PopSpec(battleid, unitid, buffinstid,"BF_SHELD")   --减伤害
	
	--sys.log("buff_124_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_124_delete  删除 减伤盾buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end

