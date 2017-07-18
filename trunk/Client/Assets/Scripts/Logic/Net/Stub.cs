public class Stub : COM_ClientToServerStub
{
    public override IWriter MethodBegin()
    {
        return NetWoking.OutgoingBuffer;
    }
    public override void MethodEnd()
    {
        NetWoking.DoWrite();
    }
}
