package okdoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var case1 = `
load('api://ctrlhack/hosts', api_hosts = 'hosts')
load('api://ctrlhack/tasks', api_tasks = 'tasks')
load('api://ctrlhack/shared_storage', api_shared_storage = 'shared_storage')

def _hosts_list(is_online=False, labels=tuple([])):
	check_point({'fn_name':'hosts.list','args':{'is_online':is_online,'labels':labels}}, debug=True)
	return api_hosts.list(is_online, labels)

def _hosts_get(host_id):
	check_point({'fn_name':'hosts.get','args':{'host_id':host_id}}, debug=True)
	return api_hosts.get(host_id)

def _hosts_force_shutdown(host_id):
	check_point({'fn_name':'hosts.force_shutdown','args':{'host_id':host_id}}, debug=True)
	return api_hosts.force_shutdown(host_id)

def _hosts_force_shutdown_by_task(task_id):
	check_point({'fn_name':'hosts.force_shutdown_by_task','args':{'task_id':task_id}}, debug=True)
	res = api_tasks.get(task_id)
	if res.result == False :
		return struct(**{'result':False, 'message':'not found'})
	return api_hosts.force_shutdown(res.task.host_id)

hosts = module(
	"hosts",
	list = _hosts_list,
	get = _hosts_get,
	force_shutdown = _hosts_force_shutdown,
	force_shutdown_by_task = _hosts_force_shutdown_by_task,
)

# find_options = {'name':'script_name', 'parent_task_id':task_id}
def _tasks_list(host_id=0, status=0, labels=tuple([]), group_key='', find_options=None):
	check_point({'fn_name':'tasks.list','args':{'host_id':host_id,'status':status,'labels':labels,'group_key':group_key,'find_options':find_options}}, debug=True)
	return api_tasks.list(host_id, status, labels, group_key, find_options)

def _tasks_get(task_id):
	check_point({'fn_name':'tasks.get','args':{'task_id':task_id}}, debug=True)
	return api_tasks.get(task_id)

def _tasks_launch(host_id, script_text, input_args, runner_type, wait=False, labels=tuple([]), group_key='', is_specific_user=False, deadline = 0, parent_task_id=0):
	check_point({'fn_name':'tasks.launch','args':{'host_id':host_id,'script_text':script_text,'input_args':input_args,'runner_type':runner_type,'wait':wait,'labels':labels,'group_key':group_key,'is_specific_user':is_specific_user,'deadline':deadline, 'parent_task_id':parent_task_id}}, debug=True)
	return api_tasks.launch(host_id, script_text, input_args, runner_type, wait, labels, group_key, is_specific_user, deadline, parent_task_id)

def _tasks_launch_id(host_id, script_id, input_args, runner_type, wait=False, labels=tuple([]), group_key='', is_specific_user=False, deadline = 0, parent_task_id=0):
	check_point({'fn_name':'tasks.launch_id','args':{'host_id':host_id,'script_id':script_id,'input_args':input_args,'runner_type':runner_type,'wait':wait,'labels':labels,'group_key':group_key,'is_specific_user':is_specific_user,'deadline':deadline, 'parent_task_id':parent_task_id}}, debug=True)
	return api_tasks.launch_id(host_id, script_id, input_args, runner_type, wait, labels, group_key, is_specific_user, deadline, parent_task_id)

def _tasks_reject(task_id, wait=False, deadline = 0):
	check_point({'fn_name':'tasks.reject','args':{'task_id':task_id,'wait':wait,'deadline':deadline}}, debug=True)
	return api_tasks.reject(task_id, wait, deadline)

def _tasks_rollback(task_id, wait=False, deadline = 0):
	check_point({'fn_name':'tasks.rollback','args':{'task_id':task_id,'wait':wait,'deadline':deadline}}, debug=True)
	return api_tasks.rollback(task_id, wait, deadline)

def _tasks_bstat(task_id):
	check_point({'fn_name':'tasks.bstat','args':{'task_id':task_id}}, debug=True)
	return api_tasks.bstat(task_id)

tasks = module(
	"tasks",
	list = _tasks_list,
	get = _tasks_get,
	launch = _tasks_launch,
	launch_id = _tasks_launch_id,
	reject = _tasks_reject,
	rollback = _tasks_rollback,
	bstat = _tasks_bstat,
	RUNNER_UNKNOWN = api_tasks.RUNNER_UNKNOWN,
	RUNNER_INSIDE = api_tasks.RUNNER_INSIDE,
	RUNNER_CHILD = api_tasks.RUNNER_CHILD,
	RUNNER_REPARENT = api_tasks.RUNNER_REPARENT,
	STRATEGY_UNKNOWN = api_tasks.STRATEGY_UNKNOWN,
	STRATEGY_ACTION = api_tasks.STRATEGY_ACTION,
	STRATEGY_ROLLBACK = api_tasks.STRATEGY_ROLLBACK,
	STATUS_UNKNOWN = api_tasks.STATUS_UNKNOWN,
	STATUS_DRAFT = api_tasks.STATUS_DRAFT,
	STATUS_NEW = api_tasks.STATUS_NEW,
	STATUS_RUNNING = api_tasks.STATUS_RUNNING,
	STATUS_RESIDENT_DONE = api_tasks.STATUS_RESIDENT_DONE,
	STATUS_FAIL = api_tasks.STATUS_FAIL,
	STATUS_COMPLETED = api_tasks.STATUS_COMPLETED,
	STATUS_RESIDENT_ACK = api_tasks.STATUS_RESIDENT_ACK,
	STATUS_REJECTED = api_tasks.STATUS_REJECTED,
)

# find_options = {'name':'script_name', 'mitre_code':'script_mitre_code', regexp: True} указать можно и по отдельности, namespace обязателен для фильтра и поддерживается 'task_scripts'
def _shared_storage_search(labels=tuple([]), page_size=0, page_token='', namespace='', find_options=None):
	check_point({'fn_name':'shared_storage.search','args':{'labels':labels,'page_size':page_size,'page_token':page_token,'namespace':namespace,'find_options':find_options}}, debug=True)
	return api_shared_storage.search(labels, page_size, page_token, namespace, find_options)

def _shared_storage_get_entry(entry_id):
	check_point({'fn_name':'shared_storage.get_entry','args':{'entry_id':entry_id}}, debug=True)
	return api_shared_storage.get_entry(entry_id)

def _shared_storage_task_script_by_name(name, regexp=False):
	check_point({'fn_name':'shared_storage.task_script_by_name','args':{'name':name, 'regexp':regexp}}, debug=True)
	return api_shared_storage.task_script_by_name(name, regexp)

def _shared_storage_task_scripts_by_labels(labels=tuple([]), page_size=0, page_token=''):
	check_point({'fn_name':'shared_storage.task_scripts_by_labels','args':{'labels':labels,'page_size':page_size,'page_token':page_token}}, debug=True)
	return api_shared_storage.search(labels=labels, page_size=page_size, page_token=page_token, namespace='task_scripts')

def _shared_storage_task_scripts_by_mitre_code(mitre_code='', regexp=False, page_size=0, page_token=''):
	check_point({'fn_name':'shared_storage.task_scripts_by_mitre_code','args':{'mitre_code':mitre_code,'regexp':regexp,'page_size':page_size,'page_token':page_token}}, debug=True)
	return api_shared_storage.search(find_options = {'mitre_code':mitre_code, 'regexp':regexp}, page_size=page_size, page_token=page_token, namespace='task_scripts')

def _shared_storage_kv_by_name(name, regexp=False):
	check_point({'fn_name':'shared_storage.kv_by_name','args':{'name':name,'regexp':regexp}}, debug=True)
	res = api_shared_storage.search(find_options = {'name':name, 'regexp':regexp}, page_size=1, page_token='', namespace='kv')
	if res.result == True and len(res.res.list_items) > 0:
		return struct(**{'result':res.result, 'message':res.message, 'res':res.res.list_items[0]})
	if res.result == True and len(res.res.list_items) == 0:
		return struct(**{'result':False, 'message':'not found', 'res':None})
	return struct(**{'result':res.result, 'message':res.message, 'res':None})

shared_storage = module(
	"shared_storage",
	search = _shared_storage_search,
	get_entry = _shared_storage_get_entry,
	task_script_by_name = _shared_storage_task_script_by_name,
	task_scripts_by_labels = _shared_storage_task_scripts_by_labels,
	task_scripts_by_mitre_code = _shared_storage_task_scripts_by_mitre_code,
	kv_by_name = _shared_storage_kv_by_name,
)

`
var gvgv = ""

