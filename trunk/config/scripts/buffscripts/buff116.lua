-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减法术属性
sys.log("buff116")

function buff_116_add(battleid, unitid, buffinstid, data) 
	--sys.log("buff_116_add "..","..battleid..","..buffinstid..","..unitid)
	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_MAGIC_ATK")  --减法术属性值  
	sys.log("buff_116_add  添加 减法术属性值buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_116_update(battleid, buffinstid, unitid)	
	buff_id = 116 --配置表中的buffid
	--sys.log("buff_116_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_116_update  更新减法术属性值buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_116_delete(battleid, unitid, buffinstid, data)

	 --sys.log("buff_116_delete"..battleid..unitid..data)

	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_MAGIC_ATK")-- 减法术属性值 
	sys.log("buff_116_delete  删除减法术属性值buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
end