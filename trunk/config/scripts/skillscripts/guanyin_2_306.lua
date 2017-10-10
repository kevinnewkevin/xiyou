sys.log("观音 SK_306_Action 开始")

-- 技能释放 传入战斗ID和释放者的ID
-- 通过释放者和battleid取得对应的目标 单体或者多个
-- 循环/直接使用接口操控战斗 类似 战斗.攻击(战斗id, 释放者id, 承受者ID, 伤害数值, 是否暴击)
-- 
-- 
-- 所需接口
--  取得对应属性
--  计算伤害数值
--  计算是否暴击
--  攻击
-- 大慈大悲。驱散我方全体所有负面效果，并且减少伤害20%，持续1回合。
-- 增加速度视作buff

function SK_306_Action(battleid, casterid)

	local skillid = 306		-- 技能id

	local  attackNum = 0   --攻击个数

	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标
	
	local mtk = Player.GetUnitMtk(battleid,casterid)
	
	local mag_atk = mtk * 0.1
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		Player.PopAllBuffByDebuff(battleid,v)
		--local  attack_damage = Player.GetUnitDamage(battleid,casterid,v)  --获取伤害
		Battle.Cure(battleid,v,0,0)
		
		Battle.AddBuff(battleid,casterid,v,124, 20)  --公式(减少20%的伤害）
		sys.log(" 观音 大慈大悲 one")
		Battle.AddBuff(battleid,casterid,v,140,mag_atk) --每当有人获得治疗，增加观音10%法术强度，可叠加。
		sys.log("观音 大慈大悲 two")
		Battle.TargetOver(battleid)
	
		
	end
	
	
	return  true
end
sys.log("观音 SK_306_Action 结束")
