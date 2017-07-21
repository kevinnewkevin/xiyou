public class EquipmentSlot{
  static string[] _NAMES = new string[]{
    "ES_MIN",
    "ES_HEAD",
    "ES_BODY",
    "ES_FEET",
    "ES_ORNAMENT",
    "ES_WEAPON",
    "ES_MAX"
  };
  public const int ES_MIN = 0;
  public const int ES_HEAD = 1;
  public const int ES_BODY = 2;
  public const int ES_FEET = 3;
  public const int ES_ORNAMENT = 4;
  public const int ES_WEAPON = 5;
  public const int ES_MAX = 6;
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
