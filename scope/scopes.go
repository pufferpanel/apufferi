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
const Oauth2Info = "oauth2.info"

//server
const ServerAdmin = "servers.admin"
const ViewServers = "servers.view"
const EditServers = "servers.edit"
const EditServerAdmin = "servers.edit.admin"
const EditServerUsers = "servers.edit.users"
const CreateServers = "servers.create"

const InstallServers = "servers.install"
const ServerConsole = "servers.console"
const ServerSendConsole = "servers.console.send"
const StopServers = "servers.stop"
const StartServers = "servers.start"
const KillServers = "servers.kill"
const StatServers = "servers.stats"
const FilesServers = "servers.files"
const GetFilesServers = "servers.files.get"
const PutFilesServers = "servers.files.put"

//node
const ViewNodes = "nodes.view"
const EditNode = "nodes.edit"
const DeployNode = "nodes.deploy"

//template
const ViewTemplates = "templates.view"

//user
const ViewUsers = "users.view"
const EditUsers = "users.edit"

func DefaultForServer() []string {
	return []string{
		ViewServers,
		ServerConsole,
		StopServers,
		StartServers,
		KillServers,
		StatServers,
		FilesServers,
		GetFilesServers,
		PutFilesServers,
		InstallServers,
		ServerSendConsole,
		EditServers,
		EditServerUsers,
	}
}
