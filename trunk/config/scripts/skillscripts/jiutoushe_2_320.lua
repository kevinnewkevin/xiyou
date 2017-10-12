sys.log("九头蛇 SK_320_Action 开始")

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
-- 群蛇乱舞。对敌方随机目标进行9次攻击（随机9次目标），每次攻击造成30%法术强度伤害，每次攻击会使目标降低10%速度，持续1回合。
-- 增加速度视作buff

function SK_320_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	local skillid = 320		-- 技能id
	local skillAttack = 10	-- 技能攻击
	--local attackNum = 0		-- 攻击个数

	local t = Player.GetMainTarget(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	local sudu = Player.GetUnitProperty(battleid, casterid, "CPT_AGILE")
	
	local magic_damage=Player.GetMagicDamage(battleid,casterid,t)
	
	local  true_damage = magic_damage*9
	
	sys.log("九头蛇 群蛇乱舞对目标9次攻击   "..t.. " 造成 法术伤害  "..true_damage )
	
	local damage = ClacDamageByAllBuff(battleid,casterid,t,true_damage)
	sys.log("九头蛇 群蛇乱舞对目标9次攻击   "..t.. " 造成 最终法术伤害  "..damage )
	
	--判断伤害
	if damage <= 0 then
		
		damage = 1
		
	end
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	local mag_damage = damage*0.3
	
	local  sudu_del = sudu*0.1
	sys.log("九头蛇 群蛇乱舞对目标9次攻击   "..t.. " 造成 法术伤害的30%  "..mag_damage )
	Battle.Attack(battleid,casterid,t,mag_damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
	
	Battle.AddBuff(battleid,casterid,t,117,sudu_del)  --每次都到伤害降低10%速度

	
	Battle.TargetOver(battleid)
	
	return 1
end

sys.log("九头蛇 SK_320_Action 结束")