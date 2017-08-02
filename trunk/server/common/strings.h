#ifndef ___STRINGS_H__
#define ___STRINGS_H__

#include "config.h"
#include "strconv.h"

namespace strings{

	int32_t IndexOf(const std::string &str, char ch);

	int32_t IndexOf(const char* cstr, char ch);

	int32_t LastIndexOf(const std::string &str, char ch);

	int32_t LastIndexOf(const char* cstr, char ch);

	std::string ToLowerCase(const std::string &str);

	std::string ToUpperCase(const std::string &str);

	std::string Trim(const std::string &str, const std::string &delims, bool left = true, bool right = true);

	std::string Replace(const std::string &str, const std::string &form, const std::string &to);

	std::vector<std::string> Split(const std::string &str, const std::string &delims, uint32_t max_splits = 0);

	template<size_t S>
	std::string Join(const std::array<std::string, S> &a, const std::string &sep = ""){
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

	std::string Join(const std::vector<std::string> &a, const std::string &sep = "");

	template<class T >
	std::string Sprintf(const std::string &a, const T & v1){
		std::string print_string = Replace(a, "%1", strconv::ToString(v1));
		return print_string;
	}
	template<class T1, class T2>
	std::string Sprintf(const std::string &a, const T1 &v1, const T2 &v2){
		std::string print_string = Replace(a, "%1", strconv::ToString(v1));
		print_string = Replace(print_string, "%2", strconv::ToString(v2));
		return print_string;
	}

	template<class T1, class T2, class T3>
	std::string Sprintf(const std::string &a, const T1 &v1, const T2 &v2, const T3 &v3){
		std::string print_string = Replace(a, "%1", strconv::ToString(v1));
		print_string = Replace(print_string, "%2", strconv::ToString(v2));
		print_string = Replace(print_string, "%3", strconv::ToString(v3));
		return print_string;
	}

	template<class T1, class T2, class T3, class T4>
	std::string Sprintf(const std::string &a, const T1 &v1, const T2 &v2, const T3 &v3, const T4 &v4){
		std::string print_string = Replace(a, "%1", strconv::ToString(v1));
		print_string = Replace(print_string, "%2", strconv::ToString(v2));
		print_string = Replace(print_string, "%3", strconv::ToString(v3));
		print_string = Replace(print_string, "%4", strconv::ToString(v4));
		return print_string;
	}

	template<class T1, class T2, class T3, class T4, class T5>
	std::string Sprintf(const std::string &a, const T1 &v1, const T2 &v2, const T3 &v3, const T4 &v4, const T5 &v5){
		std::string print_string = Replace(a, "%1", strconv::ToString(v1));
		print_string = Replace(print_string, "%2", strconv::ToString(v2));
		print_string = Replace(print_string, "%3", strconv::ToString(v3));
		print_string = Replace(print_string, "%4", strconv::ToString(v4));
		print_string = Replace(print_string, "%5", strconv::ToString(v5));
		return print_string;
	}

	template<class T1, class T2, class T3, class T4, class T5, class T6>
	std::string Sprintf(const std::string &a, const T1 &v1, const T2 &v2, const T3 &v3, const T4 &v4, const T5 &v5, const T6 &v6){
		std::string print_string = Replace(a, "%1", strconv::ToString(v1));
		print_string = Replace(print_string, "%2", strconv::ToString(v2));
		print_string = Replace(print_string, "%3", strconv::ToString(v3));
		print_string = Replace(print_string, "%4", strconv::ToString(v4));
		print_string = Replace(print_string, "%5", strconv::ToString(v5));
		print_string = Replace(print_string, "%6", strconv::ToString(v6));
		return print_string;
	}
}

#endif // ___STRINGS_H__