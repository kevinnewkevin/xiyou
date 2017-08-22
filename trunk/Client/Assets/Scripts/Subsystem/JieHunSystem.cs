using UnityEngine;
using System.Collections;

public class JieHunSystem 
{
	private static JieHunSystem _instance;
		private COM_Chapter _chapterData;

	public static JieHunSystem instance
	{
		get
		{
				if(_instance == null)
							_instance = new JieHunSystem();
				return _instance;
		}
	}


	public  COM_Chapter ChapterData
	{
		set
		{ 
			_chapterData = value;
			UIManager.SetDirty ("guanka");
		}
		get
		{
				return _chapterData;
		}
	}
 
}

