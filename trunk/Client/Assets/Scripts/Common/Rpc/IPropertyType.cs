public class IPropertyType{
  static string[] _NAMES = new string[]{
    "IPT_MIN",
    "IPT_HP",
    "IPT_PHYLE",
    "IPT_TITLE",
    "IPT_EXPERIENCE",
    "IPT_LEVEL",
    "IPT_COPPER",
    "IPT_SILVER",
    "IPT_GOLD",
    "IPT_MAX"
  };
  public const int IPT_MIN = 0;
  public const int IPT_HP = 1;
  public const int IPT_PHYLE = 2;
  public const int IPT_TITLE = 3;
  public const int IPT_EXPERIENCE = 4;
  public const int IPT_LEVEL = 5;
  public const int IPT_COPPER = 6;
  public const int IPT_SILVER = 7;
  public const int IPT_GOLD = 8;
  public const int IPT_MAX = 9;
  public static int ToId(string name){
    for(int i=0; i<NAMES.Length;++i){
      if(_NAMES[i] == name){
        return i;
      }
    }
    return -1;
  }
  public static string ToName(int id){
    if(id<0||id>=_NAMES.Length){
      return "";
    }
    return _NAMES[id];
  }
  public static string[] NAMES{ get {return _NAMES;} }
}
