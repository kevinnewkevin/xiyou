
public class Mask
{
    /** Constructor. */
    public Mask(int length){
        Bytes = new byte[length];
    }

    /** Write to next bit. */
    public void WriteBit(bool b)
    {
        if (b)
            Bytes[pos_ >> 3] |= (byte)(128 >> (int)(pos_ & 0X7));
        pos_++;
    }

    /** Read from next bit. */
    public bool ReadBit()
    {
        bool r = ((Bytes[pos_ >> 3] & (byte)(128 >> (int)(pos_ & 0X7))) != 0) ? true : false;
        pos_++;
        return r;
    }

    /** Get internal mask bits. */
   
    public byte[] Bytes = null;
    
    private uint pos_ = 0;
}

public interface IWriter
{
    void Write(byte value);
    void Write(short value);
    void Write(int value);
    void Write(long value);
    void Write(sbyte value);
    void Write(ushort value);
    void Write(uint value);
    void Write(ulong value);
    void Write(float value);
    void Write(double value);
    void Write(string value);
    void Write(byte[] value);
    void WriteSize(int value);
}

public interface IReader
{
    bool Read(out byte value);
    bool Read(out short value);
    bool Read(out int value);
    bool Read(out long value);
    bool Read(out sbyte value);
    bool Read(out ushort value);
    bool Read(out uint value);
    bool Read(out ulong value);
    bool Read(out double value);
    bool Read(out float value);
    bool Read(out string value);
    bool Read(out byte[] value);
    bool ReadSize(out int value);
}


public class Bufferd : IWriter, IReader
{
    private const int kDefaultSize = 1024;
    private int size_ = 0, write_ptr_ = 0, read_ptr_ = 0;
    private byte[] buffer_ = null;
    public Bufferd(int size)
    {
        Size = size;
    }

    public byte[] Buffer
    {
        get
        {
            return buffer_;
        }
    }
    public int GetWritePtr()
    {
        return write_ptr_;
    }
    public void SetWritePtr(int s)
    {
        write_ptr_ += s;
    }
    public int GetReadPtr()
    {
        return read_ptr_;
    }
    public void SetReadPtr(int s)
    {
        read_ptr_ += s;
    }

    public int Length
    {
        get
        {
            return write_ptr_ - read_ptr_;
        }
    }

    public int Space
    {
        get
        {
            return size_ - write_ptr_;
        }
    }

    public int Size
    {
        get
        {
            return size_;
        }
        set
        {
            if (size_ < kDefaultSize)
            {
                size_ = kDefaultSize;
            }

            if (size_ != value)
            {
                size_ = value;
                reset_buffer();
            }
        }
    }

    public void Crunch()
    {
        int length = Length;
        if (length != 0)
        {
            System.Array.Copy(buffer_, read_ptr_, buffer_, 0, length);
        }
        read_ptr_ = 0;
        write_ptr_ = length;

    }

    private void reset_buffer()
    {
        if (buffer_ == null)
        {
            buffer_ = new byte[size_];
        }
        else
        {
            int length = Length;
            byte[] t_buffer = new byte[size_];
            System.Array.Copy(buffer_, read_ptr_, t_buffer, 0, length);
            read_ptr_ = 0;
            write_ptr_ = length;
        }
    }

