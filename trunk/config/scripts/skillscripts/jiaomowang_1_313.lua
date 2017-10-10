sys.log("蛟魔王 SK_313_Action 开始")

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
-- 真龙形态。对敌方一个目标造成物理130%强度的伤害，如果这个目标死亡，则额外触发另一次真龙形态。
-- 增加速度视作buff

function SK_313_Action(battleid, casterid)

	Battle.TargetOn(battleid)
	local skillid = 313		-- 技能id
	local skillAttack = 10	-- 技能攻击
	--local attackNum = 0		-- 攻击个数
	
	local  target = Player.GetTarget(battleid,casterid)
	
	local atk = Player.GetUnitAtk(battleid,casterid)
	local atk_damage = atk * 0.1
	local  damage = Player.GetUnitDamage(battleid,casterid,target)
	sys.log("蛟魔王 真龙形态 给目标  " ..target.. " 造成 法术伤害  "..damage)
	local trueDamage = ClacDamageByAllBuff(battleid,casterid,target,damage)
	sys.log("蛟魔王 真龙形态 给目标  " ..target .." 造成 法术最终伤害  "..trueDamage)
	
	if trueDamage <= 0 then 
		
		trueDamage = 1
		
	end
	local crit = Battle.GetCrit(skillid)
	
	local Damage = trueDamage * 1.3
	
	sys.log("蛟魔王 真龙形态 给目标  " ..target .." 造成 法术最终伤害的130%  "..Damage)
	Battle.Attack(battleid,casterid,target,Damage,crit) 	--调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
	
	Battle.AddBuff(battleid,casterid,target,110,atk_damage)
	
	Battle.TargetOver(battleid)
	local  p = Player.CheckUnitDead(battleid,target)
	
	while (p == 1) do
		Battle.TargetOn(battleid)
		t = Player.GetTarget(battleid,casterid)
		sys.log("将魔王的敌对目标"..t)
		if t == -1 then 
			break
		end
		Battle.Attack(battleid,casterid,t,Damage,crit)
		Battle.AddBuff(battleid,casterid,target,110,atk_damage)
		Battle.TargetOver(battleid)
		p = Player.CheckUnitDead(battleid,t)
	end
	
	
	return  true
	
end
sys.log("蛟魔王 SK_313_Action 结束")