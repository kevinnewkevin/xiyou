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
	local skillid = 1		-- 技能id
	

	local  num = 0   --攻击个数

	local  p = Player.GetTargets(battleid,casterid,num)  --获取目标
	
	for i,v in ipairs(p) do
		
		--local  damage_buff = Battle.AddBuff(1)  --法术强度
	
		--local  recovery  = Battle.AddBuff(damage_buff*50%)      --回血 公式(法术强度的50%）
		
		local  recovery  = 13 --测试
	
		--sys.log("skill1 对id为"..v.."的目标造成"..recovery.."点回血")
	end
	
	return  true
	 
end

sys.log( "skill 13 end")