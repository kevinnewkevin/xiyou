#include "csv.h"

bool CSV::LoadFile( const char* filename)
{
	std::ifstream file;
	file.open(filename);
	if (!file.is_open()){
		return false;
	}
	
	std::string line;
	std::getline(file, line);

	parseHeader(line);

	while (!file.eof()){
		std::getline(file, line);
		parseRecord(line);
	}
	file.close();
	return true;
}

bool CSV::parseHeader(const std::string &line)
{
	std::vector<std::string> lineArr = strings::Split(line, sperator_);
	
	for (size_t i = 0; i < lineArr.size(); ++i){
		columnNames_[strings::Trim(lineArr[i], "\n\r\t\"\' ")] = i;
	}

	return true;
}

bool CSV::parseRecord(const std::string &line)
{	
	std::string line2 = strings::Trim(line, "\n\t\r ");
	std::vector<std::string> record;

	size_t q = 0, ii = 0;
	for (size_t i = 0; i < line2.size(); ++i){
		char v = line2[i];
		if (v == '\'' || v == '\"'){
			q++;
		}
		else if(v == ','){
			if ((q & 1) == 0){
				record.push_back(strings::Trim(line2.substr(ii, i-ii), "\n\r\t\"\' "));
				ii = i + 1;
			}
		}
	}
	record.push_back(strings::Trim(line2.substr(ii), "\n\r\t\"\' "));
	BOOST_ASSERT(record.size() == columnNames_.size());
	records_.push_back(record);
	return true;
}

std::string	CSV::get(uint32_t row, uint32_t col)
{
	BOOST_ASSERT(col < columnNames_.size());
	BOOST_ASSERT(row < records_.size());
	BOOST_ASSERT(col < records_[row].size());

	return (records_[row])[col];
}