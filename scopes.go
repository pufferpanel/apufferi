/*
 Copyright 2019 Padduck, LLC
  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at
  	http://www.apache.org/licenses/LICENSE-2.0
  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package apufferi

//generic
const ScopeLogin = "login"

//oauth2
const ScopeOauth2Info = "oauth2.info"

//server
const ScopeServerAdmin = "servers.admin"
const ScopeViewServers = "servers.view"
const ScopeEditServers = "servers.edit"
const ScopeEditServerAdmin = "servers.edit.admin"
const ScopeEditServerUsers = "servers.edit.users"
const ScopeCreateServers = "servers.create"

const ScopeInstallServers = "servers.install"
const ScopeServerConsole = "servers.console"
const ScopeServerSendConsole = "servers.console.send"
const ScopeStopServers = "servers.stop"
const ScopeStartServers = "servers.start"
const ScopeKillServers = "servers.kill"
const ScopeStatServers = "servers.stats"
const ScopeFilesServers = "servers.files"
const ScopeGetFilesServers = "servers.files.get"
const ScopePutFilesServers = "servers.files.put"

//node
const ScopeViewNodes = "nodes.view"
const ScopeEditNode = "nodes.edit"
const ScopeDeployNode = "nodes.deploy"

//template
const ScopeViewTemplates = "templates.view"

//user
const ScopeViewUsers = "users.view"
const ScopeEditUsers = "users.edit"

func GetDefaultUserServerScopes() []string {
	return []string{
		ScopeViewServers,
		ScopeServerConsole,
		ScopeStopServers,
		ScopeStartServers,
		ScopeKillServers,
		ScopeStatServers,
		ScopeFilesServers,
		ScopeGetFilesServers,
		ScopePutFilesServers,
	}
}
