sys.log("SK_285_Action cri 2")

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
-- 主角必杀技2技能 使自己方所有目标身上的增益buff效果提升


function SK_285_Action(battleid, casterid)
	
	local skillid = 285		-- 技能id

	local  attackNum = 0   --攻击个数
	
	local level = 1

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		
		Battle.Attack(battleid, casterid, v, 0, 0)
		
		local buffNum = Player.PopAllBuffByDebuff(battleid,v)
		
		local data = level/100
		
		if  buffNum  > 0 then 
		
			Battle.BuffChangeData(battleid,v,data)
		
		end
		
		Battle.TargetOver(battleid)
	
	
	end
	
	
	return  true
	 
end
function SK_286_Action(battleid, casterid)
	
	local skillid = 286	-- 技能id

	local  attackNum = 0   --攻击个数
	
	local level = 2

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		
		Battle.Attack(battleid, casterid, v, 0, 0)
		
		local buffNum = Player.PopAllBuffByDebuff(battleid,v)
		
		local data = level/100
		
		if  buffNum  > 0 then 
		
			Battle.BuffChangeData(battleid,v,data)
		
		end
		
		Battle.TargetOver(battleid)
	
	end
	
	
	return  true
	 
end
function SK_287_Action(battleid, casterid)
	
	local skillid = 287	-- 技能id

	local  attackNum = 0   --攻击个数
	
	local level = 3

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		
		Battle.Attack(battleid, casterid, v, 0, 0)
		
		local buffNum = Player.PopAllBuffByDebuff(battleid,v)
		
		local data = level/100
		
		if  buffNum  > 0 then 
		
			Battle.BuffChangeData(battleid,v,data)
		
		end
		
		Battle.TargetOver(battleid)
	
	
	end
	
	
	return  true
	 
end
function SK_288_Action(battleid, casterid)
	
	local skillid = 288		-- 技能id

	local  attackNum = 0   --攻击个数
	
	local level = 4

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		Battle.Attack(battleid, casterid, v, 0, 0)
		
		local buffNum = Player.PopAllBuffByDebuff(battleid,v)
		
		local data = level/100
		
		if  buffNum  > 0 then 
		
			Battle.BuffChangeData(battleid,v,data)
		
		end
		
		Battle.TargetOver(battleid)
	

	end
	
	
	return  true
	 
end

function SK_289_Action(battleid, casterid)
	
	local skillid = 289		-- 技能id

	local  attackNum = 0   --攻击个数
	
	local level = 5

	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		
		Battle.Attack(battleid, casterid, v, 0, 0)
		
		local buffNum = Player.PopAllBuffByDebuff(battleid,v)
		
		local data = level/100
		
		if  buffNum  > 0 then 
		
			Battle.BuffChangeData(battleid,v,data)
		
		end
		
		Battle.TargetOver(battleid)
	
		
	end
	
	
	return  true
	 
end