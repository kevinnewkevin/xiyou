-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff168")

function buff_168_add(battleid, unitid, buffinstid,data) 
	Player.ChangeSpecial(battleid, unitid, buffinstid,-data,"CPT_DEF")  --加 减物理防御
	--Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_STRONG")  --

	
	--sys.log("buff_110_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_168_add  添加减物理防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_168_update(battleid, buffinstid, unitid)	
	buff_id = 168 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_110_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_168_update  更新减物理防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_168_delete(battleid, unitid, buffinstid,data)

	Player.PopSpec(battleid, unitid, buffinstid,data,"CPT_DEF")   --减增伤
	--Player.PopSpec(battleid, unitid, buffinstid,"BF_STRONG")   --减输出伤害
	
	--sys.log("buff_110_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_168_delete  删除减物理防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end

