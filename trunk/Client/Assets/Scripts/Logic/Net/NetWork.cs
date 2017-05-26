
public class NetWoking 
{
    static Stub stub_ = new Stub();
    static Proxy proxy_ = new Proxy();
    static System.Net.Sockets.Socket socket_ = new System.Net.Sockets.Socket(System.Net.Sockets.AddressFamily.InterNetwork, System.Net.Sockets.SocketType.Stream, System.Net.Sockets.ProtocolType.Tcp);
    static io.Bufferd incoming_buffer_ = new io.Bufferd(65536);
    static io.Bufferd outgoing_buffer_ = new io.Bufferd(65536);

    public static Stub S
    {
        get
        {
            return stub_;
        }
    }
    public static bool Open(string host, int port) {
        try
        {
            socket_.Connect(host, port);
            socket_.Blocking = false;
            socket_.NoDelay = true;
            return true;
        }
        catch (System.Exception ex)
        {
            UnityEngine.Debug.LogError(ex.Message);
            return false;
        }
    }

    public static void Close()
    {
        socket_.Close();
        socket_.Shutdown(System.Net.Sockets.SocketShutdown.Both);
    }


    public static io.Bufferd OutgoingBuffer{
        get{
            return outgoing_buffer_;
        }
    }

    public static void DoWrite()
    {
        if (socket_.Poll(0, System.Net.Sockets.SelectMode.SelectWrite))
        {
            try
            {
                if (outgoing_buffer_.Length != 0)
                {
                    int sended = socket_.Send(outgoing_buffer_.Buffer, outgoing_buffer_.GetReadPtr(), outgoing_buffer_.Length, System.Net.Sockets.SocketFlags.None);
                    outgoing_buffer_.SetReadPtr(sended);
                    outgoing_buffer_.Crunch();
                }
            }
            catch (System.Exception ex)
            {
                UnityEngine.Debug.LogError(ex.Message);
            }
        }
    }

    public static void DoRead()
    {
        if (socket_.Poll(0, System.Net.Sockets.SelectMode.SelectRead))
        {
            try
            {
                int recved = socket_.Receive(incoming_buffer_.Buffer, incoming_buffer_.GetWritePtr(), incoming_buffer_.Space, System.Net.Sockets.SocketFlags.None);
                incoming_buffer_.SetWritePtr(recved);
            }
            catch (System.Exception ex)
            {
                UnityEngine.Debug.LogError(ex.Message);
            }
        }
    }

    public static void DoDispatch()
    {
        if (incoming_buffer_.Length >= 2)
        {
            protocol.COM_ServerToClient.Dispatch.Execute(incoming_buffer_, proxy_);
            incoming_buffer_.Crunch();
        }
    }

    public static void SetupNetFPS()
    {
        DoWrite();
        DoRead();
        DoDispatch();
    }

}
