
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
    bool Read(ref byte value);
    bool Read(ref short value);
    bool Read(ref int value);
    bool Read(ref long value);
    bool Read(ref sbyte value);
    bool Read(ref ushort value);
    bool Read(ref uint value);
    bool Read(ref ulong value);
    bool Read(ref double value);
    bool Read(ref float value);
    bool Read(ref string value);
    bool Read(ref byte[] value);
    bool ReadSize(ref int value);
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

    public void Write(byte[] bytes)
    {
        System.Array.Copy(bytes, 0, buffer_, write_ptr_, bytes.Length);
        write_ptr_ += bytes.Length;
    }
    ///
    public void Write(byte value)
    {
        if (Space < 1)
        {
            return;
        }
        buffer_[write_ptr_++] = value;

    }
    public void Write(short value)
    {
        if (Space < 2)
        {
            return;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 2);
        write_ptr_ += 2;
    }
    public void Write(int value)
    {
        if (Space < 4)
        {
            return;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 4);
        write_ptr_ += 4;
    }
    public void Write(long value)
    {
        if (Space < 8)
        {
            return;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 8);
        write_ptr_ += 8;
    }
    public void Write(sbyte value)
    {
        if (Space < 1)
        {
            return;
        }
        buffer_[write_ptr_++] = (byte)value;
    }
    public void Write(ushort value)
    {
        if (Space < 2)
        {
            return;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 2);
        write_ptr_ += 2;
    }
    public void Write(uint value)
    {
        if (Space < 4)
        {
            return;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 4);
        write_ptr_ += 4;
    }
    public void Write(ulong value)
    {
        if (Space < 8)
        {
            return;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 8);
        write_ptr_ += 8;
    }
    public void Write(float value)
    {
        if (Space < 4)
        {
            return;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 4);
        write_ptr_ += 4;
    }
    public void Write(double value)
    {
        if (Space < 8)
        {
            return;
        }
        byte[] t_bytes = System.BitConverter.GetBytes(value);
        System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, 8);
        write_ptr_ += 8;
    }
    public void Write(string value)
    {
        byte[] t_bytes = System.Text.Encoding.UTF8.GetBytes(value);

        if (Space < t_bytes.Length + 2)
        {
            return;
        }
        WriteSize(t_bytes.Length);
        if (t_bytes.Length > 0)
        {
            System.Array.Copy(t_bytes, 0, buffer_, write_ptr_, t_bytes.Length);
            write_ptr_ += t_bytes.Length;
        }
    }
    public void WriteSize(int s)
    {
        byte[] b = System.BitConverter.GetBytes(s);
        int n = 0;
        if (s <= 0X3F)
            n = 0;
        else if (s <= 0X3FFF)
            n = 1;
        else if (s <= 0X3FFFFF)
            n = 2;
        else if (s <= 0X3FFFFFFF)
            n = 3;
        b[n] |= (byte)(n << 6);
        for (int i = n; i >= 0; i--)
            Write(b[i]);
    }
    ///
    public bool Read(ref byte[] value)
    {
        if (Length < value.Length)
        {
            return false;
        }
        System.Array.Copy(buffer_, read_ptr_, value, 0, value.Length);
        read_ptr_ += value.Length;
        return true;
    }

    public bool Read(ref byte value)
    {
        if (Length < 1)
        {
            value = 0;
            return false;
        }
        value = buffer_[read_ptr_++];
        return true;
    }
    public bool Read(ref short value)
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
    public bool Read(ref int value)
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
    public bool Read(ref long value)
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

    public bool Read(ref sbyte value)
    {
        if (Length < 1)
        {
            value = 0;
            return false;
        }
        value = (sbyte)buffer_[read_ptr_++];
        return true;
    }
    public bool Read(ref ushort value)
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
    public bool Read(ref uint value)
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
    public bool Read(ref ulong value)
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
    public bool Read(ref double value)
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
    public bool Read(ref float value)
    {
        if (Length < 4)
        {
            value = 0.0F;
            return false;
        }
        value = System.BitConverter.ToSingle(buffer_, read_ptr_);
        read_ptr_ += 4;
        return true;
    }
    public bool Read(ref string value)
    {
        if (Length < 2)
        {
            value = "";
            return false;
        }
        int len = 0;
        bool check = ReadSize(ref len);
        if (!check)
        {
            value = "";
            return false;
        }
        value = System.Text.Encoding.UTF8.GetString(buffer_, read_ptr_, (int)len);
        read_ptr_ += (int)len;
        return true;
    }

    public bool ReadSize(ref int s)
    {
        s = 0;
        byte b = 0;
        if (!Read(ref b))
            return false;
        int n = (int)((b & 0XC0) >> 6);
        s = (int)(b & 0X3F);
        for (int i = 0; i < n; i++)
        {
            if (!Read(ref b))
                return false;
            s = (s << 8) | b;
        }
        return true;
    }
}
