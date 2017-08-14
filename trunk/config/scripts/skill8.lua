
sys.log(" skill 8 start")

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
-- 哪吒2号技能 掷出乾坤圈，击中敌人后选择下一个敌人弹射，弹射3次，对每个敌人造成50%物理强度的伤害。每击中一个敌人增加自己10%的物理强度，持续1回合

-- 物理强度视作buff Battle.buff

function SK_107_Action(battleid, casterid)

	local skillid = 2		-- 技能id

	local  num = 3   --攻击个数

	local  p = Player.GetTargets(battleid,casterid,num)  --获取目标
	
	local  _property = Player.GetUnitProperty(battleid,casterid,"CPT_ATK")  --获取攻击者属性
	
	--local  battleid_buff = Battle.AddBuff(1)  --释放者物理强度 
	
	for i,v in ipairs(p) do
	
		--local  del_buff = Battle.AddBuff(1)  --敌对方物理强度
		
		local defender_def = Player.GetUnitProperty(battleid, v, "CPT_DEF")
	
		--local  demage  = del_buff*0.5-defender_def  --伤害 公式（50%的物理伤害 减 防御 ）
		
		--battleid_buff = battleid_buff + battleid_buff*o.1
		
		local  damage  = 8--测试
		
		--判断伤害
		if damage <= 0 then 
		
			damage = 1
		
		end
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		--sys.log("skill1 对id为"..v.."的目标造成"..damage.."点伤害")
	end
	
	return  true
	 
end

sys.log( "skill 8 end")