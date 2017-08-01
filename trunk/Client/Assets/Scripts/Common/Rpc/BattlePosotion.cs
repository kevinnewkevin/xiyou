public class BattlePosotion{
  static string[] _NAMES = new string[]{
    "BP_RED_1",
    "BP_RED_2",
    "BP_RED_3",
    "BP_RED_4",
    "BP_RED_5",
    "BP_RED_6",
    "BP_BLUE_1",
    "BP_BLUE_2",
    "BP_BLUE_3",
    "BP_BLUE_4",
    "BP_BLUE_5",
    "BP_BLUE_6",
    "BP_MAX"
  };
  public const int BP_RED_1 = 0;
  public const int BP_RED_2 = 1;
  public const int BP_RED_3 = 2;
  public const int BP_RED_4 = 3;
  public const int BP_RED_5 = 4;
  public const int BP_RED_6 = 5;
  public const int BP_BLUE_1 = 6;
  public const int BP_BLUE_2 = 7;
  public const int BP_BLUE_3 = 8;
  public const int BP_BLUE_4 = 9;
  public const int BP_BLUE_5 = 10;
  public const int BP_BLUE_6 = 11;
  public const int BP_MAX = 12;
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
