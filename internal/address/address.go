package address

import (
"fmt"
"path/filepath"
"strings"
)

// AddrStr represents an address string in the form:
// "path/to/file.ato:Entry.Path::instance.path"
type AddrStr string

// AddressError represents an error related to address parsing
type AddressError struct {
Message string
}

func (e *AddressError) Error() string {
return e.Message
}

// NewAddressError creates a new address error
func NewAddressError(message string) *AddressError {
return &AddressError{Message: message}
}

// FilePath returns the file path portion of the address
func (a AddrStr) FilePath() string {
return GetFile(a)
}

// EntrySection returns the entry section of the address
func (a AddrStr) EntrySection() (string, error) {
if entry := GetEntrySection(a); entry != "" {
return entry, nil
}
return "", NewAddressError("No entry section in address")
}

// FromParts creates an AddrStr from component parts
func FromParts(file string, entry string, instance string) AddrStr {
var parts []string

if file != "" {
parts = append(parts, file)
}

if entry != "" {
if len(parts) > 0 {
parts = append(parts, ":"+entry)
} else {
parts = append(parts, entry)
}
}

if instance != "" {
if entry != "" {
parts = append(parts, "::"+instance)
}
}

return AddrStr(strings.Join(parts, ""))
}

// GetFile extracts the file path from an address
func GetFile(address AddrStr) string {
parts := strings.Split(string(address), ":")
if len(parts) > 0 {
return parts[0]
}
return ""
}

// GetEntry extracts the root path from an address
func GetEntry(address AddrStr) string {
parts := strings.Split(string(address), "::")
if len(parts) > 0 {
return parts[0]
}
return ""
}

// GetEntrySection extracts the entry section from an address
func GetEntrySection(address AddrStr) string {
parts := strings.Split(string(address), ":")
if len(parts) >= 2 {
return parts[1]
}
return ""
}

// GetInstanceSection extracts the instance section from an address
func GetInstanceSection(address AddrStr) string {
parts := strings.Split(string(address), ":")
if len(parts) >= 4 {
return parts[3]
}
return ""
}

// GetName extracts the name from the end of the address
func GetName(address AddrStr) string {
str := string(address)
parts := strings.Split(str, ":")
if len(parts) > 0 {
lastPart := parts[len(parts)-1]
nameParts := strings.Split(lastPart, ".")
if len(nameParts) > 0 {
return nameParts[len(nameParts)-1]
}
}
return ""
}

// AddInstance adds an instance to an address
func AddInstance(address AddrStr, instance string) (AddrStr, error) {
if instance == "" {
return address, nil
}

currentInstance := GetInstanceSection(address)
entrySection := GetEntrySection(address)

if currentInstance != "" {
return AddrStr(string(address) + "." + instance), nil
} else if entrySection != "" {
return AddrStr(string(address) + "::" + instance), nil
}

return "", NewAddressError("Cannot add instance to something without an entry section")
}

// AddInstances adds multiple instances to an address
func AddInstances(address AddrStr, instances []string) (AddrStr, error) {
var err error
for _, instance := range instances {
address, err = AddInstance(address, instance)
if err != nil {
return "", err
}
}
return address, nil
}

// AddEntry adds an entry to an address
func AddEntry(address AddrStr, entry string) (AddrStr, error) {
if GetInstanceSection(address) != "" {
return "", NewAddressError("Cannot add entry to an instance address")
}

if GetEntrySection(address) == "" {
return AddrStr(string(address) + ":" + entry), nil
}

return AddrStr(string(address) + "." + entry), nil
}

// AddEntries adds multiple entries to an address
func AddEntries(address AddrStr, entries []string) (AddrStr, error) {
var err error
for _, entry := range entries {
address, err = AddEntry(address, entry)
if err != nil {
return "", err
}
}
return address, nil
}

// GetRelativeAddrStr returns the relative address from a base path
func GetRelativeAddrStr(address AddrStr, basePath string) (AddrStr, error) {
file := GetFile(address)
if file == "" {
return address, nil
}

relFile, err := filepath.Rel(basePath, file)
if err != nil {
return "", fmt.Errorf("failed to get relative path: %w", err)
}

return FromParts(relFile, GetEntrySection(address), GetInstanceSection(address)), nil
}

// String returns the string representation of the address
func (a AddrStr) String() string {
return string(a)
}
