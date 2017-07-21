public class CPropertyType{
  static string[] _NAMES = new string[]{
    "CPT_MIN",
    "CPT_HP",
    "CPT_ATK",
    "CPT_DEF",
    "CPT_MAGIC_ATK",
    "CPT_MAGIC_DEF",
    "CPT_AGILE",
    "CPT_KILL",
    "CPT_CRIT",
    "CPT_COUNTER_ATTACK",
    "CPT_SPUTTERING",
    "CPT_DOUBLE_HIT",
    "CPT_RECOVERY",
    "CPT_REFLEX",
    "CPT_SUCK_BLOOD",
    "CPT_INCANTER",
    "CPT_RESISTANCE",
    "CPT_MAX"
  };
  public const int CPT_MIN = 0;
  public const int CPT_HP = 1;
  public const int CPT_ATK = 2;
  public const int CPT_DEF = 3;
  public const int CPT_MAGIC_ATK = 4;
  public const int CPT_MAGIC_DEF = 5;
  public const int CPT_AGILE = 6;
  public const int CPT_KILL = 7;
  public const int CPT_CRIT = 8;
  public const int CPT_COUNTER_ATTACK = 9;
  public const int CPT_SPUTTERING = 10;
  public const int CPT_DOUBLE_HIT = 11;
  public const int CPT_RECOVERY = 12;
  public const int CPT_REFLEX = 13;
  public const int CPT_SUCK_BLOOD = 14;
  public const int CPT_INCANTER = 15;
  public const int CPT_RESISTANCE = 16;
  public const int CPT_MAX = 17;
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
