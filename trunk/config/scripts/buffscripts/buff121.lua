-- buff测试用脚本 修改属性 -- 最后一个参数是字符串格式的属性介绍
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data

--受伤后回血
sys.log("buff121")

function buff_121_add(battleid, unitid, buffinstid,data) 
	sys.log("buff_121_add "..","..battleid..","..buffinstid..","..unitid)
	 Player.ChangeSpecial(battleid, unitid,buffinstid, "BF_CURE")  --受伤后 回血
end

function buff_121_update(battleid, buffinstid, unitid)	
	buff_id = 121 --配置表中的buffid
	
	Battle.BuffCureHp(battleid, unitid, buffinstid)
	
	sys.log("buff_121_update "..","..battleid..","..buffinstid..","..unitid)
end

function buff_121_delete(battleid, unitid, buffinstid,data)

	 sys.log("buff_121_delete"..battleid..unitid)

	 Player.ChangeSpecial(battleid, unitid,buffinstid, "BF_WFHFSM")-- -受伤后 回血
	-- Player.ChangeSheld(battleid, unit, -data)						 	

	sys.log("buff_121_delete "..","..battleid..","..buffinstid..","..unitid)
	
end