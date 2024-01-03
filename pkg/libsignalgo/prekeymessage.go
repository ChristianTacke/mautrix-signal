// mautrix-signal - A Matrix-signal puppeting bridge.
// Copyright (C) 2023 Sumner Evans
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package libsignalgo

/*
#cgo LDFLAGS: -lsignal_ffi -ldl -lm
#include "./libsignal-ffi.h"
*/
import "C"
import "runtime"

type PreKeyMessage struct {
	nc  noCopy
	ptr *C.SignalPreKeySignalMessage
}

func wrapPreKeyMessage(ptr *C.SignalPreKeySignalMessage) *PreKeyMessage {
	preKeyMessage := &PreKeyMessage{ptr: ptr}
	runtime.SetFinalizer(preKeyMessage, (*PreKeyMessage).Destroy)
	return preKeyMessage
}

func DeserializePreKeyMessage(serialized []byte) (*PreKeyMessage, error) {
	var m *C.SignalPreKeySignalMessage
	signalFfiError := C.signal_pre_key_signal_message_deserialize(&m, BytesToBuffer(serialized))
	runtime.KeepAlive(serialized)
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return wrapPreKeyMessage(m), nil
}

func (m *PreKeyMessage) Clone() (*PreKeyMessage, error) {
	var cloned *C.SignalPreKeySignalMessage
	signalFfiError := C.signal_pre_key_signal_message_clone(&cloned, m.ptr)
	runtime.KeepAlive(m)
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return wrapPreKeyMessage(cloned), nil
}

func (m *PreKeyMessage) Destroy() error {
	m.CancelFinalizer()
	return wrapError(C.signal_pre_key_signal_message_destroy(m.ptr))
}

func (m *PreKeyMessage) CancelFinalizer() {
	runtime.SetFinalizer(m, nil)
}

func (m *PreKeyMessage) Serialize() ([]byte, error) {
	var serialized C.SignalOwnedBuffer = C.SignalOwnedBuffer{}
	signalFfiError := C.signal_pre_key_signal_message_serialize(&serialized, m.ptr)
	runtime.KeepAlive(m)
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return CopySignalOwnedBufferToBytes(serialized), nil
}

func (m *PreKeyMessage) GetVersion() (uint32, error) {
	var version C.uint
	signalFfiError := C.signal_pre_key_signal_message_get_version(&version, m.ptr)
	runtime.KeepAlive(m)
	if signalFfiError != nil {
		return 0, wrapError(signalFfiError)
	}
	return uint32(version), nil
}

func (m *PreKeyMessage) GetRegistrationID() (uint32, error) {
	var registrationID C.uint
	signalFfiError := C.signal_pre_key_signal_message_get_registration_id(&registrationID, m.ptr)
	runtime.KeepAlive(m)
	if signalFfiError != nil {
		return 0, wrapError(signalFfiError)
	}
	return uint32(registrationID), nil
}

func (m *PreKeyMessage) GetPreKeyID() (*uint32, error) {
	var preKeyID C.uint
	signalFfiError := C.signal_pre_key_signal_message_get_pre_key_id(&preKeyID, m.ptr)
	runtime.KeepAlive(m)
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	if preKeyID == C.uint(0xffffffff) {
		return nil, nil
	}
	return (*uint32)(&preKeyID), nil
}

func (m *PreKeyMessage) GetSignedPreKeyID() (uint32, error) {
	var signedPreKeyID C.uint
	signalFfiError := C.signal_pre_key_signal_message_get_signed_pre_key_id(&signedPreKeyID, m.ptr)
	runtime.KeepAlive(m)
	if signalFfiError != nil {
		return 0, wrapError(signalFfiError)
	}
	return uint32(signedPreKeyID), nil
}

func (m *PreKeyMessage) GetBaseKey() (*PublicKey, error) {
	var publicKey *C.SignalPublicKey
	signalFfiError := C.signal_pre_key_signal_message_get_base_key(&publicKey, m.ptr)
	runtime.KeepAlive(m)
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return wrapPublicKey(publicKey), nil
}

func (m *PreKeyMessage) GetIdentityKey() (*IdentityKey, error) {
	var publicKey *C.SignalPublicKey
	signalFfiError := C.signal_pre_key_signal_message_get_identity_key(&publicKey, m.ptr)
	runtime.KeepAlive(m)
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return &IdentityKey{wrapPublicKey(publicKey)}, nil
}

func (m *PreKeyMessage) GetSignalMessage() (*Message, error) {
	var message *C.SignalMessage
	signalFfiError := C.signal_pre_key_signal_message_get_signal_message(&message, m.ptr)
	runtime.KeepAlive(m)
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return wrapMessage(message), nil
}
