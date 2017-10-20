-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
----增加法术属性
sys.log("buff147")

function buff_147_add(battleid, unitid, buffinstid, data) 
	--sys.log("buff_147_add "..","..battleid..","..buffinstid..","..unitid)
	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_MAGIC_ATK")  --加属性值  法术
	sys.log("buff_147_add  添加 增加法术属性 buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_147_update(battleid, buffinstid, unitid)	
	buff_id = 105 --配置表中的buffid
	--sys.log("buff_147_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_147_update  更新增加法术属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_147_delete(battleid, unitid, buffinstid, data)

	 --sys.log("buff_147_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_MAGIC_ATK")-- 减属性值 法术
	-- Player.ChangeSheld(battleid, unit, -data)						 	

	---sys.log("buff_147_delete "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_147_delete  删除 增加法术属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
end