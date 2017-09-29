sys.log("skill 18 start")

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
-- 子牙 杏黄旗。为所有的友方单位分别提供一个buff，提升所有攻击20%。这个buff不会消失，可以叠加。每次释放增加20%，直到被驱散或者姜子牙死亡，buff消失
-- 增加速度视作buff

function SK_297_Action(battleid, casterid)
	local skillid = 297		-- 技能id
	local skillAttack = 10	-- 技能攻击
	local attackNum = 0		-- 攻击个数
	
	local t = Player.GetFriends(battleid, casterid, attackNum)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	--local caster_attack = Player.GetUnitProperty(battleid, casterid, "CPT_ATK")	-- 获取到攻击者的属性
	
	
	
	for i,v in ipairs(t)	do
		Battle.TargetOn(battleid)
		local mtk = Player.GetUnitMtk(battleid,v)
		local atk = Player.GetUnitAtk(battleid,v)
		
		local per = 0.2
		
		local mtk_damage = mtk * per
		local atk_damage = atk * per
		sys.log("子牙杏黄旗给友军加buff")
		Battle.cure(battleid, v, 0, 0)
		sys.log("子牙杏黄旗给友军加物理攻击")
		Battle.AddBuff(battleid,casterid,v,102,atk_damage)
		sys.log("子牙杏黄旗给友军加法术攻击")
		Battle.AddBuff(battleid,casterid,v,105,mtk_damage)
		
		Battle.TargetOver(battleid)
		

	end
	sys.log("子牙杏黄旗")
	return 1
end

sys.log("skill 18 end")