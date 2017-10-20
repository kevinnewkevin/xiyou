-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--加输出伤害
sys.log("buff122")

function buff_122_add(battleid, unitid, buffinstid,data) 
	--Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_WEAK")  --加增伤
	Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_STRONG")  --加输出伤害

	
	--sys.log("buff_122_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_122_add  添加 加输出伤害buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_122_update(battleid, buffinstid, unitid)	
	buff_id = 122 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_122_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_122_update  更新加输出伤害buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_122_delete(battleid, unitid, buffinstid,data)

	--Player.PopSpec(battleid, unitid, buffinstid,"BF_WEAK")   --减增伤
	Player.PopSpec(battleid, unitid, buffinstid,"BF_STRONG")   --减输出伤害
	
	--sys.log("buff_122_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_122_delete  删除 加输出伤害buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end