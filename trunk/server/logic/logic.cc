#include "logic.h"
#include "session.h"
#include "battle.h"
#include "skill_table.h"
#include "unit_table.h"
#include "script.h"
class Accepter : public boost::enable_shared_from_this < Accepter > {
public:
	Accepter(boost::asio::io_service &ioService, boost::asio::ip::tcp::endpoint &endpoint);
	void AcceptHander(boost::shared_ptr<Session> newSession, const boost::system::error_code &errorCode);
	void Svc();
	void Open();
private:
	boost::asio::io_service &ioService_;
	boost::asio::ip::tcp::acceptor acceptor_;
};


Accepter::Accepter(boost::asio::io_service &ioService, boost::asio::ip::tcp::endpoint &endpoint)
	: ioService_(ioService)
	, acceptor_(ioService_, endpoint){

}

void Accepter::Open(){
	boost::shared_ptr<Session> newSession(new Session(ioService_));
	acceptor_.async_accept(newSession->GetSocket(),
		boost::bind(
		&Accepter::AcceptHander,
		shared_from_this(),
		newSession,
		boost::asio::placeholders::error
		)
		);
}

void Accepter::AcceptHander(boost::shared_ptr<Session> newSession, const boost::system::error_code &errorCode){
	if (errorCode){
		return;
	}

	newSession->Open();

	boost::shared_ptr<Session> nextSession(new Session(ioService_));
	acceptor_.async_accept(nextSession->GetSocket(),
		boost::bind(
		&Accepter::AcceptHander,
		shared_from_this(),
		nextSession,
		boost::asio::placeholders::error
		)
		);
}

void Accepter::Svc(){
	boost::thread thr(
		boost::bind(
		&boost::asio::io_service::run,
		boost::ref(ioService_)
		)
		);
}

boost::shared_ptr<class Script> Logic::script_;

void Logic::Init(){

	script_ = boost::make_shared<Script>();
	script_->Init();
	
	script_->LoadFile("../config/script/main.lua");

	UnitData::Load("../config/tables/entity.csv");
	SkillData::Load("../config/tables/skill.csv");


}
void Logic::Fini(){

}
void Logic::Update(){
	Battle::UpdateBattleList();
}


void LogicUpdate(const boost::system::error_code& e, boost::asio::deadline_timer* t){
	
	Logic::Update();
	t->expires_at(t->expires_at() + boost::posix_time::seconds(1));
	t->async_wait(boost::bind(LogicUpdate, boost::asio::placeholders::error, t));
}


void main(){

	Logic::Init();

	boost::asio::io_service ioService;
	boost::asio::io_service::work woker(ioService);
	boost::asio::ip::tcp::endpoint endpoint(boost::asio::ip::address_v4(), 1099);
	boost::shared_ptr<Accepter>  accepter = boost::make_shared<Accepter>(ioService, endpoint);
	accepter->Open();
	accepter->Svc();	

	boost::asio::deadline_timer t(ioService, boost::posix_time::seconds(1));
	t.async_wait(boost::bind(LogicUpdate, boost::asio::placeholders::error, &t));

	ioService.run();
}