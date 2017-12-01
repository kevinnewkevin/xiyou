sys.log(" 魔礼红 SK_369_Action 开始")

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
--抵御对方攻击，给己方所有目标增加一个吸收最大生命值50%伤害的护盾

function SK_369_Action(battleid,casterid)
	
	local skillid = 369

	local attackNum = 0

	local t = Player.GetFriends(battleid,casterid,attackNum)

	for i,v in ipairs(t)do

		Battle.TargetOn(battleid)
	
		local max_hp = Player.GetUnitProperty(battleid,v,"CPT_HP")

		sys.log(max_hp)

		max_hp = max_hp * 0.5

		Battle.Cure(battleid,v,0,0)

		sys.log("魔礼红 11111")

		Battle.AddBuff(battleid,casterid,v,184,max_hp)
		sys.log("魔礼红 22222")
		Battle.TargetOver(battleid)
	end
	return true
end
sys.log(" 魔礼红 SK_369_Action 结束")