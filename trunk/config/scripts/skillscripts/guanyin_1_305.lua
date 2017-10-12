sys.log(" 观音 SK_305_Action 开始")

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
-- 放下屠刀。对一个敌方目标造成法术强度的伤害，并降低目标40%的物理强度和法术强度，将其转移给离自己最近的一个友方目标。持续1回合
-- 增加速度视作buff

function SK_305_Action(battleid, casterid)

	Battle.TargetOn(battleid)

	local skillid = 305		-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 
	
	local  caster_attack = Player.GetUnitAtk(battleid,casterid)  --获取攻击者属性  物理
	
	local  caster_magic = Player.GetUnitMtk(battleid,casterid)  --获取攻击者属性   法术
	
	--local defender_def = Player.GetCalcMagicDef(battleid, t)  --获取防御属性
	
	
	local  truedamage  = Player.GetMagicDamage(battleid,casterid,t)    --伤害 公式（）
	
	sys.log("观音对目标造成的法术伤害   ".. truedamage)
	
	
	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)
	sys.log("观音对目标造成的最终法术伤害   ".. damage)
	
	--判断伤害
	if damage <= 0 then 
	
		damage = 0
	
	end
	
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	local pvalue = 0.4
	
	local mag_pro = caster_magic*pvalue
	
	local atk_pro = caster_attack*pvalue
	
	Battle.Attack(battleid,casterid,t,damage,crit)   --调用服务器   （伤害）(战斗者，释放者，承受者，伤害，暴击）
	sys.log("观音对目标  降低目标"..t.."    40%法术强度 ".. mag_pro)
	Battle.AddBuff(battleid,casterid, t,116, mag_pro)     --降低目标40%法术强度
	sys.log("观音对目标  降低目标"..t.."   40%物理强度 ".. atk_pro)
	Battle.AddBuff(battleid,casterid, t,115, atk_pro)     --降低目标40%物理强度
	sys.log("观音对目标  降低目标"..t.."   40%法术强度  物理强度完成")
	Battle.TargetOver(battleid)

	Battle.TargetOn(battleid)

	local  p = Player.GetFriend(battleid,casterid)  --获取目标 
	Battle.Cure(battleid,p,0,0)
	sys.log("观音对目标  增加己方"..p.."   40%法术强度 ".. mag_pro)
	Battle.AddBuff(battleid,casterid, p,105, mag_pro)
	sys.log("观音对目标  增加己方"..p.."   40%物理强度 ".. atk_pro)
	Battle.AddBuff(battleid,casterid, p,102, atk_pro)
	sys.log("观音对目标  增加己方"..p.."   40%法术 法术强度 物理强度  强度完成")
	--Battle.AddBuff(battleid,casterid, t, 115,caster_attack*0.4)     --降低目标40%物理强度
	
	Battle.TargetOver(battleid)
	
	return  true
	 
end
sys.log(" 观音 SK_305_Action 结束")