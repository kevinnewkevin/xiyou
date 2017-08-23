sys.log(" skill 12 start")

-- 技能释放 传入战斗ID和释放者的ID
-- 通过释放者和battleid取得对应的目标 单体或者多个
-- 循环/直接使用接口操控战斗 类似 战斗.攻击(战斗id, 释放者id, 承受者ID, 伤害数值, 是否暴击)
-- 
-- 
-- 所需接口
--	取得目标 （GetTarget（）  单   GetTargets（）  复）
--  取得对应属性 GetUnitProperty()
--  计算伤害数值 demage
--  计算是否暴击
--  攻击
-- 托塔李天王3号技能 敌方单体1回合无法行动，并对目标造成法术强度的伤害。

-- 物理强度视作buff Battle.buff

function SK_111_Action(battleid, casterid)
	local skillid = 111		-- 技能id
	local  t = Player.GetTarget(battleid,casterid)  --获取目标
	
	local  caster_attack = Player.GetUnitProperty(battleid,casterid,"CPT_ATK")  --获取攻击者属性
	
	
		--local  del_buff = Battle.AddBuff(1)  --一回合无法行动
		
		--local  damage_buff = Battle.AddBuff(1)  --敌对方法术强度
		
		local defender_def = Player.GetUnitProperty(battleid, t, "CPT_DEF")  -- 防御
		sys.log(1)
	
		--local  demage  = damage_buff-defender_def        --伤害 公式(法术强度的伤害  -   防御）
		
		local  damage  = 12 --测试
		sys.log(2)
	
		--判断伤害
		if damage <= 0 then 
		
			damage = 1
		
		end
		sys.log(3)
		
		local crit = Battle.GetCrit(skillid)   --是否暴击
		sys.log(4)
		Battle.Attack(battleid,casterid,t,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		sys.log("skill12 对id为"..t.."的目标造成"..damage.."点伤害")

	
	return  true
	 
end

sys.log( "skill 12 end")