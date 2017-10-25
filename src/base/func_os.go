package base

import (
	//"errors"
	"os"
	"github.com/skycoin/skycoin/src/cipher/encoder"
)

var openFiles map[string]*os.File = make(map[string]*os.File, 0)

func os_Create(fileName *CXArgument) error {
	if err := checkType("os.Create", "str", fileName); err == nil {
		var name string
		encoder.DeserializeRaw(*fileName.Value, &name)
		
		if file, err := os.Create(name); err == nil {
			openFiles[name] = file
		} else {
			return err
		}

		return nil
	} else {
		return err
	}
}

func os_Open(fileName *CXArgument) error {
	if err := checkType("os.Open", "str", fileName); err == nil {
		var name string
		encoder.DeserializeRaw(*fileName.Value, &name)
		
		if file, err := os.Open(name); err == nil {
			openFiles[name] = file
		} else {
			return err
		}

		return nil
	} else {
		return err
	}
}

func os_Close (fileName *CXArgument) error {
	if err := checkType("os.Open", "str", fileName); err == nil {
		var name string
		encoder.DeserializeRaw(*fileName.Value, &name)
		
		if file, ok := openFiles[name]; ok {
			if err := file.Close(); err != nil {
				return err
			}
		}
		return nil
	} else {
		return err
	}
}
