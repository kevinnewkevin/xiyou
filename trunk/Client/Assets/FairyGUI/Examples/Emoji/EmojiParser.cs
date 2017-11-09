using FairyGUI;
using FairyGUI.Utils;

public class EmojiParser : UBBParser
{
	static EmojiParser _instance;
	public new static EmojiParser inst
	{
		get
		{
			if (_instance == null)
				_instance = new EmojiParser();
			return _instance;
		}
	}

//	private static string[] TAGS = new string[]
//		{ "88","am","bs","bz","ch","cool","dhq","dn","fd","gz","han","hx","hxiao","hxiu" };
    public void RegistEmojiTags ()
	{
        string tagstr = Define.GetStr("EmojiTags");
        string[] tags = tagstr.Split(new char[]{','}, System.StringSplitOptions.RemoveEmptyEntries);
        foreach (string ss in tags)
		{
			this.handlers[":"+ss] = OnTag_Emoji;
		}
	}

	string OnTag_Emoji(string tagName, bool end, string attr)
	{
		return "<img src='" + UIPackage.GetItemURL("liaotian", tagName.Substring(1).ToLower()) + "'/>";
	}
}