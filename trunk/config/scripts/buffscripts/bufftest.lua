-- buff测试用户脚本
-- buff格式 buff_id_update, buff_id_delete
-- 参数暂定为 battleid targetid data
sys.log("buff1")

function buff_1_add(battleid, unitid, buffinstid) 
	-- 当需要增加属性的时候调用这个接口,用来增加对应的已知属性
	-- 正常属性通过prpc来寻找,护盾直接调用专门的借口就可以
	
	Player.ChangeUnitProperty(battleid, unitid, data, "CPT_ATK") -- 增加属性
	-- Player.AddSheld(battleid, unit, data)						-- 增加护盾
	-- Player.ChangeSpecial(battleid, unit, buffinstid, "BF_JUMP")					-- 修改特殊属性 必然暴击啊昏厥啊之类的
	
end

function buff_1_update(battleid, buffinstid, unitid)	
	-- 加血减血之类的都在这里处理 其他的在go中处理
	buff_id = 1 --配置表中的buffid
	
	Battle.BuffMintsHp(battleid, unitid, buffinstid)
	
	sys.log("buff_1_update "..","..battleid..","..buffinstid..","..unitid)
	
end

function buff_1_delete(battleid, unitid, data)

-- 删除对应unit上的buff
-- 如果是有行为的buff就单纯删掉
-- 如果属于加属性的 这里需要删掉对应的属性
-- 加什么属性是已知的
-- 护盾类的debuff需要重新判定

	Player.ChangeUnitProperty(battleid, unitid, -data, "CPT_ATK") 	-- 修改属性
	-- Player.ChangeSheld(battleid, unit, -data)						 	-- 减去护盾

end