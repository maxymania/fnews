/*
Copyright (c) 2018 Simon Schmidt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/


package config

type AuthMethod struct {
	Method string   `inn:"$method"`
	Pwhash string   `inn:"$pwd-hash"`
	Hosts  []string `inn:"@host"`
	Dbname string   `inn:"$dbname"`
	Username string `inn:"$user"`
	Password string `inn:"$pass"`
}

type AuthNoAuth struct {
	NoPosting bool `inn:"$no-posting"`
	NoReading bool `inn:"$no-reading"`
}

type AuthCfg struct {
	Enable bool        `inn:"$enable"`
	Method *AuthMethod `inn:"$method"`
	NoAuth AuthNoAuth  `inn:"$no-auth"`
}


