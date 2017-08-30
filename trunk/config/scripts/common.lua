
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