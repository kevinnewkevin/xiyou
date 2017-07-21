public class ErrorNo{
  static string[] _NAMES = new string[]{
    "EN_MIN",
    "EN_CREATE_PLAYER_SAME_NAME",
    "EN_CREATE_PLAYER_ILLEGAL_NAME",
    "EN_CREATE_PLAYER_SECOND_TIME",
    "EN_MAX"
  };
  public const int EN_MIN = 0;
  public const int EN_CREATE_PLAYER_SAME_NAME = 1;
  public const int EN_CREATE_PLAYER_ILLEGAL_NAME = 2;
  public const int EN_CREATE_PLAYER_SECOND_TIME = 3;
  public const int EN_MAX = 4;
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
