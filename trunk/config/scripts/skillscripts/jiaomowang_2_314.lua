sys.log("蛟魔王 SK_314_Action 开始")

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
-- 孽龙。自身增加一个buff，受到的伤害提升30%，造成的伤害提升50%，并且攻击敌方单个目标一次。Buff持续2回合
-- 增加速度视作buff

function SK_314_Action(battleid, casterid)

	Battle.TargetOn(battleid)
	local skillid = 314		-- 技能id
	local skillAttack = 10	-- 技能攻击
	--local attackNum = 0		-- 攻击个数
	--Battle.Cure(battleid,casterid,0,0)
	sys.log("蛟魔王 孽龙 给自己  加buff110")
	Battle.AddSkillBuff(battleid,casterid, casterid, 110,50)  -- 造成的伤害增加50%
	sys.log("蛟魔王 孽龙 给自己   加buff122")
	Battle.AddSkillBuff(battleid,casterid, casterid, 122,30)  -- 受到的伤害提升30%
	sys.log("蛟魔王 孽龙 给自己  加buff完成")
	local  t = Player.GetTarget(battleid,casterid)  --获取目标
	local atk = Player.GetUnitAtk(battleid,casterid)
	local atk_damage = atk * 0.1
	local damage = Player.GetUnitDamage(battleid,casterid,t)  -- 获取伤害
	sys.log("蛟魔王 孽龙 给目标  " ..t.. " 造成 物理伤害  "..damage)
	local caster_damage = ClacDamageByAllBuff(battleid,casterid,t,damage)
	sys.log("蛟魔王 孽龙 给目标  " ..t.. " 造成 法术最终伤害  "..caster_damage)
	--判断伤害
	if caster_damage <= 0 then 
	
		caster_damage = 0
	
	end
	
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	
	Battle.Attack(battleid,casterid,t,caster_damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
	Battle.AddBuff(battleid,casterid,t,110,atk_damage)
	Battle.TargetOver(battleid)
	
	return  true
	
end

sys.log("蛟魔王 SK_314_Action 结束")