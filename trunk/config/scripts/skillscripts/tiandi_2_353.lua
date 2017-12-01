sys.log(" 天帝 SK_353_Action 开始")

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
-- 天帝 2号技能。真龙天子：集合所有己方目标的法术攻击力（作为法术攻击力计算），对目标单体造成100%法术伤害
-- 增加速度视作buff

function SK_353_Action(battleid,casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 353
	
	local attackNum = 0 
	
	local p = Player.GetTarget(battleid,casterid)
	
	local t = Player.GetFriends(battleid,casterid,attackNum)
	
	local mtk_damage = 0 
	
	for i,v in  ipairs(t) do
	
		local truedamage = Player.GetMagicDamage(battleid,v,p)

		sys.log("天帝 真龙天子及己方对目标 "..p.."造成的法术伤害 "..truedamage)
		mtk_damage = truedamage + mtk_damage
		
	end
	sys.log("天帝 真龙天子及己方对目标 "..p.."造成的法术伤害总值 "..mtk_damage)
	
	local damage = ClacDamageByAllBuff(battleid,casterid,p,mtk_damage)
	
	sys.log("天帝 真龙天子对目标 "..p.."造成的法术伤害总值 "..damage)

	if damage <= 0 then 

		damage = 0 

	end

	local crit = Battle.GetCrit(skillid)

	Battle.Attack(battleid,casterid,p,damage,crit)

	Battle.TargetOver(battleid)
	return true
end

sys.log(" 天帝 SK_353_Action 结束")