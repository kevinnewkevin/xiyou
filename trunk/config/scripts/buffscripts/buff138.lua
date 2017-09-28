-- buff测试用脚本 
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
--石猴。暴击后下一次受到的伤害减半。
sys.log("buff138")

function buff_138_add(battleid, unitid, buffinstid,data) 
	Player.ChangeSpecial(battleid, unit, buffinstid,"BF_SHELD")  --减伤盾
	sys.log("buff_138_add "..","..battleid..","..buffinstid..","..unitid)
end

function buff_138_update(battleid, buffinstid, unitid)	
	buff_id = 138 --配置表中的buffid
	
	--Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_138_update "..","..battleid..","..buffinstid..","..unitid)

	
end

function buff_138_delete(battleid, unitid, buffinstid,data)

	Player.PopSpec(battleid, unitid, buffinstid,"BF_SHELD")						 	-- 减去护盾
	sys.log("buff_138_delete "..","..battleid..","..buffinstid..","..unitid)
	
	
end