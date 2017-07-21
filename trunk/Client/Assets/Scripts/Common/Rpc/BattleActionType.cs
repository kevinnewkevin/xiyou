public class BattleActionType{
  static string[] _NAMES = new string[]{
    "BAT_MIN",
    "BAT_CRIT",
    "BAT_SUCK",
    "BAT_RECOVERY",
    "BAT_ADD_STATE",
    "BAT_DEL_STATE",
    "BAT_MAX"
  };
  public const int BAT_MIN = 0;
  public const int BAT_CRIT = 1;
  public const int BAT_SUCK = 2;
  public const int BAT_RECOVERY = 3;
  public const int BAT_ADD_STATE = 4;
  public const int BAT_DEL_STATE = 5;
  public const int BAT_MAX = 6;
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
