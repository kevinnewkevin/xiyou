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
	public int chapterID = 0;
	public int chapterBox = 0;
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

    public void UpdateChapterData(COM_Chapter chapter)
    {
        for (int i = 0; i < chapteList.Count; i++)
        {
            if (chapteList[i].ChapterId == chapter.ChapterId)
            {
                chapteList[i] = chapter ;
                break;
            }
        }
        for (int j = 0; j < chapteEasyList.Count; j++)
        {
            if (chapteEasyList[j].ChapterId == chapter.ChapterId)
            {
                chapteEasyList[j] = chapter;
                break;
            }
        }
        for (int k = 0; k < chapteHardList.Count; k++)
        {
            if (chapteHardList[k].ChapterId == chapter.ChapterId)
            {
                chapteHardList[k] = chapter;
                break;
            }
        }
        UIManager.SetDirty("jiehun");
        UIManager.SetDirty("guanka");
    }
	
	public void AddChapterData(COM_Chapter chapter)
	{
		if (HeroStroyData.GetData (chapter.ChapterId) == null)
			return;
		chapteList.Add (chapter);
		InitEasy ();		
		UIManager.SetDirty("jiehun");
		UIManager.SetDirty("guanka");		
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

	public  void  UpdataChapterRewardData ( int id,int boxId)
		{
			for (int i = 0; i < chapteList.Count; i++) 
				{
				if (chapteList [i].ChapterId == id) 
				{
				if (chapteList [i].StarReward == null)
					{
							chapteList [i].StarReward = new int[3];
							chapteList [i].StarReward [0] = (boxId + 1) * 10;
					}
					else if (chapteList [i].StarReward.Length < 3)
					{
						int[] num = chapteList [i].StarReward;
						chapteList [i].StarReward = new int[3];
						for (int j = 0; j < num.Length; j++) 
						{
							chapteList [i].StarReward [j] = num [j]; 
						}
						for (int k = 0; k < chapteList [i].StarReward.Length; k++) 
						{
							if (chapteList [i].StarReward [k] == 0) 
							{
									chapteList [i].StarReward [k] = (boxId + 1) * 10;
							}
						}
					}				
			}

			}
		}

}