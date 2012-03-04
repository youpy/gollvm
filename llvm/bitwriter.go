package llvm

/*
#include <llvm-c/BitWriter.h>
#include <stdlib.h>
*/
import "C"
import "os"
import "errors"

var writeBitcodeToFileErr = errors.New("Failed to write bitcode to file")

func WriteBitcodeToFile(m Module, file *os.File) error {
	result := C.LLVMWriteBitcodeToFD(m.C, C.int(file.Fd()), C.int(0), C.int(0))
	if result == 0 {
		return nil
	}
	return writeBitcodeToFileErr
}

// TODO(nsf): Figure out way how to make it work with io.Writer
