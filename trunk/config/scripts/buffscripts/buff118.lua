-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data

--加物理防御
sys.log("buff118")

function buff_118_add(battleid, unitid, buffinstid, data) 
	--sys.log("buff_118_add "..","..battleid..","..buffinstid..","..unitid)
	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_DEF")  --加属性值 物理防御
	sys.log("buff_118_add  添加 加物理防御 buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_118_update(battleid, buffinstid, unitid)	
	buff_id = 118 --配置表中的buffid
	--sys.log("buff_118_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_118_update  更新加物理防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_118_delete(battleid, unitid, buffinstid, data)

	 --sys.log("buff_118_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_DEF")
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减属性值 物理防御

	--sys.log("buff_118_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_118_delete  删除 加物理防御buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end