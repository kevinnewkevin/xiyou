-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff148")

function buff_148_add(battleid, unitid, buffinstid,data) 
	Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_JUMP")  --加冰冻
	
	--sys.log("buff_148_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_148_add  添加特殊效果 冰冻"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end


function buff_148_update(battleid, buffinstid, unitid)	
	buff_id = 148 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_104_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_148_update  更新特殊效果 冰冻"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_148_delete(battleid, unitid, buffinstid,data)

	Player.PopSpec(battleid, unitid, buffinstid,"BF_JUMP")   --减眩晕
	
	--sys.log("buff_104_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_148_delete 删除特殊效果 冰冻"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end