sys.log(" skill 20 start")

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
-- 黄牙老象2号技能 为一个友方目标回复法术强度的生命值，并且使其每次受到伤害都会回复10%法术强度的生命值，持续1回合。

-- 法术强度视作buff  Battle.buff

function SK_119_Action(battleid, casterid)
	local skillid = 119		-- 技能id
	
	local  p = Player.GetTarget(battleid,casterid)  --获取目标 

	--local  add_buff = Battle.AddBuff(battleid,del_buff*0.4)    --降低伤害40%
	
	
	local  damage  = 4 --伤害 公式（攻击属性） --测试
	
	sys.log("skill20 对id为"..p.."的目标造成"..damage.."点伤害")
	
	return  true
	 
end

sys.log( "skill 20 end")