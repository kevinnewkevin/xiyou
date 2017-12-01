sys.log("如来 SK_361_Action 开始")

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
--如来神掌：一只大手掌砸向对面，对敌方全体造成100%法术伤害
-- 增加速度视作buff

function SK_361_Action(battleid, casterid)
	local skillid = 361		-- 技能id
	local skillAttack = 10	-- 技能攻击
	local attackNum = 0		-- 攻击个数
	
	local t = Player.GetTargets(battleid, casterid, attackNum)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	for i,v in ipairs(t) do

		Battle.TargetOn(battleid)
		local truedamage = Player.GetMagicDamage(battleid, casterid, v)
		sys.log("如来 如来神掌对目标   "..v.. " 造成 法术伤害  "..truedamage )
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("如来 如来神掌对目标   "..v.. " 造成 最终法术伤害  "..damage )
		if damage <= 0 then
			damage = 0
		end
		
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		Battle.Attack(battleid, casterid, v, damage, crit)

		Battle.TargetOver(battleid)
		
	end
	
	return 1
end
sys.log("如来 SK_361_Action 结束")