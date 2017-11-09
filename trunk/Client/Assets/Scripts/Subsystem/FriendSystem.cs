using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class FriendSystem 
{
	public  static List<COM_Friend> friendList = new List<COM_Friend>();
	public  static List<string> applyFriendList = new List<string>();
	public static Dictionary<int, List<COM_Friend>> friendRecvList;

		public static void InitFriends( COM_Friend[] friends)
	{
		for (int i = 0; i < friends.Length; i++) 
		{
				friendList.Add (friends [i]);
		}
		//friendList = new List<COM_Friend> (friends);
	}

	public static void	ApplyFriend(ref string name)
	{
		applyFriendList.Add (name);
	}

	public  static bool	RecvFriend(ref COM_Friend friend)
	{
				
		return true;
	}

		public static int GetFriendNum()
		{
				return 	friendList.Count;	
		}

}

