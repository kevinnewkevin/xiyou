-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff178")

function buff_178_add(battleid, unitid, buffinstid,data) 
	Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_JUMP")  --加无法行动晕
	
	--sys.log("buff_104_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_178_add  添加特殊效果 无法行动"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end


function buff_178_update(battleid, buffinstid, unitid)	
	buff_id = 178 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_104_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_178_update  更新特殊效果 无法行动"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_167_delete(battleid, unitid, buffinstid,data)

	Player.PopSpec(battleid, unitid, buffinstid,"BF_JUMP")   --减眩晕
	
	--sys.log("buff_104_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_178_delete 删除特殊效果 无法行动"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end