var case2 = `
load("foo", bar = "baz")


	"""











			main docs


	"""
def main1(
	foo,
	bar = "baz",
	**kwargs):

	def main2(
		foo,
		bar = "baz",
		**kwargs):

		foo("bar") #inline comments
		foo = "bar"
		foo = [1,2,3]
		foo = (1,2,3)
		foo = {(1,2,3): "qwd"}
		foo = [
			1,
			2,
			3]
		foo = (
			1,
			2,3)
		"""
		multi line
		"""
		foo = {
			(1,2,3): "qwd"
		}
		module("fafa", abc = "qdwwq")
		return

    # inline comment

foo
		"""
		multi line
		"""
bar

tasks = module(
	"tasks",
	list = _tasks_list,
	get = _tasks_get,
	launch = _tasks_launch,
	launch_id = _tasks_launch_id,
	reject = _tasks_reject,
	rollback = _tasks_rollback,
	bstat = _tasks_bstat,
	RUNNER_UNKNOWN = api_tasks.RUNNER_UNKNOWN,
	RUNNER_INSIDE = api_tasks.RUNNER_INSIDE,
	RUNNER_CHILD = api_tasks.RUNNER_CHILD,
	RUNNER_REPARENT = api_tasks.RUNNER_REPARENT,
	STRATEGY_UNKNOWN = api_tasks.STRATEGY_UNKNOWN,
	STRATEGY_ACTION = api_tasks.STRATEGY_ACTION,
	STRATEGY_ROLLBACK = api_tasks.STRATEGY_ROLLBACK,
	STATUS_UNKNOWN = api_tasks.STATUS_UNKNOWN,
	STATUS_DRAFT = api_tasks.STATUS_DRAFT,
	STATUS_NEW = api_tasks.STATUS_NEW,
	STATUS_RUNNING = api_tasks.STATUS_RUNNING,
	STATUS_RESIDENT_DONE = api_tasks.STATUS_RESIDENT_DONE,
	STATUS_FAIL = api_tasks.STATUS_FAIL,
	STATUS_COMPLETED = api_tasks.STATUS_COMPLETED,
	STATUS_RESIDENT_ACK = api_tasks.STATUS_RESIDENT_ACK,
	STATUS_REJECTED = api_tasks.STATUS_REJECTED,
)
`

func TestPyParse(t *testing.T) {
	// pyDebug = 1
	// pyErrorVerbose = true

	t.Run("case1", func(t *testing.T) {
		res, err := PyParse([]byte(case1))
		assert.NoError(t, err)
		assert.Len(t, res.Imports, 3)
		assert.Len(t, res.Comments, 2)
		assert.Len(t, res.Functions, 17)
		assert.Len(t, res.Exports, 3)
	})

	t.Run("case2", func(t *testing.T) {
		res, err := PyParse([]byte(case2))
		assert.NoError(t, err)
		assert.Len(t, res.Imports, 1)
		assert.Len(t, res.Comments, 5)
		assert.Len(t, res.Functions, 2)
		assert.NotEmpty(t, res.Functions[0].Comment)
		assert.Empty(t, res.Functions[1].Comment)
		assert.Len(t, res.Exports, 2)
	})

}
