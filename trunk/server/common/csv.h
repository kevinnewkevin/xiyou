#ifndef	__CSV_H__
#define	__CSV_H__
#include "config.h"
#include "strconv.h"
#include "strings.h"
#include <boost/assert.hpp>
class CSV
{
public:
	CSV(std::string sperator = ",")
	:sperator_(sperator){}
		
	bool LoadFile( const char* filename);

	/*template<class T> T Get(uint32_t row, const char* columnname){
		checkColunmKey(columnname);
		return strconv::ToType<T>(get(row, columnNames_[columnname]));
	}*/

	template<class T> T Get(uint32_t row, const std::string &columnname){
		checkColunmKey(columnname.c_str());
		return strconv::ToType<T>(get(row, columnNames_[columnname]));
	}

	size_t GetRowMax(){ return records_.size(); }
	size_t GetColumnMax(){ return columnNames_.size(); }
protected:

	bool parseHeader(const std::string &line);
	bool parseRecord(const std::string &line);

	void checkColunmKey(char const * columnname)
	{	
		BOOST_ASSERT_MSG(columnNames_.end() != columnNames_.find(columnname), "columnname error");
	}
	
	std::string	get(uint32_t  row, uint32_t  col);
private:
	std::map<std::string, uint32_t >		columnNames_;
	std::vector< std::vector<std::string> >				records_;
	std::string							sperator_;
};

#endif
