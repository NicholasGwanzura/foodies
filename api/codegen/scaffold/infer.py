import re,sys
import scaffold as sc
PAST={'submit':'submitted','approve':'approved','close':'closed','complete':'completed','verify':'verified','review':'reviewed','acknowledge':'acknowledged','issue':'issued','return':'returned','revoke':'revoked','release':'released','cancel':'cancelled','suspend':'suspended','reject':'rejected','resolve':'resolved','receive':'received','confirm':'confirmed','execute':'executed','decide':'decided','restore':'restored','seal':'sealed','hand-over':'handed_over','dispatch':'dispatched','accept':'accepted','post':'posted','hold':'on_hold','match':'matched','validate':'validated','reimburse':'reimbursed','retire':'retired','publish':'published','activate':'live','reinstate':'active','lock':'locked','unlock':'active','disable':'disabled','reopen':'open'}
def actions():
    out=[]
    for t,r in sc.RES.items():
        own=r.get('get',{}).get('path')
        if not own: continue
        for o in sc.OPS:
            if o['method']=='post' and o['path'].startswith(own+'/') and o['path'].count('/')==own.count('/')+1 and not o['path'].endswith('}') and not o['operationId'].startswith('create'):
                out.append((t,o,o['path'].rsplit('/',1)[1]))
    return out
if __name__=='__main__':
    st=sc.T
    for t,o,v in actions():
        print(sc.T[t]['section'] if 0 else '', t, v, o['operationId'], 'req:',o.get('req'), 'resp:',o['resp'], '|', ' '.join(st[t]['status']))
