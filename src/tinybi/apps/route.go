// Copyright (C)2018 by Lei Peng <pyp126@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
package apps

import (
	"tinybi/webcore"
)

var WebRoutes map[string]webcore.WebApp

func init() {
	WebRoutes = make(map[string]webcore.WebApp)
	loadRoutes()
}

func loadRoutes() {
	//Register web routes here;
	WebRoutes["/"] = IndexApp{}
	WebRoutes["/index.html"] = IndexApp{}
	WebRoutes["/login.html"] = LoginApp{}
	WebRoutes["/concurrentTasks.html"] = ConcurrentTasksApp{}
	WebRoutes["/userProfile.html"] = UserProfileApp{}
	WebRoutes["/users.html"] = UsersApp{}
	WebRoutes["/roles.html"] = RolesApp{}
	WebRoutes["/API"] = ApiApp{}
}