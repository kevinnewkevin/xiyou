sys.log("skill 24 start")

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
-- 炼石补天。为所有友方单位回复相当于法术强度50%的生命值。
-- 增加速度视作buff

function SK_303_Action(battleid, casterid)

	
	local skillid = 303		-- 技能id
	local skillAttack = 10	-- 技能攻击
	local attackNum = 0		-- 攻击个数
	
	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标 
	
	local mtk = Player.GetUnitMtk(battleid,casterid)
	
	local mag_atk = mtk * 0.5
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
	
		local crit = Battle.GetCrit(skillid)   --是否暴击
		sys.log("女娲对目标回复的血量   "..mag_atk)
		Battle.Cure(battleid,v,mag_atk,crit)
		sys.log("女娲对目标回复的血量   w完成")
		Battle.TargetOver(battleid)
	
		
	end
	return  true
	
end

sys.log("skill 24 end")