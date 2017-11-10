using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class FriendSystem 
{
	public static List<COM_Friend> friendList = new List<COM_Friend>();
	public static List<COM_Friend> blackList = new List<COM_Friend>();
	public static List<long> latelyList = new List<long>();
	public static List<COM_Friend > applyFriendList = new List<COM_Friend >();
	public static COM_Friend[] randomFriends;
	public static COM_Friend findFriend;
	public static bool isApplyFriend = false;
	public static List<string> newCahtList = new List<string>();
	

	public static Dictionary<long, List<COM_Chat>> friendRecvList = new Dictionary<long, List<COM_Chat>>();
	public static Dictionary<string, List<COM_Chat>> friendRecvListStr = new Dictionary<string, List<COM_Chat>>();
	public static void InitFriends( COM_Friend[] friends)
	{
		if (friends == null)
			return;
		for (int i = 0; i < friends.Length; i++) 
		{
			friendList.Add (friends [i]);
		}
	}
	public static void InitBlacks( COM_Friend[] black)
	{
		if (black == null)
				return;
		if (black == null)
			return;
		for (int i = 0; i < black.Length; i++) 
		{
			blackList.Add (black [i]);
		}
	}

	public static void AddFriend( COM_Friend friends)
	{
		if(!friendList.Contains(friends))
			friendList.Add (friends );
		DelApplyFriend (friends.InstId);
	}

	public static void AddLatelyFriend(long  instId)
	{
		if(!latelyList.Contains(instId))
			latelyList.Add (instId );
	}

	public static void DelLatelyFriend(long  instId)
	{
			if(latelyList.Contains(instId))
				latelyList.Remove(instId );
	}

	public static void AddBlack( COM_Friend black)
	{
		blackList.Add (black );
	}

	public static void	ApplyFriend(COM_Friend friend)
	{
		for (int i = 0; i < applyFriendList.Count; i++) 
		{
				if (applyFriendList [i].Name == friend.Name) 
				{
						return;
				}
		}
		applyFriendList.Add (friend);
		isApplyFriend = true;
	}

	public  static bool	chatFriend(long instId,COM_Chat friend)
	{
		if(!friendRecvList.ContainsKey(instId))
				friendRecvList[instId] = new List<COM_Chat>();
		friendRecvList [instId].Add (friend);
		return true;
	}

		public  static bool	chatFriendStr(string name,COM_Chat friend)
		{
				if(!friendRecvListStr.ContainsKey(name))
						friendRecvListStr[name] = new List<COM_Chat>();
				friendRecvListStr [name].Add (friend);
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

	public static int GetLatelyListNum()
	{
		return 	latelyList.Count;	
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
				break;
			}
		}
	}


		public static void	DelApplyFriend(long id)
		{
			for (int i = 0; i < applyFriendList.Count; i++) 
			{
				if (applyFriendList [i].InstId == id) 
				{
					applyFriendList.Remove (applyFriendList [i]);
					break;
				}
			}
		}


	public static void	DelBlack(long id)
	{
		for (int i = 0; i < blackList.Count; i++) 
		{
			if (blackList [i].InstId == id) 
			{	
				blackList.Remove (blackList [i]);
				break;
			}
		}
	}

	public static COM_Friend GetFriend(long id)
	{
		for (int i = 0; i < friendList.Count; i++) 
		{
			if (friendList [i].InstId == id) 
			{
				return friendList [i];
			}
		}
		return null;
	}

	public static List<COM_Chat> GetFriendChat(long InstId)
	{
		if (!friendRecvList.ContainsKey(InstId))
			return new List<COM_Chat>();
		return friendRecvList [InstId];
	}

	public static List<COM_Chat> GetFriendChatStr(string name)
	{
		if (!friendRecvListStr.ContainsKey(name))
				return new List<COM_Chat>();
		return friendRecvListStr [name];

	}

	public static void AddNewCahtList(string name)
	{
		if(!newCahtList.Contains(name))
			newCahtList.Add (name );
		UIManager.SetDirty ("zhujiemian");
	}
	public static void DelNewCahtList(string name)
	{
		if(newCahtList.Contains(name))
			newCahtList.Remove(name );
		UIManager.SetDirty ("zhujiemian");
	}

	public static int GetNewCahtListNum()
	{
		return 	newCahtList.Count;	
	}

	public static bool IsNewCaht(string name)
	{
		return 	newCahtList.Contains (name);
	}


	public static bool IsInBlack(long instId)
	{
		for (int i = 0; i < blackList.Count; i++) 
		{
			if (blackList [i].InstId == instId) 
			{	
					return true;
			}
		}
		return false;
	}

}


