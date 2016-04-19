// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript

import (
	"errors"
	"fmt"
)

var (
	// ErrStackShortScript is returned if the script has an opcode that is
	// too long for the length of the script.
	ErrStackShortScript = errors.New("execute past end of script")

	// ErrStackLongScript is returned if the script has an opcode that is
	// too long for the length of the script.
	ErrStackLongScript = errors.New("script is longer than maximum allowed")

	// ErrStackUnderflow is returned if an opcode requires more items on the
	// stack than is present.f
	ErrStackUnderflow = errors.New("stack underflow")

	// ErrStackInvalidArgs is returned if the argument for an opcode is out
	// of acceptable range.
	ErrStackInvalidArgs = errors.New("invalid argument")

	// ErrStackOpDisabled is returned when a disabled opcode is encountered
	// in the script.
	ErrStackOpDisabled = errors.New("disabled opcode")

	// ErrStackVerifyFailed is returned when one of the OP_VERIFY or
	// OP_*VERIFY instructions is executed and the conditions fails.
	ErrStackVerifyFailed = errors.New("verify failed")

	// ErrStackNumberTooBig is returned when the argument for an opcode that
	// should be an offset is obviously far too large.
	ErrStackNumberTooBig = errors.New("number too big")

	// ErrStackInvalidOpcode is returned when an opcode marked as invalid or
	// a completely undefined opcode is encountered.
	ErrStackInvalidOpcode = errors.New("invalid opcode")

	// ErrStackReservedOpcode is returned when an opcode marked as reserved
	// is encountered.
	ErrStackReservedOpcode = errors.New("reserved opcode")

	// ErrStackEarlyReturn is returned when OP_RETURN is executed in the
	// script.
	ErrStackEarlyReturn = errors.New("script returned early")

	// ErrStackNoIf is returned if an OP_ELSE or OP_ENDIF is encountered
	// without first having an OP_IF or OP_NOTIF in the script.
	ErrStackNoIf = errors.New("OP_ELSE or OP_ENDIF with no matching OP_IF")

	// ErrStackNoWhile is returned if an OP_ENDWHILE is encountered
	// without first having an OP_WHILE in the script.
	ErrStackNoWhile = errors.New("OP_ENDWHILE with no matching OP_WHILE")

	// ErrStackMissingEnd is returned if the end of a script is reached
	// without an OP_ENDIF or OP_ENDWHILE to correspond to a conditional
	// expression.
	ErrStackMissingEnd = fmt.Errorf("conditional opened but not closed")

	// ErrStackTooManyPubKeys is returned if an OP_CHECKMULTISIG is
	// encountered with more than MaxPubKeysPerMultiSig pubkeys present.
	ErrStackTooManyPubKeys = errors.New("invalid pubkey count in OP_CHECKMULTISIG")

	// ErrStackTooManyOperations is returned if a script has more than
	// MaxOpsPerScript opcodes that do not push data.
	ErrStackTooManyOperations = errors.New("too many operations in script")

	// ErrStackElementTooBig is returned if the size of an element to be
	// pushed to the stack is over MaxScriptElementSize.
	ErrStackElementTooBig = errors.New("element in script too large")

	// ErrStackUnknownAddress is returned when ScriptToAddrHash does not
	// recognise the pattern of the script and thus can not find the address
	// for payment.
	ErrStackUnknownAddress = errors.New("non-recognised address")

	// ErrStackScriptFailed is returned when at the end of a script the
	// boolean on top of the stack is false signifying that the script has
	// failed.
	ErrStackScriptFailed = errors.New("execute fail, fail on stack")

	// ErrStackScriptUnfinished is returned when CheckErrorCondition is
	// called on a script that has not finished executing.
	ErrStackScriptUnfinished = errors.New("error check when script unfinished")

	// ErrStackEmptyStack is returned when the stack is empty at the end of
	// execution. Normal operation requires that a boolean is on top of the
	// stack when the scripts have finished executing.
	ErrStackEmptyStack = errors.New("stack empty at end of execution")

	// ErrStackP2SHNonPushOnly is returned when a Pay-to-Script-Hash
	// transaction is encountered and the ScriptSig does operations other
	// than push data (in violation of bip16).
	ErrStackP2SHNonPushOnly = errors.New("pay to script hash with non " +
		"pushonly input")

	// ErrStackInvalidParseType is an internal error returned from
	// ScriptToAddrHash ony if the internal data tables are wrong.
	ErrStackInvalidParseType = errors.New("internal error: invalid parsetype found")

	// ErrStackInvalidAddrOffset is an internal error returned from
	// ScriptToAddrHash ony if the internal data tables are wrong.
	ErrStackInvalidAddrOffset = errors.New("internal error: invalid offset found")

	// ErrStackInvalidIndex is returned when an out-of-bounds index was
	// passed to a function.
	ErrStackInvalidIndex = errors.New("invalid script index")

	// ErrStackNonPushOnly is returned when ScriptInfo is called with a
	// pkScript that peforms operations other that pushing data to the stack.
	ErrStackNonPushOnly = errors.New("SigScript is non pushonly")

	// ErrStackOverflow is returned when stack and altstack combined depth
	// is over the limit.
	ErrStackOverflow = errors.New("stack overflow")

	// ErrStackInvalidLowSSignature is returned when the ScriptVerifyLowS
	// flag is set and the script contains any signatures whose S values
	// are higher than the half order.
	ErrStackInvalidLowSSignature = errors.New("invalid low s signature")

	// ErrStackInvalidPubKey is returned when the ScriptVerifyScriptEncoding
	// flag is set and the script contains invalid pubkeys.
	ErrStackInvalidPubKey = errors.New("invalid strict pubkey")

	// ErrStackCleanStack is returned when the ScriptVerifyCleanStack flag
	// is set and after evalution the stack does not contain only one element,
	// which also must be true if interpreted as a boolean.
	ErrStackCleanStack = errors.New("stack is not clean")

	// ErrStackMinimalData is returned when the ScriptVerifyMinimalData flag
	// is set and the script contains push operations that do not use
	// the minimal opcode required.
	ErrStackMinimalData = errors.New("non-minimally encoded script number")

	// ErrContractHashLength is returned when trying to create a
	// contract hash from the wrong number of bytes.
	ErrContractHashLength = errors.New("wrong number of bytes for contract hash")

	// ErrAssetIDLength is returned when trying to create an asset
	// id from the wrong number of bytes.
	ErrAssetIDLength = errors.New("wrong number of bytes for asset id")

	// ErrEarlyTimestamp is returned when OP_CHECKLOCKTIMEVERIFY is
	// given a timestamp below 5e8.
	ErrEarlyTimestamp = errors.New("early CHECKLOCKTIMEVERIFY timestamp")

	// ErrDivideByZero is returned when OP_DIV or OP_MOD tries to divide
	// by zero.
	ErrDivideByZero = errors.New("divide by zero")

	// ErrMulOverflow is returned when OP_MUL or OP_2MUL might overflow.
	ErrMulOverflow = errors.New("overflow in MUL")

	// ErrNegativeShift is when you try to shift a number by a negative
	// number of bits.
	ErrNegativeShift = errors.New("negative shift")

	// ErrShiftOverflow is when you try to shift a number left too far.
	ErrShiftOverflow = errors.New("shift overflow")

	// ErrScriptVersion is when a restricted opcode tries to run under
	// an unsuitable script version.
	ErrScriptVersion = errors.New("script version")

	// ErrScriptFormat is for various script-parsing errors.
	ErrScriptFormat = errors.New("incorrect script format")
)

var (
	// ErrInvalidFlags is returned when the passed flags to NewScript
	// contain an invalid combination.
	ErrInvalidFlags = errors.New("invalid flags combination")

	// ErrInvalidIndex is returned when the passed input index for the
	// provided transaction is out of range.
	ErrInvalidIndex = errors.New("invalid input index")

	// ErrUnsupportedAddress is returned when a concrete type that
	// implements a btcutil.Address is not a supported type.
	ErrUnsupportedAddress = errors.New("unsupported address type")

	// ErrBadNumRequired is returned from MultiSigScript when nrequired is
	// larger than the number of provided public keys.
	ErrBadNumRequired = errors.New("more signatures required than keys present")
)