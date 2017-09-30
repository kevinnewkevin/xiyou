sys.log("SK_304_Action")

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
-- 造物。使己方血量最少的单位免疫一回合会导致其死亡的伤害。（不会吸收伤害，只免疫一次死亡，攻击导致死亡则触发，触发后这一回合将不受伤害，不导致死亡则不触发，正常掉血，触发后buff消失）持续2回合
-- 增加速度视作buff

function SK_304_Action(battleid, casterid)

	Battle.TargetOn(battleid)
	local skillid = 304		-- 技能id
	local skillAttack = 10	-- 技能攻击
	--local attackNum = 0		-- 攻击个数
	
	local t = minPropertyOne(battleid, casterid,"CPT_CHP")	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择  己方血量最少的
	
	Battle.Cure(battleid,t,0,0)
	sys.log("女娲对己方血量最少的单位免疫一回合 1")
	Battle.AddBuff(battleid,casterid,t,106,1)   --（暂时么有这个函数） 造物。使一个友方单位免疫一次会导致其死亡的伤害
	sys.log("女娲对己方血量最少的单位免疫一回合 2")
	Battle.TargetOver(battleid)
	
	return  true
		
	
end

