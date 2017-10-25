using System.Net;

public class NetWoking 
{
    static Stub stub_ = new Stub();
    static Proxy proxy_ = new Proxy();
    static System.Net.Sockets.Socket socket_;
    static Bufferd incoming_buffer_ = new Bufferd(65536);
    static Bufferd outgoing_buffer_ = new Bufferd(65536);

    static public int _LastErrorCode;

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
            _LastErrorCode = 0;
            socket_ = new System.Net.Sockets.Socket(System.Net.Sockets.AddressFamily.InterNetwork, System.Net.Sockets.SocketType.Stream, System.Net.Sockets.ProtocolType.Tcp);
			socket_.SetSocketOption(System.Net.Sockets.SocketOptionLevel.Tcp, System.Net.Sockets.SocketOptionName.NoDelay, true);
            socket_.Connect(host, port);
            return true;
        }
        catch (System.Net.Sockets.SocketException ex)
        {
            _LastErrorCode = ex.ErrorCode;
            UnityEngine.Debug.LogError(ex.Message + "1:" + ex.ErrorCode);
            return false;
        }
    }

    public static void Close()
    {
        if (socket_ != null)
        {
            //socket_.Shutdown(System.Net.Sockets.SocketShutdown.Both);
            socket_.Close();
        }
        incoming_buffer_ = new Bufferd(65536);
        outgoing_buffer_ = new Bufferd(65536);
    }


    public static Bufferd OutgoingBuffer{
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
            catch (System.Net.Sockets.SocketException ex)
            {
                _LastErrorCode = ex.ErrorCode;
                UnityEngine.Debug.LogError(ex.Message + "2:" + ex.ErrorCode);
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
            catch (System.Net.Sockets.SocketException ex)
            {
                _LastErrorCode = ex.ErrorCode;
                UnityEngine.Debug.LogError(ex.Message + "3:" + ex.ErrorCode);
            }
        }
    }

    public static void DoDispatch()
    {
        if (incoming_buffer_.Length >= 2)
        {
            COM_ServerToClientDispatcher.Execute(incoming_buffer_, proxy_);
            incoming_buffer_.Crunch();
        }
    }

    public static void SetupNetFPS()
    {
        if (socket_ == null)
            return;
		if (!socket_.Connected)
            return;		
        DoWrite();
        DoRead();
        DoDispatch();
    }

    public static bool ReConnect()
    {
        if (socket_ != null && socket_.RemoteEndPoint != null)
            Close();
        string ipadd = Proxy4Lua._ServerIP;//Define.GetStr("DebugServerAddress");
        int port = Define.GetInt("DebugServerPort");
        return Open(ipadd, port);
    }
}
