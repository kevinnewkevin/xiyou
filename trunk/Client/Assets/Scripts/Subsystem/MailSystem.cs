using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class MailSystem 
{
	static public List<COM_Mail> mailList = new List<COM_Mail>();
	
	static public void AppendMail( COM_Mail[] mails)
	{
		for (int i = 0; i < mails.Length; i++) 
		{
			mailList.Add (mails [i]);
		}
	}

	static public void DelMail( int id)
	{
		for (int i = 0; i < mailList.Count; i++) 
		{
			if (mailList [i].MailId == id) 
			{
				mailList.Remove (mailList [i]);
				break;
			}
		}
	}

	static public void UpdateMail( COM_Mail mail)
	{
		for (int i = 0; i < mailList.Count; i++) 
		{
			if (mailList [i].MailId == mail.MailId) 
			{
				mailList[i] = mail;
				break;
			}
		}
	}
	
		static public COM_Mail GetMail(int id)
		{
			for (int i = 0; i < mailList.Count; i++) 
			{
				if (mailList [i].MailId == id) 
				{
					return	mailList[i];
				}
			}
			return	null;
		}

}

