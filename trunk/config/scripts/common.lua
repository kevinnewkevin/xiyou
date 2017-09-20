
function sys_log(val)
	if type(val) == "table" then
		sys.log(table.concat(val,","))
	elseif type(val) == "number" then
		sys.log(""..val)
	else 
		sys.log(val)
	end
end

function set_random_seed()
	local time = os.GetTime()
	sys.log("percent gettime ".. time)
	math.randomseed(tostring(time):reverse():sub(1, 7))
end

function percent()
	return math.random(1, 100)

end

--获取人物属性值(property)最大的一个人物
function max_property_one(battleid,unitid,property)

	local t = Player.GetTargets(battleid,unitid)
	
	
	local arr_property = {}    --人物属性
	
	
	local arr_pro = {}   -- 人物
	
	
	for i,v in ipairs(t) do
	
		local dangqian_property = Player.GetUnitProperty(battleid,v,property) --获取当前属性质量
		
		arr_property[i] = dangqian_property
		
		arr_pro[i] = v
		
	end
	
	local maxOfT = math.max(unpack(arr_property)) --最大数性值
	
	
	local idx = 0
	
	for i=1,#arr_property,1 do
	
		if arr_property[i] == maxOfT then
			idx = i
		end
		
	end
	
	return arr_pro[idx]   --返回最大血的人物  单个
	
end


--获取人物属性值(property)最少的一个人
function min_property_one(battleid,unitid,property)

	local t = Player.GetTargets(battleid,unitid)
	
	local arr_property = {}    --人物属性
	
	
	local arr_pro = {}   -- 人物
	
	for i,v in ipairs(t) do
	
		local dangqian_property = Player.GetUnitProperty(battleid,v,property) --获取当前属性质量
		
		arr_property[i] = dangqian_property
		
		arr_pro[i] = v
		
	end
	
	
	local minOfT = math.min(unpack(arr_property)) -- 最小属性值
	
	
	local idx = 0
	
	for i=1,#arr_property,1 do
	
		if arr_property[i] == minOfT then         
			idx = i
		end
		   
	end
	
	return arr_pro[idx]   --返回最大血的人物  单个
	
end

-- 计算最终伤害数值 战斗id ,释放者ID, 目标ID
function ClacDamageByAllBuff(battleid,casterid,targetid,damage) 
	-- 取得释放者的所有伤害增加数字
	-- 取得目标的所有减伤数字
	-- 然后计算
	
	local sheld = 0			-- 获取到的护盾数值
	local max_per = 75		-- 最大减伤百分比 
	local per = 0			-- 获取到的减伤百分比
	
	--判断释放者是否有增加输出伤   buff  百分比
	local  _BoolStrongBuff = Player.GetCheckSpec(battleid,casterid,"BF_STRONG")
	
	if _BoolStrongBuff == 1 then 
	
		_strongPer = Player.ClacStrongPer(battleid,casterid)
		
		sys.log("ClacDamageByAllBuff 增加输出伤害的百分比" .. _strongPer)
		
		sys.log("ClacDamageByAllBuff 这是传参 damage" ..",".. damage)
		 
		sys.log("ClacDamageByAllBuff 增加输出伤伤害" .. damage* _strongPer)
		
		damage =damage + damage* _strongPer
		 
		sys.log("ClacDamageByAllBuff 增加输出伤后得到的伤害值" .. damage)
		 
	end
	
	
	
	--------------------------------------------------------------------------
	--------------------------------------------------------------------------
	
	--判断承受者有无增加受到伤害的buff  百分比
	local _BoolWeakBuff = Player.GetCheckSpec(battleid,targetid,"BF_WEAK")
	
	if  _BoolWeakBuff == 1 then
	
		_weakPer = Player.ClacWeakPer(battleid,targetid)
		
		sys.log("ClacDamageByAllBuff 增加受到伤害的百分比" .. _weakPer)
		
		sys.log("ClacDamageByAllBuff 这是传参 damage"..",".. damage)
		
		sys.log("ClacDamageByAllBuff 增加受到伤害" .. damage * _weakPer)
		
		damage = damage + damage * _weakPer
		
		sys.log("ClacDamageByAllBuff 增加受到伤害后得到的伤害值" .. damage)
		
	end
	
	
	
	--判断承受者有无减伤害的buff   百分比
	local _boolSheldBuff = Player.GetCheckSpec(battleid,targetid,"BF_SHELD")

	
	if _boolSheldBuff == 1 then 
		
		_sheldPer = Player.ClacSheld(battleid,targetid)
		
		sys.log("ClacDamageByAllBuff 减伤伤害的百分比" .. _sheldPer)
		
		sys.log("ClacDamageByAllBuff 这是传参 damage"..",".. damage)
		
		sys.log("ClacDamageByAllBuff 减伤伤害" .. damage * _sheldPer)
		
		damage = damage - damage * _sheldPer
		
		sys.log("ClacDamageByAllBuff 减伤伤害得到的值" .. damage)
		
		Player.ChangeBuffTimes(battleid,targetid)   --开始前清理数据
	end
	
	------------------------------------------------------------------------------
	------------------------------------------------------------------------------
	
	--取得目标的护盾值
	sheld = Player.GetOneSheld(battleid,targetid)
	
	Player.DownSheld(battleid,targetid,damage)  --减護盾值
	
	if sheld >= damage then 
	
		return 0
	
	end
	
	--最终伤害
	
	damage = damage - sheld
	
	sys.log("ClacDamageByAllBuff 最终" .. damage)
	
	CheckAttackCure(battleid, targetid)
	
	-- 注   例如：造成150%伤害 这150% 会在技能skill的attack 里乘以0.15
	
	return damage
	
end


--针对于特殊受伤后回血用的
function CheckAttackCure(battleid, unitid)
	--获取到有没有121 受击回血
	local  buffid = Player.GetCheckSpec(battleid,unitid,"BF_ATKCURE")
	
	if buffid == 1 then
	
		Battle.BuffUpdate(battleid, unitid, buffid)
		
	end
	
end

--主角主动技能   num 是 技能前面乘的数字
function skill_user_atk(battleid,casterid,targetid,level,buffid,num)
	-- jia buff
	 
	local data = num * level
	
	Battle.Attack(battleid,casterid,targetid,0,0)
	
	Battle.AddBuff(battleid,casterid,targetid,buffid,data)
	
end

