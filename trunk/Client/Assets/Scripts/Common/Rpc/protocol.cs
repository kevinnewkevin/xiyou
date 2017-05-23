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
  public long InstanceId = 0;
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
  public long InstanceId = 0;
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
  public long InstanceId = 0;
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
namespace COM_ClientToServer{
  class PID{
    public const ushort kMin = 0;
    public const ushort kLogin = 1;
    public const ushort kCreatePlayer = 2;
    public const ushort kMax = 3;
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
  } // end abstract class Stub
  public interface Proxy{
    bool Login(COM_LoginInfo info);
    bool CreatePlayer(int template_id,string player_name);
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
    public const ushort kMax = 4;
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
  } // end abstract class Stub
  public interface Proxy{
    bool ErrorMessage(int err,string msg);
    bool LoginSuccess(COM_AccountInfo info);
    bool CreatePlayerSuccess(COM_PlayerInstance player);
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
        default:{
          return false;
        }
      }
    }
  } //end class Dispatch
} // end namespace COM_ServerToClient
} //end namespace protocol
