#ifndef __PRPC_H__
#define __PRPC_H__
#include "config.h"

template<size_t LEN>
class PRPCMask
{
public:
	enum{
		LENGTH = LEN
	};
public:
	PRPCMask()
		:pos_(0)
	{
		for (int i = 0; i < LENGTH; ++i)
			masks_[i] = 0;
	}

	void WriteBit(bool b){
		if (b)
			masks_[pos_ >> 3] |= (128 >> (pos_ & 0X7));
		pos_++;
	}

	bool ReadBit(){
		size_t p = pos_++;
		return !!(masks_[p >> 3] & (128 >> (p & 0X7)));
	}

	unsigned char* Bytes(){
		return masks_;
	}

private:
	uint8_t	masks_[LENGTH];
	size_t	pos_;
};

class PRPCBuffer{
public:
	virtual ~PRPCBuffer(){

	}
	template<class T>
	void Write(T value){
		WriteBytes((uint8_t*)&value, sizeof(T));
	}

	void Write(const std::string & value){
		WriteSize(value.size());
		if (!value.empty())
			WriteBytes((uint8_t*)&value[0], value.size());
	}

	void WriteSize(size_t size){
		uint8_t* p = (uint8_t*)(&size);
		uint8_t n = 0;
		if (size <= 0X3F)
			n = 0;
		else if (size <= 0X3FFF)
			n = 1;
		else if (size <= 0X3FFFFF)
			n = 2;
		else if (size <= 0X3FFFFFFF)
			n = 3;
		p[n] |= (n << 6);
		for (int32_t i = (int32_t)n; i >= 0; --i)
			Write(p[i]);
	}

	template<class T>
	bool Read(T &val){
		return ReadBytes((uint8_t*)&val, sizeof(T));
	}

	bool Read(std::string &val){
		size_t s = 0;
		if (!ReadSize(s)){
			return false;
		}
		if (s != 0){
			val.resize(s);
			if (!ReadBytes((uint8_t*)&val[0], val.size())){
				return false;
			}
		}
		return true;
	}

	
	bool ReadSize(size_t &size){
		size = 0;
		uint8_t b;
		if (!Read(b))
			return false;
		int32_t n = (b & 0XC0) >> 6;
		size = (b & 0X3F);
		for (int32_t i = 0; i < n; ++i){
			if (!Read(b))
				return false;
			size = (size << 8) | b;
		}
		return true;
	}

	virtual bool ReadBytes(uint8_t* b, size_t size) = 0;
	virtual void WriteBytes(const uint8_t* b, size_t size) = 0;
};

template< size_t SZ >
class Buffered : public PRPCBuffer{
public:
	enum{
		SIZE = SZ,
	};
public:
	Buffered()
		: readPointer_(0)
		, writePointer_(0){
	}

	Buffered(const uint8_t* buffer, size_t size)
		: readPointer_(0)
		, writePointer_(size){
		if (size > SIZE)
			size = SIZE;
		std::copy(buffer, buffer + size, buffer_.begin());
	}

	size_t GetLength(){
		return writePointer_ - readPointer_;
	}

	size_t GetSpace(){
		return SIZE - writePointer_;
	}

	bool IsEmpty(){
		return !GetLength();
	}

	void Crunch(){
		size_t length = GetLength();
		if (length != 0){
			std::copy(buffer_.begin() + readPointer_, buffer_.begin() + writePointer_, buffer_.begin());
		}
		readPointer_ = 0;
		writePointer_ = length;
	}

	void SetWritePtr(size_t sz){
		writePointer_ += sz;
	}

	void SetReadPtr(size_t sz){
		readPointer_ += sz;
	}

	uint8_t *GetWritePtr(){
		return &buffer_[writePointer_];
	}

	uint8_t *GetReadPtr(){
		return &buffer_[readPointer_];
	}

	void WriteBytes(const uint8_t* b, size_t size){
		std::copy(b, b + size, buffer_.begin() + writePointer_);
		writePointer_ += size;
	}

	bool ReadBytes(uint8_t* b, size_t size){
		uint8_t *bb = b;
		std::copy(&buffer_[readPointer_], &buffer_[readPointer_ + size], &b[0]);
		readPointer_ += size;
		return true;
	}

private:
	size_t						readPointer_;
	size_t						writePointer_;
	std::array<uint8_t, SIZE>	buffer_;
};



#endif