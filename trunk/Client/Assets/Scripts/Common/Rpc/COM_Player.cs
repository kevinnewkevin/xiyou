public class COM_Player{
  public virtual void Serialize(IWriter w){
    Mask mask = new Mask(1);
    mask.WriteBit(InstId!=0);
    mask.WriteBit(Name!=null&&Name.Length!=0&&Name!="");
    mask.WriteBit(true); // Unit
    mask.WriteBit(Employees!=null&&Employees.Length!=0);
    mask.WriteBit(Chapters!=null&&Chapters.Length!=0);
    w.Write(mask.Bytes);
    //S InstId
    if(InstId!=0){
      w.Write(InstId);
    }
    //S Name
    if(Name!=null&&Name.Length!=0&&Name!=""){
      w.Write(Name);
    }
    //S Unit
    {
      Unit.Serialize(w);
    }
    //S Employees
    if(Employees!=null&&Employees.Length!=0){
      w.WriteSize(Employees.Length);
      for(int i=0; i<Employees.Length; ++i){
        Employees[i].Serialize(w);
      }
    }
    //S Chapters
    if(Chapters!=null&&Chapters.Length!=0){
      w.WriteSize(Chapters.Length);
      for(int i=0; i<Chapters.Length; ++i){
        Chapters[i].Serialize(w);
      }
    }
  }
  public virtual bool Deserialize(IReader r){
    Mask mask = new Mask(1);
    if(!r.Read(ref mask.Bytes)){
      return false;
    }
    //D InstId
    if(mask.ReadBit()){
      if(!r.Read(ref InstId)){
        return false;
      }
    }
    //D Name
    if(mask.ReadBit()){
      if(!r.Read(ref Name)){
        return false;
      }
    }
    //D Unit
    if(mask.ReadBit()){
      if(Unit==null){
        Unit = new COM_Unit();
      }
      if(!Unit.Deserialize(r)){
        return false;
      }
    }
    //D Employees
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      Employees = new COM_Unit[size];
      for(int i=0; i<size; ++i){
        Employees[i] = new COM_Unit();
        if(!Employees[i].Deserialize(r)){
          return false;
        }
      }
    }
    //D Chapters
    if(mask.ReadBit()){
      int size = 0;
      if(!r.ReadSize(ref size) || size > 255){
        return false;
      }
      Chapters = new COM_Chapter[size];
      for(int i=0; i<size; ++i){
        Chapters[i] = new COM_Chapter();
        if(!Chapters[i].Deserialize(r)){
          return false;
        }
      }
    }
    return true;
  }
  public long InstId;
  public string Name;
  public COM_Unit Unit;
  public COM_Unit[] Employees;
  public COM_Chapter[] Chapters;
}
