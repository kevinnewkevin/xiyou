-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff107")

function buff_107_add(battleid, unitid, buffinstid,data) 
	 Player.ChangeSpecial(battleid, unitid, buffinstid,"BF_THORNS")  --加荆棘

	
	--sys.log("buff_107_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_107_add  添加荆棘buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_107_update(battleid, buffinstid, unitid)	
	buff_id = 107 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_107_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_107_update  更新荆棘buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_107_delete(battleid, unitid, buffinstid,data)

	Player.PopSpec(battleid, unitid, buffinstid,"BF_THORNS")   --减荆棘
	
	--sys.log("buff_106_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_107_delete  删除荆棘buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end