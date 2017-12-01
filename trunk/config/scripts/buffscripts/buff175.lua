-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff175")

function buff_175_add(battleid, unitid, buffinstid,data) 
	Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_SILENT")  --加沉默
	
	--sys.log("buff_104_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_175_add  添加特殊效果 沉默"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end


function buff_175_update(battleid, buffinstid, unitid)	
	buff_id = 175 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_104_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_175_update  更新特殊效果 沉默"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_175_delete(battleid, unitid, buffinstid,data)

	Player.PopSpec(battleid, unitid, buffinstid,"BF_SILENT")   --减沉默
	
	--sys.log("buff_104_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_175_delete 删除特殊效果 沉默"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end