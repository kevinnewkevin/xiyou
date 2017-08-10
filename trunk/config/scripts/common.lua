
function sys_log(val)
	if type(val) == "table" then
		sys.log(table.concat(val,","))
	elseif type(val) == "number" then
		sys.log(""..val)
	else 
		sys.log(val)
	end
end