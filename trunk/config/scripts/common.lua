
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