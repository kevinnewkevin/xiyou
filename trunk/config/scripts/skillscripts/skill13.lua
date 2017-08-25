sys.log(" skill 13 start")

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
-- 女娲1号技能 为所有友方单位回复相当于法术强度50%的生命值。

-- 物理强度视作buff Battle.buff

function SK_112_Action(battleid, casterid)

	Battle.TargetOn(battleid) --清空数据
	local skillid = 112		-- 技能id

	local  attackNum = 0   --攻击个数

	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标
	
	local  caster_attack = Player.GetUnitMtk(battleid,casterid)  --获取攻击者属性
	
	for i,v in ipairs(t) do
	
		local crit = Battle.GetCrit(skillid)   --是否暴击
	
		Battle.Cure(battleid,v,caster_attack*0.5,crit)      --回血 公式(法术强度的50%）
		
		sys.log("skill13 ")
	end
	Battle.TargetOver(battleid) --赋给下一个目标
	return  true
	 
end

sys.log( "skill 13 end")