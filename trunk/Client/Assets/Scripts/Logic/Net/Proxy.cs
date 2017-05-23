
class Proxy : protocol.COM_ServerToClient.Proxy
{
    public bool ErrorMessage(int err, string msg)
    {
        return true;
    }

    public bool LoginSuccess(protocol.COM_AccountInfo info)
    {
        return true;
    }


    public bool CreatePlayerSuccess(protocol.COM_PlayerInstance player)
    {
         return true;
    }
    
}