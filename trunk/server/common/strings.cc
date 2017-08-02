#include "config.h"
#include "strings.h"

namespace strings{
	
	int32_t IndexOf(const std::string &str, char ch){
		size_t pos = str.find_first_of(ch);
		return (pos == std::string::npos ? -1 : pos);
	}

	int32_t IndexOf(const char* cstr, char ch){
		return IndexOf(std::string(cstr), ch);
	}

	int32_t LastIndexOf(const std::string &str, char ch){
		size_t pos = str.find_last_of(ch);
		return (pos == std::string::npos ? -1 : pos);
	}

	int32_t LastIndexOf(const char* cstr, char ch){
		return LastIndexOf(std::string(cstr), ch);
	}
	
	std::string ToLowerCase(const std::string &str){
		std::string result = str;
		std::transform(result.begin(), result.end(), result.begin(), tolower);
		return result;
	}
	
	std::string ToUpperCase(const std::string &str){
		std::string result = str;
		std::transform(result.begin(), result.end(), result.begin(), toupper);
		return result;
	}
	
	std::string Trim(const std::string &str, const std::string &delims, bool left, bool right){
		std::string result = str;
		if (right){
			result.erase(result.find_last_not_of(delims) + 1);
		}
		if (left){
			result.erase(0, result.find_first_not_of(delims));
		}
		return result;
	}
	
	std::string Replace(const std::string &str, const std::string &form, const std::string &to){
		std::string result = str;
		size_t start = 0;
		int32_t pos;
		do {
			pos = result.find(form, start);
			if (pos == -1){
				break;
			}
			result.replace(result.begin() + pos, result.begin() + pos + form.size(), to.begin(), to.end());
			start = pos + form.size();
			if (start >= result.size()){
				break;
			}

		} while (pos != std::string::npos);
		return result;
	}
	
	std::vector<std::string > Split(const std::string &str, const std::string &delims, uint32_t max_splits){
		std::vector<std::string> result;
		result.reserve(max_splits ? max_splits + 1 : 1000);

		unsigned int num_splits = 0;
		size_t start = 0, pos;
		do
		{
			pos = str.find_first_of(delims, start);
			if (pos == start)
			{
				start = pos + 1;
			}
			else if (pos == std::string::npos || (max_splits && num_splits == max_splits))
			{
				result.push_back(str.substr(start));
				break;
			}
			else
			{
				result.push_back(str.substr(start, pos - start));
				start = pos + 1;
			}
			start = str.find_first_not_of(delims, start);
			++num_splits;

		} while (pos != std::string::npos);
		return result;
	}
	
	std::string Join(const std::vector<std::string> &a, const std::string &sep){
		if (a.empty()){
			return "";
		}
		std::stringstream ss;
		ss << a.front();
		for (size_t i = 1; i < a.size(); ++i){
			ss << sep << a[i];
		}
		return ss.str();
	}
}