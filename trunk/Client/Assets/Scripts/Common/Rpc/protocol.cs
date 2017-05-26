//This is rpc generate file.
//Donot change it by self.
//The desc source file is protocol.xml version 1.0.0 package protocol
namespace protocol{
public class Configure{
  const int kMajorVersionNumber = 1;
  const int kMinorVersionNumber = 0;
  const int kRevisionVersionNumber = 0;
  const int kBuildNumber = 0;
  const int kVersionNumber = (kMajorVersionNumber << 24) | (kMinorVersionNumber << 16) | (kRevisionVersionNumber << 8) | kBuildNumber ;
  const string kSignatureCode = "";
} // end class Configure
public class ErrorNo{
  private static readonly string[] Strings = {"EN_MIN","EN_CREATE_PLAYER_SAME_NAME","EN_CREATE_PLAYER_ILLEGAL_NAME","EN_CREATE_PLAYER_SECOND_TIME","EN_MAX"};
  public const int EN_MIN = 0;
  public const int EN_CREATE_PLAYER_SAME_NAME = 1;
  public const int EN_CREATE_PLAYER_ILLEGAL_NAME = 2;
  public const int EN_CREATE_PLAYER_SECOND_TIME = 3;
  public const int EN_MAX = 4;
  public static string ToString(int enumer){
    switch(enumer){
      case EN_MIN:{
        return Strings[EN_MIN];
      }
      case EN_CREATE_PLAYER_SAME_NAME:{
        return Strings[EN_CREATE_PLAYER_SAME_NAME];
      }
      case EN_CREATE_PLAYER_ILLEGAL_NAME:{
        return Strings[EN_CREATE_PLAYER_ILLEGAL_NAME];
      }
      case EN_CREATE_PLAYER_SECOND_TIME:{
        return Strings[EN_CREATE_PLAYER_SECOND_TIME];
      }
      case EN_MAX:{
        return Strings[EN_MAX];
      }
    default:{
      return "";
    }
  }
}
public static int ToEnumer(string str_enumer){
  if(str_enumer == Strings[EN_MIN]){
    return EN_MIN;
  } else 
  if(str_enumer == Strings[EN_CREATE_PLAYER_SAME_NAME]){
    return EN_CREATE_PLAYER_SAME_NAME;
  } else 
  if(str_enumer == Strings[EN_CREATE_PLAYER_ILLEGAL_NAME]){
    return EN_CREATE_PLAYER_ILLEGAL_NAME;
  } else 
  if(str_enumer == Strings[EN_CREATE_PLAYER_SECOND_TIME]){
    return EN_CREATE_PLAYER_SECOND_TIME;
  } else 
  if(str_enumer == Strings[EN_MAX]){
    return EN_MAX;
  } else 
  { return 0;}
}
} //end class ErrorNo
public class IPropertyType{
  private static readonly string[] Strings = {"IPT_MIN","IPT_PHYLE","IPT_TITLE","IPT_EXPERIENCE","IPT_LEVEL","IPT_COPPER","IPT_SILVER","IPT_GOLD","IPT_MAX"};
  public const int IPT_MIN = 0;
  public const int IPT_PHYLE = 1;
  public const int IPT_TITLE = 2;
  public const int IPT_EXPERIENCE = 3;
  public const int IPT_LEVEL = 4;
  public const int IPT_COPPER = 5;
  public const int IPT_SILVER = 6;
  public const int IPT_GOLD = 7;
  public const int IPT_MAX = 8;
  public static string ToString(int enumer){
    switch(enumer){
      case IPT_MIN:{
        return Strings[IPT_MIN];
      }
      case IPT_PHYLE:{
        return Strings[IPT_PHYLE];
      }
      case IPT_TITLE:{
        return Strings[IPT_TITLE];
      }
      case IPT_EXPERIENCE:{
        return Strings[IPT_EXPERIENCE];
      }
      case IPT_LEVEL:{
        return Strings[IPT_LEVEL];
      }
      case IPT_COPPER:{
        return Strings[IPT_COPPER];
      }
      case IPT_SILVER:{
        return Strings[IPT_SILVER];
      }
      case IPT_GOLD:{
        return Strings[IPT_GOLD];
      }
      case IPT_MAX:{
        return Strings[IPT_MAX];
      }
    default:{
      return "";
    }
  }
}
public static int ToEnumer(string str_enumer){
  if(str_enumer == Strings[IPT_MIN]){
    return IPT_MIN;
  } else 
  if(str_enumer == Strings[IPT_PHYLE]){
    return IPT_PHYLE;
  } else 
  if(str_enumer == Strings[IPT_TITLE]){
    return IPT_TITLE;
  } else 
  if(str_enumer == Strings[IPT_EXPERIENCE]){
    return IPT_EXPERIENCE;
  } else 
  if(str_enumer == Strings[IPT_LEVEL]){
    return IPT_LEVEL;
  } else 
  if(str_enumer == Strings[IPT_COPPER]){
    return IPT_COPPER;
  } else 
  if(str_enumer == Strings[IPT_SILVER]){
    return IPT_SILVER;
  } else 
  if(str_enumer == Strings[IPT_GOLD]){
    return IPT_GOLD;
  } else 
  if(str_enumer == Strings[IPT_MAX]){
    return IPT_MAX;
  } else 
  { return 0;}
}
} //end class IPropertyType
public class CPropertyType{
  private static readonly string[] Strings = {"CPT_MIN","CPT_HP","CPT_ATK","CPT_DEF","CPT_MAGIC_ATK","CPT_MAGIC_DEF","CPT_AGILE","CPT_KILL","CPT_CRIT","CPT_COUNTER_ATTACK","CPT_SPUTTERING","CPT_DOUBLE_HIT","CPT_RECOVERY","CPT_REFLEX","CPT_SUCK_BLOOD","CPT_INCANTER","CPT_RESISTANCE","CPT_MAX"};
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
  public static string ToString(int enumer){
    switch(enumer){
      case CPT_MIN:{
        return Strings[CPT_MIN];
      }
      case CPT_HP:{
        return Strings[CPT_HP];
      }
      case CPT_ATK:{
        return Strings[CPT_ATK];
      }
      case CPT_DEF:{
        return Strings[CPT_DEF];
      }
      case CPT_MAGIC_ATK:{
        return Strings[CPT_MAGIC_ATK];
      }
      case CPT_MAGIC_DEF:{
        return Strings[CPT_MAGIC_DEF];
      }
      case CPT_AGILE:{
        return Strings[CPT_AGILE];
      }
      case CPT_KILL:{
        return Strings[CPT_KILL];
      }
      case CPT_CRIT:{
        return Strings[CPT_CRIT];
      }
      case CPT_COUNTER_ATTACK:{
        return Strings[CPT_COUNTER_ATTACK];
      }
      case CPT_SPUTTERING:{
        return Strings[CPT_SPUTTERING];
      }
      case CPT_DOUBLE_HIT:{
        return Strings[CPT_DOUBLE_HIT];
      }
      case CPT_RECOVERY:{
        return Strings[CPT_RECOVERY];
      }
      case CPT_REFLEX:{
        return Strings[CPT_REFLEX];
      }
      case CPT_SUCK_BLOOD:{
        return Strings[CPT_SUCK_BLOOD];
      }
      case CPT_INCANTER:{
        return Strings[CPT_INCANTER];
      }
      case CPT_RESISTANCE:{
        return Strings[CPT_RESISTANCE];
      }
      case CPT_MAX:{
        return Strings[CPT_MAX];
      }
    default:{
      return "";
    }
  }
}
public static int ToEnumer(string str_enumer){
  if(str_enumer == Strings[CPT_MIN]){
    return CPT_MIN;
  } else 
  if(str_enumer == Strings[CPT_HP]){
    return CPT_HP;
  } else 
  if(str_enumer == Strings[CPT_ATK]){
    return CPT_ATK;
  } else 
  if(str_enumer == Strings[CPT_DEF]){
    return CPT_DEF;
  } else 
  if(str_enumer == Strings[CPT_MAGIC_ATK]){
    return CPT_MAGIC_ATK;
  } else 
  if(str_enumer == Strings[CPT_MAGIC_DEF]){
    return CPT_MAGIC_DEF;
  } else 
  if(str_enumer == Strings[CPT_AGILE]){
    return CPT_AGILE;
  } else 
  if(str_enumer == Strings[CPT_KILL]){
    return CPT_KILL;
  } else 
  if(str_enumer == Strings[CPT_CRIT]){
    return CPT_CRIT;
  } else 
  if(str_enumer == Strings[CPT_COUNTER_ATTACK]){
    return CPT_COUNTER_ATTACK;
  } else 
  if(str_enumer == Strings[CPT_SPUTTERING]){
    return CPT_SPUTTERING;
  } else 
  if(str_enumer == Strings[CPT_DOUBLE_HIT]){
    return CPT_DOUBLE_HIT;
  } else 
  if(str_enumer == Strings[CPT_RECOVERY]){
    return CPT_RECOVERY;
  } else 
  if(str_enumer == Strings[CPT_REFLEX]){
    return CPT_REFLEX;
  } else 
  if(str_enumer == Strings[CPT_SUCK_BLOOD]){
    return CPT_SUCK_BLOOD;
  } else 
  if(str_enumer == Strings[CPT_INCANTER]){
    return CPT_INCANTER;
  } else 
  if(str_enumer == Strings[CPT_RESISTANCE]){
    return CPT_RESISTANCE;
  } else 
  if(str_enumer == Strings[CPT_MAX]){
    return CPT_MAX;
  } else 
  { return 0;}
}
} //end class CPropertyType
public class EquipmentSlot{
  private static readonly string[] Strings = {"ES_MIN","ES_HEAD","ES_BODY","ES_FEET","ES_ORNAMENT","ES_WEAPON","ES_MAX"};
  public const int ES_MIN = 0;
  public const int ES_HEAD = 1;
  public const int ES_BODY = 2;
  public const int ES_FEET = 3;
  public const int ES_ORNAMENT = 4;
  public const int ES_WEAPON = 5;
  public const int ES_MAX = 6;
  public static string ToString(int enumer){
    switch(enumer){
      case ES_MIN:{
        return Strings[ES_MIN];
      }
      case ES_HEAD:{
        return Strings[ES_HEAD];
      }
      case ES_BODY:{
        return Strings[ES_BODY];
      }
      case ES_FEET:{
        return Strings[ES_FEET];
      }
      case ES_ORNAMENT:{
        return Strings[ES_ORNAMENT];
      }
      case ES_WEAPON:{
        return Strings[ES_WEAPON];
      }
      case ES_MAX:{
        return Strings[ES_MAX];
      }
    default:{
      return "";
    }
  }
}
public static int ToEnumer(string str_enumer){
  if(str_enumer == Strings[ES_MIN]){
    return ES_MIN;
  } else 
  if(str_enumer == Strings[ES_HEAD]){
    return ES_HEAD;
  } else 
  if(str_enumer == Strings[ES_BODY]){
    return ES_BODY;
  } else 
  if(str_enumer == Strings[ES_FEET]){
    return ES_FEET;
  } else 
  if(str_enumer == Strings[ES_ORNAMENT]){
    return ES_ORNAMENT;
  } else 
  if(str_enumer == Strings[ES_WEAPON]){
    return ES_WEAPON;
  } else 
  if(str_enumer == Strings[ES_MAX]){
    return ES_MAX;
  } else 
  { return 0;}
}
} //end class EquipmentSlot
public class BattleActionType{
  private static readonly string[] Strings = {"BAT_MIN","BAT_CRIT","BAT_SUCK","BAT_RECOVERY","BAT_ADD_STATE","BAT_DEL_STATE","BAT_MAX"};
  public const int BAT_MIN = 0;
  public const int BAT_CRIT = 1;
  public const int BAT_SUCK = 2;
  public const int BAT_RECOVERY = 3;
  public const int BAT_ADD_STATE = 4;
  public const int BAT_DEL_STATE = 5;
  public const int BAT_MAX = 6;
  public static string ToString(int enumer){
    switch(enumer){
      case BAT_MIN:{
        return Strings[BAT_MIN];
      }
      case BAT_CRIT:{
        return Strings[BAT_CRIT];
      }
      case BAT_SUCK:{
        return Strings[BAT_SUCK];
      }
      case BAT_RECOVERY:{
        return Strings[BAT_RECOVERY];
      }
      case BAT_ADD_STATE:{
        return Strings[BAT_ADD_STATE];
      }
      case BAT_DEL_STATE:{
        return Strings[BAT_DEL_STATE];
      }
      case BAT_MAX:{
        return Strings[BAT_MAX];
      }
    default:{
      return "";
    }
  }
}
public static int ToEnumer(string str_enumer){
  if(str_enumer == Strings[BAT_MIN]){
    return BAT_MIN;
  } else 
  if(str_enumer == Strings[BAT_CRIT]){
    return BAT_CRIT;
  } else 
  if(str_enumer == Strings[BAT_SUCK]){
    return BAT_SUCK;
  } else 
  if(str_enumer == Strings[BAT_RECOVERY]){
    return BAT_RECOVERY;
  } else 
  if(str_enumer == Strings[BAT_ADD_STATE]){
    return BAT_ADD_STATE;
  } else 
  if(str_enumer == Strings[BAT_DEL_STATE]){
    return BAT_DEL_STATE;
  } else 
  if(str_enumer == Strings[BAT_MAX]){
    return BAT_MAX;
  } else 
  { return 0;}
}
} //end class BattleActionType
public class COM_LoginInfo{
  public string Username = "";
  public string Password = "";
  public bool Package(io.IWriter writer){
    bool check = true;
    {
      check = writer.Write(Username);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(Password);
      if(!check){
        return check;
      }
    }
    return check;
  }
  public bool Unpackage(io.IReader reader){
    bool check = true;
    {
      check = reader.Read(out Username);
      if(!check){
        return check;
      }
    }
    {
      check = reader.Read(out Password);
      if(!check){
        return check;
      }
    }
    return check;
  }
} //end class COM_LoginInfo
public class COM_AccountInfo{
  public string SessionCode = "";
  public bool Package(io.IWriter writer){
    bool check = true;
    {
      check = writer.Write(SessionCode);
      if(!check){
        return check;
      }
    }
    return check;
  }
  public bool Unpackage(io.IReader reader){
    bool check = true;
    {
      check = reader.Read(out SessionCode);
      if(!check){
        return check;
      }
    }
    return check;
  }
} //end class COM_AccountInfo
public class COM_ItemInstance{
  public int ItemId = 0;
  public ulong InstanceId = 0;
  public bool Package(io.IWriter writer){
    bool check = true;
    {
      check = writer.Write(ItemId);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(InstanceId);
      if(!check){
        return check;
      }
    }
    return check;
  }
  public bool Unpackage(io.IReader reader){
    bool check = true;
    {
      check = reader.Read(out ItemId);
      if(!check){
        return check;
      }
    }
    {
      check = reader.Read(out InstanceId);
      if(!check){
        return check;
      }
    }
    return check;
  }
} //end class COM_ItemInstance
public class COM_EntityInstance{
  public ulong InstanceId = 0;
  public System.Collections.Generic.List<int> IProperty = new System.Collections.Generic.List<int>();
  public System.Collections.Generic.List<float> CProperty = new System.Collections.Generic.List<float>();
  public System.Collections.Generic.List<COM_ItemInstance> Equipments = new System.Collections.Generic.List<COM_ItemInstance>();
  public bool Package(io.IWriter writer){
    bool check = true;
    {
      check = writer.Write(InstanceId);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(IProperty.Count);
      if(!check){
        return check;
      }
      for(int i=0; i<IProperty.Count; ++i){
        check = writer.Write(IProperty[i]);
        if(!check){
          return check;
        }
      }
    }
    {
      check = writer.Write(CProperty.Count);
      if(!check){
        return check;
      }
      for(int i=0; i<CProperty.Count; ++i){
        check = writer.Write(CProperty[i]);
        if(!check){
          return check;
        }
      }
    }
    {
      check = writer.Write(Equipments.Count);
      if(!check){
        return check;
      }
      for(int i=0; i<Equipments.Count; ++i){
        check = Equipments[i].Package(writer);
        if(!check){
          return check;
        }
      }
    }
    return check;
  }
  public bool Unpackage(io.IReader reader){
    bool check = true;
    {
      check = reader.Read(out InstanceId);
      if(!check){
        return check;
      }
    }
    {
      int size = 0;
      check = reader.Read(out size);
      if(!check){
        return check;
      }
      IProperty.Clear();
      for(int i=0; i<size; ++i){
        int __IProperty;
        check = reader.Read(out __IProperty);
        if(!check){
          return check;
        }
        IProperty.Add(__IProperty);
      }
    }
    {
      int size = 0;
      check = reader.Read(out size);
      if(!check){
        return check;
      }
      CProperty.Clear();
      for(int i=0; i<size; ++i){
        float __CProperty;
        check = reader.Read(out __CProperty);
        if(!check){
          return check;
        }
        CProperty.Add(__CProperty);
      }
    }
    {
      int size = 0;
      check = reader.Read(out size);
      if(!check){
        return check;
      }
      Equipments.Clear();
      for(int i=0; i<size; ++i){
        COM_ItemInstance __Equipments = new COM_ItemInstance();
        check = __Equipments.Unpackage(reader);
        if(!check){
          return check;
        }
        Equipments.Add(__Equipments);
      }
    }
    return check;
  }
} //end class COM_EntityInstance
public class COM_PlayerInstance{
  public ulong InstanceId = 0;
  public string PlayerName = "";
  public COM_EntityInstance PlayerEntity = new COM_EntityInstance();
  public System.Collections.Generic.List<COM_EntityInstance> Employees = new System.Collections.Generic.List<COM_EntityInstance>();
  public bool Package(io.IWriter writer){
    bool check = true;
    {
      check = writer.Write(InstanceId);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(PlayerName);
      if(!check){
        return check;
      }
    }
    {
      check = PlayerEntity.Package(writer);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(Employees.Count);
      if(!check){
        return check;
      }
      for(int i=0; i<Employees.Count; ++i){
        check = Employees[i].Package(writer);
        if(!check){
          return check;
        }
      }
    }
    return check;
  }
  public bool Unpackage(io.IReader reader){
    bool check = true;
    {
      check = reader.Read(out InstanceId);
      if(!check){
        return check;
      }
    }
    {
      check = reader.Read(out PlayerName);
      if(!check){
        return check;
      }
    }
    {
      check = PlayerEntity.Unpackage(reader);
      if(!check){
        return check;
      }
    }
    {
      int size = 0;
      check = reader.Read(out size);
      if(!check){
        return check;
      }
      Employees.Clear();
      for(int i=0; i<size; ++i){
        COM_EntityInstance __Employees = new COM_EntityInstance();
        check = __Employees.Unpackage(reader);
        if(!check){
          return check;
        }
        Employees.Add(__Employees);
      }
    }
    return check;
  }
} //end class COM_PlayerInstance
public class COM_BattleTarget{
  public ulong InstanceId = 0;
  public int ActionType = new int();
  public int ActionValue = 0;
  public bool Package(io.IWriter writer){
    bool check = true;
    {
      check = writer.Write(InstanceId);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(ActionType);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(ActionValue);
      if(!check){
        return check;
      }
    }
    return check;
  }
  public bool Unpackage(io.IReader reader){
    bool check = true;
    {
      check = reader.Read(out InstanceId);
      if(!check){
        return check;
      }
    }
    {
      check = reader.Read(out ActionType);
      if(!check){
        return check;
      }
    }
    {
      check = reader.Read(out ActionValue);
      if(!check){
        return check;
      }
    }
    return check;
  }
} //end class COM_BattleTarget
public class COM_BattleUnit{
  public long InstanceId = 0;
  public int DisplayId = 0;
  public string Name = "";
  public int HP = 0;
  public bool Package(io.IWriter writer){
    bool check = true;
    {
      check = writer.Write(InstanceId);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(DisplayId);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(Name);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(HP);
      if(!check){
        return check;
      }
    }
    return check;
  }
  public bool Unpackage(io.IReader reader){
    bool check = true;
    {
      check = reader.Read(out InstanceId);
      if(!check){
        return check;
      }
    }
    {
      check = reader.Read(out DisplayId);
      if(!check){
        return check;
      }
    }
    {
      check = reader.Read(out Name);
      if(!check){
        return check;
      }
    }
    {
      check = reader.Read(out HP);
      if(!check){
        return check;
      }
    }
    return check;
  }
} //end class COM_BattleUnit
public class COM_BattlePosition{
  public ulong InstanceId = 0;
  public sbyte PosotionId = 0;
  public bool Package(io.IWriter writer){
    bool check = true;
    {
      check = writer.Write(InstanceId);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(PosotionId);
      if(!check){
        return check;
      }
    }
    return check;
  }
  public bool Unpackage(io.IReader reader){
    bool check = true;
    {
      check = reader.Read(out InstanceId);
      if(!check){
        return check;
      }
    }
    {
      check = reader.Read(out PosotionId);
      if(!check){
        return check;
      }
    }
    return check;
  }
} //end class COM_BattlePosition
public class COM_BattleAction{
  public ulong InstanceId = 0;
  public int SkillId = 0;
  public System.Collections.Generic.List<COM_BattleTarget> BattleTarget = new System.Collections.Generic.List<COM_BattleTarget>();
  public bool Package(io.IWriter writer){
    bool check = true;
    {
      check = writer.Write(InstanceId);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(SkillId);
      if(!check){
        return check;
      }
    }
    {
      check = writer.Write(BattleTarget.Count);
      if(!check){
        return check;
      }
      for(int i=0; i<BattleTarget.Count; ++i){
        check = BattleTarget[i].Package(writer);
        if(!check){
          return check;
        }
      }
    }
    return check;
  }
  public bool Unpackage(io.IReader reader){
    bool check = true;
    {
      check = reader.Read(out InstanceId);
      if(!check){
        return check;
      }
    }
    {
      check = reader.Read(out SkillId);
      if(!check){
        return check;
      }
    }
    {
      int size = 0;
      check = reader.Read(out size);
      if(!check){
        return check;
      }
      BattleTarget.Clear();
      for(int i=0; i<size; ++i){
        COM_BattleTarget __BattleTarget = new COM_BattleTarget();
        check = __BattleTarget.Unpackage(reader);
        if(!check){
          return check;
        }
        BattleTarget.Add(__BattleTarget);
      }
    }
    return check;
  }
} //end class COM_BattleAction
public class COM_BattleReport{
  public System.Collections.Generic.List<COM_BattleUnit> BattleUnit = new System.Collections.Generic.List<COM_BattleUnit>();
  public System.Collections.Generic.List<COM_BattleAction> BattleAction = new System.Collections.Generic.List<COM_BattleAction>();
  public bool Package(io.IWriter writer){
    bool check = true;
    {
      check = writer.Write(BattleUnit.Count);
      if(!check){
        return check;
      }
      for(int i=0; i<BattleUnit.Count; ++i){
        check = BattleUnit[i].Package(writer);
        if(!check){
          return check;
        }
      }
    }
    {
      check = writer.Write(BattleAction.Count);
      if(!check){
        return check;
      }
      for(int i=0; i<BattleAction.Count; ++i){
        check = BattleAction[i].Package(writer);
        if(!check){
          return check;
        }
      }
    }
    return check;
  }
  public bool Unpackage(io.IReader reader){
    bool check = true;
    {
      int size = 0;
      check = reader.Read(out size);
      if(!check){
        return check;
      }
      BattleUnit.Clear();
      for(int i=0; i<size; ++i){
        COM_BattleUnit __BattleUnit = new COM_BattleUnit();
        check = __BattleUnit.Unpackage(reader);
        if(!check){
          return check;
        }
        BattleUnit.Add(__BattleUnit);
      }
    }
    {
      int size = 0;
      check = reader.Read(out size);
      if(!check){
        return check;
      }
      BattleAction.Clear();
      for(int i=0; i<size; ++i){
        COM_BattleAction __BattleAction = new COM_BattleAction();
        check = __BattleAction.Unpackage(reader);
        if(!check){
          return check;
        }
        BattleAction.Add(__BattleAction);
      }
    }
    return check;
  }
} //end class COM_BattleReport
public class COM_BattleResult{
  public bool Package(io.IWriter writer){
    bool check = true;
    return check;
  }
  public bool Unpackage(io.IReader reader){
    bool check = true;
    return check;
  }
} //end class COM_BattleResult
namespace COM_ClientToServer{
  class PID{
    public const ushort kMin = 0;
    public const ushort kLogin = 1;
    public const ushort kCreatePlayer = 2;
    public const ushort kSetBattleEmployee = 3;
    public const ushort kBattleJoin = 4;
    public const ushort kBattleSetup = 5;
    public const ushort kMax = 6;
  } // end class PID
  namespace Package{
    public class Login{
      public COM_LoginInfo info = new COM_LoginInfo();
      public bool Package(io.IWriter writer){
        bool check = true;
        {
          check = info.Package(writer);
          if(!check){
            return check;
          }
        }
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        {
          check = info.Unpackage(reader);
          if(!check){
            return check;
          }
        }
        return check;
      }
    } //end class Login
    public class CreatePlayer{
      public int template_id = 0;
      public string player_name = "";
      public bool Package(io.IWriter writer){
        bool check = true;
        {
          check = writer.Write(template_id);
          if(!check){
            return check;
          }
        }
        {
          check = writer.Write(player_name);
          if(!check){
            return check;
          }
        }
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        {
          check = reader.Read(out template_id);
          if(!check){
            return check;
          }
        }
        {
          check = reader.Read(out player_name);
          if(!check){
            return check;
          }
        }
        return check;
      }
    } //end class CreatePlayer
    public class SetBattleEmployee{
      public long inst_id = 0;
      public bool Package(io.IWriter writer){
        bool check = true;
        {
          check = writer.Write(inst_id);
          if(!check){
            return check;
          }
        }
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        {
          check = reader.Read(out inst_id);
          if(!check){
            return check;
          }
        }
        return check;
      }
    } //end class SetBattleEmployee
    public class BattleJoin{
      public bool Package(io.IWriter writer){
        bool check = true;
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        return check;
      }
    } //end class BattleJoin
    public class BattleSetup{
      public System.Collections.Generic.List<COM_BattlePosition> positions = new System.Collections.Generic.List<COM_BattlePosition>();
      public bool Package(io.IWriter writer){
        bool check = true;
        {
          check = writer.Write(positions.Count);
          if(!check){
            return check;
          }
          for(int i=0; i<positions.Count; ++i){
            check = positions[i].Package(writer);
            if(!check){
              return check;
            }
          }
        }
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        {
          int size = 0;
          check = reader.Read(out size);
          if(!check){
            return check;
          }
          positions.Clear();
          for(int i=0; i<size; ++i){
            COM_BattlePosition __positions = new COM_BattlePosition();
            check = __positions.Unpackage(reader);
            if(!check){
              return check;
            }
            positions.Add(__positions);
          }
        }
        return check;
      }
    } //end class BattleSetup
  } // end namespace Package
  public abstract class Stub{
    protected abstract io.IWriter PackageBegin();
    protected abstract bool PackageEnd();
    public bool Login(COM_LoginInfo info){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kLogin);
      if(!check){
        return check;
      }
      Package.Login login = new Package.Login();
      login.info = info;
      check = login.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
    public bool CreatePlayer(int template_id,string player_name){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kCreatePlayer);
      if(!check){
        return check;
      }
      Package.CreatePlayer createplayer = new Package.CreatePlayer();
      createplayer.template_id = template_id;
      createplayer.player_name = player_name;
      check = createplayer.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
    public bool SetBattleEmployee(long inst_id){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kSetBattleEmployee);
      if(!check){
        return check;
      }
      Package.SetBattleEmployee setbattleemployee = new Package.SetBattleEmployee();
      setbattleemployee.inst_id = inst_id;
      check = setbattleemployee.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
    public bool BattleJoin(){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kBattleJoin);
      if(!check){
        return check;
      }
      Package.BattleJoin battlejoin = new Package.BattleJoin();
      check = battlejoin.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
    public bool BattleSetup(System.Collections.Generic.List<COM_BattlePosition> positions){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kBattleSetup);
      if(!check){
        return check;
      }
      Package.BattleSetup battlesetup = new Package.BattleSetup();
      battlesetup.positions = positions;
      check = battlesetup.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
  } // end abstract class Stub
  public interface Proxy{
    bool Login(COM_LoginInfo info);
    bool CreatePlayer(int template_id,string player_name);
    bool SetBattleEmployee(long inst_id);
    bool BattleJoin();
    bool BattleSetup(System.Collections.Generic.List<COM_BattlePosition> positions);
  } //end interface Proxy
  public class Dispatch{
    public static bool Execute(io.IReader reader, Proxy proxy){
      ushort p = PID.kMin;
      bool check = reader.Read(out p);
      if(!check){
        return check;
      }
      switch(p){
        case PID.kLogin:{
          Package.Login login = new Package.Login();
          check = login.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.Login(login.info);
        }
        case PID.kCreatePlayer:{
          Package.CreatePlayer createplayer = new Package.CreatePlayer();
          check = createplayer.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.CreatePlayer(createplayer.template_id,createplayer.player_name);
        }
        case PID.kSetBattleEmployee:{
          Package.SetBattleEmployee setbattleemployee = new Package.SetBattleEmployee();
          check = setbattleemployee.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.SetBattleEmployee(setbattleemployee.inst_id);
        }
        case PID.kBattleJoin:{
          Package.BattleJoin battlejoin = new Package.BattleJoin();
          check = battlejoin.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.BattleJoin();
        }
        case PID.kBattleSetup:{
          Package.BattleSetup battlesetup = new Package.BattleSetup();
          check = battlesetup.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.BattleSetup(battlesetup.positions);
        }
        default:{
          return false;
        }
      }
    }
  } //end class Dispatch
} // end namespace COM_ClientToServer
namespace COM_ServerToClient{
  class PID{
    public const ushort kMin = 0;
    public const ushort kErrorMessage = 1;
    public const ushort kLoginSuccess = 2;
    public const ushort kCreatePlayerSuccess = 3;
    public const ushort kSetBattleEmployeeSuccess = 4;
    public const ushort kBattleEnter = 5;
    public const ushort kBattleAddUnit = 6;
    public const ushort kBattleReport = 7;
    public const ushort kBattleExit = 8;
    public const ushort kMax = 9;
  } // end class PID
  namespace Package{
    public class ErrorMessage{
      public int err = new int();
      public string msg = "";
      public bool Package(io.IWriter writer){
        bool check = true;
        {
          check = writer.Write(err);
          if(!check){
            return check;
          }
        }
        {
          check = writer.Write(msg);
          if(!check){
            return check;
          }
        }
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        {
          check = reader.Read(out err);
          if(!check){
            return check;
          }
        }
        {
          check = reader.Read(out msg);
          if(!check){
            return check;
          }
        }
        return check;
      }
    } //end class ErrorMessage
    public class LoginSuccess{
      public COM_AccountInfo info = new COM_AccountInfo();
      public bool Package(io.IWriter writer){
        bool check = true;
        {
          check = info.Package(writer);
          if(!check){
            return check;
          }
        }
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        {
          check = info.Unpackage(reader);
          if(!check){
            return check;
          }
        }
        return check;
      }
    } //end class LoginSuccess
    public class CreatePlayerSuccess{
      public COM_PlayerInstance player = new COM_PlayerInstance();
      public bool Package(io.IWriter writer){
        bool check = true;
        {
          check = player.Package(writer);
          if(!check){
            return check;
          }
        }
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        {
          check = player.Unpackage(reader);
          if(!check){
            return check;
          }
        }
        return check;
      }
    } //end class CreatePlayerSuccess
    public class SetBattleEmployeeSuccess{
      public long inst_id = 0;
      public bool Package(io.IWriter writer){
        bool check = true;
        {
          check = writer.Write(inst_id);
          if(!check){
            return check;
          }
        }
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        {
          check = reader.Read(out inst_id);
          if(!check){
            return check;
          }
        }
        return check;
      }
    } //end class SetBattleEmployeeSuccess
    public class BattleEnter{
      public bool Package(io.IWriter writer){
        bool check = true;
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        return check;
      }
    } //end class BattleEnter
    public class BattleAddUnit{
      public System.Collections.Generic.List<COM_BattleUnit> units = new System.Collections.Generic.List<COM_BattleUnit>();
      public bool Package(io.IWriter writer){
        bool check = true;
        {
          check = writer.Write(units.Count);
          if(!check){
            return check;
          }
          for(int i=0; i<units.Count; ++i){
            check = units[i].Package(writer);
            if(!check){
              return check;
            }
          }
        }
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        {
          int size = 0;
          check = reader.Read(out size);
          if(!check){
            return check;
          }
          units.Clear();
          for(int i=0; i<size; ++i){
            COM_BattleUnit __units = new COM_BattleUnit();
            check = __units.Unpackage(reader);
            if(!check){
              return check;
            }
            units.Add(__units);
          }
        }
        return check;
      }
    } //end class BattleAddUnit
    public class BattleReport{
      public COM_BattleReport report = new COM_BattleReport();
      public bool Package(io.IWriter writer){
        bool check = true;
        {
          check = report.Package(writer);
          if(!check){
            return check;
          }
        }
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        {
          check = report.Unpackage(reader);
          if(!check){
            return check;
          }
        }
        return check;
      }
    } //end class BattleReport
    public class BattleExit{
      public COM_BattleResult result = new COM_BattleResult();
      public bool Package(io.IWriter writer){
        bool check = true;
        {
          check = result.Package(writer);
          if(!check){
            return check;
          }
        }
        return check;
      }
      public bool Unpackage(io.IReader reader){
        bool check = true;
        {
          check = result.Unpackage(reader);
          if(!check){
            return check;
          }
        }
        return check;
      }
    } //end class BattleExit
  } // end namespace Package
  public abstract class Stub{
    protected abstract io.IWriter PackageBegin();
    protected abstract bool PackageEnd();
    public bool ErrorMessage(int err,string msg){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kErrorMessage);
      if(!check){
        return check;
      }
      Package.ErrorMessage errormessage = new Package.ErrorMessage();
      errormessage.err = err;
      errormessage.msg = msg;
      check = errormessage.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
    public bool LoginSuccess(COM_AccountInfo info){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kLoginSuccess);
      if(!check){
        return check;
      }
      Package.LoginSuccess loginsuccess = new Package.LoginSuccess();
      loginsuccess.info = info;
      check = loginsuccess.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
    public bool CreatePlayerSuccess(COM_PlayerInstance player){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kCreatePlayerSuccess);
      if(!check){
        return check;
      }
      Package.CreatePlayerSuccess createplayersuccess = new Package.CreatePlayerSuccess();
      createplayersuccess.player = player;
      check = createplayersuccess.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
    public bool SetBattleEmployeeSuccess(long inst_id){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kSetBattleEmployeeSuccess);
      if(!check){
        return check;
      }
      Package.SetBattleEmployeeSuccess setbattleemployeesuccess = new Package.SetBattleEmployeeSuccess();
      setbattleemployeesuccess.inst_id = inst_id;
      check = setbattleemployeesuccess.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
    public bool BattleEnter(){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kBattleEnter);
      if(!check){
        return check;
      }
      Package.BattleEnter battleenter = new Package.BattleEnter();
      check = battleenter.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
    public bool BattleAddUnit(System.Collections.Generic.List<COM_BattleUnit> units){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kBattleAddUnit);
      if(!check){
        return check;
      }
      Package.BattleAddUnit battleaddunit = new Package.BattleAddUnit();
      battleaddunit.units = units;
      check = battleaddunit.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
    public bool BattleReport(COM_BattleReport report){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kBattleReport);
      if(!check){
        return check;
      }
      Package.BattleReport battlereport = new Package.BattleReport();
      battlereport.report = report;
      check = battlereport.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
    public bool BattleExit(COM_BattleResult result){
      io.IWriter writer= PackageBegin();
      bool check = writer.Write(PID.kBattleExit);
      if(!check){
        return check;
      }
      Package.BattleExit battleexit = new Package.BattleExit();
      battleexit.result = result;
      check = battleexit.Package(writer);
      if(!check){
        return check;
      }
      return PackageEnd();
    }
  } // end abstract class Stub
  public interface Proxy{
    bool ErrorMessage(int err,string msg);
    bool LoginSuccess(COM_AccountInfo info);
    bool CreatePlayerSuccess(COM_PlayerInstance player);
    bool SetBattleEmployeeSuccess(long inst_id);
    bool BattleEnter();
    bool BattleAddUnit(System.Collections.Generic.List<COM_BattleUnit> units);
    bool BattleReport(COM_BattleReport report);
    bool BattleExit(COM_BattleResult result);
  } //end interface Proxy
  public class Dispatch{
    public static bool Execute(io.IReader reader, Proxy proxy){
      ushort p = PID.kMin;
      bool check = reader.Read(out p);
      if(!check){
        return check;
      }
      switch(p){
        case PID.kErrorMessage:{
          Package.ErrorMessage errormessage = new Package.ErrorMessage();
          check = errormessage.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.ErrorMessage(errormessage.err,errormessage.msg);
        }
        case PID.kLoginSuccess:{
          Package.LoginSuccess loginsuccess = new Package.LoginSuccess();
          check = loginsuccess.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.LoginSuccess(loginsuccess.info);
        }
        case PID.kCreatePlayerSuccess:{
          Package.CreatePlayerSuccess createplayersuccess = new Package.CreatePlayerSuccess();
          check = createplayersuccess.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.CreatePlayerSuccess(createplayersuccess.player);
        }
        case PID.kSetBattleEmployeeSuccess:{
          Package.SetBattleEmployeeSuccess setbattleemployeesuccess = new Package.SetBattleEmployeeSuccess();
          check = setbattleemployeesuccess.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.SetBattleEmployeeSuccess(setbattleemployeesuccess.inst_id);
        }
        case PID.kBattleEnter:{
          Package.BattleEnter battleenter = new Package.BattleEnter();
          check = battleenter.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.BattleEnter();
        }
        case PID.kBattleAddUnit:{
          Package.BattleAddUnit battleaddunit = new Package.BattleAddUnit();
          check = battleaddunit.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.BattleAddUnit(battleaddunit.units);
        }
        case PID.kBattleReport:{
          Package.BattleReport battlereport = new Package.BattleReport();
          check = battlereport.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.BattleReport(battlereport.report);
        }
        case PID.kBattleExit:{
          Package.BattleExit battleexit = new Package.BattleExit();
          check = battleexit.Unpackage(reader);
          if(!check){
            return check;
          }
          return proxy.BattleExit(battleexit.result);
        }
        default:{
          return false;
        }
      }
    }
  } //end class Dispatch
} // end namespace COM_ServerToClient
} //end namespace protocol
