-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减 物理防御
sys.log("buff170")

function buff_170_add(battleid, unitid, buffinstid, data) 
	--sys.log("buff_112_add "..","..battleid..","..buffinstid..","..unitid)
	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_DEF")  --  减物理防御
	sys.log("buff_170_add  添加减物理防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_170_update(battleid, buffinstid, unitid)	
	buff_id = 170 --配置表中的buffid
	--sys.log("buff_112_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_170_update  更新减物理防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_170_delete(battleid, unitid, buffinstid, data)

	 --sys.log("buff_112_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_DEF")
	-- Player.ChangeSheld(battleid, unit, -data)						 	--  减物理防御

	--sys.log("buff_112_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_170_delete  删除减物理防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
end