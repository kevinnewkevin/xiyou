using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class GuildSystem 
{
	public static COM_Guild myGuild; 
	public static List<COM_GuildMember> guildMemberList = new List<COM_GuildMember>();
	public static List<COM_GuildViewerData> guildViewerList = new List<COM_GuildViewerData>();
	public static COM_GuildDetails searchData;
	
	public static void InitGuildMember(COM_GuildMember[] guildMember)
	{
		if (guildMember == null)
				return;
		for (int i = 0; i < guildMember.Length; i++)
		{
			guildMemberList.Add (guildMember [i]); 
		}
	}

	public static void InitViewer(COM_GuildViewerData[] guilds)
	{
		if (guilds == null)
			return;
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




	

}

