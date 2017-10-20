-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--减 所有属性
sys.log("buff127")

function buff_127_add(battleid, unitid, buffinstid,data)
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_DEF")  --物理防御
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_MAGIC_DEF")  --魔法防御
	
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_HP")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_CHP")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_ATK")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_MAGIC_ATK")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_AGILE")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_KILL") 
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_CRIT")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_COUNTER_ATTACK")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_SPUTTERING")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_DOUBLE_HIT")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_RECOVERY")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_REFLEX")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_SUCK_BLOOD")  
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_INCANTER") 
 
	Player.ChangeIptProperty(battleid, unitid,-data,"IPT_HP")  
	
	--sys.log("buff_127_add "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_127_add  添加 减所有属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)
end

function buff_127_update(battleid, buffinstid, unitid)	
	buff_id = 127 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	--sys.log("buff_127_update "..","..battleid..","..buffinstid..","..unitid)
	sys.log("buff_127_update  更新减所有属性buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid)
	
end

function buff_127_delete(battleid, unitid, buffinstid,data)

	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_DEF")  --物理防御
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_MAGIC_DEF")  --魔法防御
	 
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_HP")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_CHP")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_ATK")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_MAGIC_ATK")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_AGILE")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_KILL") 
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_CRIT")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_COUNTER_ATTACK")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_SPUTTERING")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_DOUBLE_HIT")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_RECOVERY")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_REFLEX")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_SUCK_BLOOD")  
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_INCANTER") 
	Player.ChangeUnitProperty(battleid, unitid,data,"CPT_RESISTANCE")  
	
	Player.ChangeIptProperty(battleid, unitid,data,"IPT_HP")  
	 
	
	--sys.log("buff_126_delete "..","..battleid..","..buffinstid..","..data)
	sys.log("buff_127_delete  删除 减所有属性 buff"..",battleid是"..battleid..",buffid是"..buffinstid..",目标"..unitid..",数据是"..data)

end