    ///
    public bool Write(byte value)
    {
        if (Space < 1)
        {
            return false;
        }
        buffer_[write_ptr_++] = value;
        return true;
    }
    public bool Write(short value)
    {
        if (Space < 2)
        {
            return false;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 2);
        write_ptr_ += 2;
        return true;
    }
    public bool Write(int value)
    {
        if (Space < 4)
        {
            return false;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 4);
        write_ptr_ += 4;
        return true;
    }
    public bool Write(long value)
    {
        if (Space < 8)
        {
            return false;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 8);
        write_ptr_ += 8;
        return true;
    }
    public bool Write(sbyte value)
    {
        if (Space < 1)
        {
            return false;
        }
        buffer_[write_ptr_++] = (byte)value;
        return true;
    }
    public bool Write(ushort value)
    {
        if (Space < 2)
        {
            return false;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 2);
        write_ptr_ += 2;
        return true;
    }
    public bool Write(uint value)
    {
        if (Space < 4)
        {
            return false;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 4);
        write_ptr_ += 4;
        return true;
    }
    public bool Write(ulong value)
    {
        if (Space < 8)
        {
            return false;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 8);
        write_ptr_ += 8;
        return true;
    }
    public bool Write(float value)
    {
        if (Space < 4)
        {
            return false;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 4);
        write_ptr_ += 4;
        return true;
    }
    public bool Write(double value)
    {
        if (Space < 8)
        {
            return false;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 8);
        write_ptr_ += 8;
        return true;
    }
    public bool Write(string value)
    {
        byte[] t_bytes = System.Text.Encoding.UTF8.GetBytes(value);

        if (Space < t_bytes.Length + 2)
        {
            return false;
        }
        bool check = Write((ushort)t_bytes.Length);
        if (!check)
        {
            return check;
        }
        if (t_bytes.Length > 0)
        {
            System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, t_bytes.Length);
            write_ptr_ += t_bytes.Length;
        }
        return true;
    }
    ///

    public bool Read(out byte value)
    {
        if (Length < 1)
        {
            value = 0;
            return false;
        }
        value = buffer_[read_ptr_++];
        return true;
    }
    public bool Read(out short value)
    {
        if (Length < 2)
        {
            value = 0;
            return false;
        }
        value = System.BitConverter.ToInt16(buffer_, read_ptr_);
        read_ptr_ += 2;
        return true;
    }
    public bool Read(out int value)
    {
        if (Length < 4)
        {
            value = 0;
            return false;
        }
        value = System.BitConverter.ToInt32(buffer_, read_ptr_);
        read_ptr_ += 4;
        return true;
    }
    public bool Read(out long value)
    {
        if (Length < 8)
        {
            value = 0;
            return false;
        }
        value = System.BitConverter.ToInt64(buffer_, read_ptr_);
        read_ptr_ += 8;
        return true;
    }
    public bool Read(out sbyte value)
    {
        if (Length < 1)
        {
            value = 0;
            return false;
        }
        value = (sbyte)buffer_[read_ptr_++];
        return true;
    }
    public bool Read(out ushort value)
    {
        if (Length < 2)
        {
            value = 0;
            return false;
        }
        value = System.BitConverter.ToUInt16(buffer_, read_ptr_);
        read_ptr_ += 2;
        return true;
    }
    public bool Read(out uint value)
    {
        if (Length < 4)
        {
            value = 0;
            return false;
        }
        value = System.BitConverter.ToUInt32(buffer_, read_ptr_);
        read_ptr_ += 4;
        return true;
    }
    public bool Read(out ulong value)
    {
        if (Length < 8)
        {
            value = 0;
            return false;
        }
        value = System.BitConverter.ToUInt64(buffer_, read_ptr_);
        read_ptr_ += 8;
        return true;
    }
    public bool Read(out double value)
    {
        if (Length < 8)
        {
            value = 0.0;
            return false;
        }
        value = System.BitConverter.ToDouble(buffer_, read_ptr_);
        read_ptr_ += 8;
        return true;
    }
    public bool Read(out float value)
    {
        if (Length < 8)
        {
            value = 0.0F;
            return false;
        }
        value = System.BitConverter.ToSingle(buffer_, read_ptr_);
        read_ptr_ += 4;
        return true;
    }
    public bool Read(out string value)
    {
        if (Length < 2)
        {
            value = "";
            return false;
        }
        ushort len = 0;
        bool check = Read(out len);
        if (!check)
        {
            value = "";
            return false;
        }
        value = System.Text.Encoding.UTF8.GetString(buffer_, read_ptr_, (int)len);
        read_ptr_ += (int)len;
        return true;
    }
}
