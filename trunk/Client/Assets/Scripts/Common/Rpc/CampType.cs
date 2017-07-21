public class CampType{
  static string[] _NAMES = new string[]{
    "CT_RED",
    "CT_BLUE",
    "CT_MAX"
  };
  public const int CT_RED = 0;
  public const int CT_BLUE = 1;
  public const int CT_MAX = 2;
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
