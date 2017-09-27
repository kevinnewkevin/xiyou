using UnityEngine;
using System.Collections;

public class ShopSystem
{
	public static int buyType; 	
	public static COM_ItemInst[] _buyItems;
	public static COM_ItemInst[] BuyItems
	{
		set
		{
			_buyItems = value;
		}
		get
		{ 
			return _buyItems;
		}
	}

}

