package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	BinaryType uint8 = iota + 1
	StringType
)

type Payload interface {
	io.WriterTo
	io.ReaderFrom
	Bytes() []byte
	String() string
}

type Binary []byte

func(m Binary) Bytes() []byte{
	return m
}

func (m Binary) String() string {
	return string(m)
}

func (m Binary) WriteTo(w io.Writer)(int64, error){
	//BigEndian merupakan metode encode
	//BigEndian dilakukan dengan cara memasukkan Byte terbesar ke dalam alamat nya terlebih dahulu
	err := binary.Write(w, binary.BigEndian, BinaryType) // 1 Byte
	if err != nil {
		return 0, err
	}

	err = binary.Write(w, binary.BigEndian, uint32(len(m))) // 4 Byte
	if err != nil{
		return 0, err
	}

	n, err := w.Write(m)
	if err != nil {
		return 0,  err
	}

	return int64(n + 5), err
}

func (m *Binary) ReadFrom(r io.Reader) (int64, error){
	var typ uint8 
	err := binary.Read(r, binary.BigEndian, &typ) //1 Byte

	if err != nil {
		return 0, err
	}

	if typ != BinaryType {
		return 1, nil
	}
	
	var n uint32 = 1
	err = binary.Read(r, binary.BigEndian, &n) // 4 Byte
	if err != nil {
		return 0, err
	}

	*m = make([]byte, n)
	o, err := r.Read(*m)

	return int64(o+5), err
}


type String string

func(m String) Bytes() []byte{
	return []byte(m)
}

func (m String) String() string {
	return string(m)
}

func (m String) WriteTo(w io.Writer)(int64, error){
	//BigEndian merupakan metode encode
	//BigEndian dilakukan dengan cara memasukkan Byte terbesar ke dalam alamat nya terlebih dahulu
	err := binary.Write(w, binary.BigEndian, StringType) // 1 Byte
	if err != nil {
		return 0, err
	}

	data := []byte(m)
	err = binary.Write(w, binary.BigEndian, uint32(len(m))) // 4 Byte
	if err != nil{
		return 0, err
	}

	n, err := w.Write(data)

	return int64(n + 5), err
}

func (m *String) ReadFrom(r io.Reader) (int64, error){
	var typ uint8 
	err := binary.Read(r, binary.BigEndian, &typ) //1 Byte

	if err != nil {
		return 0, err
	}

	if typ != StringType {
		return 1, nil
	}
	
	var n uint32
	err = binary.Read(r, binary.BigEndian, &n) // 4 Byte
	if err != nil {
		return 0, err
	}

	buff := make([]byte, n)

	o, err := r.Read(buff)
	*m = String(buff)

	return int64(o+5), err
}

func Decode(r io.Reader) (Payload, error){
	var typ uint8  
	err := binary.Read(r, binary.BigEndian, &typ)

	if err != nil {
		return nil, err
	}

	var payload Payload 
	switch typ {
	case BinaryType:
		fmt.Println("Binary Type")
		payload = new(Binary)
	case StringType:
		fmt.Println("String Type")
		payload = new(String)
	}

	_, err = payload.ReadFrom(io.MultiReader(bytes.NewReader([]byte{typ}), r))
	if err != nil {
		return nil, err
	}

	return payload, nil
}