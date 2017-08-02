#include "config.h"
#include "strconv.h"

namespace strconv{
	const char* Itoa(int i){
		return ToString(i).c_str();
	}
	const char* Ftoa(float f){
		return ToString(f).c_str();
	}
	const char* Dtoa(double d){
		return ToString(d).c_str();
	}

	int Atoi(const char* a){
		return ToInt32(a);
	}
	float Atof(const char* a){
		return ToFloat(a);
	}
	double Atod(const char* a){
		return ToDouble(a);
	}


	int Atoi(const std::string &a){
		return ToInt32(a);
	}
	float Atof(const std::string &a){
		return ToFloat(a);
	}
	double Atod(const std::string &a){
		return ToDouble(a);
	}
}