using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class GuildSystem 
{
	public static COM_Guild myGuild; 
	public static List<COM_GuildMember> guildMemberList = new List<COM_GuildMember>();
	public static List<COM_GuildViewerData> guildViewerList = new List<COM_GuildViewerData>();
	public static COM_GuildDetails searchData;
	public static  bool IsShowSearch = false;
    public static List<COM_GuildRequestData> requestList = new List<COM_GuildRequestData>();
	
	public static void InitGuildMember(COM_GuildMember[] guildMember)
	{
		if (guildMember == null)
				return;
		guildMemberList.Clear();
		for (int i = 0; i < guildMember.Length; i++)
		{
			guildMemberList.Add (guildMember [i]); 
		}
	}

	public static void InitViewer(COM_GuildViewerData[] guilds)
	{
		if (guilds == null)
			return;
		guildViewerList.Clear();
		for (int i = 0; i < guilds.Length; i++)
		{
			guildViewerList.Add (guilds [i]); 
		}
	}

	public static void UpdateGuildMember(COM_GuildMember member)
	{
		if (member == null)
				return;
		for (int i = 0; i < guildMemberList.Count; i++)
		{
			if (guildMemberList [i].RoleId == member.RoleId) 
			{
				guildMemberList [i] = member;
				break;
			}
		}
	}

	public static void AddGuildMember(COM_GuildMember member)
	{
		if (member == null)
			return;
		if (guildMemberList.Contains (member))
			return;		
		guildMemberList.Add(member);
	}



	public static void LeaveGuildMember( string name,bool b)
	{
		for (int i = 0; i < guildMemberList.Count; i++)
		{
			if (guildMemberList [i].RoleName == name) 
			{
				guildMemberList.Remove( guildMemberList [i]);
				break;
			}
		}
	}


	public static COM_GuildViewerData findGuildViewer(string name)
	{
		for (int i = 0; i < guildViewerList.Count; i++)
		{
			if (guildViewerList [i].GuildName == name) 
			{
				return	guildViewerList [i];
			}
		}
		return null;
	}

	public static COM_GuildViewerData GetGuildViewer(int id)
	{
		for (int i = 0; i < guildViewerList.Count; i++)
		{
			if (guildViewerList [i].GuildId == id) 
			{
					return	guildViewerList [i];
			}
		}
		return null;
	}

 	
	public static void AddGuildRequest(COM_GuildRequestData data)
    {
        requestList.Add(data);
    }

	 public static void DeleteGuildRequest(long playerid)
    {
        if(requestList == null)
            return;
        
        for(int i=0; i < requestList.Count; ++i)
        {
            if (requestList [i].RoleId == playerid)
            {
                requestList.RemoveAt(i);
                break;
            }
        }
    }

    public static int MyJob()
    {
        for(int i=0; i < guildMemberList.Count; ++i)
        {
            if (GamePlayer.IsMe(guildMemberList [i].RoleId))
                return guildMemberList [i].Job;
        }
        return -1;
    }
}

