// Package my provides a collection of utility functions for Go applications.
//
// This package includes various utility functions organized into several categories:
//
// String Utilities:
//   - IsEmpty: Checks if a string is empty
//   - Left, Right, Mid: Extract substrings from different positions
//   - Len: Count Unicode code points in a string
//   - Space: Generate a string of spaces
//
// Type Conversion:
//   - CStr: Convert various types to string
//   - CInt: Convert string to integer
//   - CDate: Parse string to time.Time
//   - BytesToString, StringToBytes: Convert between []byte and string
//
// Date and Time:
//   - Now: Get current time as formatted string
//   - FormatDateTime: Format time.Time according to specified pattern
//   - FriendlyTime: Get human-friendly time elapsed description
//
// Cryptography:
//   - MD5, SHA1, SHA256: Calculate hash values
//   - HMACSHA1, HMACSHA256: Generate keyed hash values
//
// File Operations:
//   - AppPath: Get application path
//   - FolderExist, FileExist: Check existence of folders and files
//   - MakeDir: Create directory
//   - ReadText, WriteText, AppendText: Read and write text files
//
// Network Operations:
//   - GetURL, PostURL: Make HTTP requests
//   - GetJSON, PostJSON: Handle JSON data over HTTP
//   - DownloadFile: Download file from URL
//   - RecoverWrap: HTTP handler with panic recovery
//   - RemoteIP, ClientIP: Get client IP addresses
//
// Configuration:
//   - LoadJSONConfig, LoadXMLConfig: Load configuration from files
//   - MustLoadConfig: Load configuration or panic
//
// Random Generation:
//   - RndNumber, RndAlpha, RndString: Generate random strings
//   - RndFilename: Generate random filenames
//
// Regular Expression:
//   - Test: Validate strings against common patterns
//
// The package aims to simplify common programming tasks by providing
// ready-to-use functions that handle frequent operations in Go applications.
package my
