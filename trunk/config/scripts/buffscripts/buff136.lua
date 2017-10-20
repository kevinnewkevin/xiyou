-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加防御力
sys.log("buff136")

function buff_136_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_DEF")  
	--sys.log("buff_136_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_136_add  添加 增加防御力 被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_136_update(battleid, buffinstid, unitid)	
	buff_id = 136 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_136_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_136_update  更新增加防御力 被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_136_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_DEF")  
	 
	
	--sys.log("buff_136_delete "..","..battleid..","..buffinstid..","..data)
	sys.log("buff_136_delete  删除 增加防御力 被动buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end