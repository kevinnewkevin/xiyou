public class Stub : protocol.COM_ClientToServer.Stub
{
    protected override io.IWriter PackageBegin()
    {
        return NetWoking.OutgoingBuffer;
    }
    protected override bool PackageEnd()
    {
        NetWoking.DoWrite();
        return true;
    }
}
