sys.log("test.lua load start")
-- sys.err("123123123123123123123"..sys.test(12,23))


sys.log("1111111111111")
function Print_Ln()					
	sys.log("hello world")
end

function Print_Ln2(a,b,c,d)					
	sys.log("Print_Ln2"..","..a..","..b..","..c..","..d)
	return 1,2,3,4,5,6,7,8,9
end


function Action(battleid, casterid)
	local targetids = Player.GetEnemy(battleid, casterid)
	sys.log("targetids= ".. targetids)
end

sys.log("test.lua load done")

--Player.GetStrings()
