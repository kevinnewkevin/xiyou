#ifndef __SESSION_H__
#define __SESSION_H__
#include "logic.h"

class Session 
	: public boost::enable_shared_from_this < Session >
	, public COM_ServerToClientStub
	, public COM_ClientToServerProxy
	{
public:
	Session(boost::asio::io_service &service);
	bool IsOpen() const ;
	void Open();
	void Close();
	
	boost::asio::ip::tcp::socket &GetSocket();
public:
#include "COM_ClientToServer.fun"
protected:
	PRPCBuffer* MethodBegin(){
		return &sendBuffer_;
	}
	void MethodEnd();

	void parseFrame(){
		COM_ClientToServerDispatcher::Execute(&recvBuffer_, this);
	}
private:
	void receiveHandler(const boost::system::error_code &errorCode, std::size_t bytesTransferred);
	void sendHandler(const boost::system::error_code &errorCode, std::size_t bytesTransferred);
private:

	bool isOpen_;
	boost::asio::ip::tcp::socket socket_;
	enum{ MSG_MAX = 4096 };
	Buffered<MSG_MAX>  recvBuffer_;
	Buffered<MSG_MAX>  sendBuffer_;

private:
	std::string username_;
	std::string password_;
	boost::shared_ptr<class GamePlayer> player_;
};


#endif

