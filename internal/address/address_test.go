package address

import (
"testing"
)

func TestFromParts(t *testing.T) {
tests := []struct {
name     string
file     string
entry    string
instance string
expected string
}{
{"File only", "test.ato", "", "", "test.ato"},
{"File and entry", "test.ato", "Module", "", "test.ato:Module"},
{"Full address", "test.ato", "Module", "instance", "test.ato:Module::instance"},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
result := FromParts(tt.file, tt.entry, tt.instance)
if string(result) != tt.expected {
t.Errorf("FromParts() = %q; want %q", result, tt.expected)
}
})
}
}

func TestGetFile(t *testing.T) {
addr := AddrStr("path/to/file.ato:Entry::instance")
expected := "path/to/file.ato"
result := GetFile(addr)

if result != expected {
t.Errorf("GetFile() = %q; want %q", result, expected)
}
}

func TestGetEntrySection(t *testing.T) {
addr := AddrStr("file.ato:Module.Entry::instance")
expected := "Module.Entry"
result := GetEntrySection(addr)

if result != expected {
t.Errorf("GetEntrySection() = %q; want %q", result, expected)
}
}

func TestGetInstanceSection(t *testing.T) {
addr := AddrStr("file.ato:Entry::instance.path")
expected := "instance.path"
result := GetInstanceSection(addr)

if result != expected {
t.Errorf("GetInstanceSection() = %q; want %q", result, expected)
}
}

func TestGetName(t *testing.T) {
tests := []struct {
address  AddrStr
expected string
}{
{"file.ato:Entry::instance.path", "path"},
{"file.ato:Module", "Module"},
{"simple", "simple"},
}

for _, tt := range tests {
t.Run(string(tt.address), func(t *testing.T) {
result := GetName(tt.address)
if result != tt.expected {
t.Errorf("GetName(%q) = %q; want %q", tt.address, result, tt.expected)
}
})
}
}

func TestAddInstance(t *testing.T) {
addr := AddrStr("file.ato:Entry")
result, err := AddInstance(addr, "instance1")

if err != nil {
t.Errorf("AddInstance() returned error: %v", err)
}

expected := "file.ato:Entry::instance1"
if string(result) != expected {
t.Errorf("AddInstance() = %q; want %q", result, expected)
}
}

func TestAddEntry(t *testing.T) {
addr := AddrStr("file.ato")
result, err := AddEntry(addr, "Module")

if err != nil {
t.Errorf("AddEntry() returned error: %v", err)
}

expected := "file.ato:Module"
if string(result) != expected {
t.Errorf("AddEntry() = %q; want %q", result, expected)
}
}

func TestAddressError(t *testing.T) {
err := NewAddressError("test error")
if err.Error() != "test error" {
t.Errorf("AddressError.Error() = %q; want %q", err.Error(), "test error")
}
}
