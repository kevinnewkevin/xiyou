sys.log(" 主角 主动技能  5 开始")

--主角5技能  自身加护盾

function SK_104_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 104	-- 技能id

	local level = 1
	local  t = Player.GetFriend(battleid,casterid)  --获取目标
	
	local hp = Player.GetUnitProperty(battleid, casterid, "CPT_HP")	-- 获取到攻击者的属性
	
	local damage = hp * 0.1
	
	Battle.cure(battleid, casterid, 0, 0)
	
	local buffid = 103
	
	Battle.AddBuff(battleid,casterid,battleid,buffid,damage)  --给zishen提供一个盾牌

	Battle.TargetOver(battleid)

end

sys.log(" 主角 主动技能  5 结束")