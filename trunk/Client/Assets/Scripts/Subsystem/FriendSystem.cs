using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class FriendSystem 
{
	public  static List<COM_Friend> friendList = new List<COM_Friend>();
	public  static List<COM_Friend> blackList = new List<COM_Friend>();
	public  static List<string> applyFriendList = new List<string>();
	public static  COM_Friend[] randomFriends;
	public static  COM_Friend findFriend;

	public static Dictionary<long, List<COM_Friend>> friendRecvList = new Dictionary<long, List<COM_Friend>>();
	
	public static void InitFriends( COM_Friend[] friends)
	{
		for (int i = 0; i < friends.Length; i++) 
		{
			friendList.Add (friends [i]);
		}
	}
	public static void InitBlacks( COM_Friend[] black)
	{
		if (black == null)
			return;
		for (int i = 0; i < black.Length; i++) 
		{
			blackList.Add (black [i]);
		}
	}

	public static void AddFriend( COM_Friend friends)
	{
		friendList.Add (friends );
	}

	public static void AddBlack( COM_Friend black)
	{
		blackList.Add (black );
	}

	public static void	ApplyFriend(ref string name)
	{
		applyFriendList.Add (name);
	}

	public  static bool	chatFriend(ref COM_Friend friend)
	{
		if(!friendRecvList.ContainsKey(friend.InstId))
				friendRecvList[friend.InstId] = new List<COM_Friend>();
		friendRecvList [friend.InstId].Add (friend);
		return true;
	}

		public static int GetFriendNum()
		{
				return 	friendList.Count;	
		}

		public static int GetBalckNum()
		{
				return 	blackList.Count;	
		}

		public static int GetApplyNum()
		{
			return 	applyFriendList.Count;	
		}


		public static void	DelFriend(long id)
		{
			for (int i = 0; i < friendList.Count; i++) 
			{
				if (friendList [i].InstId == id) 
				{
					friendList.Remove (friendList [i]);
				}
			}
		}
}

