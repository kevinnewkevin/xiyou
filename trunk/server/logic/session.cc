#include "logic.h"
#include "player.h"
#include "session.h"

void Session::MethodEnd(){

	if (!isOpen_)
		return;

	socket_.async_write_some(
		boost::asio::buffer(sendBuffer_.GetReadPtr(), sendBuffer_.GetLength()),
		boost::bind(
		&Session::sendHandler,
		shared_from_this(),
		boost::asio::placeholders::error,
		boost::asio::placeholders::bytes_transferred
		)
		);
}

Session::Session(boost::asio::io_service &service)
	:socket_(service){

}

bool Session::IsOpen()const{
	return isOpen_;
}

void Session::Open(){
	isOpen_ = true;
	
	socket_.async_receive(
		boost::asio::buffer(recvBuffer_.GetWritePtr(), recvBuffer_.GetSpace()),
		boost::bind(
			&Session::receiveHandler,
			shared_from_this(),
			boost::asio::placeholders::error,
			boost::asio::placeholders::bytes_transferred
		)
	);
}

void Session::receiveHandler(const boost::system::error_code &errorCode, std::size_t bytesTransferred){
	if (errorCode){
		return;
	}
	if (!isOpen_){
		return;
	}
	
	recvBuffer_.SetWritePtr(bytesTransferred);

	if (recvBuffer_.GetLength() > sizeof(uint16_t)){
		parseFrame();
	}

	socket_.async_receive(
		boost::asio::buffer(recvBuffer_.GetWritePtr(), recvBuffer_.GetSpace()),
		boost::bind(
			&Session::receiveHandler,
			shared_from_this(),
			boost::asio::placeholders::error,
			boost::asio::placeholders::bytes_transferred
		)
	);

}

void Session::Close(){
	if (!isOpen_){
		return;
	}
	isOpen_ = false;
	socket_.close();
}


void Session::sendHandler(const boost::system::error_code &errorCode, std::size_t bytesTransferred){
	if (errorCode){
		return;
	}
	if (!isOpen_)
		return;
	
	sendBuffer_.SetReadPtr(bytesTransferred);
	sendBuffer_.Crunch();

	socket_.async_write_some(
		boost::asio::buffer(sendBuffer_.GetReadPtr(), sendBuffer_.GetLength()),
		boost::bind(
		&Session::sendHandler,
		shared_from_this(),
		boost::asio::placeholders::error,
		boost::asio::placeholders::bytes_transferred
		)
	);
}

boost::asio::ip::tcp::socket &Session::GetSocket() {
	return socket_;
}

//////////////////////////////////////////////////////////////////////////
bool Session::Login(COM_LoginInfo &info){
	username_ = info.Username;
	password_ = info.Password;
	return true;
} // 0
bool Session::CreatePlayer(int32_t tempId, std::string &playerName){
	if (player_ != nullptr)
		return true;
	player_ = GamePlayer::CreatePlayer(shared_from_this(), playerName, tempId);
	COM_Player inst;
	player_->Save(inst);
	CreatePlayerOK(inst);
	return true;
}// 1
bool Session::SetBattleUnit(int64_t instId){
	if (player_ == nullptr)
		return true;
	player_->SetBattleUnit(instId);
	return true;
}// 2
bool Session::JoinBattle(){
	if (player_ == nullptr)
		return true;
	player_->JoinBattle();
	return true;
} // 3
bool Session::SetupBattle(std::vector<COM_BattlePosition> &positionList){
	if (player_ == nullptr)
		return true;
	player_->SetupBattle(positionList);
	return true;
} // 4