sys.log("skill 153 start")

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
--  普通单体攻击技能  法术
-- 增加速度视作buff

function SK_152_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	local skillid = 152		-- 技能id
	
	local t = Player.GetTarget(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	sys.log("GetTarget  ".. t)
	
	local truedamage = Player.GetMagicDamage(battleid, casterid, t)
	
	sys.log("SK_100_Action 的伤害".. truedamage)

	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)
	
	if damage <= 0 then 
		damage = 1
	end
	
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	sys.log("是否暴击"..crit)

	Battle.Attack(battleid, casterid, t, damage, crit)
	
	
	Battle.TargetOver(battleid)
	
	sys.log("skil53 对id为"..t.."的目标造成"..damage.."点伤害")	
	-- 只给游戏返回 对谁造成了多少伤害
	-- 并不参与计算
	return true
end

sys.log("skill 53 end")