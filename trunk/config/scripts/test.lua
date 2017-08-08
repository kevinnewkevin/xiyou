sys.log("test.lua load")
-- sys.err("123123123123123123123"..sys.test(12,23))


sys.log("1111111111111")
function Print_Ln()					
	sys.log("hello world")
end

function Print_Ln2(a,b,c,d)					
	sys.log(""..","..a..","..b..","..c..","..d)
	return 1,2,3,4,5,6,7,8,9
end

local a = 231231
local b = "我看你咋加"
local c = a + b


function Action(battleid, casterid)
	local targetids = Player.GetEnemy(battleid, casterid)
	sys.log("targetids= ".. targetids)
end

sys.log(""..c)
