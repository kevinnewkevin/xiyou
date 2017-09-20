function Energy_liquid(casterid,itemId)
	myEnergy = Player.GetMyUnitIProperty(casterid, "IPT_ENERGY");
	if myEnergy >= 1000 then
		return "EN_CREATE_PLAYER_SAME_NAME"
	end
	Player.AddMyUnitEnergy(casterid,10);
	sys.log("1111111111111====Energy_liquid====1111111111")
end
