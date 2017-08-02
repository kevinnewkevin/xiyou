#ifndef ___STRCONV_H__
#define ___STRCONV_H__

#include "config.h"

namespace strconv{

	//×ª»¯×Ö·û´®
#define __CHUNK(RESULT,VALUE) do {\
	std::stringstream ss;\
	ss << VALUE;\
	ss >> RESULT;\
} while (0);

#define __TO_STRING(BASE_TYPE) inline std::string ToString(BASE_TYPE value){\
		std::string result;\
		__CHUNK(result,value);\
		return result;\
}
	__TO_STRING(int8_t);
	__TO_STRING(int16_t);
	__TO_STRING(int32_t);
	__TO_STRING(int64_t);

	__TO_STRING(uint8_t);
	__TO_STRING(uint16_t);
	__TO_STRING(uint32_t);
	__TO_STRING(uint64_t);

	__TO_STRING(float);
	__TO_STRING(double);

	template<class T>
	inline std::string ToString(const T & value){
		std::string result;
		__CHUNK(result, value);
		return result;
	}
	template<class T> T ToType(const std::string &a){
		T result;
		__CHUNK(result, a);
		return result;
	}

#define __CONVERT_STRING(BASE_TYPE,NAME) inline BASE_TYPE To##NAME(const std::string &value){\
	BASE_TYPE result;\
	__CHUNK(result,value);\
	return result;\
		}
	__CONVERT_STRING(int8_t, Int8);
	__CONVERT_STRING(int16_t, Int16);
	__CONVERT_STRING(int32_t, Int32);
	__CONVERT_STRING(int64_t, Int64);

	__CONVERT_STRING(uint8_t, UInt8);
	__CONVERT_STRING(uint16_t, UInt16);
	__CONVERT_STRING(uint32_t, UInt32);
	__CONVERT_STRING(uint64_t, UInt64);

	__CONVERT_STRING(float, Float);
	__CONVERT_STRING(double, Double);

#undef __CHUNK
#undef __TO_STRING
#undef __CONVERT_STRING

	const char* Itoa(int i);
	const char* Ftoa(float f);
	const char* Dtoa(double d);

	int Atoi(const char* a);
	float Atof(const char* a);
	double Atod(const char* a);


	int Atoi(const std::string &a);
	float Atof(const std::string &a);
	double Atod(const std::string &a);

	
}

#endif //___STRCONV_H__