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

package scope

//generic
const Login = "login"

//oauth2
const OAuth2Info = "oauth2.info"

//server
const ServersAdmin = "servers.admin"
const ServersView = "servers.view"
const ServersEdit = "servers.edit"
const ServersEditAdmin = "servers.edit.admin"
const ServersEditUsers = "servers.edit.users"
const ServersCreate = "servers.create"
const ServersDelete = "servers.delete"

const ServersInstall = "servers.install"
const ServersConsole = "servers.console"
const ServersConsoleSend = "servers.console.send"
const ServersStop = "servers.stop"
const ServersStart = "servers.start"
const ServersStat = "servers.stats"
const ServersFiles = "servers.files"
const ServersFilesGet = "servers.files.get"
const ServersFilesPut = "servers.files.put"

//node
const NodesView = "nodes.view"
const NodesEdit = "nodes.edit"
const NodesDeploy = "nodes.deploy"

//template
const TemplatesView = "templates.view"

//user
const UsersView = "users.view"
const UsersEdit = "users.edit"

func ServersDefaultUser() []string {
	return []string{
		ServersView,
		ServersEdit,
		ServersInstall,
		ServersConsole,
		ServersConsoleSend,
		ServersStop,
		ServersStart,
		ServersStat,
		ServersFiles,
		ServersFilesGet,
		ServersFilesPut,
		ServersEditUsers,
	}
}

func DefaultAdmin() []string {
	return []string{
		ServersAdmin,
		NodesView,
		NodesEdit,
		NodesDeploy,
		TemplatesView,
		UsersView,
		UsersEdit,
		ServersCreate,
		ServersDelete,
		ServersEditAdmin,
	}
}
