-- buff测试用脚本 加特殊效果 如百分比减伤 眩晕等特殊效果
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--增加  所有属性
sys.log("buff126")

function buff_126_add(battleid, unitid, buffinstid,data)
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

	sys.log("buff_126_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_126_update(battleid, buffinstid, unitid)	
	buff_id = 126 --配置表中的buffid
	
	-- Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_126_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_126_delete(battleid, unitid, buffinstid,data)

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
	Player.ChangeUnitProperty(battleid, unitid,-data,"CPT_RESISTANCE")  
	 
	Player.ChangeIptProperty(battleid, unitid,-data,"IPT_HP")  
	 
	
	sys.log("buff_126_delete "..","..battleid..","..buffinstid..","..data)

end