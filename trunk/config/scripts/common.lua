
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
	
	if _BoolStrongBuff then 
	
		 _strongPer = Player.ClacStrongPer(battleid,casterid)
		 
	end
	
	damage =damage + damage* _strongPer
	
	
	
	--------------------------------------------------------------------------
	--------------------------------------------------------------------------
	
	--判断承受者有无增加伤害的buff  百分比
	
	local _BoolWeakBuff = Player.GetCheckSpec(battleid,targetid,"BF_WEAK")
	
	
	if  _BoolWeakBuff then
	
		_weakPer = Player.ClacWeakPer(battleid,targetid)
		
	
	end
	
	damage = damage + damage * _weakPer
	
	--判断承受者有无减伤害的buff   百分比
	
	local _boolSheldBuff = Player.GetCheckSpec(battleid,targetid,"BF_SHELD")

	
	if _boolSheldBuff then 
		
		_sheldPer = Player.ClacSheld(battleid,targetid)
		
		Player.ChangeBuffTimes(battleid,targetid)   --开始前清理数据
	end
	
	damage = damage - damage * _sheldPer
	
	
	------------------------------------------------------------------------------
	------------------------------------------------------------------------------
	
	--取得目标的护盾值
	
	sheld = Player.GetOneSheld(battleid,targetid)
	
	Player.DownSheld(battleid,targetid,damage)  --减護盾值
	
	if sheld >= damage then 
	
		damage = 0
	
	end
	
	--最终伤害
	
	damage = damage - sheld
	
	
	return damage
	
end