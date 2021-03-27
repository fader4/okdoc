
# @Method(
#     arguments = [
#        [ctx, "ctx param"],
#        [ctx.foo, "ctx.foo param"]
#     ]
# )
def main(ctx):

    if res.result == False:
        result.update(message = "mark_as_failed: can't get psexec.exe")
        # @Return("if not found file")
        return { "result": result, "state": state }

    cmd = psexec_run.format(e = psexec_local, h = host_t, u = user_t, p = pass_t, c = command_t)
    print("process.run: ", cmd)
    res = process.run(cmd = cmd, wait = True)

    if getattr(res, "log", None):
        result.update(log = res.log)

    if res.result:
        loot = log_parse(res.log)
        result.update(data = loot)
        state.update(result = True)
        result.update(message = "mark_as_success: psexec is ok")
        result.update(result = True)

    # @Return("success")
    return { "result": result, "state": state }

def rollback(ctx):
    result = {
        "got_state": ctx.rollback_state,
        "message": "mark_as_failed: rollback failed",
        "result": False
        }

    if ctx.rollback_state.get("result", False):
        # todo: drop created token
        result.update(message = "mark_as_success: all is fine")
        result.update(result = True)

    return { "result": result }

def log_parse(log):

    return log
