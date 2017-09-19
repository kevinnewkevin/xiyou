using UnityEngine;
using System.Collections;
using System.Collections.Generic;
public class JieHunSystem 
{
	private static JieHunSystem _instance;
	private COM_Chapter _chapterData;
    private COM_Chapter[] Chapters;
    private List<COM_Chapter> chapteList = new List<COM_Chapter>();
    private List<COM_Chapter> chapteEasyList = new List<COM_Chapter>();
    private List<COM_Chapter> chapteHardList = new List<COM_Chapter>();
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

    public void InitChapterData(COM_Chapter[] data)
    {
        chapteList.Clear();
        for (int i = 0; i < data.Length; i++)
        {
            chapteList.Add(data[i]);
        }
        InitEasy();
    }

    private void InitEasy()
    {
        chapteHardList.Clear();
        chapteEasyList.Clear();
        for (int i = 0; i < chapteList.Count; i++)
        {
            HeroStroyData hData = HeroStroyData.GetData(chapteList[i].ChapterId);
            if (hData.Type_ == 1)
            {
                chapteEasyList.Add(chapteList[i]);
            }
            else if (hData.Type_ == 2)
            {
                chapteHardList.Add(chapteList[i]);
            }
        }
    }

    public List<COM_Chapter> ChapterDataList
    {
        get
        {
            return chapteList;
        }
    }

    public List<COM_Chapter> ChapterEasyDataList
    {
        get
        {
            return chapteEasyList;
        }
    }

    public List<COM_Chapter> ChapterHardDataList
    {
        get
        {
            return chapteHardList;
        }
    }

    public  COM_Chapter GetChapterData ( int id )
    {
        for (int i = 0; i < chapteList.Count; i++)
        {
            if (chapteList[i].ChapterId == id)
                return chapteList[i];
        }
        return null;
    }

}

