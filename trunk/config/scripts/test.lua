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

local u1 = Player.GetPlayerUnit(100,2222)
sys.log(u1)

local u2 = Player.GetUnitProperty(10)
local u3 = Player.GetUnitProperty(1)

sys.log("sssssssss "..u2.." aaaaaaaaaaaaa "..u3)

function Action(battleid, casterid)
	local targetids = Player.GetEnemy(battleid, casterid)
	sys.log("targetids= ".. targetids)
end

sys.log("test.lua loaded")
