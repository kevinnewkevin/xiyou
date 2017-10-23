-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff154")

function buff_154_add(battleid, unitid, buffinstid, data) 
	sys.log("buff_154_add  添加加物理属性buff "..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_ATK")  --加属性值  物理
	sys.log("buff_154_add "..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_154_update(battleid, buffinstid, unitid)	
	buff_id = 154 --配置表中的buffid
	sys.log("buff_154_update  更新加物理属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
end

function buff_154_delete(battleid, unitid, buffinstid, data)

	 sys.log("buff_154_delete 加物理属性, battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是.."..data)

	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK")
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减属性值 物理

	sys.log("buff_154_delete  删除增加物理属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
	
end