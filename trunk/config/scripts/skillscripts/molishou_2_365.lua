sys.log("魔礼寿 SK_365_Action 开始")

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
-- 删除己方一个目标身上的所有负面状态，并且提升该目标双防20%

-- 法术强度视作buff  Battle.buff

function SK_365_Action(battleid, casterid)

	
	
	local skillid = 365

	local t = Player.GetFriend(battleid,casterid)
	
	Battle.TargetOn(battleid)

	--local isdelbuff = Battle.HasDebuff(battleid,t)

	--if  isdelbuff == true then

		Player.PopAllBuffByDebuff(battleid,t)

	--end
	--Battle.TargetOver(battleid)

	--Battle.TargetOn(battleid)

	local atkdef = Player.GetUnitProperty(battleid,t,"CPT_DEF")

	local magdef = Player.GetUnitProperty(battleid,t,"CPT_MAGIC_DEF")

	local per = 0.2

	atkdef = atkdef * per

	magdef = magdef * per

	Battle.Attack(battleid,casterid,t,0,0)

	Battle.AddBuff(battleid,casterid,t,181,atkdef)

	Battle.AddBuff(battleid,casterid,t,180,magdef)

	Battle.TargetOver(battleid)	
			

	return ture 
	 
end
sys.log("魔礼寿 SK_365_Action 结束")
