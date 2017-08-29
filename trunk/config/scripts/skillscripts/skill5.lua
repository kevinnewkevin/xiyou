sys.log(" skill 5 start")

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
-- 猴子2号技能 为周围的友方单位分别提供一个护盾，护盾抵挡的伤害值相当于姜子牙法术强度的10%。

-- 法术强度视作buff Battle.buff

function SK_104_Action(battleid, casterid)



	local skillid = 104		-- 技能id

	local  attackNum = 0  -- 攻击个数
	

	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标
	
	local  damage = Player.GetUnitMtk(battleid,casterid)  --获取法术强度百分比
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid) --清空数据
		
		Battle.Cure(battleid, v,0, 0)
		
		Battle.AddBuff(battleid,casterid,v,4,damage*0.1)  --给友方分别提供一个盾牌
		
		Battle.TargetOver(battleid)  --赋给下个目标
	
	end

	
	return  true
	 
end

sys.log( "skill 5 end")